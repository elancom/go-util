package sqln

import (
	"database/sql"
)

// FindOne
// List
// Count

// ScanStruct
// ScanMap

type DB struct {
	*sql.DB
}

func (d DB) Prepare(query string) (*Stmt, error) {
	prepare, err := d.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	return &Stmt{prepare}, nil
}

type Stmt struct {
	*sql.Stmt
}

func (s *Stmt) Select(inst any, query string, args map[string]any) {
	// 命名处理
	if len(args) > 0 {
		//named := NewNamed(query)
		//named.Value()
	}
}
