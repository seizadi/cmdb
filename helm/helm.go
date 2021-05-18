package helm

import (
	"bytes"
	"github.com/seizadi/cmdb/pkg/pb"
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
		devMode:         viper.GetBool("development"),
		runStreamingCmd: utils.RunStreamingCmd,
		runCmd:          utils.RunCmd,
		path: viper.GetString("helm.path"),
	}

	return &k, nil
}

func (h *HelmCmd) CreateManifest(app pb.ApplicationInstance) (string, error) {

	// Use appInstance to create the values file in tmp
	valuesFile := "tmp/values_file"

	// Use app to find the ChartVersion for this Inatnace
	chartVersion := "chart_repo/app_chart:chart_version"
	kopsCmdStr := h.path +
		" template" +
		" --values " + valuesFile + chartVersion

	out, err := h.runCmd(kopsCmdStr)
	if err != nil {
		return nil, err
	}
	artifact := string(out.Bytes())
	return artifact, nil
}
