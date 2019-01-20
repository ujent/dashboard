package server

import "github.com/ujent/dashboard/api/data"

// LoginRQ is a request struct for api/login
type LoginRQ struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// StatsRS is a response struct for api/dashboard/stats
type StatsRS struct {
	Stats []*data.StatsCard `json:"stats"`
}

// UserChartRS is a response struct for api/dashboard/userchart
type UserChartRS struct {
	Series [][]int `json:"series"`
}

// ActivityChartRS is a response struct for api/dashboard/activitychart
type ActivityChartRS struct {
	Series [][]int `json:"series"`
}
