FROM golang:1.21rc2-alpine3.18 

WORKDIR /app

COPY go.* ./
COPY *.go ./
COPY api/ ./api
COPY config/ ./config
COPY database/ ./database
COPY migrations/ ./migrations

RUN go mod download && go build main.go

EXPOSE 80

CMD [ "/app/main" ]