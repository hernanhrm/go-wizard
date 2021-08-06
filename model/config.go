package model

type Config struct {
	ModuleName string   `yaml:"module_name"`
	TableName  string   `yaml:"table_name"`
	Fields     []string `yaml:"fields"`
	Layers     []string `yaml:"layers"`
}
