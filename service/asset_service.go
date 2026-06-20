package service

import (
	"errors"
	"github.com/google/uuid"
	"strings"
	"time"

	"asset-service/model"
	"asset-service/storage"
)

type AssetService struct {
	store storage.Storage
}

func NewAssetService(
	s storage.Storage,
) *AssetService {

	return &AssetService{
		store: s,
	}
}

func validateType(t string) bool {

	return t == "domain" ||
		t == "ip" ||
		t == "service"

}

// CREATE

func (s *AssetService) Create(
	req model.CreateAssetRequest,
) (model.Asset, error) {

	if req.Name == "" {
		return model.Asset{}, errors.New("name required")
	}

	if !validateType(req.Type) {
		return model.Asset{}, errors.New("invalid type")
	}

	a := model.Asset{

		ID:        uuid.New().String(),
		Name:      req.Name,
		Type:      req.Type,
		Status:    "active",
		CreatedAt: time.Now(),
	}

	s.store.Create(a)

	return a, nil

}

// BATCH CREATE

func (s *AssetService) BatchCreate(
	req model.BatchCreateRequest,
) ([]model.Asset, error) {

	if len(req.Assets) > 100 {

		return nil, errors.New("max 100 assets")
	}

	temp := make([]model.Asset, 0)

	// validate all

	for _, x := range req.Assets {

		if x.Name == "" ||
			!validateType(x.Type) {

			return nil,
				errors.New("validation failed")
		}

		temp = append(temp, model.Asset{

			ID:        uuid.New().String(),
			Name:      x.Name,
			Type:      x.Type,
			Status:    "active",
			CreatedAt: time.Now(),
		})
	}

	// insert all

	s.store.BatchCreate(temp)

	return temp, nil
}

// Bài 1 Stats

func (s *AssetService) Stats() map[string]interface{} {

	assets := s.store.GetAll()

	result := map[string]interface{}{}

	byType := map[string]int{}
	byStatus := map[string]int{}

	for _, a := range assets {

		byType[a.Type]++

		byStatus[a.Status]++

	}

	result["total"] = len(assets)
	result["by_type"] = byType
	result["by_status"] = byStatus

	return result

}

func (s *AssetService) Count(
	t string,
	status string,
) int {

	count := 0

	for _, a := range s.store.GetAll() {

		if t != "" &&
			a.Type != t {
			continue
		}

		if status != "" &&
			a.Status != status {
			continue
		}

		count++

	}

	return count

}

// Bài 3 Delete

func (s *AssetService) BatchDelete(
	ids []string,
) (int, int) {

	deleted := 0
	notFound := 0

	for _, id := range ids {

		if s.store.Delete(id) {

			deleted++

		} else {

			notFound++

		}

	}

	return deleted, notFound
}

// Search

func (s *AssetService) Search(
	q string,
) []model.Asset {

	result := []model.Asset{}

	q = strings.ToLower(q)

	for _, a := range s.store.GetAll() {

		if strings.Contains(
			strings.ToLower(a.Name),
			q,
		) {

			result = append(result, a)
		}

		if len(result) >= 100 {
			break
		}

	}

	return result

}

func (s *AssetService) TotalAssets() int {

	return len(
		s.store.GetAll(),
	)

}

func (s *AssetService) ListAssets(
	page int,
	limit int,
	assetType string,
	status string,
) ([]model.Asset, int, int) {

	assets := s.store.GetAll()

	// FILTER

	filtered := make(
		[]model.Asset,
		0,
	)

	for _, asset := range assets {

		if assetType != "" &&
			asset.Type != assetType {

			continue
		}

		if status != "" &&
			asset.Status != status {

			continue
		}

		filtered = append(
			filtered,
			asset,
		)

	}

	total := len(filtered)

	// PAGINATION

	start := (page - 1) * limit

	end := start + limit

	if start > total {

		return []model.Asset{}, total, 0

	}

	if end > total {

		end = total

	}

	return filtered[start:end],
		total,
		end - start
}
