# album

Implementation of a rest api named album with the help of gin framework

![album rest api swagger ui](./image/ScreenshotalbumAPIwithJWT.png)


### Note: 
    - Before running the api, first customize the .env file.
    - In this project, two Dockerfiles are defined, one of which is for multi stage.
## usage:
```bash
# Build images
make build-images

# Use image by docker-compose.
make compose-up

# For down docker-compose.
make compose-down

# For show logs
make compose-logs
```
- Run album REST API without docker
    - First install mysql using [here](https://dev.mysql.com/doc/refman/8.0/en/linux-installation.html)
    - ```bash
    make run
    ```

- Run Unit Test for this project
```bash
make test
```

- For access Swagger page
click [here](http://localhost:8070/api/docs/index.html)
