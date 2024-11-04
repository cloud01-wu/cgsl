package redisx

import (
	"context"
	"time"

	"github.com/gomodule/redigo/redis"
)

const (
	dbIndex      = 0
	pingInterval = 30 * time.Second
)

var (
	pool *redis.Pool
)

func Init(address string, password string, maxIdle int, maxActive int) {
	if nil != pool {
		return
	}

	pool = &redis.Pool{
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		Wait:        true,
		IdleTimeout: 5 * 60 * time.Second,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", address,
				redis.DialPassword(password),
				redis.DialConnectTimeout(3*time.Second),
				redis.DialReadTimeout(pingInterval+10*time.Second),
				redis.DialWriteTimeout(5*time.Second))

			return conn, err
		},
		// Use the TestOnBorrow function to check the health of an idle connection
		// before the connection is returned to the application.
		TestOnBorrow: func(conn redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}

			_, err := conn.Do("PING")
			return err
		},
	}
}

func New() (redis.Conn, error) {
	conn, err := pool.GetContext(context.Background())
	if nil != err {
		return nil, err
	}

	_, err = conn.Do("SELECT", dbIndex)
	if nil != err {
		conn.Close()
		return nil, err
	}

	return conn, nil
}

func NewWithContext(ctx context.Context) (redis.Conn, error) {
	conn, err := pool.GetContext(ctx)
	if nil != err {
		return nil, err
	}

	_, err = conn.Do("SELECT", dbIndex)
	if nil != err {
		conn.Close()
		return nil, err
	}

	return conn, nil
}

func Close() {
	if nil != pool {
		pool.Close()
		pool = nil
	}
}
