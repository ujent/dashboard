package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/ujent/dashboard/api/data"
)

const dashboardRoute = "/api/dashboard"

type server struct {
	http   *http.Server
	logger *log.Logger
}

func newServer(logger *log.Logger) *server {
	s := server{logger: logger}
	router := http.NewServeMux()

	router.HandleFunc("/api/login", s.handleLogin)
	router.HandleFunc("/api/dashboard/stats", s.handleDashboardStats)
	router.HandleFunc("/api/dashboard/userchart", s.handleDashboardUserChart)
	router.HandleFunc("/api/dashboard/activitychart", s.handleDashboardActivityChart)

	s.http = &http.Server{Handler: s.loggingMiddleware(router.ServeHTTP)}
	return &s
}

func (s *server) handleLogin(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	req := LoginRQ{}
	json.Unmarshal(body, &req)

	//{ "email": "john@smith.com","password":}
	if req.Email == "john@smith.com" && req.Password == "mypassword" {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusForbidden)
	}
}

// api/dashboard/stats/?from=2018-12-01&to=2018-12-31

func (s *server) handleDashboardStats(w http.ResponseWriter, r *http.Request) {

	from, to, err := parsDateRQ(w, r)
	if err != nil {
		return
	}

	cards := data.StatsCards(*from, *to)
	s.writeJSON(w, http.StatusOK, &StatsRS{Stats: cards})
}

func (s *server) handleDashboardUserChart(w http.ResponseWriter, r *http.Request) {

	from, to, err := parsDateRQ(w, r)
	if err != nil {
		return
	}

	userChart := data.UserChart(*from, *to)
	s.writeJSON(w, http.StatusOK, &UserChartRS{Series: userChart})
}

func (s *server) handleDashboardActivityChart(w http.ResponseWriter, r *http.Request) {

	from, to, err := parsDateRQ(w, r)
	if err != nil {
		return
	}

	activityChart := data.ActivityChart(*from, *to)
	s.writeJSON(w, http.StatusOK, &ActivityChartRS{Series: activityChart})
}

func parsDateRQ(w http.ResponseWriter, r *http.Request) (*time.Time, *time.Time, error) {
	if r.Method != http.MethodGet {
		txt := "Method Not Allowed"
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(txt))

		return nil, nil, errors.New(txt)
	}

	q := r.URL.Query()

	from, err := time.Parse(time.RFC3339, q.Get("from"))
	if err != nil {
		txt := "Bad Request: from"
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(txt))

		return nil, nil, errors.New(txt)
	}

	to, err := time.Parse(time.RFC3339, q.Get("to"))
	if err != nil {
		txt := "Bad Request: to"
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(txt))

		return nil, nil, errors.New(txt)
	}

	return &from, &to, nil
}

func (s *server) loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.logger.Printf("%s %s %s", r.Method, r.URL.Path, r.RemoteAddr)
		next(w, r)
	}

}

func (s *server) writeJSON(w http.ResponseWriter, statusCode int, payload interface{}) {

	json, err := json.Marshal(payload)
	if err != nil {
		s.logger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(json)
}
