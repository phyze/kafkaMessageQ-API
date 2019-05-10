package serverConfig

type serving struct {
	Spec        string        `yaml:"spec"`
	Application []Application `yaml:"application"`
}

type ServeConfig struct {
	Serving serving `yaml:"serving"`
}

type Application struct {
	Spec string `yaml:"spec"`
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}
