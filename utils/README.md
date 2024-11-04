# util
## Dependency
[uuid](https://github.com/google/uuid)  used in RandomUUIDString

## Table of contexts
* [AES](#AES)
    * [AesEncryptCBC](#AesEncryptCBC)
    * [AesDecryptCBC](#AesDecryptCBC)
* [AlgorithmUtil](#AlgorithmUtil)
    * [RandString](#RandString)
    * [RandomUUIDString](#RandomUUIDString)
    * [NewURN](#NewURN)
    * [ParseURN](#ParseURN)
* [ArgumentUtil](#ArgumentUtil) 
    * [GetAsByte](#GetAsByte)
    * [GetAsShort](#GetAsShort)
    * [GetAsInt](#GetAsInt)
    * [GetAsLong](#GetAsLong)
    * [GetAsFloat](#GetAsFloat)
    * [GetAsDouble](#GetAsDouble)
    * [GetAsString](#GetAsString)
    * [GetAsObject](#GetAsObject)
* [DumpStacks](#DumpStacks)
    * [DumpStacks](#DumpStacks)
    * [CurrentFunctionName](#CurrentFunctionName)
    * [CurrentCallerName](#CurrentCallerName)
* [misc](#misc)
    * [GetCurrentDirectory](#GetCurrentDirectory)
    * [GetAppName](#GetAppName)

---
## Documentation

## AES
### AesEncryptCBC
#### Description
Using this function to encrypt byte by AES algorithm using CBC mode
```go=
func AesEncryptCBC(origData []byte, key []byte) (encrypted []byte)
```
| Parameter |                       Description                        |
| --------- |:--------------------------------------------------------:|
| origData  |             Byte data which want to encrypt              |
| key       | key of encrypt algorithm(key must be 16, 24 or 32 bytes) |

| Response  |        Description         |
| --------- |:--------------------------:|
| encrypted | byte data after encrypted  |
| err       | go error(nil when success) |

#### Usage
```go=
encryptByte, err := util.AesEncryptCBC([]byte("test"), []byte("1234567890123456"))
```
---
### AesDecryptCBC
#### Description
Using this function to decrypt byte which encrypt by AES algorithm using CBC mode
```go=
func AesDecryptCBC(encrypted []byte, key []byte) (decrypted []byte)
```
| Parameter |                    Description                     |
| --------- |:--------------------------------------------------:|
| encrypted | byte which encrypt by AES algorithm using CBC mode |
| key       |            key which used when encrypt             |

| Response  |        Description         |
| --------- |:--------------------------:|
| decrypted | byte data after decrypted  |
| err       | go error(nil when success) |

#### Usage
```go=
// to get encryptByte
encryptByte, _ := util.AesEncryptCBC([]byte("test"), []byte("1234567890123456")) 
// decryptByte should be "test"
decryptByte, err := util.AesDecryptCBC(encryptByte, []byte("1234567890123456"))
```
---
## AlgorithmUtil
### RandString
#### Description
Use this function to get random string of given length
```go=
func RandString(n int) string
```
| Parameter |                Description                |
| --------- |:-----------------------------------------:|
| n         | length of random string which want to get |

| Response |          Description          |
| -------- |:-----------------------------:|
| string   | random string of given length |

#### Usage
```go=
length := 8
randomString := util.RandString(length)
```
---
### RandomUUIDString
#### Description
Use this function to generate a random UUID-V4 string
```go=
func RandomUUIDString() string
```
| Response |      Description      |
| -------- |:---------------------:|
| string   | random UUID-V4 string |

#### Usage
```go=
uuidString := util.RandomUUIDString()
```
---
### NewURN
#### Description
Use this function to generate a urn string
```go=
func NewURN(service string, resource string, tenant string, name string) string
```
| Parameter |   Description    |
| --------- |:----------------:|
| service   | name of service  |
| resource name | name of resource |
| account UUID    | UUID of account  |
| object UUID      | name of resource  |

**NOTICE**: The specific tenant with name parameters should be unique of entire system  

| Response | Description |
| -------- |:-----------:|
| string   | urn string  |

#### Usage
```go=
urnString := util.NewURN(
    "mss",
    "video",
    "9089f025-d49f-46a2-a62e-329febaffcfc",
    "56971a12-7200-4076-a97c-c88e2fa4d745"))
```
---
### ParseURN
#### Description
Parse urn string to urn object
```go=
type Urn struct {
	Service  string
	Resource string
    Account  string
	Object   string
}

func ParseURN(urn string) (Urn, error)
```
| Parameter | Description |
| --------- |:-----------:|
| urn       | urn string  |

| Response |        Description         |
| -------- |:--------------------------:|
| urn      |         urn object         |
| err      | go error(nil when success) |

#### Usage
```go=
urnString := util.NewURN("jms", "user", "1c07a064-d4c1-4e53-a137-f1fb2a515ab2"))
urnObject := util.ParseURN(urnString)
```
---
## ArgumentUtil
### GetAsByte
#### Description 
Use this function to get arguement in slice of interface with assertion type(byte)
```go=
func GetAsByte(arguments []interface{}, idx int, defaultValue int8) int8
```
| Parameter    |       Description       |
| ------------ |:-----------------------:|
| arguments    |  slice of interface{}   |
| idx          | index of target element |
| defaultValue |  return value if error  |

| Response |       Description       |
| -------- |:-----------------------:|
| value    | value of assertion type |

#### Usage
```go=
testSlice := []interface{}{int8(0), "1"} // example
getByte := util.GetAsByte(testSlice, 0, int8(1))
```
___
### GetAsShort
#### Description 
Use this function to get arguement in slice of interface with assertion type(short)
```go=
func GetAsShort(arguments []interface{}, idx int, defaultValue int16) int16
```
| Parameter    |       Description       |
| ------------ |:-----------------------:|
| arguments    |  slice of interface{}   |
| idx          | index of target element |
| defaultValue |  return value if error  |

| Response |       Description       |
| -------- |:-----------------------:|
| value    | value of assertion type |

#### Usage
```go=
testSlice := []interface{}{int16(0), "1"} // example
getShort := util.GetAsShort(testSlice, 0, int16(1))
```
___
### GetAsInt
#### Description 
Use this function to get arguement in slice of interface with assertion type(int)
```go=
func GetAsInt(arguments []interface{}, idx int, defaultValue int32) int32
```
| Parameter    |       Description       |
| ------------ |:-----------------------:|
| arguments    |  slice of interface{}   |
| idx          | index of target element |
| defaultValue |  return value if error  |

| Response |       Description       |
| -------- |:-----------------------:|
| value    | value of assertion type |

#### Usage
```go=
testSlice := []interface{}{int32(0), "1"} // example
getInt := util.GetAsInt(testSlice, 0, int32(1))
```
___

### GetAsLong
#### Description 
Use this function to get arguement in slice of interface with assertion type(long)
```go=
func GetAsLong(arguments []interface{}, idx int, defaultValue int64) int64
```
| Parameter    |       Description       |
| ------------ |:-----------------------:|
| arguments    |  slice of interface{}   |
| idx          | index of target element |
| defaultValue |  return value if error  |

| Response |       Description       |
| -------- |:-----------------------:|
| value    | value of assertion type |

#### Usage
```go=
testSlice := []interface{}{int64(0), "1"} // example
getLong := util.GetAsLong(testSlice, 0, int64(1))
```
___

### GetAsFloat
#### Description 
Use this function to get arguement in slice of interface with assertion type(float)
```go=
func GetAsFloat(arguments []interface{}, idx int, defaultValue float32) float32
```
| Parameter    |       Description       |
| ------------ |:-----------------------:|
| arguments    |  slice of interface{}   |
| idx          | index of target element |
| defaultValue |  return value if error  |

| Response |       Description       |
| -------- |:-----------------------:|
| value    | value of assertion type |

#### Usage
```go=
testSlice := []interface{}{float32(0.5), "1"} // example
getFloat := util.GetAsFloat(testSlice, 0, float32(0.0))
```
---
### GetAsDouble
#### Description 
Use this function to get arguement in slice of interface with assertion type(double)
```go=
func GetAsDouble(arguments []interface{}, idx int, defaultValue float64) float64
```
| Parameter    |       Description       |
| ------------ |:-----------------------:|
| arguments    |  slice of interface{}   |
| idx          | index of target element |
| defaultValue |  return value if error  |

| Response |       Description       |
| -------- |:-----------------------:|
| value    | value of assertion type |

#### Usage
```go=
testSlice := []interface{}{float64(0.5), "1"} // example
getDouble := util.GetAsDouble(testSlice, 0, float64(0.0))
```
---
### GetAsString
Use this function to get arguement in slice of interface with assertion type(string)
```go=
func GetAsString(arguments []interface{}, idx int, defaultValue string) string
```
| Parameter    |       Description       |
| ------------ |:-----------------------:|
| arguments    |  slice of interface{}   |
| idx          | index of target element |
| defaultValue |  return value if error  |

| Response |       Description       |
| -------- |:-----------------------:|
| value    | value of assertion type |

#### Usage
```go=
testSlice := []interface{}{"0", 1} // example
getString := util.GetAsString(testSlice, 0, "0")
```
---
### GetAsObject
Use this function to get arguement in slice of interface with assertion type(customized struct)
```go=
func GetAsObject(arguments []interface{}, idx int, defaultValue interface{}) interface{}
```
| Parameter    |       Description       |
| ------------ |:-----------------------:|
| arguments    |  slice of interface{}   |
| idx          | index of target element |
| defaultValue |  return value if error  |

| Response |       Description       |
| -------- |:-----------------------:|
| value    | value of assertion type |

#### Usage
```go=
type test struct {
	value string
}

testSlice := []interface{}{test{value: "test"}, "1"} //example
object := util.GetAsObject(testSlice, 0, test{})
```
---
## DumpStacks
### DumpStacks
#### Description
Use this function to get dump of stacks
```go=
func DumpStacks() string
```
| Response |      Description      |
| -------- |:---------------------:|
| string   | string of dump stacks |

#### Usage
```go=
dumpStacks := util.DumpStacks()
fmt.Printf("dumpStacks: %s \n", dumpStacks)
```
---
### CurrentFunctionName
#### Description
Use this function to get current function name
```go=
func CurrentFunctionName() string
```
| Response |           Description           |
| -------- |:-------------------------------:|
| string   | string of current function name |

#### Usage
```go=
name := util.CurrentFunctionName()
fmt.Printf("CurrentFunctionName: %s \n", name)
```
___
### CurrentCallerName
#### Description
Use this function to get current caller name
```go=
func CurrentCallerName() string
```
| Response |           Description           |
| -------- |:-------------------------------:|
| string   | string of current caller name |

#### Usage
```go=
name := util.CurrentCallerName()
fmt.Printf("CurrentCallerName: %s \n", name)
```
___
## misc
### GetCurrentDirectory
#### Description
Use this function to get the directory of go exe
```go=
func GetCurrentDirectory() (string, error)
```
| Response |         Description         |
| -------- |:---------------------------:|
| string   | current directory of go exe |
| err      | go error(nil when success)  |

#### Usage
```go=
directory, err := util.GetCurrentDirectory()
if err != nil {
    fmt.Println(err.Error())
}
fmt.Printf("Get current directory success, directory: %s \n", directory)
```
---
### GetAppName
#### Description
Use this function to get the name of go application
```go=
func GetAppName() string
```
| Response |      Description       |
| -------- |:----------------------:|
| string   | name of go application |

#### Usage
```go=
appName := util.GetAppName()
fmt.Printf("Get app name success, appName: %s \n", appName)
```