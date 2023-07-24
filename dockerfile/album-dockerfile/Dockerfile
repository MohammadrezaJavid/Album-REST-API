FROM golang:1.21rc2-alpine3.18 

WORKDIR /app

COPY ../../go.*             ./
RUN go mod download

COPY ../../*.go             ./
COPY ../../api/             ./api
COPY ../../authentication   ./authentication
COPY ../../config/          ./config
COPY ../../database/        ./database
COPY ../../docs/            ./docs
COPY ../../middlewares      ./middlewares
COPY ../../swagger          ./swagger

EXPOSE 8070

CMD [ "go", "run", "." ]
