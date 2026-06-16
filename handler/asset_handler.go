package handler

import (
	"asset-service/model"
	"asset-service/service"
	"encoding/json"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Handler struct {
	service *service.AssetService

	start time.Time
}

func NewHandler(
	s *service.AssetService,
) *Handler {

	return &Handler{
		service: s,
		start:   time.Now(),
	}

}

func (h *Handler) Create(
	w http.ResponseWriter,
	r *http.Request,
) {

	var req model.CreateAssetRequest

	json.NewDecoder(
		r.Body,
	).Decode(&req)

	a, err := h.service.Create(req)

	if err != nil {

		http.Error(w, err.Error(), 400)
		return
	}

	json.NewEncoder(w).Encode(a)

}

func (h *Handler) BatchCreate(
	w http.ResponseWriter,
	r *http.Request,
) {

	var req model.BatchCreateRequest

	json.NewDecoder(
		r.Body,
	).Decode(&req)

	result, err := h.service.BatchCreate(req)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	json.NewEncoder(w).Encode(result)

}

func (h *Handler) Stats(
	w http.ResponseWriter,
	r *http.Request,
) {

	json.NewEncoder(w).
		Encode(
			h.service.Stats(),
		)

}

func (h *Handler) Count(
	w http.ResponseWriter,
	r *http.Request,
) {

	q := r.URL.Query()

	count := h.service.Count(
		q.Get("type"),
		q.Get("status"),
	)

	json.NewEncoder(w).Encode(
		map[string]interface{}{
			"count": count,
		},
	)
}

func (h *Handler) BatchDelete(
	w http.ResponseWriter,
	r *http.Request,
) {

	ids := strings.Split(
		r.URL.Query().Get("ids"),
		",",
	)

	d, n := h.service.BatchDelete(ids)

	json.NewEncoder(w).Encode(
		map[string]int{
			"deleted":   d,
			"not_found": n,
		},
	)
}

func (h *Handler) Health(
	w http.ResponseWriter,
	r *http.Request,
) {

	json.NewEncoder(w).Encode(
		map[string]interface{}{

			"status": "ok",

			"uptime_seconds": int(time.Since(h.start).Seconds()),
		},
	)
}

func (h *Handler) ListAssets(
	w http.ResponseWriter,
	r *http.Request,
) {

	query := r.URL.Query()

	page, _ := strconv.Atoi(
		query.Get("page"),
	)

	limit, _ := strconv.Atoi(
		query.Get("limit"),
	)

	// default

	if page <= 0 {

		page = 1

	}

	if limit <= 0 {

		limit = 20

	}

	// max 100

	if limit > 100 {

		limit = 100

	}

	t := query.Get("type")

	status := query.Get("status")

	data, total, _ :=
		h.service.ListAssets(
			page,
			limit,
			t,
			status,
		)

	totalPages :=
		int(
			math.Ceil(
				float64(total) /
					float64(limit),
			),
		)

	json.NewEncoder(w).Encode(
		map[string]interface{}{

			"data": data,

			"pagination": map[string]interface{}{

				"page": page,

				"limit": limit,

				"total": total,

				"total_pages": totalPages,
			},
		},
	)

}

func (h *Handler) Search(
	w http.ResponseWriter,
	r *http.Request,
) {

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	q := r.URL.Query().Get(
		"q",
	)

	if q == "" {

		http.Error(
			w,
			"query required",
			http.StatusBadRequest,
		)

		return

	}

	result := h.service.Search(q)

	json.NewEncoder(w).Encode(
		result,
	)

}
