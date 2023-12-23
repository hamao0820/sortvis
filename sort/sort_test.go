package sort

import (
	"math/rand"
	"reflect"
	"testing"
)

type args struct {
	arr []int
}

type TestCase struct {
	name string
	args args
	want []int
}

func Series() []int {
	series := make([]int, 99)
	for i := 0; i < 99; i++ {
		series[i] = i
	}

	return series
}

func Shuffled() []int {
	rand.NewSource(42)
	shuffled := Series()
	rand.Shuffle(len(shuffled), func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})

	return shuffled
}

var tests = []TestCase{
	{"test1", args{[]int{3, 2, 1}}, []int{1, 2, 3}},
	{"test2", args{[]int{3, 2, 1, 4, 5, 6, 7, 8, 9, 10}}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	{"test3", args{[]int{3, 7, 2, 8, 1, 6, 4, 0, -5, 6, 9}}, []int{-5, 0, 1, 2, 3, 4, 6, 6, 7, 8, 9}},
	{"test4", args{[]int{2, 1}}, []int{1, 2}},
	{"test5", args{[]int{1}}, []int{1}},
	{"test6", args{[]int{1, 1, 1, 1, 1}}, []int{1, 1, 1, 1, 1}},
	{"test7", args{[]int{1, 1, 1, 1}}, []int{1, 1, 1, 1}},
	{"test8", args{Shuffled()}, Series()},
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

func TestBubbleSortAsync(t *testing.T) {
	for _, tt := range tests {
		bubble := Sorters["bubble"]
		bubble.Init(tt.args.arr)
		t.Run(tt.name, func(t *testing.T) {
			for bubble.Next() {
			}
			if !reflect.DeepEqual(tt.args.arr, tt.want) {
				t.Errorf("BubbleSortAsync() = %v, want %v", tt.args.arr, tt.want)
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

func TestMergeSortAsync(t *testing.T) {
	for _, tt := range tests {
		merge := NewSorter(MergeSortAsync)
		merge.Init(tt.args.arr)
		t.Run(tt.name, func(t *testing.T) {
			for merge.Next() {
			}
			if !reflect.DeepEqual(tt.args.arr, tt.want) {
				t.Errorf("MergeSortAsync() = %v, want %v", tt.args.arr, tt.want)
			}
		})
	}
}

func TestHeapSort(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HeapSort(tt.args.arr)
			if !reflect.DeepEqual(tt.args.arr, tt.want) {
				t.Errorf("BubbleSort() = %v, want %v", tt.args.arr, tt.want)
			}
		})
	}
}

func TestHeapSortAsync(t *testing.T) {
	for _, tt := range tests {
		heap := NewSorter(HeapSortAsync)
		heap.Init(tt.args.arr)
		t.Run(tt.name, func(t *testing.T) {
			for heap.Next() {
			}
			if !reflect.DeepEqual(tt.args.arr, tt.want) {
				t.Errorf("HeapSortAsync() = %v, want %v", tt.args.arr, tt.want)
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

func TestQuickSortAsync(t *testing.T) {
	for _, tt := range tests {
		quick := NewSorter(QuickSortAsync)
		quick.Init(tt.args.arr)
		t.Run(tt.name, func(t *testing.T) {
			for quick.Next() {
			}
			if !reflect.DeepEqual(tt.args.arr, tt.want) {
				t.Errorf("QuickSortAsync() = %v, want %v", tt.args.arr, tt.want)
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

func TestSelectionSortAsync(t *testing.T) {
	for _, tt := range tests {
		selection := NewSorter(SelectionSortAsync)
		selection.Init(tt.args.arr)
		t.Run(tt.name, func(t *testing.T) {
			for selection.Next() {
			}
			if !reflect.DeepEqual(tt.args.arr, tt.want) {
				t.Errorf("SelectionSortAsync() = %v, want %v", tt.args.arr, tt.want)
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

func TestBucketSortAsync(t *testing.T) {
	for _, tt := range tests {
		bucket := NewSorter(BucketSortAsync)
		bucket.Init(tt.args.arr)
		t.Run(tt.name, func(t *testing.T) {
			for bucket.Next() {
			}
			if !reflect.DeepEqual(tt.args.arr, tt.want) {
				t.Errorf("BucketSortAsync() = %v, want %v", tt.args.arr, tt.want)
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

func TestInsertionSortAsync(t *testing.T) {
	for _, tt := range tests {
		insertion := NewSorter(InsertionSortAsync)
		insertion.Init(tt.args.arr)
		t.Run(tt.name, func(t *testing.T) {
			for insertion.Next() {
			}
			if !reflect.DeepEqual(tt.args.arr, tt.want) {
				t.Errorf("InsertionSortAsync() = %v, want %v", tt.args.arr, tt.want)
			}
		})
	}
}

func TestShellSort(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ShellSort(tt.args.arr)
			if !reflect.DeepEqual(tt.args.arr, tt.want) {
				t.Errorf("BubbleSort() = %v, want %v", tt.args.arr, tt.want)
			}
		})
	}
}

func TestShellSortAsync(t *testing.T) {
	for _, tt := range tests {
		shell := NewSorter(ShellSortAsync)
		shell.Init(tt.args.arr)
		t.Run(tt.name, func(t *testing.T) {
			for shell.Next() {
			}
			if !reflect.DeepEqual(tt.args.arr, tt.want) {
				t.Errorf("ShellSortAsync() = %v, want %v", tt.args.arr, tt.want)
			}
		})
	}
}
