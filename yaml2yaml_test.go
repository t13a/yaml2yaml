package main

import (
	"strings"
	"testing"
)

func assertFormatJSON(t *testing.T, in string, want string) {
	assertFormat(t, FormatJSON, in, want)
}

func assertFormatYAML(t *testing.T, in string, want string) {
	assertFormat(t, FormatYAML, in, want)
}

func assertFormat(t *testing.T, formatFunc func([]byte) ([]byte, error), in string, want string) {
	got, err := formatFunc([]byte(in))
	if err != nil {
		t.Errorf("%s", err)
	}
	if string(got) != want {
		t.Errorf("Unexpected format string\ngot:\n%s\n\nwant:\n%s", string(got), want)
	}
}

func TestFormatJSON_Boolean(t *testing.T) {
	assertFormatJSON(t, "False", "false\n")
	assertFormatJSON(t, "No", "false\n")
	assertFormatJSON(t, "True", "true\n")
	assertFormatJSON(t, "Yes", "true\n")
	assertFormatJSON(t, "false", "false\n")
	assertFormatJSON(t, "no", "false\n")
	assertFormatJSON(t, "true", "true\n")
	assertFormatJSON(t, "yes", "true\n")
}

func TestFormatJSON_List(t *testing.T) {
	var in = `
[1, 2, 3, 4]
`
	var want = strings.TrimLeft(`
[
  1,
  2,
  3,
  4
]
`, "\n")
	assertFormatJSON(t, in, want)
}

func TestFormatJSON_Null(t *testing.T) {
	assertFormatJSON(t, "null", "null\n")
}

func TestFormatJSON_Object(t *testing.T) {
	var in = `
{ foo: 1, bar: 2, baz: 3, qux: 4 }
`
	var want = strings.TrimLeft(`
{
  "bar": 2,
  "baz": 3,
  "foo": 1,
  "qux": 4
}
`, "\n")
	assertFormatJSON(t, in, want)
}

func TestFormatYAML_Boolean(t *testing.T) {
	assertFormatYAML(t, "False", "false\n")
	assertFormatYAML(t, "No", "false\n")
	assertFormatYAML(t, "True", "true\n")
	assertFormatYAML(t, "Yes", "true\n")
	assertFormatYAML(t, "false", "false\n")
	assertFormatYAML(t, "no", "false\n")
	assertFormatYAML(t, "true", "true\n")
	assertFormatYAML(t, "yes", "true\n")
}

func TestFormatYAML_List(t *testing.T) {
	var in = `
[1, 2, 3, 4]
`
	var want = strings.TrimLeft(`
- 1
- 2
- 3
- 4
`, "\n")
	assertFormatYAML(t, in, want)
}

func TestFormatYAML_Null(t *testing.T) {
	assertFormatYAML(t, "null", "null\n")
}

func TestFormatYAML_Object(t *testing.T) {
	var in = `
{ foo: 1, bar: 2, baz: 3, qux: 4 }
`
	var want = strings.TrimLeft(`
bar: 2
baz: 3
foo: 1
qux: 4
`, "\n")
	assertFormatYAML(t, in, want)
}

func TestFormatYAML_String(t *testing.T) {
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
	assertFormatYAML(t, in, want)
}
