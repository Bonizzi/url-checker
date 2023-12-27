# URL-CHECKER

To check 60 urls status from the file `url.txt`.
By default the output file name is `url_list_status.txt` located in the `tmp/` folder of the project, it's possible to specify a different file name by using the flag `-file`, in case the file doesn't exist, it'll be created. This option cannot be used with the flag `-split`.
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

To specify a different file name

```
go run . -file=<FILENAME>
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

To specify a different file name

```
go run . -async -file=<FILENAME>
```
