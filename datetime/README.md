# Datetime
## Dependency
[jodaTime](https://github.com/vjeantet/jodaTime)

## Documentation
This package use the following struct
```go=
type DateTime struct {
    timeObject time.Time
}
```
---
### Now
#### Description
Use this function to get DateTime struct included current time

```go=
func Now() *DateTime
```

| Response          |        Description         |
| ----------------- |:--------------------------:|
| DateTime(pointer) | Pointer of DateTime struct |

#### Usage

```go=
dateTime := Now()
```
---

### FromString
#### Description
Use this function to get DateTime struct from string which format is "YYYY-MM-ddTHH:mm:ss.SSSZ"

```go=
func FromString(timestamp string) (*DateTime, error)
```

| Parameter |         Description          |
| --------- |:----------------------------:|
| timestamp | string which want to convert |

| Response          |        Description         |
| ----------------- |:--------------------------:|
| DateTime(pointer) | Pointer of DateTime struct |
| err               | go error(nil when success) |

#### Usage
```go=
dateReceive, err := FromString("2021-07-01T12:00:00.000+0800")
```
---

### FromStringWithFormat
#### Description
Use this function to get DateTime struct from string with customized format

```go=
func FromStringWithFormat(timestamp string, format string) (*DateTime, error)
```

| Parameter |         Description          |
| --------- |:----------------------------:|
| timestamp | string which want to convert |
| format    |      customized format       |

| Response          |        Description         |
| ----------------- |:--------------------------:|
| DateTime(pointer) | Pointer of DateTime struct |
| err               | go error(nil when success) |

### Usage
```go=
dateReceive, err := FromStringWithFormat("2021-07-01T04:00:00", "YYYY-MM-ddTHH:mm:ss")
```

---

## FromUnixTime
### Description
Use this function to get DateTime struct from Unix time stamp
```go=
func FromUnixTime(unixTime int64) *DateTime
```

| Parameter |  Description   |
| --------- |:--------------:|
| unixTime  | Unix timestamp |

| Response          |        Description         |
| ----------------- |:--------------------------:|
| DateTime(pointer) | Pointer of DateTime struct |

### Usage

```go=
dateReceive := FromUnixTime(1625112000) // 2021/7/12 12:00:00
```
---
## String
### Description
Use this function to get string which format is "YYYY-MM-ddTHH:mm:ss.SSSZ" from DataTime struct

```go=
func (dateTime *DateTime) String() string
```
| Response |                      Description                       |
| -------- |:------------------------------------------------------:|
| string   | Time string which format is "YYYY-MM-ddTHH:mm:ss.SSSZ" |

### Usage
```go=
dateTime := &DateTime{
    timeObject: time.Date(2021, time.July, 1, 12, 0, 0, 0, time.UTC),
} // Example for DateTime struct
dateTimeString := dateTime.String() //Get "2021-07-01T12:00:00.000+0800"
```
---

## StringWithFormat
### Description
Use this function to get string with customized format from DataTime struct

```go=
(dateTime *DateTime) StringWithFormat(format string) string
```

| Parameter |         Description          |
| --------- |:----------------------------:|
| format    |      customized format       |

| Response |            Description             |
| -------- |:----------------------------------:|
| string   | Time string with customized format |

### Usage
```go=
dateTime := &DateTime{
    timeObject: time.Date(2021, time.July, 1, 12, 0, 0, 0, time.UTC),
} // Example for DateTime struct
dateTimeString := dateTime.StringWithFormat("YYYY-MM-ddTHH:mm:ss") //Get "2021-07-01T12:00:00"
```
---
## EpochInSecond
### Description
Use this function to get epoch in second from DateTime struct
```go=
func (dateTime *DateTime) EpochInSecond() int64
```
| Response |   Description   |
| -------- |:---------------:|
| epoch    | epoch in second |

### Usage
```go=
dateTime := &DateTime{
    timeObject: time.Date(2021, time.July, 1, 12, 0, 0, 0, time.UTC),
} // Example for DateTime struct
second := dateTime.EpochInSecond() //Get 1625140800
```

---
## EpochInMilli
### Description
Use this function to get epoch in milliseconds from DateTime struct
```go=
func (dateTime *DateTime) EpochInMilli() int64 
```

| Response |      Description      |
| -------- |:---------------------:|
| epoch    | epoch in milliseconds |

### Usage
```go=
dateTime := &DateTime{
    timeObject: time.Date(2021, time.July, 1, 12, 0, 0, 0, time.UTC),
} // Example for DateTime struct
second := dateTime.EpochInMilli() //Get 1625140800000
```
---
## WeekDay
### Description
Use this function to get week day of DateTime
```go=
func (dateTime *DateTime) WeekDay() int
```

| Response |           Description            |
| -------- |:--------------------------------:|
| weekDay  | weekDay of DateTime(Sunday -> 0) |

### Usage
```go=
dateTime := &DateTime{
    timeObject: time.Date(2021, time.July, 1, 12, 0, 0, 0, time.UTC),
} // Example for DateTime struct
weekDay := dateTime.WeekDay() //Get 4
```
---

## AddSecond
### Description
Use this function to add seconds for DateTime struct
```go=
func (dateTime *DateTime) AddSecond(value int) *DateTime
```

| Parameter |     Description     |
| --------- |:-------------------:|
| value     | seconds want to add |

| Response          |        Description         |
| ----------------- |:--------------------------:|
| DateTime(pointer) | Pointer of DateTime struct |

### Usage
```go=
dateTime := &DateTime{
    timeObject: time.Date(2021, time.July, 1, 12, 0, 0, 0, time.UTC),
} // Example for DateTime struct

dateTime = dateTime.AddSecond(10)
```
---
## AddMinute
### Description
Use this function to add minutes for DateTime struct
```go=
func (dateTime *DateTime) AddMinute(value int) *DateTime
```

| Parameter |     Description      |
| --------- |:--------------------:|
| value     | minutes want to add |

| Response          |        Description         |
| ----------------- |:--------------------------:|
| DateTime(pointer) | Pointer of DateTime struct |

### Usage
```go=
dateTime := &DateTime{
    timeObject: time.Date(2021, time.July, 1, 12, 0, 0, 0, time.UTC),
} // Example for DateTime struct

dateTime = dateTime.AddMinute(10)
```
---

## AddHour
### Description
Use this function to add hours for DateTime struct
```go=
func (dateTime *DateTime) AddHour(value int) *DateTime
```

| Parameter |    Description    |
| --------- |:-----------------:|
| value     | hours want to add |

| Response          |        Description         |
| ----------------- |:--------------------------:|
| DateTime(pointer) | Pointer of DateTime struct |

### Usage
```go=
dateTime := &DateTime{
    timeObject: time.Date(2021, time.July, 1, 12, 0, 0, 0, time.UTC),
} // Example for DateTime struct

dateTime = dateTime.AddHour(10)
```
---
## AddDay
### Description
Use this function to add days for DateTime struct
```go=
func (dateTime *DateTime) AddDay(value int) *DateTime
```

| Parameter |   Description    |
| --------- |:----------------:|
| value     | days want to add |

| Response          |        Description         |
| ----------------- |:--------------------------:|
| DateTime(pointer) | Pointer of DateTime struct |

### Usage
```go=
dateTime := &DateTime{
    timeObject: time.Date(2021, time.July, 1, 12, 0, 0, 0, time.UTC),
} // Example for DateTime struct

dateTime = dateTime.AddDay(10)
```
---
## AddMonth
### Description
Use this function to add months for DateTime struct
```go=
func (dateTime *DateTime) AddMonth(value int) *DateTime
```

| Parameter |    Description     |
| --------- |:------------------:|
| value     | months want to add |

| Response          |        Description         |
| ----------------- |:--------------------------:|
| DateTime(pointer) | Pointer of DateTime struct |

### Usage
```go=
dateTime := &DateTime{
    timeObject: time.Date(2021, time.July, 1, 12, 0, 0, 0, time.UTC),
} // Example for DateTime struct

dateTime = dateTime.AddMonth(10)
```
---
## AddYear
### Description
Use this function to add years for DateTime struct
```go=
func (dateTime *DateTime) AddYear(value int) *DateTime
```

| Parameter |    Description     |
| --------- |:------------------:|
| value     | years want to add |

| Response          |        Description         |
| ----------------- |:--------------------------:|
| DateTime(pointer) | Pointer of DateTime struct |

### Usage
```go=
dateTime := &DateTime{
    timeObject: time.Date(2021, time.July, 1, 12, 0, 0, 0, time.UTC),
} // Example for DateTime struct

dateTime = dateTime.AddYear(10)
```
---
