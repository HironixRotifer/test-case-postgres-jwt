FROM postgres:latest

ENV DB_USER=rotifer
ENV DB_PASSWORD=rotifer
ENV DB_NAME=users
ENV HOST=localhost
ENV DB_HOST=db
ENV PORT=8080
ENV DB_PORT=5432
ENV DB_SSLMODE=disable
ENV DB_MIGRATIONTABLE=users

COPY /migrations /docker-entrypoint-initdb.d/

EXPOSE 5431

FROM golang:1.23.1-alpine as builder

WORKDIR /test-case-postgres-jwt

RUN apk --no-cache add bash git make gcc gettext musl-dev

COPY ["./go.mod", "./go.sum", "./"]
RUN  go mod download

COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -o . ./cmd/app
RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -o . ./cmd/migrator
RUN chmod +x wait-for-postgres.sh

EXPOSE 8080

CMD ["./app"]