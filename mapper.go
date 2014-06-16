package rowmapper

import "database/sql"

// A Mapper maps columns to variables.
type Mapper struct {
	cols map[string]interface{}
}

// New creates and returns a new Mapper.
func New() *Mapper {
	return &Mapper{
		cols: make(map[string]interface{}),
	}
}

// Map maps the column named col to the variable pointed to by dest.
// dest must be a pointer.
func (m *Mapper) Map(col string, dest interface{}) {
	m.cols[col] = dest
}

var null = new(interface{})

// Do performs the mapping with the given sql.Rows. Columns where no
// mapping is defined are ignored.
func (m *Mapper) Do(rows *sql.Rows) error {
	cols, err := rows.Columns()
	if err != nil {
		return err
	}

	var values []interface{}

	for _, col := range cols {
		if dest, exists := m.cols[col]; exists {
			values = append(values, dest)
		} else {
			values = append(values, null)
		}
	}

	return rows.Scan(values...)
}

// DoOne works like Do but processes at most one row. If there are no rows
// sql.ErrNoRows is returned. It works analogous to QueryRow.
func (m *Mapper) DoOne(rows *sql.Rows) error {
	if !rows.Next() {
		return sql.ErrNoRows
	}

	err := m.Do(rows)
	rows.Close()
	return err
}
