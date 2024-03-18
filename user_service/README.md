# User Service

Simple http service for authentication and management of user data.

OpenAPI specification can be found [here](openapi.yaml).

### How to build

You can build locally with `cargo` or build [docker image](Dockerfile).

### Deploy options

**Private key** for JWT authentication can be provided by `JWT_KEY` environment variable.

Other options are described in `--help`.
