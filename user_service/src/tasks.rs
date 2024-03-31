use crate::common::{AppClaims, AppError, AppStateRef};
use crate::proto::tasks_service as ts;
use axum::{extract::State, response::Result, Json};
use jwt_simple::prelude::*;
use log::info;
use serde::Deserialize;

pub async fn create_task_handler(
    State(state): State<AppStateRef>,
    claims: AppClaims,
    Json(req): Json<CreateTaskRequest>,
) -> Result<Json<Task>, AppError> {
    info!("create_task_handler: handling create task request");
    let request = tonic::Request::new(ts::CreateTaskRequest {
        user_id: claims.username,
        title: req.title,
        description: req.description,
    });
    let response = state
        .tasks_service
        .write()
        .await
        .create_task(request)
        .await
        .map_err(|_| AppError::IncorrectRequest)?;
    match response.into_inner().response.unwrap() {
        ts::task_response::Response::Task(x) => Ok(Json(Task {
            id: x.id,
            created_at: x.created_at.unwrap().to_string(),
            title: x.title,
            description: x.description,
            status: ts::TaskStatus::try_from(x.status)
                .unwrap()
                .as_str_name()
                .to_string(),
        })),
        ts::task_response::Response::Error(_) => Err(AppError::IncorrectRequest),
    }
}

pub async fn get_task_handler(
    State(state): State<AppStateRef>,
    claims: AppClaims,
    Json(req): Json<GetTaskRequest>,
) -> Result<Json<Task>, AppError> {
    info!("get_task_handler: handling get task request");
    let request = tonic::Request::new(ts::GetTaskRequest {
        user_id: claims.username,
        task_id: req.task_id,
    });
    let response = state
        .tasks_service
        .write()
        .await
        .get_task(request)
        .await
        .map_err(|_| AppError::IncorrectRequest)?;
    match response.into_inner().response.unwrap() {
        ts::task_response::Response::Task(x) => Ok(Json(Task {
            id: x.id,
            created_at: x.created_at.unwrap().to_string(),
            title: x.title,
            description: x.description,
            status: ts::TaskStatus::try_from(x.status)
                .unwrap()
                .as_str_name()
                .to_string(),
        })),
        ts::task_response::Response::Error(_) => Err(AppError::IncorrectRequest),
    }
}

pub async fn update_task_handler(
    State(state): State<AppStateRef>,
    claims: AppClaims,
    Json(req): Json<UpdateTaskRequest>,
) -> Result<Json<Task>, AppError> {
    info!("update_task_handler: handling update task request");
    let request = tonic::Request::new(ts::UpdateTaskRequest {
        user_id: claims.username,
        task_id: req.task_id,
        new_title: req.new_title,
        new_description: req.new_description,
    });
    let response = state
        .tasks_service
        .write()
        .await
        .update_task(request)
        .await
        .map_err(|_| AppError::IncorrectRequest)?;
    match response.into_inner().response.unwrap() {
        ts::task_response::Response::Task(x) => Ok(Json(Task {
            id: x.id,
            created_at: x.created_at.unwrap().to_string(),
            title: x.title,
            description: x.description,
            status: ts::TaskStatus::try_from(x.status)
                .unwrap()
                .as_str_name()
                .to_string(),
        })),
        ts::task_response::Response::Error(_) => Err(AppError::IncorrectRequest),
    }
}

pub async fn delete_task_handler(
    State(state): State<AppStateRef>,
    claims: AppClaims,
    Json(req): Json<DeleteTaskRequest>,
) -> Result<Json<Task>, AppError> {
    info!("delete_task_handler: handling delete task request");
    let request = tonic::Request::new(ts::DeleteTaskRequest {
        user_id: claims.username,
        task_id: req.task_id,
    });
    let response = state
        .tasks_service
        .write()
        .await
        .delete_task(request)
        .await
        .map_err(|_| AppError::IncorrectRequest)?;
    match response.into_inner().response.unwrap() {
        ts::task_response::Response::Task(x) => Ok(Json(Task {
            id: x.id,
            created_at: x.created_at.unwrap().to_string(),
            title: x.title,
            description: x.description,
            status: ts::TaskStatus::try_from(x.status)
                .unwrap()
                .as_str_name()
                .to_string(),
        })),
        ts::task_response::Response::Error(_) => Err(AppError::IncorrectRequest),
    }
}

pub async fn get_task_page_handler(
    State(state): State<AppStateRef>,
    claims: AppClaims,
    Json(req): Json<GetTaskPageRequest>,
) -> Result<Json<TaskPage>, AppError> {
    info!("get_task_page_handler: handling get task page request");
    let request = tonic::Request::new(ts::GetTaskPageRequest {
        user_id: claims.username,
        start_id: req.start_id,
        page_size: req.page_size,
    });
    let response = state
        .tasks_service
        .write()
        .await
        .get_task_page(request)
        .await
        .map_err(|_| AppError::IncorrectRequest)?;
    match response.into_inner().response.unwrap() {
        ts::task_page_response::Response::TaskPage(mut x) => Ok(Json(TaskPage {
            tasks: x
                .tasks
                .iter_mut()
                .map(|t| Task {
                    id: t.id.clone(),
                    created_at: t.created_at.clone().unwrap().to_string(),
                    title: t.title.clone(),
                    description: t.description.clone(),
                    status: ts::TaskStatus::try_from(t.status)
                        .unwrap()
                        .as_str_name()
                        .to_string(),
                })
                .collect(),
        })),
        ts::task_page_response::Response::Error(_) => Err(AppError::IncorrectRequest),
    }
}

#[derive(Deserialize)]
#[serde(deny_unknown_fields)]
pub struct CreateTaskRequest {
    title: String,
    description: String,
}

#[derive(Deserialize)]
#[serde(deny_unknown_fields)]
pub struct GetTaskRequest {
    task_id: String,
}

#[derive(Deserialize)]
#[serde(deny_unknown_fields)]
pub struct UpdateTaskRequest {
    task_id: String,
    new_title: Option<String>,
    new_description: Option<String>,
}

#[derive(Deserialize)]
#[serde(deny_unknown_fields)]
pub struct DeleteTaskRequest {
    task_id: String,
}

#[derive(Deserialize)]
#[serde(deny_unknown_fields)]
pub struct GetTaskPageRequest {
    start_id: i32,
    page_size: i32,
}

#[derive(Serialize)]
pub struct Task {
    id: String,
    created_at: String,
    title: String,
    description: String,
    status: String,
}

#[derive(Serialize)]
pub struct TaskPage {
    tasks: Vec<Task>,
}
