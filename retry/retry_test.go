package retry_test

import (
	"errors"
	"testing"
	"time"

	"github.com/gilang-anggara/go-retry/retry"
	"github.com/stretchr/testify/assert"
)

func Test_WithRetry_Success(t *testing.T) {
	errRetryable := errors.New("retryable")
	errNotRetryable := errors.New("not retryable")

	retryer := retry.New(retry.RetryConfig{
		MaxRetry:              0,
		MinBackoffDelayMillis: 1000,
		RetryableErrors:       []error{errRetryable},
	})

	callCount := 0
	toBeCalled := func() error {
		callCount += 1

		return errNotRetryable
	}

	start := time.Now()
	err := retryer.WithRetry(toBeCalled)
	duration := time.Since(start)

	assert.ErrorIs(t, err, errNotRetryable)
	assert.Equal(t, callCount, 1)
	assert.True(t, duration < time.Duration(1000)*time.Millisecond)
}

func Test_WithRetry_MaxRetries(t *testing.T) {
	errRetryable := errors.New("retryable")

	retryer := retry.New(retry.RetryConfig{
		MaxRetry:              10,
		MinBackoffDelayMillis: 100,
		MaxBackoffDelayMillis: 1000,
		RetryableErrors:       []error{errRetryable},
	})

	callCount := 0
	toBeCalled := func() error {
		callCount += 1

		return errRetryable
	}

	start := time.Now()
	err := retryer.WithRetry(toBeCalled)
	duration := time.Since(start)

	assert.ErrorIs(t, err, errRetryable)
	assert.Equal(t, callCount, 11)
	assert.True(t, duration > time.Duration(5000)*time.Millisecond)
}
