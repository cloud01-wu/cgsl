# Redisx
## Dependency
[Redigo](https://github.com/gomodule/redigo)

## Documentation
* [Init](#Init)
* [New](#New)
* [Close](#Close)
* [Ping](#Ping)
* String type
    * [Get](#Get)
    * [Set](#Set)
* Hash type
    * [HGet](#HGet)
    * [HSet](#HSet)
* [Delete](#Delete)
* [Scan](#Scan)
* [GetKeys](#GetKeys)
* [Expire](#Expire)
* [Incr](#Incr)
### Init
#### Description
Use this function to creating and configuring a connection.

```go=
func Init(address string, password string, maxIdle int, maxActive int)
```

| Parameter |                                                                  Description                                                                   |
| --------- |:----------------------------------------------------------------------------------------------------------------------------------------------:|
| address   |                                                              Address of database                                                               |
| password  |                                                              Password of database                                                              |
| maxIdle   |                                                 Maximum number of idle connections in the pool                                                 |
| maxActive | Maximum number of connections allocated by the pool at a given time.<br>When zero, there is no limit on the number of connections in the pool. |

#### Usage
```go=
Init("127.0.0.1", "password", 16, 16)
```
---
### New
#### Description
Use this function to get a connection
<font color=#FF0000>Must call [Init](#Init) Once</font>
```go=
func New(db int) (redis.Conn, error)
```

| Parameter |               Description               |
| --------- |:---------------------------------------:|
| db        | Number of Database which want to select |

| Response |        Description         |
| -------- |:--------------------------:|
| conn     |         connection         |
| err      | go error(nil when success) |

#### Usage
```go=
// Call Init Once
conn, err := New(3)
```

---

### Close
#### Description
Use this function to close connection

```go=
func Close()
```

#### Usage
```go=
Close()
```
---
### Ping
#### Description
Use this function to test if connection success
<font color=#FF0000>Must call [Init](#Init) Once</font>

```go=
func Ping(db int) error
```

| Parameter |               Description               |
| --------- |:---------------------------------------:|
| db        | Number of Database which want to select |


| Response |        Description         |
| -------- |:--------------------------:|
| err      | go error(nil when success) |

#### Usage
```go=
// Call Init Once
err := Ping(3)
```
---

### Get
#### Description
Use this function to get value of string type
<font color=#FF0000>Must call [Init](#Init) Once</font>

```go=
func Get(db int, key string) ([]byte, error)
```
| Parameter |               Description               |
| --------- |:---------------------------------------:|
| db        | Number of Database which want to select |
| key       |          key which want to get          |


| Response |        Description         |
| -------- |:--------------------------:|
| result   |   result of instruction    |
| err      | go error(nil when success) |

#### Usage
```go=
// Call Init Once
result, err := Get(3, "Test")
```
---
### Set 
#### Description
Use this function to set value of string type
<font color=#FF0000>Must call [Init](#Init) Once</font>

```go=
func Set(db int, key string, value []byte) error
```
| Parameter |               Description               |
| --------- |:---------------------------------------:|
| db        | Number of Database which want to select |
| key       |        key which want to insert         |
| value     |       value which want to insert        |


| Response |        Description         |
| -------- |:--------------------------:|
| err      | go error(nil when success) |

#### Usage
```go=
// Call Init Once 
err := Set(3, "Test", byte[]("test"))
if err != nil {
    panic(err)
}
```
---
### Delete
#### Description
Use this function to delete key in database
<font color=#FF0000>Must call [Init](#Init) Once</font>

```go=
func Delete(db int, key string) error
```

| Parameter |               Description               |
| --------- |:---------------------------------------:|
| db        | Number of Database which want to select |
| key       |        key which want to delete         |


| Response |        Description         |
| -------- |:--------------------------:|
| err      | go error(nil when success) |

#### Usage
```go=
// Call Init Once
err := Delete(3, "test")
```
---
### HSet
#### Description
Use this function to set value of hash type
<font color=#FF0000>Must call [Init](#Init) Once</font>

```go=
func HSet(db int, args ...interface{}) error
```

| Parameter |                               Description                               |
| --------- |:-----------------------------------------------------------------------:|
| db        |                 Number of Database which want to select                 |
| args      | key, filed and value which want to select(filed and value must in pair) |

| Response |        Description         |
| -------- |:--------------------------:|
| err      | go error(nil when success) |

#### Usage
```go=
// Call Init Once
err := HSET(3, "book", "name", "test, "price", 100)
```
---
### HGet
#### Description
Use this function to get value of hash type
<font color=#FF0000>Must call [Init](#Init) Once</font>

```go=
func HGet(db int, args ...interface{}) ([]byte, error)
```
| Parameter |               Description               |
| --------- |:---------------------------------------:|
| db        | Number of Database which want to select |
| args      |     key and filed which want to Get     |

| Response |        Description         |
| -------- |:--------------------------:|
| result   |   result of instruction    |
| err      | go error(nil when success) |

#### Usage
```go=
result, err := HGet(3, "book", "name")
```
---
### Scan
#### Description
This function are used in order to incrementally iterate over a collection of elements
<font color=#FF0000>Must call [Init](#Init) Once</font>

| Parameter  |                              Description                              |
| ---------- |:---------------------------------------------------------------------:|
| db         |                Number of Database which want to select                |
| cursor     | [cursor of iterator](https://redis.io/commands/scan#scan-basic-usage) |
| keyPattern |                          Regular Expression                           |
| size       |                         size of the iterated                          |

| Response         |                Description                 |
| ---------------- |:------------------------------------------:|
| cursor           | cursor of iterator(0 when iterator finish) |
| result([]string) |         keys which satisfy pattern         |
| err              |         go error(nil when success)         |

#### Usage
---

### Expire
#### Description
Use this function to set a expire for a key
<font color=#FF0000>Must call [Init](#Init) Once</font>
```go=
func Expire(db int, key string, seconds int) error
```
| Parameter |               Description               |
| --------- |:---------------------------------------:|
| db        | Number of Database which want to select |
| key       |        key want to set a expire         |
| seconds   |           duration of expire            |

| Response |        Description         |
| -------- |:--------------------------:|
| err      | go error(nil when success) |

#### Usage
```go=
// set a key
err := Set(3, "expire", []byte("expire"))
if err != nil {
    fmt.Println(err.Error())
}
// set expire
err = Expire(3, "expire", 3)
if err != nil {
    fmt.Println(err.Error())
}
```
---
### GetKeys
#### Description
Use this function to get keys in database
<font color=#FF0000>Must call [Init](#Init) Once</font>

```go=
func GetKeys(db int, pattern string) ([]string, error)
```

| Parameter |               Description               |
| --------- |:---------------------------------------:|
| db        | Number of Database which want to select |
| pattern   |           Regular Expression            |

| Response         |        Description         |
| ---------------- |:--------------------------:|
| result([]string) | keys which satisfy pattern |
| err              | go error(nil when success) |

#### Usage
```go=
result, err := GetKeys(3, "*")
```
---
### Incr
#### Description
Use this function to increments the number stored at key by one.
<font color=#FF0000>Must call [Init](#Init) Once</font>

```go=
func Incr(db int, counterKey string) (int, error)
```
| Parameter |                Description                |
| --------- |:-----------------------------------------:|
| db        |  Number of Database which want to select  |
| key       | the key of number which want to increment |

| Response     |        Description         |
| ------------ |:--------------------------:|
| result(int)) | the number after increment |
| err          | go error(nil when success) |

#### Usage
```go=
//set a key of number
err := Set(3, "incr", []byte("1"))
if err != nil {
    fmt.Println(err.Error())
}
result, err := Incr(3, "incr") // result will be 2
if err != nil {
    fmt.Println(err.Error())
}
```