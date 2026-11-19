package helper

import (
	"log"
	"time"
)

func LogF5Request[T any](name string, fn func() (T, error)) (T, error) {
	start := time.Now()

	log.Printf("[F5] START %s", name)

	result, err := fn()

	duration := time.Since(start)

	if err != nil {
		log.Printf("[F5] END %s | duration=%s | status=failed | error=%v",
			name, duration, err)
		return result, err
	}

	log.Printf("[F5] END %s | duration=%s | status=success",
		name, duration)

	return result, nil
}
