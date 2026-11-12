#!/bin/sh
# server-sdk/codegen/openapi.json 으로부터 portonev2/ 의 *.gen.go 를 재생성합니다.
# 요구 사항: deno, gofmt
set -e
cd "$(dirname "$0")"
deno run --no-lock --config server-sdk/codegen/deno.json \
  --allow-read --allow-write --allow-run=gofmt \
  codegen/main.ts
