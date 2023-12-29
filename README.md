# URL-CHECKER

To check 60 urls status from the file `url.txt`.
By default the output file name is `url_list_status.txt` located in the `tmp/` folder of the project, it's possible to specify a different file name by using the flag `-file`, in case the file doesn't exist, it'll be created. This option cannot be used with the flag `-split`.
The check can be performed in a sync mode which is the default way or in an async mode by specifying the flag `-async`.
For each mode, it's possible to create a file for each domain by using the flag `-split`.

## SYNC

The result of each domain in a single file

```
go run cmd/url-check/main.go
```

Create a file for each domain

```
go run cmd/url-check/main.go -split
```

To specify a different file name

```
go run cmd/url-check/main.go -file=<FILENAME>
```

## ASYNC

The result of each domain in a single file

```
go run cmd/url-check/main.go -async
```

Create a file for each domain

```
go run cmd/url-check/main.go -async -split
```

To specify a different file name

```
go run cmd/url-check/main.go -async -file=<FILENAME>
```

## Web Server

To have better control over the URL answers, we've defined a custom web server with the following details. Those are the technical specs of our server:

1. Hostname: `localhost:8000`
1. HTTP verb to use: `GET`
1. Endpoints:
   1. `/health` endpoint that immediately replies with the `200` status code
   1. `/slowdown?wait=<numberOfSecondsToWait>` endpoint that replies with the `200` status code after the amount of time specified in the query string is elapsed
   1. `/broken` endpoint that immediately replies with the `500` status code

To run the server:

```
go run cmd/web-server/main.go
```
