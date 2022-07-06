package statistics

import (
	"testing"
)

func TestAvg(t *testing.T) {
	// верный результат
	nums := []float64{1, 2, 3}
	want := 2.0
	got := Avg(nums)
	if got != want {
		t.Errorf("получено %f, ожидалось %f\n", got, want)
	}
	// ошибка
	nums = []float64{1, 2, 3}
	want = 3.0
	got = Avg(nums)
	if got == want {
		t.Errorf("получено %f, ожидалось %f\n", got, want)
	}
}

func TestMaxN(t *testing.T) {
	// верный результат
	nums := []float64{4, 12, 120, 1}
	want := 120.0
	got := MaxN(nums)
	if got != want {
		t.Errorf("получено %f, ожидалось %f\n", got, want)
	}
	// ошибка
	nums = []float64{4, 12, 120, 1}
	want = 1.0
	got = MaxN(nums)
	if got == want {
		t.Errorf("получено %f, ожидалось %f\n", got, want)
	}
}
