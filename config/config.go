package config

type Server struct {
	Redis Redis `mapstructure:"redis" json:"redis" yaml:"redis"`
}
