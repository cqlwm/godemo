package video

import (
	"tryon01/net/once/dao/video"
)

func QueryAll() (*[]DTO, error) {
	resultDB, err := video.QueryAll()
	if err != nil {
		return nil, err
	}
	videos := make([]DTO, 0)
	for _, v := range *resultDB {
		videos = append(videos, DTO{v.Id, v.Name})
	}
	return &videos, nil
}