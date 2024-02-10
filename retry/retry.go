package retry

import (
	"errors"
	"time"
)

type RetryConfig struct {
	MaxRetry              int
	MinBackoffDelayMillis int
	MaxBackoffDelayMillis int
	RetryableErrors       []error
}

// WithRetry will execute with following conditions:
// target function executes RetryConfig.MaxRetry + 1 times,
// backoff is linear, calculated from RetryConfig.MinBackoffDelayMillis until MaxBackoffDelayMillis,
// MaxBackoffDelayMillis will be overwritten as max(MinBackoffDelayMillis, MaxBackoffDelayMillis)
// empty retryables means target function executes only 1 times.
func WithRetry(config RetryConfig, f func() error) error {
	var err error

	backoff, backoffIncrement := calculateBackoff(config.MinBackoffDelayMillis, config.MaxBackoffDelayMillis, config.MaxRetry)

	i := 0
	for {
		err = f()

		if err == nil || !isRetryable(err, config.RetryableErrors) {
			break
		}

		if i >= config.MaxRetry {
			break
		}

		<-time.After(time.Duration(backoff) * time.Millisecond)
		backoff += backoffIncrement
		i += 1
	}

	return err
}

func calculateBackoff(minBackoffDelayMillis int, maxBackoffDelayMillis int, maxRetry int) (initialBackoff int, backoffIncrement int) {
	initialBackoff = minBackoffDelayMillis
	backoffIncrement = 0
	if maxRetry > 0 {
		backoffIncrement = (max(maxBackoffDelayMillis, minBackoffDelayMillis) - minBackoffDelayMillis) / maxRetry
	}

	return initialBackoff, backoffIncrement
}

func isRetryable(err error, retryableErrors []error) bool {
	for i := range retryableErrors {
		if errors.Is(err, retryableErrors[i]) {
			return true
		}
	}

	return false
}
