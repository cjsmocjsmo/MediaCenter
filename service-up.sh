#!/bin/sh
cd;
docker system prune -f;
docker-compose -f /home/pi/MediaCenter/docker-compose.yml -d up;
docker-compose -f /home/pi/yts-svlt/docker-compose.yml -d up;
docker-compose -f /home/pi/ampvel/docker-compose.yml -d up;
docker start Movies-httpd;
docker start TVShows-httpd;