package config

func init() {
	CFG = LoadConfig()
	if CFG == nil {
		panic("fail to load application configuration")
	}
}
