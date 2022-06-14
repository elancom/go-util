package lang

import "os"

func IsExist(name string) bool {
	yes, err := IsExists(name)
	if err != nil {
		return false
	}
	return yes
}

func IsExists(name string) (bool, error) {
	_, err := os.Stat(name)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func IsDir(name string) bool {
	yes, err := IsDirs(name)
	if err != nil {
		return false
	}
	return yes
}

func IsDirs(name string) (bool, error) {
	stat, err := os.Stat(name)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return stat.IsDir(), nil
}
