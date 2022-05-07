package sqln

type Named struct {
	sql      string
	args     []any
	compiled bool
}

func NewNamed(sql string) *Named {
	return &Named{sql: sql}
}

func (n *Named) compile() {
	if n.compiled {
		return
	}
}

func (n *Named) Bind(map[string]any) {
	if !n.compiled {
		n.compile()
	}
}

func (n *Named) Value(name string, value any) *Named {
	if !n.compiled {
		n.compile()
	}
	return n
}

func (n *Named) GetSQL() string {
	return ""
}

func (n *Named) GetArgs() []any {
	return nil
}
