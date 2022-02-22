package types

type Env struct {
	Id  string
	Url string
}

type Config struct {
	Envs []Env
}

func GetConf() Config {
	// TODO handle default vs local
	baseEnv := Env{Id: "kube", Url: "https://pogues-back-office.dev.insee.io/swagger-ui/dist/"}
	defaultConf := Config{Envs: []Env{baseEnv}}
	return defaultConf
}
