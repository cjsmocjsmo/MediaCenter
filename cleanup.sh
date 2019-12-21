#!/bin/bash

docker-compose down && \
docker rmi app || true && \
docker rmi tvapp || true && \
docker image prune -f && \
docker volume rm docgobuild_moviegobsvol && \
docker volume rm tvgobs_tvgobsvol && \
docker volume prune -f && \
docker network prune -f && \
docker stop $(docker ps -a -q) || true && \
docker rm -v $(docker ps -a -q) || true


