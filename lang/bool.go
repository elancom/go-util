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

// IsTrue no nil && true
func IsTrue(b *bool) bool {
	return b != nil && *b
}

// IsFalse nil or false
func IsFalse(b *bool) bool {
	return !IsTrue(b)
}

func Ifs[T any](b bool, tv, fv T) T {
	if b {
		return tv
	}
	return fv
}

func Ifn[T any](b bool, tv, fv T) T {
	return Ifs[T](!b, tv, fv)
}
