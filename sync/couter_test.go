package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter thrice leaves it at 3", func(t *testing.T) {
		counter := Counter{}

		counter.Inc()
		counter.Inc()
		counter.Inc()

		if counter.Value() != 3 {
			t.Errorf("expected %d, but got %d", 3, counter.Value())
		}
	})

	t.Run("it runs safely in a concurrent environment", func(t *testing.T) {
		counter := Counter{}
		expectedCount := 1000

		var wg sync.WaitGroup
		for range expectedCount {
			wg.Add(1)
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}

		wg.Wait()

		if counter.Value() != expectedCount {
			t.Errorf("expected %d, but got %d", expectedCount, counter.Value())
		}
	})
}
