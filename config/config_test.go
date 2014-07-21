package config

import (
    "fmt"
    "testing"
    "strings"
)

func init() {
    reader := strings.NewReader(configText)
    Load(reader)
}

func TestBasicString(t *testing.T) {
    assert(t, "teststring", Get("basic_string"))
}

func TestOneDeepString(t *testing.T) {
    assert(t, "one", Get("deep:one_string"))
}

func TestTwoDeepString(t *testing.T) {
    assert(t, "two", Get("deep:two:two_string"))
}

func TestNotFoundBasic(t *testing.T) {
    assert(t, "", Get("nothing"))
}

func TestNotFoundOneDeep(t *testing.T) {
    assert(t, "", Get("nothing:here"))
}

func TestNotFoundTwoDeep(t *testing.T) {
    assert(t, "", Get("deep:nothing:here"))
}

func TestEmpty(t *testing.T) {
    assert(t, "", Get(""))
}

func assert(t *testing.T, expected, actual string) {
    if expected != actual {
        t.Errorf(fmt.Sprintf("Expected %v, got %v", expected, actual))
    }
}

var configText = `
basic_string: teststring
deep:
    one_string: one
    two:
        two_string: two
`
