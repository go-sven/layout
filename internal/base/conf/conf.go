package conf

type AppConfig struct {
	Web      *Web      `json:"web" yaml:"web"`
	Database *Database `json:"database" yaml:"database"`
	Redis    *Redis    `json:"redis" yaml:"redis"`
}

type Web struct {
	Mode    string `json:"mode" yaml:"mode"`
	Addr    string `json:"addr" yaml:"addr"`
	Default bool   `json:"default" yaml:"default"`
}

type Database struct {
	Connection string `json:"connection" yaml:"connection"`
}

type Redis struct {
	Host     string `json:"host" yaml:"host"`
	Password string `json:"password" yaml:"password"`
	Db       int    `json:"db" yaml:"db"`
}

func (c *AppConfig) SetDefault() {
	c.Web = &Web{
		Mode:    "debug",
		Addr:    ":8888",
		Default: true,
	}
}

func Load(path string) (c *AppConfig, err error) {
	c = &AppConfig{
		Web:      nil,
		Database: nil,
		Redis:    nil,
	}
	c.SetDefault()
	return c, nil
}
