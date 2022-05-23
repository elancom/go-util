package collection

func Clone[K string | int, V any](d map[K]V) map[K]V {
	nd := make(map[K]V, len(d))
	for k, v := range d {
		nd[k] = v
	}
	return nd
}

func Keys[K string | int, V any](d map[K]V) []K {
	keys := make([]K, 0, len(d))
	for k := range d {
		keys = append(keys, k)
	}
	return keys
}
