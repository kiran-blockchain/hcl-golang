

---

## Environment
To edit environment variables, create a file with name `.env` and copy the contents from `.env.default` to start with.

| Var Name  | Type  | Default | Description  |
|---|---|---|---|
| JWT_SECRET  | string  | `secret` |JWT secret to verify  |
|  PORT | number  | `8080` | Port to run the API server on |
|  MONGO_URL | string  | `mongodb://localhost:27017/db` | URL for MongoDB |

<!-- ## Logging
The application uses [winston](https://github.com/winstonjs/winston) as the default logger. The configuration file is at `src/logger.ts`.
* All logs are saved in `./logs` directory and at `/logs` in the docker container.
* The `docker-compose` file has a volume attached to container to expose host directory to the container for writing logs.
* Console messages are prettified
* Each line in error log file is a stringified JSON. -->


### Directory Structure

```
+-- controllers
|   +-- personController.go
+-- db
|   +-- db.go
+-- handlers
|   +-- config.go
|   +-- logs.go
|   +-- response.go
|   +-- verifyJWT.go
+-- models
|   +-- models.go
+-- validators
|   +-- validators.go
+-- tests
|   +-- api_test.go
+-- routes
|   +-- routes.go
+-- uploaded
+-- vendor
+-- nginx
|   +-- dev.conf.d
|   |   +-- nginx.conf
+-- .env
+-- .env.default
+-- .gitignore
+-- docker-compose.yml
+-- Dockerfile
+-- go.mod
+-- go.sum
+-- main.go
+-- README.md
```