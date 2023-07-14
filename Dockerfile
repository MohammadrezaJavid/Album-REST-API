FROM golang:1.21rc2-alpine3.18 

WORKDIR /app

COPY go.* ./
COPY *.go ./
COPY service/ ./service
COPY infrastructure/ ./infrastructure


RUN go mod download && go build main.go

EXPOSE 80

CMD [ "/app/main" ]
