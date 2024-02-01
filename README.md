# MySQL Driver in GO
It is an implementation of Go's ```database/sql/driver``` interface


## To Use
--> first intialize a repository [Optional]
<br>
--> install mysql package

## Intialize

```go mod init```
or specified a name to your repository like this
```go mod init github.com/shubhammishra-1```

## Install mysql driver package

```bash
go get -u github.com/go-sql-driver/mysql
```

## import it into your main file

``` _ github.com/go-sql-driver/mysql```

## Connection String

to connect with mysql database a string like URL must be parse

```username:password@tcp(127.0.0.1:3306)/database_name```

writing databse name is optional so, this also right you can swtich letter to any DB.

```username:password@tcp(127.0.0.1:3306)/```


After setuping see the code how to perform SQL operations









