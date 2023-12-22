package sort

import (
	"reflect"
	"testing"
)

type args struct {
	arr []int
}

var tests = []struct {
	name string
	args args
	want []int
}{
	{"test1", args{[]int{3, 2, 1}}, []int{1, 2, 3}},
	{"test2", args{[]int{3, 2, 1, 4, 5, 6, 7, 8, 9, 10}}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	{"test3", args{[]int{3, 7, 2, 8, 1, 6, 4, 0, -5, 6, 9}}, []int{-5, 0, 1, 2, 3, 4, 6, 6, 7, 8, 9}},
	{"test4", args{[]int{2, 1}}, []int{1, 2}},
	{"test5", args{[]int{1}}, []int{1}},
	{"test6", args{[]int{1, 1, 1, 1, 1}}, []int{1, 1, 1, 1, 1}},
	{"test7", args{[]int{1, 1, 1, 1}}, []int{1, 1, 1, 1}},
}

func TestBubbleSort(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BubbleSort(tt.args.arr)
			if !reflect.DeepEqual(tt.args.arr, tt.want) {
				t.Errorf("BubbleSort() = %v, want %v", tt.args.arr, tt.want)
			}
		})
	}
}

func TestMergeSort(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MergeSort(tt.args.arr)
			if !reflect.DeepEqual(tt.args.arr, tt.want) {
				t.Errorf("BubbleSort() = %v, want %v", tt.args.arr, tt.want)
			}
		})
	}
}

func TestHeapSort(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Heapsort(tt.args.arr)
			if !reflect.DeepEqual(tt.args.arr, tt.want) {
				t.Errorf("BubbleSort() = %v, want %v", tt.args.arr, tt.want)
			}
		})
	}
}

func TestQuickSort(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			QuickSort(tt.args.arr)
			if !reflect.DeepEqual(tt.args.arr, tt.want) {
				t.Errorf("BubbleSort() = %v, want %v", tt.args.arr, tt.want)
			}
		})
	}
}

func TestSelectionSort(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SelectionSort(tt.args.arr)
			if !reflect.DeepEqual(tt.args.arr, tt.want) {
				t.Errorf("BubbleSort() = %v, want %v", tt.args.arr, tt.want)
			}
		})
	}
}

func TestBucketSort(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BucketSort(tt.args.arr)
			if !reflect.DeepEqual(tt.args.arr, tt.want) {
				t.Errorf("BubbleSort() = %v, want %v", tt.args.arr, tt.want)
			}
		})
	}
}

func TestInsertionSort(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InsertionSort(tt.args.arr)
			if !reflect.DeepEqual(tt.args.arr, tt.want) {
				t.Errorf("BubbleSort() = %v, want %v", tt.args.arr, tt.want)
			}
		})
	}
}
