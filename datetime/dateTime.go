package datetime

import (
	"time"

	"github.com/vjeantet/jodaTime"
)

const (
	TimeFormatMilli  = "YYYY-MM-ddTHH:mm:ss.SSSZ"
	TimeFormatSecond = "YYYY-MM-ddTHH:mm:ssZ"
)

type DateTime struct {
	timeObject time.Time
}

func Now() *DateTime {
	return &DateTime{time.Now()}
}

func FromString(timestamp string) (*DateTime, error) {
	var dateTime DateTime
	timeObject, err := jodaTime.Parse(TimeFormatMilli, timestamp)
	if nil == err {
		dateTime = DateTime{timeObject}
	}

	return &dateTime, err
}

func FromStringWithFormat(timestamp string, format string) (*DateTime, error) {
	var dateTime DateTime
	timeObject, err := jodaTime.Parse(format, timestamp)
	if nil == err {
		dateTime = DateTime{timeObject}
	}

	return &dateTime, err
}

func FromStringWithLayout(timestamp string, layout string) (*DateTime, error) {
	timeObject, err := time.Parse(layout, timestamp)
	if err != nil {
		return nil, err
	}

	return &DateTime{timeObject}, nil
}

func FromUnixTime(unixTime int64) *DateTime {
	var dateTime DateTime
	timeObject := time.Unix(unixTime, 0)
	dateTime = DateTime{timeObject}

	return &dateTime
}

func FromTime(time time.Time) *DateTime {
	return &DateTime{
		timeObject: time,
	}
}

func (dateTime *DateTime) SetTZ(tz string) error {
	loc, err := time.LoadLocation(tz)
	if err != nil {
		return err
	}

	dateTime.timeObject = dateTime.timeObject.In(loc)
	return nil
}

func (dateTime *DateTime) GetTime() time.Time {
	return dateTime.timeObject
}

func (dateTime *DateTime) String() string {
	return jodaTime.Format(TimeFormatSecond, dateTime.timeObject)
}

func (dateTime *DateTime) StringWithFormat(format string) string {
	return jodaTime.Format(format, dateTime.timeObject)
}

func (dateTime *DateTime) EpochInSecond() int64 {
	return dateTime.timeObject.Unix()
}

func (dateTime *DateTime) EpochInMilli() int64 {
	return dateTime.timeObject.UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))
}

func (dateTime *DateTime) WeekDay() int {
	return (int)(dateTime.timeObject.Weekday())
}

func (dateTime *DateTime) AddSecond(value int) *DateTime {
	dateTime.timeObject = dateTime.timeObject.Add(time.Second * time.Duration(value))
	return dateTime
}

func (dateTime *DateTime) AddMinute(value int) *DateTime {
	dateTime.timeObject = dateTime.timeObject.Add(time.Minute * time.Duration(value))
	return dateTime
}

func (dateTime *DateTime) AddHour(value int) *DateTime {
	dateTime.timeObject = dateTime.timeObject.Add(time.Hour * time.Duration(value))
	return dateTime
}

func (dateTime *DateTime) AddDay(value int) *DateTime {
	dateTime.timeObject = dateTime.timeObject.AddDate(0, 0, value)
	return dateTime
}

func (dateTime *DateTime) AddMonth(value int) *DateTime {
	dateTime.timeObject = dateTime.timeObject.AddDate(0, value, 0)
	return dateTime
}

func (dateTime *DateTime) AddYear(value int) *DateTime {
	dateTime.timeObject = dateTime.timeObject.AddDate(value, 0, 0)
	return dateTime
}
