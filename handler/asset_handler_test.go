package handler


import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"


	"asset-service/service"
	"asset-service/storage"


	"github.com/stretchr/testify/assert"
)



func TestCreateAssetHandler(t *testing.T){


	store :=
		storage.NewMemoryStorage()


	svc :=
		service.NewAssetService(store)


	h :=
		NewHandler(svc)



	req :=
		httptest.NewRequest(
			"POST",
			"/assets",
			strings.NewReader(
				`{"name":"test.com","type":"domain"}`,
			),
		)


	req.Header.Set(
		"Content-Type",
		"application/json",
	)


	rr :=
		httptest.NewRecorder()


	h.Create(rr,req)


	assert.Equal(
		t,
		http.StatusOK,
		rr.Code,
	)

}
