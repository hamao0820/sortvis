package gui

import (
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
)

func (a Algorithm) String() string {
	return string(a)
}

func (a Algorithm) Pretty() string {
	name := cases.Title(language.English).String(a.String() + " sort")
	return name
}
