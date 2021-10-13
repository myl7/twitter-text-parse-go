#!/bin/bash
set -euo pipefail

mkdir -p lib

target=x86_64-unknown-linux-gnu
(cd twitter-text-parse-go && cargo build --release --target $target)
cp twitter-text-parse-go/target/$target/release/libtwitter_text_parse_go.a lib/libtwitter_text_parse_go-$target.a

target=x86_64-unknown-linux-musl
docker build -t build -f build/musl/Dockerfile .
docker run --name build build
docker cp build:/app/libtwitter_text_parse_go.a lib/libtwitter_text_parse_go-$target.a
