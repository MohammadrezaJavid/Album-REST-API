build-images:
	@ docker build -t album:latest -f ./dockerfile/album-dockerfile/Dockerfile .
	@ docker build -t my-mysql:8.0 -f ./dockerfile/mysql-dockerfile/Dockerfile .

compose-up:
	@ docker-compose up -d

compose-down:
	@ docker-compose down -v

compose-logs:
	@ docker-compose logs -f

run:
	@ go run main.go

test:
	@ go test -v -cover ./unitTest/

mysql-up:
	@ docker-compose -f ./mysql-docker-compose.yml up -d

mysql-down:
	@ docker-compose -f ./mysql-docker-compose.yml down -v

push-images:
	# @ TODO tags images
	# @ TODO push images