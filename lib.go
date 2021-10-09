package twitter_text_parse_go

// #cgo CPPFLAGS: -I${SRCDIR}/include
// #cgo LDFLAGS: -L${SRCDIR}/lib -ltwitter_text_parse_go -Wl,--no-as-needed -ldl -lm
// #include "twitter_text_parse_go.h"
import "C"
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

var e = errors.New("twitter-text parse failed")

func Parse(s string) (Res, error) {
	res := C.tw_text_parse(C.CString(s))
	if res == nil {
		return Res{}, e
	}

	return Res{
		WeightedLength: int(res.weighted_length),
		Permillage:     int(res.permillage),
		IsValid:        res.is_valid != 0,
		DisplayTextRange: Range{
			Start: int(res.display_text_range_start),
			End:   int(res.display_text_range_end),
		},
		ValidTextRange: Range{
			Start: int(res.valid_text_range_start),
			End:   int(res.valid_text_range_end),
		},
	}, nil
}
