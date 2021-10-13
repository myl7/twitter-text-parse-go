FROM rust:alpine
COPY twitter-text-parse-go /app
WORKDIR /app
ENV target=x86_64-unknown-linux-musl
RUN apk add musl-dev gcc
RUN cargo build --release --target $target
RUN cp twitter-text-parse-go/target/$target/release/libtwitter_text_parse_go.a .
