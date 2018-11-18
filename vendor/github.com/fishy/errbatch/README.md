[![GoDoc](https://godoc.org/github.com/fishy/errbatch?status.svg)](https://godoc.org/github.com/fishy/errbatch)
[![Go Report Card](https://goreportcard.com/badge/github.com/fishy/errbatch)](https://goreportcard.com/report/github.com/fishy/errbatch)

# ErrBatch

ErrBatch is a [Go](https://golang.org)
[library](https://godoc.org/github.com/fishy/errbatch) that can be used to
compile multiple errors into a single error.

Example code:

```go
type worker func() error

func runWorksParallel(works []worker) error {
	errChan := make(chan error, len(works))
	var wg sync.WaitGroup
	wg.Add(len(works))

	for _, work := range works {
		go func(work worker) {
			errChan <- work()
		}(work)
	}

	wg.Wait()
	var batch errbatch.ErrBatch
	for err := range errChan {
		// nil errors will be auto skipped
		batch.Add(err)
	}
	// If all works succeeded, Compile() returns nil.
	// If only one work failed, Compile() returns that error directly
	// instead of wrapping it inside ErrBatch.
	return batch.Compile()
}
```
