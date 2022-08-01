package lang

const Dev = "dev"
const Prod = "prod"
const Test = "test"

var env = ""

// SetEnv 由启动程序设置
func SetEnv(name string) {
	env = name
}

func GetEnv() string {
	if env == "" {
		return Dev
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
