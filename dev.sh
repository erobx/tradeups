#!/usr/bin/env bash

start_bun() {
    bun run dev
}

start_server() {
    docker compose up -d backend
    docker logs --follow tradeups_api
}

start_server
