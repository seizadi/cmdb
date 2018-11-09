package model

type Artifact struct {
	Model
	Name string
	Description string
	Repo string
	CommitId string
	VersionTag VersionTag
}