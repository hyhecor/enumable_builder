package enumable_test

import (
	"fmt"
	"testing"

	enumable "github.com/hyhecor/enumable_builder/test"
)

func TestAppend(t *testing.T) {

	var (
		item_0 int
		item_1 int
		item_2 int
		item_3 int
		item_4 int
		item_5 int
		item_6 int
		item_7 int
		item_8 int
		item_9 int
	)

	arr := enumable.SliceInt{
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
	}

	arr = arr.Append(item_0)
	arr = arr.Append(item_1)
	arr = arr.Append(item_2)
	arr = arr.Append(item_3)
	arr = arr.Append(item_4)
	arr = arr.Append(item_5)
	arr = arr.Append(item_6)
	arr = arr.Append(item_7)
	arr = arr.Append(item_8)
	arr = arr.Append(item_9)

	if 20 != len(arr) {
		t.Errorf("Expect len=%d Actual=%d", 20, len(arr))
	}

	arrf := arr.Fold(0, func(a, b int) int {
		return a + b
	})

	if 0 != arrf {
		t.Errorf("Expect len=%d Actual=%d", 0, arrf)
	}

	arrD := arr.MapString(func(item int) string {
		return fmt.Sprintf("%d", item)
	})

	for _, v := range arrD {
		t.Log(v)
	}
}
