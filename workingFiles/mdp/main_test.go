package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

const (
	inputFile  = "./testdata/test1.md"
	goldenFile = "./testdata/test1.md.html"
)

func TestParseContent(t *testing.T) {
	result, err := parseContent(inputFile, "")
	if err != nil {
		t.Fatal(err)
	}

	expected, err := ioutil.ReadFile(goldenFile)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(expected, result) {
		t.Logf("golden:\n%s\n", expected)
		t.Logf("result:\n%s\n", result)
		t.Error("Result content does not match golden file")
	}
}

func TestRun(t *testing.T) {
	var mockStdOut bytes.Buffer
	if err := run(inputFile, "", &mockStdOut, true); err != nil {
		t.Fatal(err)
	}

	resultFile := strings.TrimSpace(mockStdOut.String())

	result, err := ioutil.ReadFile(resultFile)
	if err != nil {
		t.Fatal(err)
	}

	expected, err := ioutil.ReadFile(goldenFile)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(expected, result) {
		t.Logf("golden:\n%s\n", expected)
		t.Logf("result:\n%s\n", result)
		t.Error("Result content does not match golden file")
	}

	os.Remove(resultFile)
}
