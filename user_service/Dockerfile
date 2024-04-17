ARG TARGET_DEFAULT=x86_64-unknown-linux-gnu

FROM rust:1.76.0 AS builder

ARG TARGET_DEFAULT
ENV TARGET=${TARGET_DEFAULT}

# RUN ln -s /bin/g++ /bin/musl-g++
RUN apt-get update -y && apt-get install -y cmake clang protobuf-compiler

RUN rustup target add "$TARGET"

COPY ./user_service /app/user_service
COPY ./tasks_service/api /app/tasks_service/api
WORKDIR /app/user_service

RUN cargo build --release --target "$TARGET"

FROM ubuntu:22.04

ARG TARGET_DEFAULT
ENV TARGET=${TARGET_DEFAULT}

COPY --from=builder /app/user_service/target/${TARGET}/release/user_service /bin/user_service
ENTRYPOINT ["/bin/user_service"]
