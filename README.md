# Simple Hot Reloader

This is a lightweight hot reload service written in Go for your front-end projects, it deploys a local service that hot reloads your any file you specify in any directory you specify.

## Configuration
Under `config.json` you can find the configuration of the Hot Reloader Server.
- `location`: Here you can specify the path you want the service to listen changes to.
- `port`: Here you can update the port where the service will be running.
- `filename`: Here you can specify the name of the file you want to see served and hot-reloaded.

## Run
```bash
go run .
```

## Compile

If you want to compile the project to obtain a binary of it you can do so running the following commands:
```bash
# Creates executable
go build .

# Executes the server
./hotreload
```
