package decode

import (
	"fmt"
	"github.com/geo-data/mapfile/types"
)

func (t *Decoder) Metadata() (m types.Metadata, err error) {
	token := t.Value()
	if token != "METADATA" {
		err = fmt.Errorf("expected token METADATA, got: %s", token)
		return
	}
	t.Next()

	m = types.NewMetadata()

	for t != nil {
		var key, value types.String
		if t.Value() == "END" {
			break
		}
		if key, err = t.String(); err != nil {
			return
		}

		if t.Next().Value() == "END" {
			break
		}
		if value, err = t.String(); err != nil {
			return
		}

		m[string(key)] = string(value)

		t = t.Next()
	}

	return
}