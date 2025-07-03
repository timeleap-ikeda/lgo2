package env

import "os"
import "fmt"

type Config struct {
	OutputFormat string
}

func ProcessEnvVars() Config {
	env, _ := os.LookupEnv("OUTPUT_FORMAT")
	fmt.Printf("in env.go env=%+v\n", env)
	return Config{
		OutputFormat: env,
	}
}
