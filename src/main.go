package main

import (
	"apex-trading-platform/src/controllers"
	"apex-trading-platform/src/models"
	"apex-trading-platform/src/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	utils.InitConfig()
	http.HandleFunc("/metrics", controllers.GetMetrics)
	http.HandleFunc("/strategies", controllers.GetStrategies)
	log.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
