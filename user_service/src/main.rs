use axum::{
    async_trait,
    extract::{FromRequestParts, State},
    http::{request::Parts, StatusCode},
    response::{IntoResponse, Response, Result},
    routing::{get, post},
    Json, RequestPartsExt, Router,
};
use axum_extra::headers::authorization::Bearer;
use axum_extra::headers::Authorization;
use axum_extra::TypedHeader;
use chrono::NaiveDate;
use clap::Parser;
use jwt_simple::prelude::*;
use log::{error, info};
use serde::Deserialize;
use serde_json::json;
use std::env;
use std::{sync::Arc, time::Duration};
use tokio::time::sleep;
use tokio_postgres::{Client, NoTls};

#[tokio::main]
async fn main() {
    env_logger::init();
    let args = CliArgs::parse();

    let jwt_key = env::var("JWT_KEY")
        .map(|x| HS256Key::from_bytes(&const_hex::decode(x).unwrap()))
        .unwrap_or_else(|_| HS256Key::generate());
    let app_state = Arc::new(AppState {
        user_database: connect_user_database(&args.db_config).await,
        jwt_key,
    });

    let app = Router::new()
        .route("/", get(root_handler))
        .route("/register", post(register_handler))
        .route("/login", post(login_handler))
        .route("/update", post(update_handler))
        .with_state(app_state);

    let listener = tokio::net::TcpListener::bind(format!("{}:{}", args.host, args.port))
        .await
        .unwrap();
    info!("listening on {}", listener.local_addr().unwrap());
    axum::serve(listener, app).await.unwrap();
}

async fn connect_user_database(config: &str) -> Client {
    loop {
        match tokio_postgres::connect(config, NoTls).await {
            Ok((client, connection)) => {
                info!("connect_user_database: connected to user database");
                tokio::spawn(async move {
                    if let Err(e) = connection.await {
                        error!("user_database connection error: {}", e);
                    }
                });
                return client;
            }
            Err(e) => {
                error!(
                    "connect_user_database: couldn't connect to user_database: {}. Retrying ...",
                    e
                );
            }
        }
        sleep(Duration::from_secs(1)).await;
    }
}

async fn root_handler() -> &'static str {
    info!("root_handler: handling root request");
    "User service API"
}

async fn register_handler(
    State(state): State<AppStateRef>,
    Json(auth_info): Json<AuthInfo>,
) -> Result<StatusCode, AppError> {
    info!("register_handler: handling register request");

    let password_hash = bcrypt::hash(auth_info.password, 10).unwrap();
    state
        .user_database
        .query(
            "INSERT INTO users (username, password) VALUES ($1, $2)",
            &[&auth_info.username, &password_hash],
        )
        .await
        .map_err(|_| AppError::IncorrectRequest)?;
    Ok(StatusCode::OK)
}

async fn login_handler(
    State(state): State<AppStateRef>,
    Json(auth_info): Json<AuthInfo>,
) -> Result<Json<AccessToken>, AppError> {
    info!("login_handler: handling login request");

    let row = state
        .user_database
        .query_one(
            "SELECT password FROM users WHERE username=$1",
            &[&auth_info.username],
        )
        .await
        .map_err(|_| AppError::NonExistingUser)?;
    let password_hash: String = row.get(0);
    if !bcrypt::verify(auth_info.password, &password_hash).unwrap() {
        return Err(AppError::WrongPassword);
    }

    let claims = Claims::create(jwt_simple::prelude::Duration::from_hours(2))
        .with_subject(auth_info.username);
    let token = state.jwt_key.authenticate(claims).unwrap();
    Ok(Json(AccessToken { token }))
}

async fn update_handler(
    State(state): State<AppStateRef>,
    claims: AppClaims,
    Json(user_info): Json<UserInfo>,
) -> Result<StatusCode, AppError> {
    info!("update_handler: handling update request");

    let username = claims.username;
    let date_of_birth = match user_info.date_of_birth {
        Some(x) => Some(
            NaiveDate::parse_from_str(&x, "%Y-%m-%d").map_err(|_| AppError::IncorrectDateFormat)?,
        ),
        None => None,
    };

    state
        .user_database
        .query(
            "UPDATE 
                users 
            SET 
                first_name=COALESCE($1, first_name),
                last_name=COALESCE($2, last_name),
                date_of_birth=COALESCE($3, date_of_birth),
                email=COALESCE($4, email),
                phone_number=COALESCE($5, phone_number)
            WHERE
                username=$6",
            &[
                &user_info.first_name,
                &user_info.last_name,
                &date_of_birth,
                &user_info.email,
                &user_info.phone_number,
                &username,
            ],
        )
        .await
        .map_err(|_| AppError::IncorrectRequest)?;

    Ok(StatusCode::OK)
}

/// User service
#[derive(Parser, Debug)]
#[command(version, about, long_about = None)]
struct CliArgs {
    /// Service hostname
    #[arg(short = 'H', long, default_value = "0.0.0.0")]
    host: String,

    /// Service port
    #[arg(short, long, default_value = "3000")]
    port: u16,

    /// User database connection [config](https://docs.rs/tokio-postgres/latest/tokio_postgres/config/struct.Config.html)
    #[arg(short, long, default_value = "host=localhost user=postgres")]
    db_config: String,
}

struct AppState {
    user_database: tokio_postgres::Client,
    jwt_key: HS256Key,
}
type AppStateRef = Arc<AppState>;

struct AppClaims {
    username: String,
}

#[async_trait]
impl FromRequestParts<AppStateRef> for AppClaims {
    type Rejection = AppError;

    async fn from_request_parts(
        parts: &mut Parts,
        state: &AppStateRef,
    ) -> Result<Self, Self::Rejection> {
        // Extract the token from the authorization header
        let TypedHeader(Authorization(bearer)) = parts
            .extract::<TypedHeader<Authorization<Bearer>>>()
            .await
            .map_err(|_| AppError::InvalidToken)?;
        let jwt_claims = state
            .jwt_key
            .verify_token::<NoCustomClaims>(bearer.token(), None)
            .map_err(|_| AppError::InvalidToken)?;
        Ok(AppClaims {
            username: jwt_claims.subject.ok_or(AppError::InvalidToken)?,
        })
    }
}

#[derive(Debug)]
enum AppError {
    InvalidToken,
    NonExistingUser,
    WrongPassword,
    IncorrectRequest,
    IncorrectDateFormat,
}

impl IntoResponse for AppError {
    fn into_response(self) -> Response {
        let (status, error_message) = match self {
            AppError::InvalidToken => (StatusCode::FORBIDDEN, "Invalid token"),
            AppError::NonExistingUser => (StatusCode::BAD_REQUEST, "User doesn't exist"),
            AppError::WrongPassword => (StatusCode::FORBIDDEN, "Wrong password"),
            AppError::IncorrectRequest => (StatusCode::BAD_REQUEST, "Incorrect request"),
            AppError::IncorrectDateFormat => (
                StatusCode::BAD_REQUEST,
                "Incorrect date format. Expected YYYY-MM-DD",
            ),
        };
        let body = Json(json!({
            "error": error_message,
        }));
        (status, body).into_response()
    }
}

#[derive(Deserialize)]
#[serde(deny_unknown_fields)]
struct AuthInfo {
    username: String,
    password: String,
}

#[derive(Deserialize)]
#[serde(deny_unknown_fields)]
struct UserInfo {
    first_name: Option<String>,
    last_name: Option<String>,
    date_of_birth: Option<String>,
    email: Option<String>,
    phone_number: Option<String>,
}

#[derive(Serialize)]
struct AccessToken {
    token: String,
}
