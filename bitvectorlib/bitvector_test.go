package bitvectorlib_test

import (
	"fmt"
	"testing"

	"github.com/georgfedermann/bitvector/bitvectorlib"
)

func TestString(t *testing.T) {
	expect := "{1 2 3}"
	intSet := new(bitvectorlib.IntSet)
	intSet.Add(1)
	intSet.Add(2)
	intSet.Add(3)
	actual := intSet.String()
	if actual != expect {
		t.Fatal(fmt.Sprintf("expected: '%s' but was '%s'\n", expect, actual))
	}
}

func TestUnion(t *testing.T) {
	expect := "{0 27 51 68 177 321 404}"
	a, b := new(bitvectorlib.IntSet), new(bitvectorlib.IntSet)
	a.Add(27)
	a.Add(404)
	a.Add(321)
	b.Add(51)
	b.Add(68)
	b.Add(177)
	b.Add(321)
	b.Add(0)
	a.UnionWith(*b)
	actual := a.String()
	if actual != expect {
		t.Fatal(fmt.Sprintf("expected:\n'%s' but was\n'%s'\n", expect, actual))
	}
}
