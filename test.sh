#!/bin/bash

# `act` is a command line tool that help us test our GitHub Actions locally using Docker.
# See: https://github.com/nektos/act
act pull_request \
    --rebuild \
    -P ubuntu-latest=catthehacker/ubuntu:act-latest \
    -P ubuntu-22.04=catthehacker/ubuntu:act-22.04 \
    -P ubuntu-20.04=catthehacker/ubuntu:act-20.04 \
    -P ubuntu-18.04=catthehacker/ubuntu:act-18.04 \
    --secret-file ./.env \
    --container-architecture linux/amd64
