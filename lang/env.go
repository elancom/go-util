package lang

import "os"

const Dev = "dev"
const Prod = "prod"
const Test = "test"

var env = Dev

func SetEnv(env_ string) {
	env = env_
}

func GetEnv() string {
	if env == "" {
		env = os.Getenv("env")
	}
	return env
}

func IsDev() bool {
	return GetEnv() == Dev
}

func IsProd() bool {
	return GetEnv() == Prod
}

func IsTest() bool {
	return GetEnv() == Test
}
