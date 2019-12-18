package types

type Config struct {
	DB mysql `toml:"mysql"`
}


type mysql struct {
	Addr string `toml:"addr"`
	Port int	`toml:"port"`
	UserName string `toml:"username"`
	Password string `toml:"password"`
	Database string `toml:"database"`
}