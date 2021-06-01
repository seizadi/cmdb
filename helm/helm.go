package helm

import (
	"bytes"
	"github.com/seizadi/cmdb/utils"
	"github.com/spf13/viper"
)

type HelmCmd struct {
	devMode         bool
	runStreamingCmd func(string) error
	runCmd          func(string) (*bytes.Buffer, error)
	path string
}

func NewHelm() (*HelmCmd, error) {
	k := HelmCmd{
		runStreamingCmd: utils.RunStreamingCmd,
		runCmd:          utils.RunCmd,
		path: viper.GetString("helm.path"),
	}

	return &k, nil
}

func (h *HelmCmd) CreateManifest(repo string, version string) (string, error) {

	// Use appInstance to create the values file in tmp
	//valuesFile := "tmp/values_file"

	// Use app to find the ChartVersion for this Inatnace
	//chartVersion := "chart_repo/app_chart:chart_version"
	helmCmdStr := h.path +
		" template " + repo + " --version " + version
//		" --values " + valuesFile + chartVersion

	out, err := h.runCmd(helmCmdStr)
	if err != nil {
		return string(out.Bytes()), err
	}
	artifact := string(out.Bytes())
	return artifact, nil
}

func (h *HelmCmd) CreateValues() (string, error) {
	helmCmdStr := h.path + " template ./render"

	out, err := h.runCmd(helmCmdStr)
	if err != nil {
		return string(out.Bytes()), err
	}
	values := string(out.Bytes())
	return values, nil
}
