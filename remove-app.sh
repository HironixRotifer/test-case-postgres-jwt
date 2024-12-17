#!/bin/bash

clear

echo "Удаление контейнеров"

docker-compose stop db
docker-compose stop migrator
docker-compose stop app

docker-compose rm db migrator app
rm -rf .database/

echo "Сервис успешно остановлен и удалён"
