package enumable_test

import (
	"testing"

	enumable "github.com/hyhecor/enumable_builder/test"
)

func TestAppend(t *testing.T) {

	var (
		item_0 int = 0
		item_1 int = 1
		item_2 int = 2
		item_3 int = 3
		item_4 int = 4
		item_5 int = 5
		item_6 int = 6
		item_7 int = 7
		item_8 int = 8
		item_9 int = 9
	)

	arr := enumable.CreateLazyInt(
		item_0,
		item_1,
		item_2,
		item_3,
		item_4,
		item_5,
		item_6,
		item_7,
		item_8,
		item_9,
	)

	arr = arr.Append(enumable.CreateLazyInt(item_0))
	arr = arr.Append(enumable.CreateLazyInt(item_1))
	arr = arr.Append(enumable.CreateLazyInt(item_2))
	arr = arr.Append(enumable.CreateLazyInt(item_3))
	arr = arr.Append(enumable.CreateLazyInt(item_4))
	arr = arr.Append(enumable.CreateLazyInt(item_5))
	arr = arr.Append(enumable.CreateLazyInt(item_6))
	arr = arr.Append(enumable.CreateLazyInt(item_7))
	arr = arr.Append(enumable.CreateLazyInt(item_8))
	arr = arr.Append(enumable.CreateLazyInt(item_9))

	count := arr.Count()
	t.Logf("Count=%d\n", count)

	fold := arr.Fold(0, func(a, b int) int {
		return a + b
	})

	t.Logf("%d\n", fold)
}
