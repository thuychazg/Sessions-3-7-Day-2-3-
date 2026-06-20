package service


import (
	"testing"

	"asset-service/model"
	"asset-service/storage"

	"github.com/stretchr/testify/assert"
)


func TestCreateAsset(t *testing.T){

	store := storage.NewMemoryStorage()

	s := NewAssetService(store)


	asset, err := s.Create(
		model.CreateAssetRequest{
			Name:"google.com",
			Type:"domain",
		},
	)


	assert.NoError(t,err)

	assert.Equal(
		t,
		"google.com",
		asset.Name,
	)

}


func TestCreateInvalidType(t *testing.T){

	store := storage.NewMemoryStorage()

	s := NewAssetService(store)


	_,err :=
		s.Create(
			model.CreateAssetRequest{
				Name:"abc",
				Type:"wrong",
			},
		)


	assert.Error(t,err)

}
