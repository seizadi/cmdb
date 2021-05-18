package utils

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
)

var outErrCmdString = []string{"sh", "-c", "echo out && >&2 echo error"}

func TestStart(t *testing.T) {
	logger := logrus.StandardLogger()
	buf := new(bytes.Buffer)
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(buf)
	c := New(context.TODO(), logrus.NewEntry(logger), outErrCmdString[0], outErrCmdString[1:]...)
	if err := c.Start(); err != nil {
		t.Fatal(err)
	}
	if err := c.Wait(); err != nil {
		t.Fatal(err)
	}

	// https://github.com/sirupsen/logrus/issues/1111
	time.Sleep(50 * time.Millisecond)

	scanner := bufio.NewScanner(buf)
	for i := 0; i < 2; i++ {

		scanner.Scan()
		var jl jsonLog
		err := json.Unmarshal(scanner.Bytes(), &jl)
		if err != nil {
			t.Fatalf("%d %s", i, err)
		}

		if len(jl.Msg) == 0 || len(jl.Level) == 0 {
			t.Fatal("json log did not parse")
		}

		// order is not enforced, check consistency of level and sg
		if jl.Level == "error" && jl.Msg != "error" {
			t.Errorf("got: %s wanted: error", jl.Msg)
		}

		if jl.Level == "info" && jl.Msg != "out" {
			t.Errorf("got: %s wanted: out", jl.Msg)
		}

		if err := scanner.Err(); err != nil {
			t.Fatal(err)
		}
	}

}

type jsonLog struct {
	Level string
	Msg   string
}

func TestMultiCloser(t *testing.T) {
	c := New(context.TODO(), logrus.NewEntry(logrus.StandardLogger()), outErrCmdString[0], outErrCmdString[1:]...)
	m, err := c.StdoutStderrPipe()
	if err != nil {
		t.Fatal(err)
	}

	if err := c.Cmd.Start(); err != nil {
		t.Fatal(err)
	}

	bs, err := ioutil.ReadAll(m.stdout)
	if err != nil {
		t.Fatal(err)
	}
	if e := "out"; e != string(bytes.TrimSpace(bs)) {
		t.Errorf("got: %s wanted: %s", string(bs), e)
	}
	bs, err = ioutil.ReadAll(m.stderr)
	if err != nil {
		t.Fatal(err)
	}
	if e := "error"; e != string(bytes.TrimSpace(bs)) {
		t.Errorf("got: %s wanted: %s", string(bs), e)
	}
}
