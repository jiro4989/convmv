package db

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBindJsonToStruct(t *testing.T) {
	tests := []struct {
		desc string
		in   string
	}{
		{desc: "バインドしてエラーがでない", in: "../testdata/Map001.json"},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			b, err := ioutil.ReadFile(tt.in)
			assert.Nil(t, err, tt.desc)
			var m MapJson
			err = json.Unmarshal(b, &m)
			assert.Nil(t, err, tt.desc)
		})
	}
}
