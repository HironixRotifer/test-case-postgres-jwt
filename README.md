# App

Модуль авторизации с помощью JWT

## Программный интерфейс

- HTTP 

    http://localhost:8080/api-v1/tokens - генерирует accessToken и refreshToken для пользователя в базе данных <br>
    Cтруктура запроса:
    ```
    {
    "guid": 3
    }   
    ```

    http://localhost:8080/api-v1/refresh - генерирует новую пару accessToken и refreshToken если accessToken истекает <br>
        Cтруктура запроса:
    ```
    {
    "refresh_token": "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjAsImlwIjoiIiwianRpIjowLCJleHAiOjE3MzUwNDI1NDF9.gIlHhC02bbfiiYmq2aqsGCp04RgEZqc4dvrWokEc9vuBuIgcqNmEUo6qpL_xV5snGBO3ZC1O7Fa-RfzGbMsl2w",
    "token": "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjMsImlwIjoiMTcyLjMwLjAuMSIsImp0aSI6MTczNDQzNzc0MTY4NywiZXhwIjoxNzM0NTI0MTQxfQ.pH2P88B3upatX_-4nti8jK-PntodrIL3jOlNrbMlhHQHs89N8gUVeUtgvR4ZcLKUCJWNCZTCztkfpMJkEsL3ZA"
    }  
    ```


## Запуск

Для управления сервисом используются скрипты

**start-app.sh** для поднятия контейнеров в docker и запуска сервиса.<br>
Запуск:
```
sh start-app.sh
```

**remove-app.sh** для остановки сервиса и удаления контейнеров.<br>
Запуск:
```
sudo remove-app.sh
```