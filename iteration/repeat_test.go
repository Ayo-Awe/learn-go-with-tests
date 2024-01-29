package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("should repeat char 5 times when count is less than 0", func(t *testing.T) {
		repeated := Repeat("a", -1)
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("expected %q, but got %q", expected, repeated)
		}
	})

	t.Run("should repeat character <count> number of times", func(t *testing.T) {
		repeated := Repeat("a", 2)
		expected := "aa"

		if repeated != expected {
			t.Errorf("expected %q, but got %q", expected, repeated)
		}
	})

}

func BenchmarkRepeat(t *testing.B) {
	for i := 0; i < t.N; i++ {
		Repeat("a", 4)
	}
}
