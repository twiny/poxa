package poxa

import "testing"

func TestNewSpinner(t *testing.T) {
	tests := []struct {
		elems    []int
		expected []int
	}{
		{[]int{1, 2, 3}, []int{1, 2, 3, 1, 2}},
		{[]int{4}, []int{4, 4, 4}},
		{[]int{5, 6}, []int{5, 6, 5, 6}},
	}

	for _, tt := range tests {
		sp := NewSpinner(tt.elems...)
		for _, expected := range tt.expected {
			result := sp.Next()
			if result != expected {
				t.Errorf("Expected %v, got %v", expected, result)
			}
		}
	}
}

func TestNewSpinnerWithNoElements(t *testing.T) {
	sp := NewSpinner[int]()
	if sp != nil {
		t.Errorf("Expected nil, got %v", sp)
	}
}

func BenchmarkNext(b *testing.B) {
	sp := NewSpinner[int](1, 2, 3, 4, 5)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sp.Next()
	}
}
