package slices

import (
	"reflect"
)

//InArray
func InArray[E comparable](needle E, haystack []E) bool {
	for _, e := range haystack {
		if e == needle {
			return true
		}
	}

	return false
}

//ArrayInde
func ArrayIndex[E comparable](needle E, haystack []E) int64 {
	for index, e := range haystack {
		if e == needle {
			return int64(index)
		}
	}

	return -1
}

//ArrayUnique
func ArrayUnique[E comparable](haystack []E) []E {
	var (
		m    = make(map[E]struct{})
		resp []E
	)
	for _, e := range haystack {
		if _, ok := m[e]; ok {
			continue
		}
		m[e] = struct{}{}
		resp = append(resp, e)
	}

	return resp
}

//MaxMinParams
type MaxMinParams interface {
	int8 | int32 | int64 | float64 | string
}

//ArrayMin returns the minimum value in the array
func ArrayMin[T MaxMinParams](item ...T) (min T) {
	if len(item) == 0 {
		return min
	}

	min = item[0]
	for _, v := range item {
		if v < min {
			min = v
		}
	}
	return min
}

//ArrayMax returns the maximum value in the array
func ArrayMax[T MaxMinParams](item ...T) (max T) {
	if len(item) == 0 {
		return max
	}

	max = item[0]
	for _, v := range item {
		if v > max {
			max = v
		}
	}
	return max
}

//ArrayDiff
func ArrayDiff[T comparable](slice1, slice2 []T) []T {
	m := make(map[T]int)
	nn := make([]T, 0)
	for _, v := range slice2 {
		m[v]++
	}

	for _, val := range slice1 {
		if _, ok := m[val]; !ok {
			nn = append(nn, val)
		}
	}
	return nn
}

//ArrayIntersect
func ArrayIntersect[T comparable](nums1 []T, nums2 []T) []T {
	m := make(map[T]uint8, 0)
	for _, v := range nums1 {
		m[v] += 1
	}

	var ret []T
	for _, v := range nums2 {
		if m[v] > 0 {
			m[v] = 0
			ret = append(ret, v)
		}
	}

	return ret
}

//ArrayColumn
func ArrayColumn[T, E any](haystack []E, k string) []T {
	values := []T{}
	switch reflect.TypeOf(haystack).Elem().Kind() {
	case reflect.Slice, reflect.Array:
		for _, v := range haystack {
			values = append(values, reflect.ValueOf(v).Index(int(reflect.ValueOf(k).Int())).Interface().(T))
		}
		break
	case reflect.Map:
		for _, v := range haystack {
			if !reflect.ValueOf(v).MapIndex(reflect.ValueOf(k)).IsValid() {
				continue
			}

			values = append(values, reflect.ValueOf(v).MapIndex(reflect.ValueOf(k)).Interface().(T))
		}
		break
	case reflect.Struct:
		for _, v := range haystack {
			values = append(values, reflect.ValueOf(v).FieldByName(reflect.ValueOf(k).String()).Interface().(T))
		}
		break
	}
	return values
}
