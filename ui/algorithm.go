package gui

import (
	"fmt"

	"github.com/hamao0820/sortvis/sort/iter"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Algorithm string

const (
	Bubble    Algorithm = "bubble"
	Heap      Algorithm = "heap"
	Merge     Algorithm = "merge"
	Quick     Algorithm = "quick"
	Selection Algorithm = "selection"
	Bucket    Algorithm = "bucket"
	Insertion Algorithm = "insertion"
	Shell     Algorithm = "shell"
)

func (a Algorithm) String() string {
	return string(a)
}

func (a Algorithm) Pretty() string {
	name := cases.Title(language.English).String(a.String() + " sort")
	return name
}

func (a Algorithm) iterator() (*iter.SortIterator, error) {
	switch a {
	case Bubble:
		return iter.NewIterator(iter.BubbleSort), nil
	case Heap:
		return iter.NewIterator(iter.HeapSort), nil
	case Merge:
		return iter.NewIterator(iter.MergeSort), nil
	case Quick:
		return iter.NewIterator(iter.QuickSort), nil
	case Selection:
		return iter.NewIterator(iter.SelectionSort), nil
	case Bucket:
		return iter.NewIterator(iter.BucketSort), nil
	case Insertion:
		return iter.NewIterator(iter.InsertionSort), nil
	case Shell:
		return iter.NewIterator(iter.ShellSort), nil
	default:
		return nil, fmt.Errorf("algorithm %s is not supported", a)
	}
}
