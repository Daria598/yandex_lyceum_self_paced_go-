package main

import (
	"context"
	"sync"
)

func ParallelMapCtx(ctx context.Context, inputs []int, fn func(int) int, workers int) ([]int, error) {
	if workers <= 0 {
		workers = 1
	}

	n := len(inputs)
	output := make([]int, n)

	type job struct {
		idx int
		val int
	}

	type result struct {
		idx int
		val int
	}

	jobs := make(chan job)
	results := make(chan result)

	var wg sync.WaitGroup

	for w := 0; w < workers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := range jobs {
				select {
				case <-ctx.Done():
					return
				default:
				}

				r := fn(j.val)

				select {
				case <-ctx.Done():
					return
				case results <- result{idx: j.idx, val: r}:
				}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	go func() {
		defer close(jobs)
		for i, v := range inputs {
			select {
			case <-ctx.Done():
				return
			case jobs <- job{idx: i, val: v}:
			}
		}
	}()

	received := 0
	for received < n {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case res, ok := <-results:
			if !ok {
				if ctx.Err() != nil {
					return nil, ctx.Err()
				}
				return nil, context.Canceled
			}
			output[res.idx] = res.val
			received++
		}
	}

	return output, nil
}
