package service

import (
	"admin_panel/models"
	"admin_panel/pkg/repository"
)

func CreateSegment(segment models.Segment) error {

	return repository.CreateSegment(segment)

}
