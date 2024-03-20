package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Status struct {
	Water  int    `json:"water"`
	Wind   int    `json:"wind"`
	Status string `json:"status"`
}

var dataStatus []Status

func UpdateStatus(ctx *gin.Context) {
	var payload struct {
		Status Status `json:"status"`
	}

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	water := payload.Status.Water
	wind := payload.Status.Wind
	sts := getStatus(water, wind)
	status := sts

	fmt.Printf("Water: %d, Wind: %d, Status: %s\n", water, wind, status)

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Status updated successfully: Water=%d, Wind=%d, Status=%s", water, wind, status),
	})
}

func getStatus(water, wind int) string {
	if water < 5 || wind < 6 {
		return "Aman"
	} else if (water >= 6 && water <= 8) || (wind >= 7 && wind <= 15) {
		return "Siaga"
	} else {
		return "Bahaya"
	}
}
