# logger
## Dependency
[zap](https://github.com/uber-go/zap)
[lumberjack](https://github.com/natefinch/lumberjack)
## Documentation
### SetEnvParam
#### Description
Use this function to set logger parameter
```go=
func SetEnvParam(folderString string, rotateCount int, logLevel string)
```
| Parameter    |                                      Description                                      |
| ------------ |:-------------------------------------------------------------------------------------:|
| folderString |                    Set to path for folder which store logger file                     |
| rotateCount  |                      To set rotate count of logger file's backup                      |
| logLevel     | [Enable log levels](https://pkg.go.dev/go.uber.org/zap@v1.18.1/zapcore#Level) which want to log out |

#### Usage
```go=
logger.SetEnvParam("Test", 5, "debug")
```
---
### New
#### Description
Use this function to get a zap logger. You can write log to file and console using this logger.
```go=
func New() *Logger
```
| Response |        Description         |
| -------- |:--------------------------:|
| logger   | zap logger for further use |

#### Usage
```go=
logger.SetEnvParam("Test", 5, "debug") // if do not call this function log file will save at ./
logger := logger.New()
logger.Info("Info Test")
```
---
### WriteHttpHandlerLog
#### Description
Use this function to write logs in unified file format.
```go=
func WriteHttpHandlerLog(logLevel zapcore.Level, statusCode int, message string, params map[string]string)
```
| Parameter    |                                      Description                                      |
| ------------ |:-------------------------------------------------------------------------------------:|
| logLevel     | [log level](https://pkg.go.dev/go.uber.org/zap@v1.18.1/zapcore#Level) which level of this log |
| statusCode |  [http status](https://pkg.go.dev/net/http#pkg-constants) which status of this log|
| message | error message if needed |
| params | parameters which want to log out|

#### Usage
```go=
var params = map[string]string{}
params["Value"] = "test"
logger.WriteHttpHandlerLog(zapcore.ErrorLevel, http.StatusInternalServerError, err.Error(), params)
```