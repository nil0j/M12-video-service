package repository

import (
	"github.com/nil0j/jirafeitor/models/postgres"

	"context"
)

func CreateVideo(title string, description string) (int, error) {
	var id int
	err := conn.QueryRow(context.Background(), "INSERT INTO jirafeitor.videos (title, description, user_id) VALUES ($1, $2) RETURNING id", title, description).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func GetAllVideos() ([]postgres.Video, error) {
	rows, err := conn.Query(context.Background(), "SELECT id, title, description, user_id FROM jirafeitor.videos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var videos []postgres.Video
	for rows.Next() {
		var video postgres.Video
		if err := rows.Scan(&video.ID, &video.Title, &video.Description, &video.UserID); err != nil {
			return nil, err
		}
		videos = append(videos, video)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return videos, nil
}
