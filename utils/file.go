package utils

import (
	"bytes"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
)

func WriteTempFile(data string, pattern string) (string, error) {
	tmpDir := "/tmp"
	f, err := ioutil.TempFile(tmpDir, pattern)
	if err != nil {
		return "", err
	}
	defer func() {
		cerr := f.Close()
		if err == nil {
			err = cerr
		}
	}()
	if _, err = f.Write(bytes.NewBufferString(data).Bytes()); err != nil {
		return "", err
	}
	err = f.Sync()
	return f.Name(), nil
}

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
