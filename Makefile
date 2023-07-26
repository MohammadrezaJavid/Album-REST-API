build-images:
	@ docker build -t album:latest ./dockerfile/album-dockerfile
	@ docker build -t my-mysql:8.0 ./dockerfile/mysql-dockerfile

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