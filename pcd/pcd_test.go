package pcd

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {

	outStream := new(bytes.Buffer)
	errStream := new(bytes.Buffer)

	type pattern struct {
		args      []string
		expectMsg string
	}

	// ok test
	func() {
		reset := setTestEnv("PWD", "/path/to/test")
		defer reset()

		patterns := []pattern{
			{strings.Split("", " "), "/\n/path\n/path/to\n/path/to/test\n"},
			{strings.Split("-version", " "), fmt.Sprintf("pcd's version is %s\n", version)},
		}

		for _, p := range patterns {
			err := Run(p.args, outStream, errStream)
			if err != nil {
				t.Errorf("Error: %s", err)
			}
			if outStream.String() != p.expectMsg {
				t.Errorf("Output=%q, want %q", outStream.String(), p.expectMsg)
			}
			outStream.Reset()
		}
	}()

	// parse error test
	func() {
		reset := setTestEnv("PWD", "/path/to/test")
		defer reset()

		err := Run(strings.Split("-error", " "), outStream, errStream)
		if err == nil {
			t.Errorf("Want to Error but not error")
		}
		if !strings.Contains(fmt.Sprint(err), "args parse error") {
			t.Errorf("Want to conatains parse error in error")
		}

		outStream.Reset()
	}()
}

func TestPrintVersion(t *testing.T) {
	outStream := new(bytes.Buffer)

	err := printVersion(outStream)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	expected := fmt.Sprintf("pcd's version is %s\n", version)
	if outStream.String() != expected {
		t.Errorf("Output=%q, want %q", outStream.String(), expected)
	}
}

func TestPrintPathList(t *testing.T) {
	reset := setTestEnv("PWD", "/path/to/test")
	defer reset()

	outStream := new(bytes.Buffer)
	type pattern struct {
		pwd, expected string
	}

	// ok
	func() {

		patterns := []pattern{
			{"/path/to/test", "/\n/path\n/path/to\n/path/to/test"},
			{"C:/path/to/test", "C:/\nC:/path\nC:/path/to\nC:/path/to/test"},
		}
		for _, p := range patterns {
			reset := setTestEnv("PWD", p.pwd)
			defer reset()

			err := printPathList(outStream)
			if err != nil {
				t.Errorf("Error: %s", err)
			}

			expected := p.expected
			if !strings.Contains(outStream.String(), expected) {
				t.Errorf("Output=%q, want %q", outStream.String(), expected)
			}

			outStream.Reset()
		}
	}()

}

func setTestEnv(key, val string) func() {
	preVal := os.Getenv(key)
	os.Setenv(key, val)
	return func() {
		os.Setenv(key, preVal)
	}
}
