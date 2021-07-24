package video

import (
	"tryon01/net/once/config/datasource"
	error2 "github.com/pkg/errors"
)

const queryAll = "SELECT v.id Id, v.name Name FROM look.video1 v LIMIT 10"

func QueryAll() (*[]DB, error) {
	videos := make([]DB, 0)
	rows, err := datasource.Instance.Query(queryAll)
	defer rows.Close()
	if err != nil {
		return nil, error2.WithMessage(err, "###########################")
	}
	for rows.Next() {
		var video DB
		if err := rows.Scan(&video.Id, &video.Name); err != nil {
			return nil, error2.Wrap(err, "rows.Next()")
		}
		videos = append(videos, video)
	}
	return &videos, nil
}
