docker-login:
	@ docker login --username ${DOCKER_USER} --password ${DOCKER_PASS} private.jlinux.ir

build-images:
	@ docker build -t private.jlinux.ir/app/album:latest -f ./dockerfile/album-dockerfile/Dockerfile .
	@ docker build -t private.jlinux.ir/database/my-mysql:8.0 -f ./dockerfile/mysql-dockerfile/Dockerfile .

push-images:
	@ docker push private.jlinux.ir/app/album:latest
	@ docker push private.jlinux.ir/database/my-mysql:8.0

compose-up:
	@ docker-compose up -d

compose-down:
	@ docker-compose down -v

compose-logs:
	@ docker-compose logs -f

mysql-up:
	@ docker-compose -f ./mysql-docker-compose.yml up -d

mysql-down:
	@ docker-compose -f ./mysql-docker-compose.yml down -v

install-requirements:
	@ go mod tidy
	@ go mod download

run:
	@ go run main.go

test:
	@ go test -v -cover ./unitTest/

production-up:
	@ ssh -o StrictHostKeyChecking=no ${SERVER_HOST}@${SERVER_IP} "cd ~/var/production/ && docker login --username ${DOCKER_USER} --password ${DOCKER_PASS} private.jlinux.ir && docker-compose pull && sudo docker-compose up -d"