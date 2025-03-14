package snest_test

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/fletcharoo/snest"
)

type testConfig struct {
	URL     string `snest:"API_URL"`
	Port    int    `snest:"API_PORT"`
	Verbose bool   `snest:"API_VERBOSE"`
}

func Test_Load(t *testing.T) {
	expected := testConfig{
		URL:     "http://localhost",
		Port:    80,
		Verbose: true,
	}

	os.Setenv("API_URL", expected.URL)
	os.Setenv("API_PORT", fmt.Sprint(expected.Port))
	os.Setenv("API_VERBOSE", fmt.Sprint(expected.Verbose))

	s := new(testConfig)
	if err := snest.Load(s); err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}

	if !reflect.DeepEqual(*s, expected) {
		t.Fatalf("Unexpected result\nExpected: %v\nActual: %v", expected, *s)
	}
}

func Test_LoadWithDefault(t *testing.T) {
	expected := testConfig{
		URL:     "http://localhost",
		Port:    80,
		Verbose: true,
	}

	os.Setenv("API_URL", expected.URL)

	s := &testConfig{
		Port:    80,
		Verbose: true,
	}

	if err := snest.Load(s); err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}

	if !reflect.DeepEqual(*s, expected) {
		t.Fatalf("Unexpected result\nExpected: %v\nActual: %v", expected, *s)
	}
}
