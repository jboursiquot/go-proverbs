package proverbs_test

import (
	"reflect"
	"testing"

	"github.com/jboursiquot/go-proverbs"
)

func TestRandom(t *testing.T) {
	if p := proverbs.Random(); p == nil {
		t.Fail()
	}
}

func TestNth(t *testing.T) {
	cases := map[string]struct {
		input int
		want  string
		err   error
	}{
		"second": {input: 1, want: "Concurrency is not parallelism."},
		"fifth":  {input: 4, want: "Make the zero value useful."},
		"error":  {input: 20, want: "", err: &proverbs.ErrNthProverbNotFound{}},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			if p, err := proverbs.Nth(tc.input); err != nil {
				if reflect.TypeOf(tc.err) != reflect.TypeOf(err) {
					t.Fail()
				}
			} else {
				if tc.want != p.Saying {
					t.Fail()
				}
			}
		})
	}
}

func TestAll(t *testing.T) {
	list, err := proverbs.All()
	if err != nil {
		t.Fail()
	}
	if len(list) <= 0 {
		t.Fail()
	}
}
