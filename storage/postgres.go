package storage

import (
	"database/sql"

	"asset-service/model"
)

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage(db *sql.DB) *PostgresStorage {
	return &PostgresStorage{
		db: db,
	}
}

// Create
func (s *PostgresStorage) Create(asset model.Asset) error {

	query := `
	INSERT INTO assets(id,name,type,status,created_at)
	VALUES($1,$2,$3,$4,$5)
	`

	_, err := s.db.Exec(
		query,
		asset.ID,
		asset.Name,
		asset.Type,
		asset.Status,
		asset.CreatedAt,
	)

	return err
}

// BatchCreate
func (s *PostgresStorage) BatchCreate(
	assets []model.Asset,
) error {

	for _, a := range assets {

		err := s.Create(a)
		if err != nil {
			return err
		}

	}

	return nil
}

// Delete
func (s *PostgresStorage) Delete(
	id string,
) bool {

	result, err := s.db.Exec(
		"DELETE FROM assets WHERE id=$1",
		id,
	)

	if err != nil {
		return false
	}

	rows, _ := result.RowsAffected()

	return rows > 0
}

// GetAll
func (s *PostgresStorage) GetAll() []model.Asset {

	rows, err := s.db.Query(`
	SELECT id,name,type,status,created_at
	FROM assets
	`)

	if err != nil {
		return []model.Asset{}
	}

	defer rows.Close()

	var assets []model.Asset

	for rows.Next() {

		var asset model.Asset

		err := rows.Scan(
			&asset.ID,
			&asset.Name,
			&asset.Type,
			&asset.Status,
			&asset.CreatedAt,
		)

		if err != nil {
			continue
		}

		assets = append(assets, asset)
	}

	return assets
}

// Count
func (s *PostgresStorage) Count() int {

	var count int

	s.db.QueryRow(
		"SELECT COUNT(*) FROM assets",
	).Scan(&count)

	return count
}
