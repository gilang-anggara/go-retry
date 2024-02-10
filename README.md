# Overview

Simple retry functions by passing config & functions.

# Usage Example
## Simple Function
```go
package main

import (
	"errors"

	"github.com/gilang-anggara/go-retry/retry"
)

var (
	ErrRetryable    = errors.New("retryable")
	ErrNotRetryable = errors.New("not retryable")
)

func main() {
	config := retry.RetryConfig{
		MaxRetry:              1,
		MinBackoffDelayMillis: 1000,
		MaxBackoffDelayMillis: 30000,
		RetryableErrors:       []error{ErrRetryable},
	}

	err := retry.WithRetry(config, sampleFunction1)

	if err != nil {
		panic(err)
	}
}

func sampleFunction1() error {
	// sample
}
```

## Complex Function
```go
package main

import (
	"errors"

	"github.com/gilang-anggara/go-retry/retry"
)

var (
	ErrRetryable    = errors.New("retryable")
	ErrNotRetryable = errors.New("not retryable")
)

func main() {
	config := retry.RetryConfig{
		MaxRetry:              1,
		MinBackoffDelayMillis: 1000,
		MaxBackoffDelayMillis: 30000,
		RetryableErrors:       []error{ErrRetryable},
	}

    // wrap function to simple error & get result
    var result int // careful not to share this variable
    var err error

	exec1 := func() error {
		result, err = sampleFunction2(10)

		return err
	}

    err = retry.WithRetry(config, exec1)
}

func sampleFunction2(n int) (int, error) {
	// sample
}
```
