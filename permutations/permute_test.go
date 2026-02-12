package permutations

import (
	"fmt"
	"testing"
)

func TestPermute(t *testing.T) {
	tests := []struct {
		name      string
		input     []int
		wantCount int
		wantErr   bool
	}{
		{
			name:      "empty slice",
			input:     []int{},
			wantCount: 1,
			wantErr:   false,
		},
		{
			name:      "single element",
			input:     []int{1},
			wantCount: 1,
			wantErr:   false,
		},
		{
			name:      "two elements",
			input:     []int{1, 2},
			wantCount: 2,
			wantErr:   false,
		},
		{
			name:      "three elements",
			input:     []int{1, 2, 3},
			wantCount: 6,
			wantErr:   false,
		},
		{
			name:      "four elements",
			input:     []int{1, 2, 3, 4},
			wantCount: 24,
			wantErr:   false,
		},
		{
			name:      "nil input",
			input:     nil,
			wantCount: 0,
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Permute(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Permute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && len(got) != tt.wantCount {
				t.Errorf("Permute() got %d permutations, want %d", len(got), tt.wantCount)
			}
		})
	}
}

func TestPermuteStrings(t *testing.T) {
	input := []string{"a", "b", "c"}
	result, err := Permute(input)

	if err != nil {
		t.Fatalf("Permute() unexpected error: %v", err)
	}

	if len(result) != 6 {
		t.Errorf("Permute() got %d permutations, want 6", len(result))
	}
}

func TestPermuteIter(t *testing.T) {
	tests := []struct {
		name      string
		input     []int
		wantCount int
	}{
		{
			name:      "empty slice",
			input:     []int{},
			wantCount: 1,
		},
		{
			name:      "single element",
			input:     []int{1},
			wantCount: 1,
		},
		{
			name:      "two elements",
			input:     []int{1, 2},
			wantCount: 2,
		},
		{
			name:      "three elements",
			input:     []int{1, 2, 3},
			wantCount: 6,
		},
		{
			name:      "nil input",
			input:     nil,
			wantCount: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			count := 0
			for range PermuteIter(tt.input) {
				count++
			}
			if count != tt.wantCount {
				t.Errorf("PermuteIter() got %d permutations, want %d", count, tt.wantCount)
			}
		})
	}
}

func TestPermuteIterVsPermute(t *testing.T) {
	input := []int{1, 2, 3, 4}

	perms, err := Permute(input)
	if err != nil {
		t.Fatalf("Permute() error: %v", err)
	}

	count := 0
	for range PermuteIter(input) {
		count++
	}

	if count != len(perms) {
		t.Errorf("PermuteIter() count=%d, Permute() count=%d", count, len(perms))
	}
}

func BenchmarkPermute(b *testing.B) {
	benchmarks := []struct {
		name  string
		input []int
	}{
		{"n=5", []int{1, 2, 3, 4, 5}},
		{"n=7", []int{1, 2, 3, 4, 5, 6, 7}},
		{"n=8", []int{1, 2, 3, 4, 5, 6, 7, 8}},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = Permute(bm.input)
			}
		})
	}
}

func BenchmarkPermuteIter(b *testing.B) {
	benchmarks := []struct {
		name  string
		input []int
	}{
		{"n=5", []int{1, 2, 3, 4, 5}},
		{"n=7", []int{1, 2, 3, 4, 5, 6, 7}},
		{"n=8", []int{1, 2, 3, 4, 5, 6, 7, 8}},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for range PermuteIter(bm.input) {
				}
			}
		})
	}
}

func ExamplePermute() {
	result, _ := Permute([]int{1, 2, 3})
	fmt.Printf("Total permutations: %d\n", len(result))
	// Output: Total permutations: 6
}

func ExamplePermuteIter() {
	count := 0
	for range PermuteIter([]int{1, 2, 3}) {
		count++
	}
	fmt.Printf("Total permutations: %d\n", count)
	// Output: Total permutations: 6
}
