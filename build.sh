#!/bin/bash
set -euo pipefail

function build {
  (cd twitter-text-parse-go && cargo build --release --target $1)
  cp twitter-text-parse-go/target/$1/release/libtwitter_text_parse_go.a lib/libtwitter_text_parse_go-$1.a
}

mkdir -p lib
build x86_64-unknown-linux-gnu
build x86_64-unknown-linux-musl
