package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the coutner 3 times leaves it at 3", func(t *testing.T) {
		counter := NewCounter()
		for i := 1; i <= 3; i++ {
			counter.Inc()
		}

		assertCounter(t, counter, 3)
	})

	t.Run("it runs safe concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := NewCounter()

		var wg sync.WaitGroup	// A convenient way of synchronosing concrurent processes
		wg.Add(wantedCount)	// A counter of how many goroutines to wait for

		for i := 0; i < wantedCount; i++ {
			go func() {
				counter.Inc()
				wg.Done() // Decreases the wait counter	
			}()
		}
		wg.Wait()	// sync.WaitGroup.Wait() is used to block until all goroutines have finished

		// By using wg.Wait(), by the time we assert the counter all of our goroutines would have attempted to Inc the Counter
		assertCounter(t, counter, wantedCount)
	})
}

func assertCounter(t testing.TB, got *Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("got %d, want %d", got.Value(), want)
	}
}