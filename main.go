package main

import (
	"testing"
)

func TestGetUTFLength(t *testing.T) {
	validInput := []byte{208, 159, 209, 128, 208, 184, 208, 178, 208, 181, 209, 130}
	result, err := GetUTFLength(validInput)
	if err != nil {
		t.Errorf("Не ожидалась ошибка, но получилась: %v", err)
	}
	if result != 6 {
		t.Errorf("Ожидалось 6, но получили %d", result)
	}

	invalidInput := []byte{0xff, 0xfe, 0xfd}
	_, err = GetUTFLength(invalidInput)
	if err == nil {
		t.Errorf("Ожидалась ошибка, но ее не было")
	}
}

// type Test struct {
// 	in []byte
// 	outOne string
// 	outTwo string
// }

// var tests = []Test{
// 	{[208 159 209 128 208 184 208 178 208 181 209 130], 12, nil},
// 	{"Привет", 0, ErrInvalidUTF8},
// }

// func TestGetUTFLength(t *testing.T) {
// 	for i, test := range tests {

// 		if GetUTFLength(in) != test.outOne {
// 			t.Errorf("#%d: Size(%d)=%s; want %s", i, test.in, size, test.out)
// 		}
// 	}
// }
