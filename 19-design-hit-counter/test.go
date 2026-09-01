package hitcounter

import (
	"sync"
	"testing"
)

func RunHitCounterTests(t *testing.T) {
	t.Run("Standard LeetCode sequence", func(t *testing.T) {
		counter := Constructor()

		counter.Hit(1)
		counter.Hit(2)
		counter.Hit(3)

		if hits := counter.GetHits(4); hits != 3 {
			t.Errorf("GetHits(4) = %d, want 3", hits)
		}

		counter.Hit(300)

		if hits := counter.GetHits(300); hits != 4 {
			t.Errorf("GetHits(300) = %d, want 4", hits)
		}

		// At t=301, hit at t=1 expires (301 - 1 = 300 not < 300)
		if hits := counter.GetHits(301); hits != 3 {
			t.Errorf("GetHits(301) = %d, want 3", hits)
		}
	})

	t.Run("Multiple hits at identical timestamp", func(t *testing.T) {
		counter := Constructor()
		for range 50 {
			counter.Hit(10)
		}
		if hits := counter.GetHits(10); hits != 50 {
			t.Errorf("GetHits(10) = %d, want 50", hits)
		}
		if hits := counter.GetHits(309); hits != 50 {
			t.Errorf("GetHits(309) = %d, want 50", hits)
		}
		if hits := counter.GetHits(310); hits != 0 {
			t.Errorf("GetHits(310) = %d, want 0 (expired)", hits)
		}
	})

	t.Run("Concurrent access safe", func(t *testing.T) {
		counter := Constructor()
		var wg sync.WaitGroup

		for range 100 {
			wg.Add(1)
			go func() {
				defer wg.Done()
				counter.Hit(100)
			}()
		}

		wg.Wait()

		if hits := counter.GetHits(100); hits != 100 {
			t.Errorf("GetHits(100) = %d, want 100", hits)
		}
	})
}
