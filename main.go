package main

import (
	"log"
	"net/http"

	"asset-service/config"
	"asset-service/handler"
	"asset-service/service"
	"asset-service/storage"

	"github.com/joho/godotenv"
)

func CORSMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(
		w http.ResponseWriter,
		r *http.Request,
	) {

		w.Header().Set(
			"Access-Control-Allow-Origin",
			"*",
		)

		w.Header().Set(
			"Access-Control-Allow-Methods",
			"GET, POST, PUT, DELETE, OPTIONS",
		)

		w.Header().Set(
			"Access-Control-Allow-Headers",
			"Content-Type, Authorization",
		)

		if r.Method == http.MethodOptions {

			w.WriteHeader(http.StatusOK)
			return

		}

		next.ServeHTTP(w, r)

	})
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println(".env not found")
	}

	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	assetStorage := storage.NewPostgresStorage(db)

	assetService := service.NewAssetService(assetStorage)

	h := handler.NewHandler(assetService)

	// ===== Assets =====
	http.HandleFunc(
		"/assets",
		func(w http.ResponseWriter, r *http.Request) {

			switch r.Method {

			case http.MethodGet:
				h.ListAssets(w, r)

			case http.MethodPost:
				h.Create(w, r)

			default:
				http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			}

		},
	)

	// ===== Stats =====
	http.HandleFunc(
		"GET /stats",
		h.Stats,
	)

	// ===== Count =====
	http.HandleFunc(
		"GET /count",
		h.Count,
	)

	// ===== Batch delete =====
	http.HandleFunc(
		"DELETE /assets",
		h.BatchDelete,
	)

	// ===== Health =====
	http.HandleFunc(
		"GET /health",
		h.Health,
	)

	// ===== Search =====
	http.HandleFunc(
		"GET /search",
		h.Search,
	)

	// ===== Scan APIs =====
	http.HandleFunc(
		"POST /scan",
		handler.StartScan,
	)

	http.HandleFunc(
		"GET /scan-jobs/{id}",
		handler.GetScanJob,
	)

	http.HandleFunc(
		"GET /scan-jobs/{id}/results",
		handler.GetScanResults,
	)

	log.Println("server :8080")

	err = http.ListenAndServe(
		":8080",
		CORSMiddleware(http.DefaultServeMux),
	)

	if err != nil {
		log.Fatal(err)
	}
}
