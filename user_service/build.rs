fn main() -> Result<(), Box<dyn std::error::Error>> {
    tonic_build::configure()
        .build_server(false)
        .out_dir("src/proto")
        .compile(
            &["../tasks_service/api/tasks_service.proto"],
            &["../tasks_service/api"],
        )?;
    Ok(())
}