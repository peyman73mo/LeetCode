package ZigzagConversion

import (
	"testing"
)

type Sample struct {
	Str      string
	NumRows  int
	Expected string
}

var samples = []Sample{
	Sample{
		Str:      "PAYPALISHIRING",
		NumRows:  3,
		Expected: "PAHNAPLSIIGYIR",
	},
	Sample{
		Str:      "PAYPALISHIRING",
		NumRows:  4,
		Expected: "PINALSIGYAHRPI",
	},
}

func TestZigzag(t *testing.T) {

	for _, sample := range samples {
		if convert(sample.Str, sample.NumRows) != sample.Expected {
			t.Errorf("expected : %q , got : %q", sample.Expected, convert(sample.Str, sample.NumRows))
		}
	}
}
