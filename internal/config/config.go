package config

type AppConfig struct {
	CurrentContext 	string 				`mapstructure:"current-context"`
	Contexts 		map[string]Context	`mapstructure:"contexts"`
}

type Context struct {
	Tailnet  string	`mapstructure:"tailnet"`
}