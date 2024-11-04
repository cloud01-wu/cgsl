package datetime

import (
	"testing"
	"time"
)

var secondsEastOfUTC = int((8 * time.Hour).Seconds())
var TaiwanStandardTime = time.FixedZone("UTC+8", secondsEastOfUTC)
var date = time.Date(2021, time.July, 1, 12, 0, 0, 0, TaiwanStandardTime)

func TestNow(t *testing.T) {
	expectTime := time.Now()
	dateTime := Now()

	if expectTime.Equal(dateTime.timeObject) {
		t.Log("Test Now function success")
	} else {
		t.Error("Now function error")
		t.Errorf("Expect get %s but get %s", expectTime.Local(), dateTime.timeObject.Local())
	}
}

func TestFromString(t *testing.T) {
	dateReceive, err := FromString("2021-07-01T12:00:00.000+0800")

	if err != nil {
		t.Error(err.Error())
	}
	if dateReceive.timeObject.Equal(date) {
		t.Log("Get date object from string success")
	} else {
		t.Error("Get date object from string error")
		t.Errorf("Expect get at %s but get %s \n", date.Local(), dateReceive.timeObject.Local())
	}
}

func TestFromStringWithFormat(t *testing.T) {
	dateReceive, err := FromStringWithFormat("2021-07-01T04:00:00", "YYYY-MM-ddTHH:mm:ss")

	if err != nil {
		t.Error(err.Error())
	}
	if dateReceive.timeObject.Equal(date) {
		t.Log("Get date object from string with format success")
	} else {
		t.Error("Get date object from string with format error")
		t.Errorf("Expect get at %s but get %s \n", date.Local(), dateReceive.timeObject.Local())
	}
}

func TestFromUnixTime(t *testing.T) {
	dateReceive := FromUnixTime(1625112000)

	if dateReceive.timeObject.Equal(date) {
		t.Log("Get date object from unix time success")
	} else {
		t.Error("Get date object from unix time error")
		t.Errorf("Expect get at %s but get %s \n", date.Local(), dateReceive.timeObject.Local())
	}
}

func TestString(t *testing.T) {
	expectString := "2021-07-01T12:00:00.000+0800"
	dateTime := &DateTime{
		timeObject: date,
	}
	if dateTime.String() == expectString {
		t.Log("Get date time string success")
	} else {
		t.Error("Get date time string error")
		t.Errorf("Expect %s but get %s\n", expectString, dateTime.String())
	}
}

func TestStringWithFormat(t *testing.T) {
	expectString := "2021-07-01T12:00:00"
	dateTime := &DateTime{
		timeObject: date,
	}
	if dateTime.StringWithFormat("YYYY-MM-ddTHH:mm:ss") == expectString {
		t.Log("Get date time string with format success")
	} else {
		t.Error("Get date time string with format error")
		t.Errorf("Expect %s but get %s\n", expectString, dateTime.String())
	}
}

func TestEpochInSecond(t *testing.T) {
	var expectSecond int64 = 1625112000
	dateTime := &DateTime{
		timeObject: date,
	}
	if dateTime.EpochInSecond() == expectSecond {
		t.Log("Get epoch in second success")
	} else {
		t.Error("Get epoch in second error")
		t.Errorf("Expect %d but get %d", expectSecond, dateTime.EpochInSecond())
	}
}

func TestEpochInMilli(t *testing.T) {
	var expectMilli int64 = 1625112000000
	dateTime := &DateTime{
		timeObject: date,
	}
	if dateTime.EpochInMilli() == expectMilli {
		t.Log("Get epoch in milli success")
	} else {
		t.Error("Get epoch in milli error")
		t.Errorf("Expect %d but get %d", expectMilli, dateTime.EpochInMilli())
	}
}

func TestWeekDay(t *testing.T) {
	expectWeekDay := 4
	dateTime := &DateTime{
		timeObject: date,
	}
	if dateTime.WeekDay() == expectWeekDay {
		t.Log("Get week day success")
	} else {
		t.Error("Get week day error")
		t.Errorf("Expect %d but get %d", expectWeekDay, date.Weekday())
	}
}

func TestAddSecond(t *testing.T) {
	addSecond := 10
	dateTimeAdded := date.Add(time.Second * time.Duration(addSecond))
	dateTime := &DateTime{
		timeObject: date,
	}
	dateTime = dateTime.AddSecond(addSecond)

	if dateTime.timeObject.Equal(dateTimeAdded) {
		t.Log("Add second success")
	} else {
		t.Error("Add second error")
		t.Errorf("Expect %s but get %s", dateTimeAdded.Local(), dateTime.timeObject.Local())
	}
}
func TestAddMinute(t *testing.T) {
	addMinute := 10
	dateTimeAdded := date.Add(time.Minute * time.Duration(addMinute))
	dateTime := &DateTime{
		timeObject: date,
	}
	dateTime = dateTime.AddMinute(addMinute)

	if dateTime.timeObject.Equal(dateTimeAdded) {
		t.Log("Add minute success")
	} else {
		t.Error("Add minute error")
		t.Errorf("Expect %s but get %s", dateTimeAdded.Local(), dateTime.timeObject.Local())
	}
}

func TestAddHour(t *testing.T) {
	addHour := 1
	dateTimeAdded := date.Add(time.Hour * time.Duration(addHour))
	dateTime := &DateTime{
		timeObject: date,
	}
	dateTime = dateTime.AddHour(addHour)

	if dateTime.timeObject.Equal(dateTimeAdded) {
		t.Log("Add hour success")
	} else {
		t.Error("Add hour error")
		t.Errorf("Expect %s but get %s", dateTimeAdded.Local(), dateTime.timeObject.Local())
	}
}

func TestAddDay(t *testing.T) {
	addDay := 1
	dateTimeAdded := date.AddDate(0, 0, addDay)
	dateTime := &DateTime{
		timeObject: date,
	}
	dateTime = dateTime.AddDay(addDay)

	if dateTime.timeObject.Equal(dateTimeAdded) {
		t.Log("Add day success")
	} else {
		t.Error("Add day error")
		t.Errorf("Expect %s but get %s", dateTimeAdded.Local(), dateTime.timeObject.Local())
	}
}

func TestAddMonth(t *testing.T) {
	addMonth := 1
	dateTimeAdded := date.AddDate(0, addMonth, 0)
	dateTime := &DateTime{
		timeObject: date,
	}
	dateTime = dateTime.AddMonth(addMonth)

	if dateTime.timeObject.Equal(dateTimeAdded) {
		t.Log("Add month success")
	} else {
		t.Error("Add month error")
		t.Errorf("Expect %s but get %s", dateTimeAdded.Local(), dateTime.timeObject.Local())
	}
}

func TestAddYear(t *testing.T) {
	addYear := 1
	dateTimeAdded := date.AddDate(addYear, 0, 0)
	dateTime := &DateTime{
		timeObject: date,
	}
	dateTime = dateTime.AddYear(addYear)

	if dateTime.timeObject.Equal(dateTimeAdded) {
		t.Log("Add year success")
	} else {
		t.Error("Add year error")
		t.Errorf("Expect %s but get %s", dateTimeAdded.Local(), dateTime.timeObject.Local())
	}
}

func TestParseWithLayout(t *testing.T) {
	value := "2022-10-19T08:08:08Z"
	dateTime, err := FromStringWithLayout(value, time.RFC3339)
	if err != nil {
		t.Error(err)
	}

	err = dateTime.SetTZ("Asia/Taipei")
	if err != nil {
		t.Error(err)
	}

	t.Log(dateTime.String())
}
