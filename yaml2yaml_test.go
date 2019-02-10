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
	// https://yaml.org/type/bool.html
	assertFormatJSON(t, "y", "true\n")
	assertFormatJSON(t, "Y", "true\n")
	assertFormatJSON(t, "yes", "true\n")
	assertFormatJSON(t, "Yes", "true\n")
	assertFormatJSON(t, "YES", "true\n")
	assertFormatJSON(t, "n", "false\n")
	assertFormatJSON(t, "N", "false\n")
	assertFormatJSON(t, "no", "false\n")
	assertFormatJSON(t, "No", "false\n")
	assertFormatJSON(t, "NO", "false\n")
	assertFormatJSON(t, "true", "true\n")
	assertFormatJSON(t, "True", "true\n")
	assertFormatJSON(t, "TRUE", "true\n")
	assertFormatJSON(t, "false", "false\n")
	assertFormatJSON(t, "False", "false\n")
	assertFormatJSON(t, "FALSE", "false\n")
	assertFormatJSON(t, "on", "true\n")
	assertFormatJSON(t, "On", "true\n")
	assertFormatJSON(t, "ON", "true\n")
	assertFormatJSON(t, "off", "false\n")
	assertFormatJSON(t, "Off", "false\n")
	assertFormatJSON(t, "OFF", "false\n")
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
	// https://yaml.org/type/bool.html
	assertFormatYAML(t, "y", "true\n")
	assertFormatYAML(t, "Y", "true\n")
	assertFormatYAML(t, "yes", "true\n")
	assertFormatYAML(t, "Yes", "true\n")
	assertFormatYAML(t, "YES", "true\n")
	assertFormatYAML(t, "n", "false\n")
	assertFormatYAML(t, "N", "false\n")
	assertFormatYAML(t, "no", "false\n")
	assertFormatYAML(t, "No", "false\n")
	assertFormatYAML(t, "NO", "false\n")
	assertFormatYAML(t, "true", "true\n")
	assertFormatYAML(t, "True", "true\n")
	assertFormatYAML(t, "TRUE", "true\n")
	assertFormatYAML(t, "false", "false\n")
	assertFormatYAML(t, "False", "false\n")
	assertFormatYAML(t, "FALSE", "false\n")
	assertFormatYAML(t, "on", "true\n")
	assertFormatYAML(t, "On", "true\n")
	assertFormatYAML(t, "ON", "true\n")
	assertFormatYAML(t, "off", "false\n")
	assertFormatYAML(t, "Off", "false\n")
	assertFormatYAML(t, "OFF", "false\n")
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
