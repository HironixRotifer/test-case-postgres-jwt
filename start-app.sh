#!/bin/bash

CONTAINER_NAME=app

clear
echo "Запуск контейнера: $CONTAINER_NAME"

docker-compose up -d --build db
docker-compose up -d --build migrator

# sleep 5
sh wait-for-postgres.sh

docker-compose up -d --build $CONTAINER_NAME
    
echo "Контейнер $CONTAINER_NAME Запущен."

docker-compose logs -f $CONTAINER_NAME