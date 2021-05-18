package utils

import (
	"bytes"
	"context"
	"fmt"
	"github.com/spf13/viper"
	"io"
	"os"
	"os/exec"
	"sync"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func CopyBufferContentsToFile(srcBuff []byte, destFile string) (err error) {
	out, err := os.Create(destFile)
	if err != nil {
		return
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()
	if _, err = out.Write(srcBuff); err != nil {
		return
	}
	err = out.Sync()
	return
}

// Copies source buffer to the tmp directory
// Insures tmp directory exists and prepends path
// Temporary, used for writing Kops Manifest to file but exploring using STDIN instead
func CopyBufferContentsToTempFile(srcBuff []byte, destFile string) (err error) {
	var mode os.FileMode = 509
	err = os.MkdirAll("." + viper.GetString("kops.kube.dir"), mode)
	if err != nil {
		return err
	}

	out, err := os.Create("." + viper.GetString("kops.kube.dir") + "/" + destFile)
	if err != nil {
		return
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()
	if _, err = out.Write(srcBuff); err != nil {
		return
	}
	err = out.Sync()
	return
}

func RunCmd(cmdString string) (*bytes.Buffer, error) {
	var out bytes.Buffer

	cmd := exec.Command("echo", cmdString)
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return nil, err
	}

	var mode os.FileMode = 509
	err = os.MkdirAll("./tmp", mode)
	if err != nil {
		return nil, err
	}

	err = CopyBufferContentsToFile(out.Bytes(), "./tmp/cmd.sh")
	if err != nil {
		return nil, err
	}

	out.Reset()
	cmd = exec.Command("/bin/sh", "./tmp/cmd.sh")
	cmd.Stdout = &out
	var errout bytes.Buffer
	cmd.Stderr = &errout
	err = cmd.Run()
	if err != nil {
		CopyBufferContentsToFile(errout.Bytes(), "./tmp/error.txt")
		return &errout, err
	}

	return &out, nil
}

func RunStreamingCmd(cmdString string) error {
	var out bytes.Buffer

	cmd := exec.Command("echo", cmdString)
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return err
	}

	var mode os.FileMode = 509
	err = os.MkdirAll("./tmp", mode)
	if err != nil {
		return err
	}

	err = CopyBufferContentsToFile(out.Bytes(), "./tmp/cmd.sh")
	if err != nil {
		return err
	}

	out.Reset()
	command := New(context.TODO(), nil, "/bin/sh", "./tmp/cmd.sh")

	if err := command.Start(); err != nil {
		return err
	}
	if err := command.Wait(); err != nil {
		return err
	}

	return nil
}

func RunDockerCmd(dockerArgs []string) (*bytes.Buffer, error) {
	var out bytes.Buffer

	//cmd := exec.Command(cmdString)
	cmd := exec.Command("/usr/local/bin/docker", dockerArgs...)
	cmd.Stdout = &out
	var errout bytes.Buffer
	cmd.Stderr = &errout
	err := cmd.Run()
	if err != nil {
		CopyBufferContentsToFile(errout.Bytes(), "./tmp/error.txt")
		return &errout, err
	}

	return &out, nil
}

type Cmd struct {
	*exec.Cmd
	entry     *logrus.Entry
	cmdString []string
	m         *multiCloser
}

func New(ctx context.Context, entry *logrus.Entry, command string, arg ...string) *Cmd {
	if entry == nil {
		entry = defaultEntry
	}
	return &Cmd{
		Cmd:       exec.Command(command, arg...),
		cmdString: append([]string{command}, arg...),
		entry:     entry,
	}
}

func (c *Cmd) GetCmdString() []string {
	return c.cmdString
}

func (c *Cmd) Start() error {
	_, err := c.startAndPipe()
	return err
}

func (c *Cmd) error(err error) error {
	return fmt.Errorf("cmd: %q err: %s", c.cmdString, err)
}

func (c *Cmd) Wait() error {
	// Drain stderr first, probably used less
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		pipeToLog(c.entry.WriterLevel(logrus.ErrorLevel), c.m.stderr)
	}()
	go func() {
		defer wg.Done()
		pipeToLog(c.entry.WriterLevel(logrus.InfoLevel), c.m.stdout)
	}()
	wg.Wait()
	return c.Cmd.Wait()
}

func pipeToLog(pipe *io.PipeWriter, reader io.ReadCloser) (int, error) {
	for {
		n64, err := io.Copy(pipe, reader)
		n := int(n64)
		if err != nil {
			return n, err
		}
		if n == 0 {
			pipe.Close()
			reader.Close()
			return n, io.EOF
		}
	}
}

func (c *Cmd) startAndPipe() (*multiCloser, error) {
	m, err := c.StdoutStderrPipe()
	if err != nil {
		return nil, err
	}
	c.m = m

	return m, c.Cmd.Start()
}

func (c *Cmd) StdoutStderrPipe() (*multiCloser, error) {
	stdout, err := c.Cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	stderr, err := c.Cmd.StderrPipe()
	if err != nil {
		return nil, err
	}
	// TODO: close underlying pipes
	return &multiCloser{stdout, stderr}, nil
}

type multiCloser struct {
	stdout, stderr io.ReadCloser
}

func (m *multiCloser) Close() error {
	var err error
	if errOut := m.stdout.Close(); errOut != nil {
		err = errors.Wrap(err, errOut.Error())
	}
	errErr := m.stderr.Close()
	if errErr != nil {
		err = errors.Wrap(err, errErr.Error())
	}

	return err
}
