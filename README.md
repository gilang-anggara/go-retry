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
	retryer := retry.New(retry.RetryConfig{
		MaxRetry:              5,
		MinBackoffDelayMillis: 1000,
		MaxBackoffDelayMillis: 5000,
		RetryableErrors:       []error{ErrRetryable},
	})

	err := retryer.WithRetry(sampleFunction1)

	if err != nil {
		panic(err)
	}
}

func sampleFunction1() error {
	return ErrRetryable
}
```

## Complex Function
```go
package main

import (
	"errors"
	"fmt"

	"github.com/gilang-anggara/go-retry/retry"
)

var (
	ErrRetryable    = errors.New("retryable")
	ErrNotRetryable = errors.New("not retryable")
)

func main() {
	retryer := retry.New(retry.RetryConfig{
		MaxRetry:              5,
		MinBackoffDelayMillis: 1000,
		MaxBackoffDelayMillis: 5000,
		RetryableErrors:       []error{ErrRetryable},
	})

	// wrap function to simple error & get result
	var result int // careful not to share this variable
	var err error

	exec1 := func() error {
		result, err = sampleFunction2(10)

		return err
	}

	err = retryer.WithRetry(exec1)

	fmt.Println(result, err)
}

func sampleFunction2(n int) (int, error) {
	return 1, ErrRetryable
}
```
