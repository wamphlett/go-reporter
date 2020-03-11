package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	Consumer *Consumer
}

func StartServer(consumer *Consumer) {
	s := Server{Consumer: consumer}

	r := mux.NewRouter()

	r.HandleFunc("/report", s.NewReportHandler)
	r.HandleFunc("/stats/{id}", s.StatsHandler)
	r.HandleFunc("/reports/{id}", s.ReportsHandler)

	http.Handle("/", r)
	if err := http.ListenAndServe(":1111", r); err != nil {
		panic(err)
	}
}

func (s *Server) NewReportHandler(w http.ResponseWriter, r *http.Request) {
	var report Report

	err := json.NewDecoder(r.Body).Decode(&report)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	s.Consumer.consumeReport(&report)
}

func (s *Server) StatsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var stats *Stats
	if val, ok := s.Consumer.stats[vars["id"]]; ok {
		stats = val
	} else {
		stats = &Stats{}
	}

	js, err := json.Marshal(stats)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (s *Server) ReportsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	reports := []*Report{}
	if val, ok := s.Consumer.Reports[vars["id"]]; ok {
		reports = val
	}

	js, err := json.Marshal(reports)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
