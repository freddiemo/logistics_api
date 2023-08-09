package config

func Init() map[string]string {
	envs := GetEnvVariables()

	return envs
}
