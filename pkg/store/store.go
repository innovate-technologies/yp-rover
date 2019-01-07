package store

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // import MySQL driver
	migrate "github.com/rubenv/sql-migrate"
)

// Store allows interaction with the discover.fm database
type Store struct {
	db *sql.DB
}

// New sets up a connection to the discover.fm database
func New(dbpath string) (*Store, error) {
	db, err := sql.Open("mysql", dbpath)
	if err != nil {
		return nil, err
	}

	return &Store{
		db: db,
	}, nil
}

// Migrate perfoms a schema creation/migration
func (s *Store) Migrate() (int, error) {
	migrations := &migrate.MemoryMigrationSource{
		Migrations: []*migrate.Migration{
			&migrate.Migration{
				Id:   "1",
				Up:   []string{"CREATE TABLE `rover_sc_genres` ( `id` INT NOT NULL AUTO_INCREMENT , `name` TEXT NOT NULL , PRIMARY KEY (`id`)) ENGINE = InnoDB"},
				Down: []string{"DROP TABLE `rover_sc_genres`"},
			},
		},
	}

	return migrate.Exec(s.db, "mysql", migrations, migrate.Up)
}

// GetSHOUTcastGenres gets the roved SHOUTcast.com genres
func (s *Store) GetSHOUTcastGenres() ([]string, error) {
	rows, err := s.db.Query("SELECT name FROM rover_sc_genres")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	genres := []string{}
	for rows.Next() {
		var name string
		rows.Scan(&name)
		genres = append(genres, name)

	}

	return genres, nil
}

// AddGenre adds a roved SHOUTcast.com genre
func (s *Store) AddSHOUTcastGenre(genre string) error {
	row := s.db.QueryRow("SELECT COUNT(*) FROM rover_sc_genres WHERE name=?", genre)

	var count int
	err := row.Scan(&count)
	if err != nil {
		return err
	}

	if count > 0 {
		return nil // already exists
	}

	_, err = s.db.Exec("INSERT INTO `rover_sc_genres` (name) VALUES ( ? )", genre)

	return err
}
