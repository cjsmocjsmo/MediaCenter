#!/bin/sh
cd /home/pi/MediaCenter && docker-compose down;
cd /home/pi/yts-svlt && docker-compose down;
cd /home/pi/ampvel && docker-compose down;
cd && docker stop Movies-httpd;
docker stop TVShows-httpd;
docker image prune -f;
docker volume prune -f;
