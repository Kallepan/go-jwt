#!/bin/bash

docker-compose down -v
docker-compose -f docker-compose.yaml up --build