package redisx

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/gomodule/redigo/redis"
)

func Ping() error {
	conn, err := New()
	if nil != err {
		return err
	}
	defer conn.Close()

	_, err = redis.String(conn.Do("PING"))
	if err != nil {
		return fmt.Errorf("cannot 'PING' db: %v", err)
	}
	return nil
}

func Get(key string) ([]byte, error) {
	var data []byte

	conn, err := New()
	if nil != err {
		return data, err
	}
	defer conn.Close()

	data, err = redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return data, fmt.Errorf("error getting key %s: %v", key, err)
	}
	return data, err
}

func Set(key string, value []byte) error {
	conn, err := New()
	if nil != err {
		return err
	}
	defer conn.Close()

	_, err = conn.Do("SET", key, value)
	if err != nil {
		v := string(value)
		if len(v) > 15 {
			v = v[0:12] + "..."
		}
		return fmt.Errorf("error setting key %s to %s: %v", key, v, err)
	}
	return err
}

func HGet(args ...interface{}) ([]byte, error) {
	conn, err := New()
	if nil != err {
		return nil, err
	}
	defer conn.Close()

	data, err := redis.Bytes(conn.Do("HGET", args...))
	if err != nil {
		return nil, fmt.Errorf("error getting hash: %v", err)
	}
	return data, err
}

func HSet(args ...interface{}) error {
	conn, err := New()
	if nil != err {
		return err
	}
	defer conn.Close()

	_, err = conn.Do("HSET", args...)
	if err != nil {
		return fmt.Errorf("error setting hash: %v", err)
	}
	return err
}

func Scan(cursor int, keyPattern string, size int) (int, []string, error) {
	keys := []string{}
	conn, err := New()
	if nil != err {
		return 0, nil, err
	}
	defer conn.Close()

	arr, err := redis.Values(conn.Do("SCAN", cursor, "MATCH", keyPattern, "COUNT", size))

	if err != nil {
		return 0, nil, fmt.Errorf("error scanning: %v", err)
	}

	iter, _ := redis.Int(arr[0], nil)
	k, _ := redis.Strings(arr[1], nil)
	keys = append(keys, k...)

	return iter, keys, err
}

func Expire(key string, seconds int) error {
	conn, err := New()
	if nil != err {
		return err
	}
	defer conn.Close()

	_, err = conn.Do("EXPIRE", key, seconds)
	return err
}

func Exists(key string) (bool, error) {
	var ok bool

	conn, err := New()
	if nil != err {
		return ok, err
	}
	defer conn.Close()

	ok, err = redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return ok, fmt.Errorf("error checking if key %s exists: %v", key, err)
	}
	return ok, err
}

func Delete(key string) error {
	conn, err := New()
	if nil != err {
		return err
	}
	defer conn.Close()

	_, err = conn.Do("DEL", key)
	return err
}

func GetKeys(pattern string) ([]string, error) {
	keys := []string{}

	conn, err := New()
	if nil != err {
		return keys, err
	}
	defer conn.Close()

	iter := 0
	for {
		arr, err := redis.Values(conn.Do("SCAN", iter, "MATCH", pattern))
		if err != nil {
			return keys, fmt.Errorf("error retrieving '%s' keys", pattern)
		}

		iter, _ = redis.Int(arr[0], nil)
		k, _ := redis.Strings(arr[1], nil)
		keys = append(keys, k...)

		if iter == 0 {
			break
		}
	}

	return keys, nil
}

func Incr(counterKey string) (int, error) {
	conn, err := New()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	return redis.Int(conn.Do("INCR", counterKey))
}

func Publish(channel string, message []byte) error {
	conn, err := New()
	if err != nil {
		return err
	}
	defer conn.Close()

	if _, err = conn.Do("PUBLISH", channel, message); err != nil {
		return err
	}

	return nil
}

func Subscribe(
	ctx context.Context,
	handler func(channel string, data []byte) error,
	channels ...string) error {
	conn, err := NewWithContext(ctx)
	if err != nil {
		return err
	}

	psConn := redis.PubSubConn{
		Conn: conn,
	}

	// subscribe specific channels
	if err := psConn.Subscribe(redis.Args{}.AddFlat(channels)...); err != nil {
		return err
	}

	done := make(chan error, 1)
	defer close(done)

	waitGroup := sync.WaitGroup{}
	waitGroup.Add(1)

	go func() {
		defer psConn.Close()
		defer waitGroup.Done()

		for {
			switch message := psConn.Receive().(type) {
			case error:
				// server sent error
				done <- message
				return
			case redis.Message:
				if err := handler(message.Channel, message.Data); err != nil {
					done <- err
					return
				}
			case redis.Subscription:
				if message.Count == 0 {
					// return from the goroutine when all channels are unsubscribed.
					done <- nil
					return
				}
			}
		}
	}()
	defer waitGroup.Wait()

	// health check with ping
	ticker := time.NewTicker(pingInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			// maybe canceled by caller
			if err := psConn.Unsubscribe(); err != nil {
				return err
			}
			return nil
		case <-ticker.C:
			// by ticker for health check
			if err := psConn.Ping(""); err != nil {
				return err
			}
		case err := <-done:
			// by receiving error or unsubscribing all of the channels
			return err
		}
	}
}
