package combinations

import (
	"fmt"
	"testing"
)

func TestCombine(t *testing.T) {
	tests := []struct {
		name      string
		input     []int
		k         int
		wantCount int
		wantErr   bool
	}{
		{
			name:      "C(5,0)",
			input:     []int{1, 2, 3, 4, 5},
			k:         0,
			wantCount: 1,
			wantErr:   false,
		},
		{
			name:      "C(5,1)",
			input:     []int{1, 2, 3, 4, 5},
			k:         1,
			wantCount: 5,
			wantErr:   false,
		},
		{
			name:      "C(5,2)",
			input:     []int{1, 2, 3, 4, 5},
			k:         2,
			wantCount: 10,
			wantErr:   false,
		},
		{
			name:      "C(5,3)",
			input:     []int{1, 2, 3, 4, 5},
			k:         3,
			wantCount: 10,
			wantErr:   false,
		},
		{
			name:      "C(5,5)",
			input:     []int{1, 2, 3, 4, 5},
			k:         5,
			wantCount: 1,
			wantErr:   false,
		},
		{
			name:      "C(4,2)",
			input:     []int{1, 2, 3, 4},
			k:         2,
			wantCount: 6,
			wantErr:   false,
		},
		{
			name:      "empty slice k=0",
			input:     []int{},
			k:         0,
			wantCount: 1,
			wantErr:   false,
		},
		{
			name:      "nil input",
			input:     nil,
			k:         2,
			wantCount: 0,
			wantErr:   true,
		},
		{
			name:      "negative k",
			input:     []int{1, 2, 3},
			k:         -1,
			wantCount: 0,
			wantErr:   true,
		},
		{
			name:      "k > n",
			input:     []int{1, 2, 3},
			k:         5,
			wantCount: 0,
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Combine(tt.input, tt.k)
			if (err != nil) != tt.wantErr {
				t.Errorf("Combine() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && len(got) != tt.wantCount {
				t.Errorf("Combine() got %d combinations, want %d", len(got), tt.wantCount)
			}
		})
	}
}

func TestCombineStrings(t *testing.T) {
	input := []string{"a", "b", "c", "d"}
	result, err := Combine(input, 2)

	if err != nil {
		t.Fatalf("Combine() unexpected error: %v", err)
	}

	if len(result) != 6 {
		t.Errorf("Combine() got %d combinations, want 6", len(result))
	}
}

func TestCombineIter(t *testing.T) {
	tests := []struct {
		name      string
		input     []int
		k         int
		wantCount int
	}{
		{
			name:      "C(5,0)",
			input:     []int{1, 2, 3, 4, 5},
			k:         0,
			wantCount: 1,
		},
		{
			name:      "C(5,2)",
			input:     []int{1, 2, 3, 4, 5},
			k:         2,
			wantCount: 10,
		},
		{
			name:      "C(5,5)",
			input:     []int{1, 2, 3, 4, 5},
			k:         5,
			wantCount: 1,
		},
		{
			name:      "nil input",
			input:     nil,
			k:         2,
			wantCount: 0,
		},
		{
			name:      "k > n",
			input:     []int{1, 2, 3},
			k:         5,
			wantCount: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			count := 0
			for range CombineIter(tt.input, tt.k) {
				count++
			}
			if count != tt.wantCount {
				t.Errorf("CombineIter() got %d combinations, want %d", count, tt.wantCount)
			}
		})
	}
}

func TestCombineIterVsCombine(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	k := 3

	combs, err := Combine(input, k)
	if err != nil {
		t.Fatalf("Combine() error: %v", err)
	}

	count := 0
	for range CombineIter(input, k) {
		count++
	}

	if count != len(combs) {
		t.Errorf("CombineIter() count=%d, Combine() count=%d", count, len(combs))
	}
}

func BenchmarkCombine(b *testing.B) {
	benchmarks := []struct {
		name  string
		input []int
		k     int
	}{
		{"C(10,5)", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5},
		{"C(15,7)", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}, 7},
		{"C(20,10)", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, 10},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = Combine(bm.input, bm.k)
			}
		})
	}
}

func BenchmarkCombineIter(b *testing.B) {
	benchmarks := []struct {
		name  string
		input []int
		k     int
	}{
		{"C(10,5)", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5},
		{"C(15,7)", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}, 7},
		{"C(20,10)", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, 10},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for range CombineIter(bm.input, bm.k) {
				}
			}
		})
	}
}

func ExampleCombine() {
	result, _ := Combine([]int{1, 2, 3, 4}, 2)
	fmt.Printf("Total combinations: %d\n", len(result))
	// Output: Total combinations: 6
}

func ExampleCombineIter() {
	count := 0
	for range CombineIter([]int{1, 2, 3, 4}, 2) {
		count++
	}
	fmt.Printf("Total combinations: %d\n", count)
	// Output: Total combinations: 6
}
