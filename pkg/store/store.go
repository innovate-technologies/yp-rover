package store

import (
	"database/sql"
	"strings"

	"github.com/innovate-technologies/yp-rover/pkg/shoutcastcom"

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
				Up:   []string{"CREATE TABLE `rover_sc_genres` (`id` INT NOT NULL AUTO_INCREMENT, `name` NOT NULL, UNIQUE KEY `id` (`id`) USING BTREE) ENGINE=InnoDB;)"},
				Down: []string{"DROP TABLE `rover_sc_genres`"},
			},
			&migrate.Migration{
				Id:   "2",
				Up:   []string{"CREATE TABLE `rover_sc_stations` ( `id` INT NOT NULL AUTO_INCREMENT , `shoutcast_id` INT NOT NULL , `station_name` TEXT NOT NULL , `media_type` TEXT NULL , `bitrate` TEXT NULL , `genre` TEXT NULL , `genre_1` TEXT NULL , `genre_2` TEXT NULL , `genre_3` TEXT NULL , `genre_4` TEXT NULL , `genre_5` TEXT NULL , `logo_url` TEXT NOT NULL , `listen_urls` TEXT NOT NULL , PRIMARY KEY (`id`), INDEX `shoutcast_id` (`shoutcast_id`)) ENGINE = InnoDB;"},
				Down: []string{"DROP TABLE `rover_sc_stations`"},
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

// AddSHOUTcastGenre adds a roved SHOUTcast.com genre
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

// AddSHOUTcastStation adds a roved SHOUTcast.com station
func (s *Store) AddSHOUTcastStation(station shoutcastcom.Station) error {
	row := s.db.QueryRow("SELECT COUNT(*) FROM rover_sc_genres WHERE shoutcast_id=?", station.ID)

	var count int
	err := row.Scan(&count)
	if err != nil {
		return err
	}

	if count > 0 {
		return nil // we don't care about up to date info at the moment
	}

	_, err = s.db.Exec("INSERT INTO `rover_sc_genres` (shoutcast_id, station_name, media_type, bitrate, genre, genre_1, genre_2, genre_3, genre_4, genre_5, logo_url, listen_urls) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? )", station.ID, station.Name, station.MediaType, station.BitRate, station.Genre, station.Genre2, station.Genre3, station.Genre4, station.Genre5, station.LogoURL, strings.Join(station.ListenURLs, ","))

	return err
}
