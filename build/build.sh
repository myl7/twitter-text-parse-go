#!/bin/bash
set -euo pipefail

mkdir -p lib
mkdir -p lib/x86_64-unknown-linux-gnu
mkdir -p lib/x86_64-unknown-linux-musl

target=x86_64-unknown-linux-gnu
(cd twitter-text-parse-go && cargo build --release --target $target)
cp twitter-text-parse-go/target/$target/release/libtwitter_text_parse_go.a lib/$target/libtwitter_text_parse_go.a

target=x86_64-unknown-linux-musl
docker build -t build -f build/musl.dockerfile .
docker run --name build build
docker cp build:/app/libtwitter_text_parse_go.a lib/$target/libtwitter_text_parse_go.a
