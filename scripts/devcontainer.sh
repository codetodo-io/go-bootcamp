#!/bin/bash

if [ "${1-}" = "up" ]; then
    mkdir -p "${GO_BOOTCAMP_ROOT_DIR:-.}/volumes/vscode-extensions"
    chmod -R 777 "${GO_BOOTCAMP_ROOT_DIR:-.}/volumes"

    docker-compose -f ${GO_BOOTCAMP_ROOT_DIR:-.}/docker-compose-devcontainer.yml up -d
fi

if [ "${1-}" = "down" ]; then
    docker-compose -f ${GO_BOOTCAMP_ROOT_DIR:-.}/docker-compose-devcontainer.yml down
    rm -rf "${GO_BOOTCAMP_ROOT_DIR:-.}/volumes"
fi
