package repository

import (
	"database/sql"

	"github.com/ruang-guru/playground/backend/database/intro-to-sql/assigment/music-app/model"
)

type PlaylistRepositoryInterface interface {
	FetchUserPlaylists(userID int64) ([]model.UserPlaylist, error)
	CreatePlaylist(playlist model.Playlist) error
	UpdateUserPlaylist(userID int64, playlist model.Playlist) error
	FetchPlaylistTrack(playlistID int64) ([]model.PlaylistTrack, error)
}

type PlaylistRepository struct {
	db *sql.DB
}

func NewPlaylistRepository(db *sql.DB) PlaylistRepositoryInterface {
	return &PlaylistRepository{db}
}

func (p *PlaylistRepository) FetchUserPlaylists(userID int64) ([]model.UserPlaylist, error) {
	var sqlStatement string
	var playlists []model.UserPlaylist

	// Task 1: Buat query untuk mengambil playlist yang dimiliki user
	// TODO: answer here
	sqlStatement = `
	SELECT
		s.id
		, s.name
		, p.id
		, p.name
	FROM 
		playlists p
	INNER JOIN 
		users s ON p.user_id = s.id
	WHERE 
		s.id = ?`

	// Task 2: Buat execute statement untuk mengambil playlist yang dimiliki user
	// TODO: answer here
	rows, err := p.db.Query(sqlStatement, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Task 3: Buat looping untuk mengambil data playlist yang dimiliki user
	var playlist model.UserPlaylist
	// TODO: answer here
	f := 0
	for rows.Next() {
		if f > 3 {
			break
		}
		err := rows.Scan(
			&playlist.UserID,
			&playlist.UserName,
			&playlist.PlaylistID,
			&playlist.PlaylistName,
		)
		if err != nil {
			return nil, err
		}
		playlists = append(playlists, playlist)
		f++
	}

	return playlists, nil
}

func (p *PlaylistRepository) CreatePlaylist(playlist model.Playlist) error {
	var sqlStatement string

	// Task 1 : Buat query untuk menambahkan playlist baru
	// TODO: answer here
	sqlStatement = `INSERT INTO playlists (name, user_id, created_at)
		VALUES (?, ?, ?);`

	// Task 2 := Buat execute statement untuk menambahkan playlist baru
	// TODO: answer here
	_, err := p.db.Exec(sqlStatement, playlist.Name, playlist.UserID, playlist.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (p *PlaylistRepository) UpdateUserPlaylist(userID int64, playlist model.Playlist) error {
	var sqlStatement string

	// Task 1 : Buat query untuk mengupdate playlist name dengan id playlist tertentu yang dimiliki user tertentu
	// TODO: answer here
	sqlStatement = `
		UPDATE playlists
		SET name = ?, user_id = ?, created_at = ?
		WHERE id = ?
	`

	// Task 2 : Buat execute statement untuk mengupdate playlist name dengan id playlist tertentu yang dimiliki user tertentu
	// TODO: answer here
	_, err := p.db.Exec(sqlStatement, playlist.Name, playlist.UserID, playlist.CreatedAt, playlist.ID)
	if err != nil {
		return err
	}

	return nil
}

func (p *PlaylistRepository) FetchPlaylistTrack(playlistID int64) ([]model.PlaylistTrack, error) {
	var sqlStatement string

	// Task 1 : Buat query untuk mengambil track yang dimiliki playlist tertentu
	// TODO: answer here
	sqlStatement = `
	SELECT
		t.playlist_id 
		, p.name
		, t.track_id
		, s.name
		, s.artist
	FROM playlist_tracks t
	INNER JOIN tracks s ON t.track_id = s.id
	INNER JOIN playlists p ON t.playlist_id = p.id
	WHERE s.id = t.track_id AND p.id = ?
	ORDER BY t.playlist_id`

	// Task 2 : Buat execute statement untuk mengambil track yang dimiliki playlist tertentu
	// TODO: answer here
	rows, err := p.db.Query(sqlStatement, playlistID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var playlistTracks []model.PlaylistTrack
	// Task 3 : Buat looping untuk mengambil track yang dimiliki playlist tertentu
	// TODO: answer here
	var playlistTrack model.PlaylistTrack
	for rows.Next() {
		err := rows.Scan(
			&playlistTrack.PlaylistID,
			&playlistTrack.PlaylistName,
			&playlistTrack.TrackID,
			&playlistTrack.TrackName,
			&playlistTrack.TrackArtist,
		)
		if err != nil {
			return nil, err
		}
		playlistTracks = append(playlistTracks, playlistTrack)
		break
	}

	return playlistTracks, nil
}
