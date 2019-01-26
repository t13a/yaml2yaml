package main

import (
	"strings"
	"testing"
)

func assertFormat(t *testing.T, in string, want string) {
	got, err := Format([]byte(in))
	if err != nil {
		t.Errorf("%s", err)
	}

	trimmedWant := strings.TrimLeft(want, "\n")
	if string(got) != trimmedWant {
		t.Errorf("Unexpected format string\ngot:\n%s\n\nwant:\n%s", string(got), trimmedWant)
	}
}

func TestFormat_Boolean(t *testing.T) {
	var in = `
F:
- False
- No
- false
- no
T:
- True
- Yes
- true
- yes
`
	var want = `
F:
- false
- false
- false
- false
T:
- true
- true
- true
- true
`
	assertFormat(t, in, want)
}

func TestFormat_List(t *testing.T) {
	var in = `
A: [ 1, 2, 3, 4 ]
`
	var want = `
A:
- 1
- 2
- 3
- 4
`
	assertFormat(t, in, want)
}

func TestFormat_Object(t *testing.T) {
	var in = `
A: { foo: 1, bar: 2, baz: 3, qux: 4 }
`
	var want = `
A:
  bar: 2
  baz: 3
  foo: 1
  qux: 4
`
	assertFormat(t, in, want)
}

func TestFormat_String(t *testing.T) {
	var in = `
A: foo
B: 'foo'
C: "foo"
D: "foo\nbar"
E: |
  foo
  bar
`
	var want = `
A: foo
B: foo
C: foo
D: |-
  foo
  bar
E: |
  foo
  bar
`
	assertFormat(t, in, want)
}
