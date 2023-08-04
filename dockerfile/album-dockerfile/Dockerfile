# Stage one build go app
# FROM public.jlinux.ir/language/go/golang:1.19 AS builder

FROM public.jlinux.ir/language/go/golang:1.19
WORKDIR /app

COPY ../../go.*              ./
RUN go mod download

COPY ../../*.go              ./
COPY ../../api/              ./api
COPY ../../authentication/   ./authentication
COPY ../../config/           ./config
COPY ../../database/         ./database
COPY ../../docs/             ./docs
COPY ../../middlewares/      ./middlewares
COPY ../../swagger/          ./swagger

RUN CGO_ENABLE=0 GOOS=linux GOARCH=amd64 go build main.go

# Stage two create image 
# FROM scratch
# COPY --from=builder /app/main /

EXPOSE 8070

CMD [ "go", "run", "." ]