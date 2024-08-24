# throttle

[![Build Status](https://github.com/anilsenay/throttle/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/features/actions)

`throttle` is a Go package that provides a simple and efficient way to throttle operations in your Go applications. It allows you to limit the rate at which operations are performed, which can be useful for rate limiting API calls, controlling resource usage, or managing concurrency.

## Features

- Easy-to-use throttling mechanism via Iterators
- Supports generic slices
- Configurable operations per interval
- Uses efficient channels and tickers for throttling

## Installation

To install the throttle package, use the following command:

```bash
go get github.com/anilsenay/throttle
```

## Usage

**NOTE: You've got to have Go version 1.23 for using the Limit function**

Here's a basic example of how to use the `Limit` function:

```go
package main

import (
    "fmt"
    "time"

    "github.com/anilsenay/throttle"
)

func main() {
    numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

    // Throttle to 2 operations per second
    for i, v := range throttle.Limit(numbers, 2, time.Second) {
        fmt.Printf("Processing: %d\n", v)
    }
}
```

```go
Output:

Processing: 1
Processing: 2
// wait 1 second
Processing: 3
Processing: 4
// wait 1 second
...
```

This example will process the numbers slice at a rate of 2 items per second.

## API Reference

### throttle package

#### func [Limit](https://github.com/anilsenay/throttle/blob/master/throttle.go#L10)

`func Limit[Slice ~[]E, E any](s Slice, ops int, interval time.Duration) iter.Seq2[int, E]`

`Limit` takes a slice, the number of operations per interval, and the interval duration. It returns an iterator that yields the index and value of each element in the slice, throttled according to the specified rate.

`Limit` provides a high-level, iterator-based API that internally uses `Throttler` for the actual rate limiting mechanism.

### throttler package

#### type [Throttler](https://github.com/anilsenay/throttle/blob/master/throttler/throttler.go#L8)

```go
type Throttler struct {
    bucket      int
    bucketLimit int
    interval    time.Duration
    ticker      *time.Ticker
    done        bool
    waitCh      chan struct{}
    once        sync.Once
}
```

The Throttler struct provides low-level control over throttling. It can be used directly for more fine-grained throttling needs.

#### func [New](https://github.com/anilsenay/throttle/blob/master/throttler/throttler.go#L18)

`func New(ops int, interval time.Duration) *Throttler`

Creates a new Throttler with the specified number of operations per interval.

#### func (\*Throttler) [Allow](https://github.com/anilsenay/throttle/blob/master/throttler/throttler.go#L51)

`func (t *Throttler) Allow() bool`

Allow checks if an operation is permitted based on the current throttling state. It returns `true` if the operation is allowed, and `false` otherwise.

#### func (\*Throttler) [Stop](https://github.com/anilsenay/throttle/blob/master/throttler/throttler.go#L30)

`func (t *Throttler) Stop()`

Stops the throttler and releases associated resources.

#### func (\*Throttler) [IsDone](https://github.com/anilsenay/throttle/blob/master/throttler/throttler.go#L26)

`func (t *Throttler) IsDone() bool`

IsDone returns `true` if the Throttler has been stopped, and `false` otherwise.

##### Example:

```go
th := New(3, time.Second)

i := 0
for th.Allow() {
    if i == 10 {
        th.Stop()
    }
    i++
}
```

## Internal Workings

The Throttler uses a token bucket algorithm for rate limiting:

- The bucket is filled with tokens at the start of each interval.
- Each Allow() call consumes one token.
- If the bucket is empty, Allow() blocks until the next interval.
- A background goroutine refills the bucket at the start of each interval.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
