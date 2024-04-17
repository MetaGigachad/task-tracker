use crate::proto::tasks_service::tasks_service_client::TasksServiceClient;
use axum::{
    async_trait,
    extract::FromRequestParts,
    http::{request::Parts, StatusCode},
    response::{IntoResponse, Response, Result},
    Json, RequestPartsExt,
};
use axum_extra::headers::authorization::Bearer;
use axum_extra::headers::Authorization;
use axum_extra::TypedHeader;
use jwt_simple::prelude::*;
use serde_json::json;
use std::sync::Arc;
use tokio::sync::RwLock;
use tonic::transport::Channel;

pub struct AppState {
    pub user_database: tokio_postgres::Client,
    pub jwt_key: HS256Key,
    pub tasks_service: RwLock<TasksServiceClient<Channel>>,
}
pub type AppStateRef = Arc<AppState>;

pub struct AppClaims {
    pub username: String,
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
pub enum AppError {
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
