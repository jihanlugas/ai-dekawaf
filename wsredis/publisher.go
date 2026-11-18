package wsredis

import (
	"log"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
)

// Publisher publishes messages to the Redis channel
type Publisher struct {
	pool     *redis.Pool
	channel  string
	messages chan []byte
}

// NewPublisher creates a publisher that will use to publish
// message to the provided redis.Pool.
func NewPublisher(pool *redis.Pool, channel string) *Publisher {
	return &Publisher{
		pool:     pool,
		channel:  channel,
		messages: make(chan []byte, 10000),
	}
}

// run the main redisWriter loop that publishes incoming messages to Redis.
func (rw *Publisher) Run() error {
	conn := rw.pool.Get()
	defer conn.Close()

	for data := range rw.messages {
		if err := Publish(conn, rw.channel, data); err != nil {
			rw.Publish(data) // attempt to redeliver later
			return err
		}
	}
	return nil
}

// Publish to Redis via channel. Used by server
func Publish(conn redis.Conn, channel string, data []byte) error {
	if err := conn.Send("PUBLISH", channel, data); err != nil {
		return errors.Wrap(err, "Unable to publish message to Redis")
	}
	if err := conn.Flush(); err != nil {
		return errors.Wrap(err, "Unable to flush published message to Redis")
	}
	return nil
}

// PublishWithRetry publishes a message to Redis with retry logic.
// It obtains a fresh connection from the pool on each attempt and
// properly returns it after use. Retries up to 3 times with
// exponential backoff (500ms, 1s, 2s).
func PublishWithRetry(pool *redis.Pool, channel string, data []byte) error {
	maxRetries := 3
	backoff := 500 * time.Millisecond

	var lastErr error
	for attempt := 0; attempt <= maxRetries; attempt++ {
		if attempt > 0 {
			log.Printf("retrying websocket redis publish (attempt %d/%d) after error: %v", attempt, maxRetries, lastErr)
			time.Sleep(backoff)
			backoff *= 2
		}

		conn := pool.Get()
		err := Publish(conn, channel, data)
		conn.Close() // always return connection to pool
		if err == nil {
			return nil
		}
		lastErr = err
	}
	return lastErr
}

// Publish to Redis via channel. Used by client
func (rw *Publisher) Publish(data []byte) {
	rw.messages <- data
}
