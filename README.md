# Simple Hot Reloader

This is a lightweight hot reload service written in Go for your front-end projects, it deploys a local service that detects when files in a specified directory are updated.


## Configuration
Under `config.json` you can find the configuration of the Hot Reloader Server.
- `location`: Here you can specify the path you want the service to listen changes to.
- `port`: Here you can update the port where the service will be running.

## Run
```bash
go run .
```

## Server

After running the project go to: `http://localhost:8080/hot-reload`.
- `updated`: this attibute will change to `true` when there are changes in the folder specified in `config.json`
- `message`: this attribute will show which file changed that triggered the change.

## JavaScript Snippet

Implement the hot reload feature in any of your files by running the Go server, placing the "simple-hot-reload.js" file under your "js" folder and adding the following code to your header tag.

```html
<script>
  const HOT_RELOAD_PORT = 8080;
</script>
<script src="./js/simple-hot-reload.js"></script>
```

## Compile

If you want to compile the project to obtain a binary of it you can do so running the following commands:
```bash
# Creates executable
go build .

# Executes the server
./hotreload
```
