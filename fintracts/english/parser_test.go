package english

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var (
	successPath = "./contracts/success"
	failurePath = "./contracts/failures"
)

func TestSuccess(t *testing.T) {
	files, err := ioutil.ReadDir(successPath)
	if err != nil {
		t.Error(err)
		return
	}

	for _, f := range files {
		// Only test .txt files
		if !strings.HasSuffix(f.Name(), ".txt") {
			continue
		}

		bytes, err := os.ReadFile(filepath.Join(successPath, f.Name()))
		if err != nil {
			t.Error(err)
		}

		contract, err := Parse(bytes)
		if err != nil {
			t.Error(err)
		}

		expectedBytes, err := os.ReadFile(filepath.Join(successPath, strings.Replace(f.Name(), ".txt", ".json", -1)))
		if err != nil {
			t.Error(err)
		}

		result, err := contract.String()
		if err != nil {
			t.Error(err)
		}

		expected := strings.TrimSpace(string(expectedBytes))

		if expected != string(result) {
			t.Errorf("expected output did not match for %s", f.Name())
		}
	}
}

func TestFailures(t *testing.T) {
	files, err := ioutil.ReadDir(failurePath)
	if err != nil {
		t.Error(err)
		return
	}

	for _, f := range files {
		bytes, err := os.ReadFile(filepath.Join(failurePath, f.Name()))
		if err != nil {
			t.Error(err)
		}

		_, err = Parse(bytes)
		if err == nil {
			t.Errorf("expected error for %s", f.Name())
		}
	}
}
