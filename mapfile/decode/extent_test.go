package decode_test

import (
	"errors"
	"github.com/geo-data/mapfile/mapfile/decode"
	"github.com/geo-data/mapfile/types"
	"reflect"
	"testing"
)

var extentTests = []struct {
	input    string
	expected *types.Extent // expected result
}{
	{`
EXTENT -0.5 50.977222 0.5 51.977222`, types.NewExtent(-0.5, 50.977222, 0.5, 51.977222)},
	{`
EXTENT 0 0 0 0`, types.NewExtent(0, 0, 0, 0)},
}

var extentErrorTests = []struct {
	input    string
	expected error // expected result
}{
	{`
FOOBAR 
END`, errors.New(`expected token EXTENT, got: "FOOBAR"`)},
	{`
EXTENT 1 2 3`, decode.EndOfTokens},
	{`
EXTENT foo 2 3 4`, errors.New(`invalid syntax for X coordinate: "foo"`)},
	{`
EXTENT 1 foo 3 4`, errors.New(`invalid syntax for Y coordinate: "foo"`)},
	{`
EXTENT 1 2 foo 4`, errors.New(`invalid syntax for X coordinate: "foo"`)},
	{`
EXTENT 1 2 3 foo`, errors.New(`invalid syntax for Y coordinate: "foo"`)},
}

func TestDecodeExtent(t *testing.T) {
	for _, tt := range extentTests {
		dec, err := decode.DecodeString(tt.input)
		if err != nil {
			t.Error("For decoding:", tt.input, ", expected error:", nil, ", got:", err)
			continue
		}
		actual, err := dec.Extent()
		if err != nil {
			t.Error("For:", tt.input, ", expected error:", nil, ", got:", err)
			continue
		}

		if !reflect.DeepEqual(tt.expected, actual) {
			t.Error("For:", tt.input, ", expected:", tt.expected, ", got:", actual)
		}
	}
}

func TestDecodeExtentError(t *testing.T) {
	for _, tt := range extentErrorTests {
		dec, err := decode.DecodeString(tt.input)
		if err != nil {
			t.Error("For decoding:", tt.input, ", expected error:", nil, ", got:", err)
			continue
		}
		actual, err := dec.Extent()
		if actual != nil {
			t.Error("For:", tt.input, ", expected extent:", nil, ", got:", actual)
			continue
		}

		if err.Error() != tt.expected.Error() {
			t.Error("For:", tt.input, ", expected error:", tt.expected, ", got:", err)
		}
	}
}