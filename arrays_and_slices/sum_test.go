package arraysandslices

import (
	"reflect"

	"testing"
)

func TestSum(t *testing.T) {
	input := []int{20, 1, 3, 2, 0, 12}
	expected := 38

	sum := Sum(input)

	if sum != expected {
		t.Errorf("expected %d, got %d", expected, sum)
	}
}

func TestSumAll(t *testing.T) {

	sums := SumAll([]int{2, 4, 2}, []int{1, 1, 0, 19}, []int{9, 0, 1})
	expected := []int{8, 21, 10}

	if reflect.DeepEqual(sums, expected) == false {
		t.Errorf("expected %v, got %v", expected, sums)
	}
}

func TestSumTail(t *testing.T) {
	t.Run("should correctly sum tails", func(t *testing.T) {
		tailSums := SumTail([]int{2, 3, 1}, []int{0, 2})
		expected := []int{4, 2}
		checkSums(t, tailSums, expected)
	})

	t.Run("should safely sum empty arrays", func(t *testing.T) {
		tailSums := SumTail([]int{2, 3, 1}, []int{})
		expected := []int{4, 0}
		checkSums(t, tailSums, expected)
	})

}

func checkSums(t testing.TB, sums, expected []int) {
	t.Helper()
	if reflect.DeepEqual(sums, expected) == false {
		t.Errorf("expected %v, got %v", expected, sums)
	}
}
