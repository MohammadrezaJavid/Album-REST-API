# album

Implementation of a rest api named album with the help of gin framework

![album rest api swagger ui](./image/album_rest_api.png)

```bash
# Build image
docker build -t album-app:latest .

# Use image without docker-compose.
docker run -itd -p 8080:80 album-app:latest

# Use image by docker-compose.
docker-compose up -d

# For down docker-compose.
docker-compose down -v

# Run rest api without docker
go run main.go
```