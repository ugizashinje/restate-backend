# warrant-api

Warrant api backend

## Start app for development

```bash
# Start postgres and pgadmin
$ cd docker & docker-compose up -d

# Start api
$ go run ./cmd/main.go

# Start api with hot reload
$ gin -i --build cmd/ -a 9876 cmd/main.go 
```

## Create swagger docs

Pre-requisites:

- Install swag:

```
go install github.com/swaggo/swag/cmd/swag@latest
```

Generate docs:

```bash
# -d is the endpoints directory and -g represents where gin is initialized
$ swag init -d pkg/server/endpoints/v1/  -g ../../start.go --parseDependency --parseInternal
```

Docs are available at http://localhost:9876/swagger/index.html and are saved at /docs

## Deployment

After successfully ssh-ing into the server, run the following commands:

```bash
# Locate the directory where the app is located
$ cd warrant-api

# Pull the latest changes from the repo
$ git pull origin develop

# Build the app
$ go build ./cmd/main.go

# Stop the app if it is running
$ pkill -f main

# Run the app
$ nohup ./main > output.log 2>&1 &

# Check if the app is running and monitor the logs
$ tail -f output.log
```

if you need to restart the database, run the following commands:

```bash
# Locate the directory where the app is located
$ cd warrant-api

# Go to the docker directory
$ cd docker

# Stop the containers
$ docker-compose down

# Restart the containers
$ docker-compose up -d
```
