# Etapa de construccion
FROM rust:latest AS builder
WORKDIR /app
COPY . .
RUN cargo build --release

# Etapa de ejecucion
FROM ubuntu:22.04
WORKDIR /app

RUN mkdir -p /app

COPY --from=builder /app/target/release/redis-rust /app/redis-rust
COPY Rocket.toml /app/Rocket.toml

RUN apt-get update && apt-get install -y libssl-dev ca-certificates && apt-get clean

CMD ["./redis-rust"]
