package randomizedset

import "testing"

func RunRandomizedSetTests(t *testing.T) {
	t.Run("Standard Insert Remove and Random Flow", func(t *testing.T) {
		rSet := Constructor()

		if !rSet.Insert(1) {
			t.Errorf("expected Insert(1) to be true")
		}
		if rSet.Insert(1) {
			t.Errorf("expected duplicate Insert(1) to be false")
		}
		if !rSet.Insert(2) {
			t.Errorf("expected Insert(2) to be true")
		}

		randVal := rSet.GetRandom()
		if randVal != 1 && randVal != 2 {
			t.Errorf("expected GetRandom() to return 1 or 2, got %d", randVal)
		}

		if !rSet.Remove(1) {
			t.Errorf("expected Remove(1) to be true")
		}
		if rSet.Remove(1) {
			t.Errorf("expected removing already removed 1 to be false")
		}

		if val := rSet.GetRandom(); val != 2 {
			t.Errorf("expected only element left to be 2, got %d", val)
		}
	})

	t.Run("Multiple elements swap-and-pop correctness", func(t *testing.T) {
		rSet := Constructor()
		values := []int{10, 20, 30, 40, 50}
		for _, v := range values {
			rSet.Insert(v)
		}

		// Remove middle element (30)
		if !rSet.Remove(30) {
			t.Errorf("expected Remove(30) to be true")
		}
		if rSet.Remove(30) {
			t.Errorf("expected second Remove(30) to be false")
		}

		// Verify remaining elements can all be picked
		observed := make(map[int]bool)
		for range 200 {
			observed[rSet.GetRandom()] = true
		}

		for _, v := range []int{10, 20, 40, 50} {
			if !observed[v] {
				t.Errorf("expected value %d to be chosen by GetRandom", v)
			}
		}
		if observed[30] {
			t.Errorf("removed value 30 should never be picked")
		}
	})
}
