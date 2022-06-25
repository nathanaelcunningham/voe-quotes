package mysql

import (
	"database/sql"
	"nathanaelcunningham/voe-quotes/pkg/models"
)

type QuoteModel struct {
	DB *sql.DB
}

func (m *QuoteModel) Latest() ([]*models.Quote, error) {
	stmt := `SELECT id, quote, author, created_at
  FROM quotes
  ORDER BY created_at DESC LIMIT 10`

	//run query
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	quotes := []*models.Quote{}

	for rows.Next() {
		q := &models.Quote{}

		err = rows.Scan(&q.ID, &q.Quote, &q.Author, &q.Created)
		if err != nil {
			return nil, err
		}

		quotes = append(quotes, q)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return quotes, nil
}
