mod auth;
mod common;
mod proto;
mod tasks;

use axum::{
    routing::{get, post},
    Router,
};
use clap::Parser;
use common::AppState;
use jwt_simple::prelude::*;
use log::{error, info};
use proto::tasks_service::tasks_service_client::TasksServiceClient;
use std::env;
use std::{sync::Arc, time::Duration};
use tokio::time::sleep;
use tokio_postgres::{Client, NoTls};
use tonic::transport::Channel;

#[tokio::main]
async fn main() {
    env_logger::init();
    let args = CliArgs::parse();

    let jwt_key = env::var("JWT_KEY")
        .map(|x| HS256Key::from_bytes(&const_hex::decode(x).unwrap()))
        .unwrap_or_else(|_| HS256Key::generate());
    let app_state = Arc::new(AppState {
        user_database: connect_user_database(&args.db_config).await,
        tasks_service: tokio::sync::RwLock::new(
            connect_tasks_service(args.tasks_service_host).await,
        ),
        jwt_key,
    });

    let app = Router::new()
        .route("/", get(root_handler))
        .route("/register", post(auth::register_handler))
        .route("/login", post(auth::login_handler))
        .route("/update", post(auth::update_handler))
        .route("/createTask", post(tasks::create_task_handler))
        .route("/getTask", post(tasks::get_task_handler))
        .route("/updateTask", post(tasks::update_task_handler))
        .route("/deleteTask", post(tasks::delete_task_handler))
        .route("/getTaskPage", post(tasks::get_task_page_handler))
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

async fn connect_tasks_service(host: String) -> TasksServiceClient<Channel> {
    loop {
        match TasksServiceClient::connect(host.clone()).await {
            Ok(client) => {
                info!("connect_tasks_service: connected to tasks_service");
                return client;
            }
            Err(e) => {
                error!(
                    "connect_tasks_service: couldn't connect to tasks_service: {}. Retrying ...",
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

    /// Hostname of tasks_service
    #[arg(short, long, default_value = "tasks_service:50051")]
    tasks_service_host: String,
}
