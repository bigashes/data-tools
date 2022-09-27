package slices

import (
	"reflect"
	"testing"
)

func TestInArray(t *testing.T) {
	arr := []string{"a", "b", "c"}

	ex := InArray[string]("a", arr)
	if !ex {
		t.Errorf("'a' should in arr")
	}

	ex = InArray[string]("d", arr)
	if ex {
		t.Errorf("'d' should not in arr")
	}
}

func TestArrayIndex(t *testing.T) {
	arr := []string{"a", "b", "c"}

	index := ArrayIndex[string]("a", arr)
	if index != 0 {
		t.Errorf("'a' should index as 0")
	}
}

type P struct {
	Id   int64
	Name string
}

func TestArrayColumn(t *testing.T) {
	arr := []P{{1, "a"}, {2, "b"}, {3, "c"}, {4, "d"}}

	columnArr := ArrayColumn[int64, P](arr, "Id")
	if !reflect.DeepEqual(columnArr, []int64{1, 2, 3, 4}) {
		t.Errorf("column : %+v \n", columnArr)
	}

	mapArr := []map[string]int64{{"id": 1, "age": 18}, {"id": 1, "age": 18}, {"id": 1, "age": 18}}

	columnMap := ArrayColumn[int64, map[string]int64](mapArr, "id")
	if !reflect.DeepEqual(columnMap, []int64{1, 1, 1}) {
		t.Logf("column : %+v \n", columnMap)
	}
}
