package config

type Server struct {
	Redis    Redis  `mapstructure:"redis" json:"redis" yaml:"redis"`
	FontFile string `mapstructure:"font_file" json:"font_file" yaml:"font_file"`
}
