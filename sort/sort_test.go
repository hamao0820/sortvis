package sort

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	type args struct {
		arr []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", args{[]int{3, 2, 1}}, []int{1, 2, 3}},
		{"test2", args{[]int{3, 2, 1, 4, 5, 6, 7, 8, 9, 10}}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{"test3", args{[]int{3, 7, 2, 8, 1, 6, 4, 0, -5, 6, 9}}, []int{-5, 0, 1, 2, 3, 4, 6, 6, 7, 8, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BubbleSort(tt.args.arr)
			if !reflect.DeepEqual(tt.args.arr, tt.want) {
				t.Errorf("BubbleSort() = %v, want %v", tt.args.arr, tt.want)
			}
		})
	}
}
