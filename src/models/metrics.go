package models

import (
	"encoding/json"
)

type Metric struct {
	Name string `json:"name"`
	Value float64 `json:"value"`
}

func GetMetrics() []Metric {
	return []Metric{{Name: "CPU", Value: 50.0}, {Name: "Memory", Value: 30.0}}
}
