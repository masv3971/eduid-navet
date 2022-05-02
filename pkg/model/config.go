package model

type Cfg struct {
	Production bool `yaml:"production"`
	APIServer  struct {
		Addr string `yaml:"addr"`
	} `yaml:"api_server"`
	OrganizeNumber string `yaml:"organize_number"`
	OrderIdentity  string `yaml:"order_identity"`
}

// Config represent the complete config file structure
type Config struct {
	Navet Cfg `yaml:"navet"` // navet-service ??
}
