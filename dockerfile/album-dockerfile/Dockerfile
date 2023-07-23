FROM golang:1.21rc2-alpine3.18 

WORKDIR /app

COPY ../../go.*        ./
RUN go mod download

COPY ../../*.go        ./
COPY ../../api/        ./api
COPY ../../config/     ./config
COPY ../../database/   ./database
COPY ../../migrations/ ./migrations
COPY ../../docs/       ./docs
COPY ../../vendor/     ./vendor

EXPOSE 8070

CMD [ "go", "run", "." ]
