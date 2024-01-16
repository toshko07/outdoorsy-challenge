package configs

type DB struct {
	Host     string `required:"true"`
	Port     string `required:"true"`
	Username string `required:"true"`
	Password string `required:"true"`
	Name     string `required:"true"`
}
