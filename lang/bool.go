package lang

func GetTrue() *bool {
	var b = true
	return &b
}

func GetFalse() *bool {
	var b = false
	return &b
}

func GetBool(b bool) *bool {
	var v = false
	if b {
		v = true
	}
	return &v
}
