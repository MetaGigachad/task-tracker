use crate::common::{AppClaims, AppError, AppStateRef};
use axum::{extract::State, http::StatusCode, response::Result, Json};
use chrono::NaiveDate;
use jwt_simple::prelude::*;
use log::info;
use serde::Deserialize;

pub async fn register_handler(
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

pub async fn login_handler(
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

pub async fn update_handler(
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

#[derive(Deserialize)]
#[serde(deny_unknown_fields)]
pub struct AuthInfo {
    username: String,
    password: String,
}

#[derive(Deserialize)]
#[serde(deny_unknown_fields)]
pub struct UserInfo {
    first_name: Option<String>,
    last_name: Option<String>,
    date_of_birth: Option<String>,
    email: Option<String>,
    phone_number: Option<String>,
}

#[derive(Serialize)]
pub struct AccessToken {
    token: String,
}
