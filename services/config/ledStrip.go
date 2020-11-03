package config

type LedStripConfig struct {
	StripType  string `yaml:"type"`
	LedCount   int    `yaml:"ledCount"`
	Brightness uint8  `yaml:"brightness"`
}
