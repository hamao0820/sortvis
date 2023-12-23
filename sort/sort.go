package sort

import (
	"sync"
)

type Status int

const (
	UnInitialized Status = iota
	Initialized
	Sorting
	Done
)

type SortIterator struct {
	Arr    []int
	cStep  chan struct{}
	cDone  chan struct{}
	wg     *sync.WaitGroup
	Status Status
	sort   func(arr []int, cStep, cDone chan struct{}, wg *sync.WaitGroup)
}

var Sorters = map[string]*SortIterator{
	"bubble": NewSorter(BubbleSortAsync),
	"heap":   NewSorter(HeapSortAsync),
}

func NewSorter(sort func(arr []int, cStep, cDone chan struct{}, wg *sync.WaitGroup)) *SortIterator {
	return &SortIterator{
		Arr:    []int{},
		cStep:  make(chan struct{}),
		cDone:  make(chan struct{}),
		wg:     &sync.WaitGroup{},
		Status: UnInitialized,
		sort:   sort,
	}
}

func (s *SortIterator) Init(arr []int) {
	s.Arr = arr
	s.Status = Initialized
	s.cStep = make(chan struct{})
	s.cDone = make(chan struct{})
	s.wg = &sync.WaitGroup{}
}

func (s *SortIterator) step() {
	s.cStep <- struct{}{}
}

func (s *SortIterator) Next() bool {
	if s.Status == Done {
		return false
	}

	if s.Status == UnInitialized {
		return false
	}

	if s.Status == Initialized {
		s.Status = Sorting

		s.wg.Add(1) // for the first step
		go s.sort(s.Arr, s.cStep, s.cDone, s.wg)
		s.wg.Wait()
	}

	select {
	case <-s.cDone:
		s.Status = Done
		close(s.cStep)
		close(s.cDone)
		return false
	default:
		s.wg.Add(1)
		s.step()
		s.wg.Wait()

		select {
		case <-s.cDone:
			s.Status = Done
			close(s.cStep)
			close(s.cDone)
			return false
		default:
		}
	}

	return true
}
