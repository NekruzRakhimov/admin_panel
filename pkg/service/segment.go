package service

import (
	"admin_panel/models"
	"admin_panel/pkg/repository"
	"encoding/json"
)

func CreateSegment(segment models.Segment) error {

	product, err := json.Marshal(segment.Products)
	if err != nil {
		return err
	}
	segment.ProductStr = string(product)

	region, err := json.Marshal(segment.Region)
	if err != nil {
		return err
	}
	segment.RegionStr = string(region)

	return repository.CreateSegment(segment)

}
