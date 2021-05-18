package helm

import (
	"bytes"
	"strings"
	"testing"
)

type testCase struct {
	value string
	found bool
}

var cmd string

func mockRunStreamingCmd(cmdString string) error {
	cmd = cmdString
	return nil
}

func mockRunCmd(cmdString string) (*bytes.Buffer, error) {
	cmd = cmdString
	return nil, nil
}

func TestCreateManifest(t *testing.T) {
	k, err := NewHelm()
	if err != nil {
		t.Error("Expected no error got", err)
		return
	}

	k.runStreamingCmd = mockRunStreamingCmd
	k.runCmd = mockRunCmd

	values := []testCase{
		// Add values for testing the command

	}


	cmdValues := strings.Split(cmd, " ")
	for _, c := range cmdValues {
		for i, v := range values {
			if v.value == c {
				values[i].found = true
				break
			}
		}
	}

	for _, v := range values {
		if v.found == false {
			t.Error("Expected ", v.value, "not found")
		}
	}
}
