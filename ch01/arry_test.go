package Arryt_test

import "testing"

func TestArryInit(t *testing.T) {
	var arr [3]int
	arr1 := [4]int{1, 2, 3, 4}
	arr2 := [...]int{1, 3, 4, 5}
	t.Log(arr[1], arr[2])
	t.Log(arr1, arr2)
}

func TestArryTravel(t *testing.T) {
	arr2 := [...]int{1, 3, 4, 5}

	for i := 0; i < len(arr2); i++ {
		t.Log(arr2[i])
	}

	for _, e := range arr2 {
		t.Log(e)
	}
}
func TestArrySe(t *testing.T) {
	arr2 := [...]int{1, 3, 4, 5}
	arr2_sec := arr2[3:]
	t.Log(arr2_sec)
}