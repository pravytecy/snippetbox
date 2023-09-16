package models

import (
	"database/sql"
	"time"
)

type Snippet struct {
	id      int
	title   string
	content string
	created time.Time
	expires time.Time
}

type SnippetModel struct {
	*sql.DB
}

func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	stmt := `INSERT INTO snippets (title, content, created, expires)
			VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`
	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, nil
	}
	return int(id), nil
}

func (*SnippetModel) Get(id int) (*Snippet, error) {
	return nil, nil
}
func (m *SnippetModel) Latest() ([]*Snippet, error) {
	return nil, nil
}
