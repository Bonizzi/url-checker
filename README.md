# URL-CHECKER

To check 60 urls status from file.
By default the output path is the project folder, it's possible to specify a different path by using the flag `-path`, it's possible to specify a different local folder or a folder in a different path; In case the folder doesn't exist it'll be created.
The check can be performed in a sync mode which is the default way or in an async mode by specifying the flag `-async`.
For each mode, it's possible to create a file for each domain by using the flag `-split`.

## SYNC

The result of each domain in a single file

```
go run .
```

Create a file for each domain

```
go run . -split
```

To specify a different folder but in local path

```
go run . -path=<FOLDERNAME>
```

To specify a different path

```
go run . -path=<PATH/FOLDERNAME>
```

## ASYNC

The result of each domain in a single file

```
go run . -async
```

Create a file for each domain

```
go run . -async -split
```

To specify a different folder but in local path

```
go run . -async -path=<FOLDERNAME>
```

To specify a different path

```
go run . -async -path=<PATH/FOLDERNAME>
```
