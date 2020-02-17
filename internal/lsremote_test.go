package internal

import (
	"reflect"
	"testing"
)

func Test_sortVersions(t *testing.T) {

	cases := []struct {
		name   string
		in     []goversion
		expect []goversion
	}{
		{
			name:   "minor version",
			in:     []goversion{"1.13", "1.11", "1.12"},
			expect: []goversion{"1.11", "1.12", "1.13"},
		},
		{
			name:   "patch version",
			in:     []goversion{"1.13.1", "1.13.3", "1.13.2"},
			expect: []goversion{"1.13.1", "1.13.2", "1.13.3"},
		},
		{
			name:   "with rc",
			in:     []goversion{"1.13rc1", "1.13.1", "1.13"},
			expect: []goversion{"1.13rc1", "1.13", "1.13.1"},
		},
		{
			name:   "complex case",
			in:     []goversion{"1.13rc1", "1.13.1", "1.13", "1.1", "1.14rc1", "1.13rc2"},
			expect: []goversion{"1.1", "1.13rc1", "1.13rc2", "1.13", "1.13.1", "1.14rc1"},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			act := sortVersions(c.in)
			if !reflect.DeepEqual(act, c.expect) {
				t.Errorf("must be sorted as %v but %v", c.expect, act)
			}
		})
	}
}
