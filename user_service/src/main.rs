use axum::{
    async_trait,
    extract::FromRequestParts,
    extract::State,
    http::{request::Parts, StatusCode},
    response::{IntoResponse, Response, Result},
    routing::post,
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
use std::{sync::Arc, thread::sleep, time::Duration};
use tokio::sync::RwLock;
use tokio_postgres::{tls::NoTlsStream, Client, NoTls, Socket};

/// User service
#[derive(Parser, Debug)]
#[command(version, about, long_about = None)]
struct Args {
    /// Number of times to greet
    #[arg(short = 'H', long, default_value = "0.0.0.0")]
    host: String,

    /// Number of times to greet
    #[arg(short, long, default_value = "3000")]
    port: u16,

    /// User database connection [config](https://docs.rs/tokio-postgres/latest/tokio_postgres/config/struct.Config.html)
    #[arg(short, long, default_value = "host=localhost user=postgres")]
    db_config: String,
}

/// TODO: Bcrypt for passwords, read about db_client (maybe possible to easily avoid rwlock), update spec, submit, read article, setup vpn and chat-gpt

#[tokio::main]
async fn main() {
    env_logger::init();

    let args = Args::parse();

    let db_client: Client;
    let db_connection: tokio_postgres::Connection<Socket, NoTlsStream>;
    loop {
        let result = tokio_postgres::connect(&args.db_config, NoTls).await;
        match result {
            Ok(x) => {
                (db_client, db_connection) = x;
                break;
            }
            Err(x) => {
                info!("retrying connection to user database: {}", x);
                sleep(Duration::from_secs(1));
            }
        }
    }
    info!("connected to user database");
    tokio::spawn(async move {
        if let Err(e) = db_connection.await {
            error!("user_database connection error: {}", e);
        }
    });

    let jwt_key = HS256Key::generate();

    let app_state = Arc::new(RwLock::new(AppState {
        user_database: db_client,
        jwt_key,
    }));
    let app = Router::new()
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

async fn register_handler(
    State(state): State<AppStateRef>,
    Json(auth_info): Json<AuthInfo>,
) -> Result<StatusCode, AppError> {
    info!("register_handler: Handling register request");

    state
        .read()
        .await
        .user_database
        .query(
            "INSERT INTO users (username, password) VALUES ($1, $2)",
            &[&auth_info.username, &auth_info.password],
        )
        .await
        .map_err(|_| AppError::IncorrectRequest)?;
    Ok(StatusCode::OK)
}

async fn login_handler(
    State(state): State<AppStateRef>,
    Json(auth_info): Json<AuthInfo>,
) -> Result<Json<AccessToken>, AppError> {
    info!("login_handler: Handling login request");

    let row = state
        .read()
        .await
        .user_database
        .query_one(
            "SELECT password FROM users WHERE username=$1",
            &[&auth_info.username],
        )
        .await
        .map_err(|_| AppError::NonExistingUser)?;
    let password: String = row.get(0);
    if auth_info.password != password {
        return Err(AppError::WrongPassword);
    }

    let claims = Claims::create(jwt_simple::prelude::Duration::from_hours(2))
        .with_subject(auth_info.username);
    let token = state.read().await.jwt_key.authenticate(claims).unwrap();
    Ok(Json(AccessToken { token }))
}

async fn update_handler(
    State(state): State<AppStateRef>,
    claims: AppClaims,
    Json(user_info): Json<UserInfo>,
) -> Result<StatusCode, AppError> {
    info!("update_handler: Handling update request");

    let username = claims.username;

    let mut state_lock = state.write().await;
    let tx = state_lock.user_database.transaction().await.unwrap();
    if let Some(first_name) = user_info.first_name {
        tx.query(
            "UPDATE users SET first_name=$1 WHERE username=$2",
            &[&first_name, &username],
        )
        .await
        .map_err(|_| AppError::IncorrectRequest)?;
    }
    if let Some(last_name) = user_info.last_name {
        tx.query(
            "UPDATE users SET last_name=$1 WHERE username=$2",
            &[&last_name, &username],
        )
        .await
        .map_err(|_| AppError::IncorrectRequest)?;
    }
    if let Some(date_of_birth) = user_info.date_of_birth {
        let date = NaiveDate::parse_from_str(&date_of_birth, "%Y-%m-%d")
            .map_err(|_| AppError::IncorrectDateFormat)?;
        tx.query(
            "UPDATE users SET date_of_birth=$1 WHERE username=$2",
            &[&date, &username],
        )
        .await
        .map_err(|_| AppError::IncorrectRequest)?;
    }
    if let Some(email) = user_info.email {
        tx.query(
            "UPDATE users SET email=$1 WHERE username=$2",
            &[&email, &username],
        )
        .await
        .map_err(|_| AppError::IncorrectRequest)?;
    }
    if let Some(phone_number) = user_info.phone_number {
        tx.query(
            "UPDATE users SET phone_number=$1 WHERE username=$2",
            &[&phone_number, &username],
        )
        .await
        .map_err(|_| AppError::IncorrectRequest)?;
    }
    tx.commit().await.unwrap();
    Ok(StatusCode::OK)
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
            .read()
            .await
            .jwt_key
            .verify_token::<NoCustomClaims>(bearer.token(), None)
            .map_err(|_| AppError::InvalidToken)?;
        Ok(AppClaims {
            username: jwt_claims.subject.ok_or(AppError::InvalidToken)?,
        })
    }
}

impl IntoResponse for AppError {
    fn into_response(self) -> Response {
        let (status, error_message) = match self {
            AppError::InvalidToken => (StatusCode::BAD_REQUEST, "Invalid token"),
            AppError::NonExistingUser => (StatusCode::BAD_REQUEST, "User doesn't exist"),
            AppError::WrongPassword => (StatusCode::FORBIDDEN, "Wrong password"),
            AppError::IncorrectRequest => (StatusCode::BAD_REQUEST, "Incorrect request"),
            AppError::IncorrectDateFormat => (StatusCode::BAD_REQUEST, "Incorrect date format. Expected YYYY-MM-DD"),
        };
        let body = Json(json!({
            "error": error_message,
        }));
        (status, body).into_response()
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

struct AppState {
    user_database: tokio_postgres::Client,
    jwt_key: HS256Key,
}
type AppStateRef = Arc<RwLock<AppState>>;

struct AppClaims {
    username: String,
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
