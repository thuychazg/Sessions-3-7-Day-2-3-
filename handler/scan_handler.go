package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"asset-service/model"
	"asset-service/scanner"

	"github.com/google/uuid"
)

var ScanJobs = make(map[string]*model.ScanJob)

type ScanRequest struct {
	ScanType string `json:"scan_type"`
}

func StartScan(w http.ResponseWriter, r *http.Request) {

	var req ScanRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	target := "127.0.0.1"

	var sc scanner.Scanner

	switch req.ScanType {

	case "ip":
		sc = &scanner.IPScanner{}

	case "port":
		sc = &scanner.PortScanner{}

	case "ssl":
		sc = &scanner.SSLScanner{}

	case "tech":
		sc = &scanner.TechScanner{}

	default:
		http.Error(w, "unsupported scan type", 400)
		return
	}

	job := &model.ScanJob{
		ID:        uuid.New().String(),
		ScanType:  req.ScanType,
		Status:    "pending",
		CreatedAt: time.Now(),
	}

	ScanJobs[job.ID] = job

	go func() {

		job.Status = "running"

		log.Printf(
			"scan=%s target=%s",
			req.ScanType,
			target,
		)

		result, err := sc.Scan(target)

		if err != nil {

			job.Status = "failed"
			job.Error = err.Error()

			return
		}

		job.Results = result
		job.Status = "completed"

	}()

	w.WriteHeader(http.StatusAccepted)

	json.NewEncoder(w).Encode(job)
}

func GetScanJob(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("id")

	job, ok := ScanJobs[id]

	if !ok {
		http.Error(w, "not found", 404)
		return
	}

	json.NewEncoder(w).Encode(job)
}

func GetScanResults(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("id")

	job, ok := ScanJobs[id]

	if !ok {

		http.Error(w, "not found", 404)
		return
	}

	json.NewEncoder(w).Encode(
		map[string]interface{}{
			"job_id":    job.ID,
			"scan_type": job.ScanType,
			"results":   job.Results,
		},
	)
}
