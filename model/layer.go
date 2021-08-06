package model

import (
	"os"
	"text/template"
)

const pwdEnv = "$PWD"

type Layer struct {
	ModelName string
	TableName string
	Fields    Fields

	// ProjectPath indicates the location of the project
	// is obtained by looking on the $PWD env variable
	ProjectPath     string
	TemplateHelpers template.FuncMap
}

func NewLayerFromConf(conf Config) Layer {
	return Layer{
		ModelName:   conf.ModuleName,
		TableName:   conf.TableName,
		ProjectPath: os.Getenv(pwdEnv),
	}
}
