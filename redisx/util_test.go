package redisx

import (
	"context"
	"sync"
	"testing"
	"time"
)

var once sync.Once

func InitPool() {
	once.Do(func() {
		// init
		Init("192.168.99.233:6379", "", 16, 16)
	})
}

func TestPing(t *testing.T) {
	InitPool()
	defer Close()
	err := Ping()
	if err != nil {
		t.Error("Ping database Error")
	} else {
		t.Log("Ping database Success")
	}
}

func TestSet(t *testing.T) {
	InitPool()
	defer Close()
	err := Set("test", []byte("test"))
	if err != nil {
		t.Error("Set Database Error")
		t.Errorf("error: %s", err.Error())
	} else {
		t.Log("Set Database Success")
	}
}

func TestGet(t *testing.T) {
	InitPool()
	defer Close()
	_, err := Get("test")
	if err != nil {
		t.Error("Get Database Error")
		t.Errorf("error: %s", err.Error())
	} else {
		t.Log("Get Database Success")
	}
}

func TestDelete(t *testing.T) {
	InitPool()
	defer Close()
	err := Delete("book")
	if err != nil {
		t.Error("Delete Database Error")
		t.Errorf("error: %s", err.Error())
	} else {
		t.Log("Delete Database Success")
	}
}

func TestHSet(t *testing.T) {
	InitPool()
	defer Close()
	err := HSet(3, "book1", "name", "test", "price", 100)
	if err != nil {
		t.Error("HSET Error")
		t.Errorf("error: %s", err.Error())
	} else {
		t.Log("HSET Success")
	}
}

func TestHGet(t *testing.T) {
	InitPool()
	defer Close()
	_, err := HGet("book", "name")
	if err != nil {
		t.Error("HGet Error")
		t.Errorf("error: %s", err.Error())
	} else {
		t.Log("HGet Success")
	}
}

func TestScan(t *testing.T) {
	InitPool()
	defer Close()

	_, _, err := Scan(0, "*", 1)
	if err != nil {
		t.Error("Scan error")
		t.Errorf("error: %s", err.Error())
	}
}

func TestExpire(t *testing.T) {
	InitPool()
	defer Close()
	// set a key
	err := Set("expire", []byte("expire"))
	if err != nil {
		t.Error("Set Database Error")
		t.Errorf("error: %s", err.Error())
	}
	// set expire
	err = Expire("expire", 3)
	if err != nil {
		t.Error("Set Expire error")
	}
	time.Sleep(time.Second * 3)
	data, _ := Get("expire")
	if data != nil {
		t.Error("Set Expire error")
	} else {
		t.Log("Set expire success")
	}
}

func TestExists(t *testing.T) {
	InitPool()
	defer Close()
	//set a key
	err := Set("exists", []byte("exists"))
	if err != nil {
		t.Error("Set Database Error")
		t.Errorf("error: %s", err.Error())
	}
	result, err := Exists("exists")
	if err != nil {
		t.Error("Get if exists Error")
		t.Errorf("error: %s", err.Error())
	}
	if result == true {
		t.Log("Get if exists Error")
	} else {
		t.Logf("Check if exists error expect %t but get %t", true, result)
	}
}

func TestGetKeys(t *testing.T) {
	InitPool()
	defer Close()

	_, err := GetKeys("*")

	if err != nil {
		t.Error("Get keys error")
	} else {
		t.Log("Get keys success")
	}
}

func TestIncr(t *testing.T) {
	InitPool()
	defer Close()

	//set a key of number
	err := Set("incr", []byte("1"))
	if err != nil {
		t.Error("Set Database Error")
		t.Errorf("error: %s", err.Error())
	}
	result, err := Incr("incr")
	if err != nil {
		t.Error("Incr Error")
		t.Errorf("error: %s", err.Error())
	}
	if result == 2 {
		t.Log("Incr success")
	} else {
		t.Errorf("Incr error expect %d but get %d", 2, result)
	}
}

func TestPublish(t *testing.T) {
	InitPool()
	defer Close()

	err := Publish("test", []byte("hello"))
	if err != nil {
		t.Errorf("%v", err)
	}
}

func TestSubscribe(t *testing.T) {
	InitPool()
	defer Close()

	err := Subscribe(context.Background(), func(channel string, data []byte) error {
		return nil
	}, "test")

	if err != nil {
		t.Errorf("%v", err)
	}
}
