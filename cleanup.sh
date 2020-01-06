#!/bin/bash

docker-compose down && \
docker rmi app:1.1 || true && \
docker rmi tvapp:1.1 || true && \
docker image prune -f && \
docker volume rm mediacenter_moviegobsvol && \
docker volume rm mediacenter_tvgobsvol && \
docker volume prune -f && \
docker network prune -f && \
docker stop $(docker ps -a -q) || true && \
docker rm -v $(docker ps -a -q) || true


