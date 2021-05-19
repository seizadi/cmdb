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
		runStreamingCmd: utils.RunStreamingCmd,
		runCmd:          utils.RunCmd,
		path: viper.GetString("helm.path"),
	}

	return &k, nil
}

func (h *HelmCmd) CreateManifest(app pb.ApplicationInstance) (string, error) {

	// Use appInstance to create the values file in tmp
	//valuesFile := "tmp/values_file"

	// Use app to find the ChartVersion for this Inatnace
	//chartVersion := "chart_repo/app_chart:chart_version"
	helmCmdStr := h.path +
		" template" +
		" infobloxcto/appinfra-grafana-crds --version v0.1.0-46-gf636d12-j5"
//		" --values " + valuesFile + chartVersion

	out, err := h.runCmd(helmCmdStr)
	if err != nil {
		return string(out.Bytes()), err
	}
	artifact := string(out.Bytes())
	return artifact, nil
}
