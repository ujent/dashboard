package data

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

type StatsCard struct {
	Type       string `json:"type"`
	Icon       string `json:"icon"`
	Title      string `json:"title"`
	Value      string `json:"value"`
	FooterText string `json:"footerText"`
	FooterIcon string `json:"footerIcon"`
}

func StatsCards(from time.Time, to time.Time) []*StatsCard {
	log.Printf("StatsCards, from = %s, to = %s", from, to)

	var res []*StatsCard
	res = append(res, &StatsCard{
		Type:       "warning",
		Icon:       "ti-server",
		Title:      "Capacity",
		Value:      fmt.Sprintf("%dGB", rand.Int31n(200)),
		FooterText: "Updated now",
		FooterIcon: "ti-reload",
	})

	res = append(res, &StatsCard{
		Type:       "success",
		Icon:       "ti-wallet",
		Title:      "Revenue",
		Value:      fmt.Sprintf("$%d", rand.Int31n(2000)),
		FooterText: "Last day",
		FooterIcon: "ti-calendar",
	})

	res = append(res, &StatsCard{
		Type:       "danger",
		Icon:       "ti-pulse",
		Title:      "Errors",
		Value:      fmt.Sprintf("%d", rand.Int31n(50)),
		FooterText: "In the last hour",
		FooterIcon: "ti-timer",
	})

	res = append(res, &StatsCard{
		Type:       "info",
		Icon:       "ti-twitter-alt",
		Title:      "Followers",
		Value:      fmt.Sprintf("+%d", rand.Int31n(50)),
		FooterText: "Updated now",
		FooterIcon: "ti-reload",
	})

	return res
}

func UserChart(from time.Time, to time.Time) [][]int {
	log.Printf("UserChart, from = %s, to = %s", from, to)

	var res [][]int

	for j := 0; j < 4; j++ {
		res = append(res, nil)
		for i := 0; i < 10; i++ {
			res[j] = append(res[j], int(rand.Int31n(500)))
		}
	}

	return res
}

func ActivityChart(from time.Time, to time.Time) [][]int {
	log.Printf("ActivityChart, from = %s, to = %s", from, to)

	var res [][]int

	for j := 0; j < 3; j++ {
		res = append(res, nil)
		for i := 0; i < 13; i++ {
			res[j] = append(res[j], int(rand.Int31n(500)))
		}
	}

	return res
}
