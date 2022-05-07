package collection

import "container/list"

type List struct {
	*list.List
}

func (l *List) Random() any {
	size := l.Len()
	if size == 0 {
		return nil
	}
	return nil
}
