#!/bin/bash

clear

echo "Удаление контейнеров"

docker-compose stop db
docker-compose stop migrator
docker-compose stop app
docker-compose stop redis


docker-compose rm db migrator app redis
rm -rf .database/
rm -rf redisdata/


echo "Сервис успешно остановлен и удалён"
