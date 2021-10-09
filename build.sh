#!/bin/bash
set -euo pipefail

(cd twitter-text-parse-go && cargo build --release)
mkdir -p lib
cp twitter-text-parse-go/target/release/libtwitter_text_parse_go.a lib
