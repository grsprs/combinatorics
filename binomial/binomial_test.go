package binomial

import (
	"fmt"
	"math/big"
	"testing"
)

func TestBinomial(t *testing.T) {
	tests := []struct {
		name    string
		n       int
		k       int
		want    string
		wantErr bool
	}{
		{
			name:    "C(0,0)",
			n:       0,
			k:       0,
			want:    "1",
			wantErr: false,
		},
		{
			name:    "C(5,0)",
			n:       5,
			k:       0,
			want:    "1",
			wantErr: false,
		},
		{
			name:    "C(5,5)",
			n:       5,
			k:       5,
			want:    "1",
			wantErr: false,
		},
		{
			name:    "C(5,2)",
			n:       5,
			k:       2,
			want:    "10",
			wantErr: false,
		},
		{
			name:    "C(10,3)",
			n:       10,
			k:       3,
			want:    "120",
			wantErr: false,
		},
		{
			name:    "C(20,10)",
			n:       20,
			k:       10,
			want:    "184756",
			wantErr: false,
		},
		{
			name:    "negative n",
			n:       -1,
			k:       0,
			want:    "",
			wantErr: true,
		},
		{
			name:    "negative k",
			n:       5,
			k:       -1,
			want:    "",
			wantErr: true,
		},
		{
			name:    "k > n",
			n:       5,
			k:       10,
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Binomial(tt.n, tt.k)
			if (err != nil) != tt.wantErr {
				t.Errorf("Binomial(%d, %d) error = %v, wantErr %v", tt.n, tt.k, err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				want := new(big.Int)
				want.SetString(tt.want, 10)
				if got.Cmp(want) != 0 {
					t.Errorf("Binomial(%d, %d) = %v, want %v", tt.n, tt.k, got, want)
				}
			}
		})
	}
}

func TestBinomialSymmetry(t *testing.T) {
	// C(n, k) = C(n, n-k)
	tests := []struct {
		n int
		k int
	}{
		{5, 2},
		{10, 3},
		{20, 8},
		{15, 7},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("C(%d,%d)=C(%d,%d)", tt.n, tt.k, tt.n, tt.n-tt.k), func(t *testing.T) {
			left, err1 := Binomial(tt.n, tt.k)
			right, err2 := Binomial(tt.n, tt.n-tt.k)

			if err1 != nil || err2 != nil {
				t.Fatalf("unexpected errors: %v, %v", err1, err2)
			}

			if left.Cmp(right) != 0 {
				t.Errorf("Symmetry failed: C(%d,%d)=%v != C(%d,%d)=%v",
					tt.n, tt.k, left, tt.n, tt.n-tt.k, right)
			}
		})
	}
}

func BenchmarkBinomial(b *testing.B) {
	benchmarks := []struct {
		name string
		n    int
		k    int
	}{
		{"C(10,5)", 10, 5},
		{"C(20,10)", 20, 10},
		{"C(50,25)", 50, 25},
		{"C(100,50)", 100, 50},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = Binomial(bm.n, bm.k)
			}
		})
	}
}

func ExampleBinomial() {
	result, _ := Binomial(5, 2)
	fmt.Println(result.String())
	// Output: 10
}
