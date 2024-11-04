# dbx
## Dependency
[Go-MySQL-Driver](https://github.com/go-sql-driver/mysql)
## Documentation
### Init
#### Description
Use this function to creating and configuring a connection.
```go=
func Init(endpoint string, username string, password string, dbName string, maxOpenConns int, maxIdleConns int) error
```

| Parameter    |                  Description                   |
| ------------ |:----------------------------------------------:|
| endpoint     |              Addredd of database               |
| username     |              username of database              |
| password     |                password of user                |
| dbName       |     Name of database which want to connect     |
| maxOpenConns | Maximum number of open connections in the pool |
| maxIdleConns | Maximum number of idle connections in the pool |

| Response |        Description         |
| -------- |:--------------------------:|
| err      | go error(nil when success) |

#### Usage
```go=
Init("127.0.0.1:3306", "user", "password", "DBName", 16, 16)
```
---
### New
#### Description
Use this function to get a connection
<font color=#FF0000>Must call [Init](#Init) First</font>
```go=
func New() *sql.DB
```
| Response |        Description         |
| -------- |:--------------------------:|
| conn     |         connection         |

#### Usage
```go=
//must call Init First
Init("127.0.0.1:3306", "user", "password", "DBName", 16, 16)
//Then get a connection
conn := New()
```
---
### Close
#### Description
Use this function to close a connection
```go=
func Close() 
```

#### Usage
```go=
//must call Init First
Init("127.0.0.1:3306", "user", "password", "DBName", 16, 16)
//Then get a connection
conn := New()
defer Close()
```
---
### GetData
#### Description
Use this function to get data from database
```go=
func GetData(rows *sql.Rows, object interface{}) error
```
| Parameter |        Description         |
| -------- |:--------------------------:|
| rows      | Select result |
| object | customized struct of selected result|


| Response |        Description         |
| -------- |:--------------------------:|
| err      | go error(nil when success) |
#### Usage
```go=
type Test struct {
	Col string `db:"Col1"`
}

//must call Init First
Init("127.0.0.1:3306", "user", "password", "DBName", 16, 16)
//Then get a connection
conn := New()
defer Close()
rows, err := conn.Query("SELECT `Col1` FROM `table1`")
if err != nil {
    fmt.Println(err.Error())
}
resultList := []Test{}
result := &Test{}
defer rows.Close()
for rows.Next() {
    _ = GetData(rows, result)
    fmt.Println(result.Col)
    resultList = append(resultList, *result)
}
fmt.Println(resultList)
```