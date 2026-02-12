package factorial

import (
	"fmt"
	"math/big"
	"testing"
)

func TestFactorial(t *testing.T) {
	tests := []struct {
		name    string
		input   int
		want    string
		wantErr bool
	}{
		{
			name:    "factorial of 0",
			input:   0,
			want:    "1",
			wantErr: false,
		},
		{
			name:    "factorial of 1",
			input:   1,
			want:    "1",
			wantErr: false,
		},
		{
			name:    "factorial of 5",
			input:   5,
			want:    "120",
			wantErr: false,
		},
		{
			name:    "factorial of 10",
			input:   10,
			want:    "3628800",
			wantErr: false,
		},
		{
			name:    "factorial of 20",
			input:   20,
			want:    "2432902008176640000",
			wantErr: false,
		},
		{
			name:    "negative input",
			input:   -1,
			want:    "",
			wantErr: true,
		},
		{
			name:    "negative input -5",
			input:   -5,
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Factorial(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Factorial(%d) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				want := new(big.Int)
				want.SetString(tt.want, 10)
				if got.Cmp(want) != 0 {
					t.Errorf("Factorial(%d) = %v, want %v", tt.input, got, want)
				}
			}
		})
	}
}

func BenchmarkFactorial(b *testing.B) {
	benchmarks := []struct {
		name  string
		input int
	}{
		{"n=10", 10},
		{"n=20", 20},
		{"n=50", 50},
		{"n=100", 100},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = Factorial(bm.input)
			}
		})
	}
}

func ExampleFactorial() {
	result, _ := Factorial(5)
	fmt.Println(result.String())
	// Output: 120
}
