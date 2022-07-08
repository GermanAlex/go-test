package main

/*import (
	"fmt"
	"statistics/pkg/statistics"
)

func main() {
	// верный результат
	nums := []float64{1, 2, 3}
	want := 2.0
	got := statistics.Avg(nums)
	if got != want {
		fmt.Printf("получено %f, ожидалось %f\n", got, want)
	}
	fmt.Printf("результат: %f\n", got)
	// ошибка
	nums = []float64{1, 2, 3}
	want = 3.0
	got = statistics.Avg(nums)
	if got != want {
		fmt.Printf("получено %f, ожидалось %f\n", got, want)
	}
	fmt.Printf("результат: %f\n", got)
}
*/

import (
	"fmt"
	"statistics/pkg/storage"
	"statistics/pkg/storage/memdb"
)

var db storage.Interface

func main() {
	db = memdb.New()
	t := task(db, 0)
	fmt.Println(t)
}

func task(db storage.Interface, id int) storage.Task {
	tt, err := db.Tasks()
	if err != nil {
		return storage.Task{}
	}
	for _, t := range tt {
		if t.ID == id {
			return t
		}
	}
	return storage.Task{}
}
