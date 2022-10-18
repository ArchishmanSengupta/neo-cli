package models

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type Metadata struct {
	ID            int       `db:"id"`
	Title         string    `db:"title"`
	FolderName    string    `db:"folder_name"`
	Url           string    `db:"video_url"`
	Downloaded_At time.Time `db:"downloaded_at"`
}

func (m *Metadata) Insert(url, videoFolderName string, db *sqlx.DB) error {
	m.Downloaded_At = time.Now()

	insertQuery := `INSERT INTO youtube (title, folder_name, video_url, downloaded_at) VALUES (:title, :folder_name, :video_url, :downloaded_at)`

	// Perform insertion operation
	operation := db.MustBegin()
	_, err := operation.NamedExec(insertQuery, &m)
	if err != nil {
		return err
	}
	// Commit the transaction to the database
	_ = operation.Commit()

	return nil
}
