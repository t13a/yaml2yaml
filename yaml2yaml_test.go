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

	if string(got) != want {
		t.Errorf("Unexpected format string\ngot:\n%s\n\nwant:\n%s", string(got), want)
	}
}

func TestFormat_Boolean(t *testing.T) {
	assertFormat(t, "False", "false\n")
	assertFormat(t, "No", "false\n")
	assertFormat(t, "True", "true\n")
	assertFormat(t, "Yes", "true\n")
	assertFormat(t, "false", "false\n")
	assertFormat(t, "no", "false\n")
	assertFormat(t, "true", "true\n")
	assertFormat(t, "yes", "true\n")
}

func TestFormat_List(t *testing.T) {
	var in = `
[1, 2, 3, 4]
`
	var want = strings.TrimLeft(`
- 1
- 2
- 3
- 4
`, "\n")
	assertFormat(t, in, want)
}

func TestFormat_Null(t *testing.T) {
	assertFormat(t, "null", "null\n")
}

func TestFormat_Object(t *testing.T) {
	var in = `
{ foo: 1, bar: 2, baz: 3, qux: 4 }
`
	var want = strings.TrimLeft(`
bar: 2
baz: 3
foo: 1
qux: 4
`, "\n")
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
	var want = strings.TrimLeft(`
A: foo
B: foo
C: foo
D: |-
  foo
  bar
E: |
  foo
  bar
`, "\n")
	assertFormat(t, in, want)
}
