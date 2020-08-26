#!/bin/sh
docker-compose -f /home/pi/MediaCenter/docker-compose.yml down;
docker-compose -f /home/pi/yts-svlt/docker-compose.yml down;
docker-compose -f /home/pi/ampvel/docker-compose.yml down;
docker stop Movies-httpd;
docker stop TVShows-httpd;
docker image prune -f;
docker volume prune -f;
