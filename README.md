# URL-CHECKER

To check 60 urls status from file.
The check can be performed in a sync mode which is the default way or in an async mode by specifying the flag -async.
For each mode it's possible to create a file for each domain by using the flag -split.

## SYNC

The result of each domain in a single file

```
go run .
```

Create a file for each domain

```
go run . -split
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
