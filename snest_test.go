package snest_test

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/fletcharoo/snest"
)

func Test_Load(t *testing.T) {
	type TestStruct struct {
		URL     string `snest:"API_URL"`
		Port    int    `snest:"API_PORT"`
		Verbose bool   `snest:"API_VERBOSE"`
	}

	expected := TestStruct{
		URL:     "http://localhost",
		Port:    80,
		Verbose: true,
	}

	os.Setenv("API_URL", expected.URL)
	os.Setenv("API_PORT", fmt.Sprint(expected.Port))
	os.Setenv("API_VERBOSE", fmt.Sprint(expected.Verbose))

	s := new(TestStruct)
	if err := snest.Load(s); err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}

	if !reflect.DeepEqual(*s, expected) {
		t.Fatalf("Unexpected result\nExpected: %v\nActual: %v", expected, *s)
	}
}
