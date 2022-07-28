package musl

// #cgo CPPFLAGS: -I${SRCDIR}/../../include
// #cgo LDFLAGS: -L${SRCDIR}/../../lib/x86_64-unknown-linux-musl -ltwitter_text_parse_go -Wl,--no-as-needed -ldl -lm
// #include "twitter_text_parse_go.h"
import "C"
import "github.com/myl7/twitter-text-parse-go/pkg"

func Parse(s string) (pkg.Res, error) {
	res := C.tw_text_parse(C.CString(s))
	if res == nil {
		return pkg.Res{}, pkg.ErrParseFailed
	}

	return pkg.Res{
		WeightedLength: int(res.weighted_length),
		Permillage:     int(res.permillage),
		IsValid:        res.is_valid != 0,
		DisplayTextRange: pkg.Range{
			Start: int(res.display_text_range_start),
			End:   int(res.display_text_range_end),
		},
		ValidTextRange: pkg.Range{
			Start: int(res.valid_text_range_start),
			End:   int(res.valid_text_range_end),
		},
	}, nil
}
