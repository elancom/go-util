package lang

type Table[T any] struct {
	List  []T   `json:"list"`
	Total int64 `json:"total"`
}

func NewTable[T any](list []T, total int64) *Table[T] {
	t := new(Table[T])
	t.List = list
	t.Total = total
	return t
}

func NewTableEmpty[T any]() *Table[T] {
	return new(Table[T])
}
