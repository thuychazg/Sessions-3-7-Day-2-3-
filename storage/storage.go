package storage

import "asset-service/model"

type Storage interface {
	Create(asset model.Asset) error
	BatchCreate([]model.Asset) error
	Delete(id string) bool
	GetAll() []model.Asset
	Count() int
}
