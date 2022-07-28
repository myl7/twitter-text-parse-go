package pkg

import "errors"

type Range struct {
	Start int
	End   int
}

type Res struct {
	WeightedLength   int
	Permillage       int
	IsValid          bool
	DisplayTextRange Range
	ValidTextRange   Range
}

var ErrParseFailed = errors.New("twitter-text parse failed")
