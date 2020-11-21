package config

type Server struct {
	System System `mapstructure:"system" json:"system" yaml:"system"`
	Log    Log    `mapstructure:"log" json:"log" yaml:"log"`
}

type System struct {
	Env  string `mapstructure:"env" json:"env" yaml:"env"`
	Addr int    `mapstructure:"addr" json:"addr" yaml:"addr"`
	Flag string `mapstructure:"flag" json:"flag" yaml:"flag"`
}

type Log struct {
	Prefix  string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`
	LogFile bool   `mapstructure:"log-file" json:"logFile" yaml:"log-file"`
	Stdout  string `mapstructure:"stdout" json:"stdout" yaml:"stdout"`
	File    string `mapstructure:"file" json:"file" yaml:"file"`
}
