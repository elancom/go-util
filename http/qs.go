package http

import "net/url"

func ParseQuery(qs string) map[string]string {
	m := make(map[string]string, 0)
	query, err := url.ParseQuery(qs)
	if err != nil {
		return m
	}
	for k, v := range query {
		if len(v) > 0 {
			m[k] = v[0]
		}
	}
	return m
}
