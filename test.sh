#!/bin/bash

act pull_request \
    -P self-hosted=catthehacker/ubuntu:act-22.04 \
    --secret-file ./.env \
    --container-architecture linux/amd64
