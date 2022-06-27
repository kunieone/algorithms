package algorithms

import "testing"

func TestGcd(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			",",
			args{a: 25671, b: 49219},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Gcd(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Gcd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkGcd(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Gcd(20000, 300)
	}
	b.StopTimer()

}
