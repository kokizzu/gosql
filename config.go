package gosql

type Config struct {
	Enable       bool   `json:"enable"`
	Driver       string `json:"driver"`
	Dsn          string `json:"dsn"`
	MaxOpenConns int    `toml:"max_open_conns" json:"max_open_conns"`
	MaxIdleConns int    `toml:"max_idle_conns" json:"max_idle_conns"`
}