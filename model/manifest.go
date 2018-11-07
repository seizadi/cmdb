package model

type ManifestValues struct {
	Values map[string]string `json:"values"`
}

type ManifestPort struct {
	Name       string `json:"name"`
	Port       int    `json:"port"`
	TargetPort string `json:"targetPort"`
	NodePort   int    `json:"nodePort"`
	Protocol   string `json:"protocol"`
}

type ManifestService struct {
	Name        string         `json:"name"`
	ServiceName string         `json:"serviceName"`
	Type        string         `json:"type"`
	Ports       []ManifestPort `json:"ports"`
}


type MainfestIngressTls struct {
	SecretName string   `json:"secretName"`
	Hosts      []string `json:"hosts"`
}

type MainfestIngressRule struct {
	Host       string `json:"host"`
	SecretName string `json:"secretName"`
}

type ManifestIngress struct {
	Enabled     bool                 `json:"enabled"`
	Annotations map[string]string    `json:"annotations"`
	Tls         []MainfestIngressTls `json:"tls"`
	Hosts       []string             `json:"hosts"`
	Path        string               `json:"path"`
}

type Manifest struct {
	Model
	Name         string
	Description  string
	Repo         string
	CommitId     string
	Values       string
	Services     string
	Ingress      string
	AwsService   AwsService `gorm:"foreignkey:AwsServiceID"`
	AwsServiceID uint
	Artifact     Artifact `gorm:"foreignkey:ArtifactID"`
	ArtifactID   uint
	Vault        Vault `gorm:"foreignkey:VaultID"`
	VaultID      uint
}

