package repository

import (
	"github.com/nil0j/jirafeitor/models/postgres"

	"context"
)

func CreateVideo(title, description string) (int, error) {
	var id int
	err := conn.QueryRow(context.Background(), "insert into videos (name, description) VALUES ($1, $2) RETURNING id", title, description).Scan(&id)
	return id, err
}

func GetVideoInfo(id int) (postgres.Video, error) {
	var info postgres.Video
	err := conn.QueryRow(context.Background(), "select (id, name, description) from videos where id = $1", id).Scan(&info.ID, &info.Name, &info.Description)
	if err != nil {
		return postgres.Video{}, err
	}

	return info, nil
}

func GetRecentVideos(limit int) ([]postgres.Video, error) {
	rows, err := conn.Query(
		context.Background(),
		"SELECT id, title, description, user_id FROM videos ORDER BY user_id DESC LIMIT $1",
		limit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var videos []postgres.Video
	for rows.Next() {
		var video postgres.Video
		if err := rows.Scan(&video.ID, &video.Name, &video.Description, &video.UserID); err != nil {
			return nil, err
		}
		videos = append(videos, video)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return videos, nil
}
