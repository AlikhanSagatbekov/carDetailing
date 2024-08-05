package tests

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/AlikhanSagatbekov/carDetailing/internal/jsonlog"
)

func TestJSONLogger_PrintError(t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "test_jsonlog_*.log")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name()) // Clean up

	logger := jsonlog.New(tmpfile, jsonlog.LevelError)

	logger.PrintError(errors.New("this is an error message"), nil)

	content, err := ioutil.ReadFile(tmpfile.Name())
	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(string(content), `"level":"ERROR"`) {
		t.Error("Expected ERROR log level, got:", string(content))
	}
}

func TestJSONLogger_MessageInformationInLogs(t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "test_jsonlog_*.log")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name()) // Clean up

	logger := jsonlog.New(tmpfile, jsonlog.LevelError)

	logger.PrintError(errors.New("this is an error message"), nil)

	content, err := ioutil.ReadFile(tmpfile.Name())
	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(string(content), "message") {
		t.Error("Expected message information in the log, got:", string(content))
	}
}

func TestJSONLogger_TraceInformationInLogs(t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "test_jsonlog_*.log")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name()) // Clean up

	logger := jsonlog.New(tmpfile, jsonlog.LevelError)

	logger.PrintError(errors.New("this is an error message"), nil)

	content, err := ioutil.ReadFile(tmpfile.Name())
	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(string(content), "trace") {
		t.Error("Expected trace information in the log, got:", string(content))
	}
}
