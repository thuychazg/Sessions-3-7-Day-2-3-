package main

import (
	"net/http"

	"asset-service/handler"
	"asset-service/service"
	"asset-service/storage"
)

func main() {

	store := storage.NewMemoryStorage()

	svc := service.NewAssetService(store)

	h := handler.NewHandler(svc)

	rl := middleware.NewRateLimiter()

	http.HandleFunc(
		"/assets",
		rl.Middleware(
			h.ListAssets,
		),
		func(
			w http.ResponseWriter,
			r *http.Request,
		) {

			if r.Method == "GET" {

				h.ListAssets(w, r)

			} else if r.Method == "POST" {

				h.Create(w, r)

			}

		},
	)

	http.HandleFunc(
		"/assets/batch",
		func(w http.ResponseWriter, r *http.Request) {

			if r.Method == "POST" {

				h.BatchCreate(w, r)

			} else {

				h.BatchDelete(w, r)
			}
		},
	)

	http.HandleFunc(
		"/assets/stats",
		h.Stats,
	)

	http.HandleFunc(
		"/assets/count",
		h.Count,
	)

	http.HandleFunc(
		"/health",
		h.Health,
	)
	http.HandleFunc(
		"/assets/search",
		h.Search,
	)
	http.ListenAndServe(
		":8080",
		nil,
	)

}
