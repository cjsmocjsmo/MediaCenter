#!/bin/sh
docker system prune -f;
cd /home/pi/MediaCenter && docker-compose up -d;
cd /home/pi/yts-svlt && docker-compose up -d;
cd /home/pi/ampvel && docker-compose up -d;
docker start Movies-httpd;
docker start TVShows-httpd;