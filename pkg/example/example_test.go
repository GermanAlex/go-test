package example

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

func TestCountLetters(t *testing.T) {
	type args struct {
		s string
		l byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				s: "Hello World",
				l: 'l',
			},
			want: 3,
		},
		{
			name: "Test 2",
			args: args{
				s: "",
				l: 'T',
			},
			want: 0,
		},
		{
			name: "Test 3",
			args: args{
				s: "Hello World",
				l: ' ',
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountLetters(tt.args.s, tt.args.l); got != tt.want {
				t.Errorf("CountLetters() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkSomeMath(b *testing.B) {
	for i := 0; i < b.N; i++ {
		n := rand.Float64()
		SomeMath(n)
	}
}

func sampleData() []int {
	rand.Seed(time.Now().UnixNano())
	var data []int
	for i := 0; i < 1_000; i++ {
		data = append(data, rand.Intn(1000))
	}

	sort.Slice(data, func(i, j int) bool { return data[i] < data[j] })
	return data
}

func BenchmarkSortInts(b *testing.B) {
	data := sampleData()
	for i := 0; i < b.N; i++ {
		SortInts(data)
	}
}

func TestNineSqrt(t *testing.T) {
	// верный результат
	num := 9.0
	want := 3.0
	got := NineSqrt(num)
	if got != want {
		t.Errorf("получено %f, ожидалось %f\n", got, want)
	}
	// ошибка
	num = 9.0
	want = 2.0
	got = NineSqrt(num)
	if got != want {
		t.Errorf("получено %f, ожидалось %f\n", got, want)
	}
}
