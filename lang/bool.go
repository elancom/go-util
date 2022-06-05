package lang

func GetTrue() *bool {
	var b = true
	return &b
}

func GetFalse() *bool {
	var b = false
	return &b
}
