package controllers

import (
	"assignment-3/helper"
	"assignment-3/model"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func GetStatus(ctx *gin.Context) {
	helper.UpdateJSON()

	jsonData, err := os.Open("data.json")

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)

		return
	}

	bytes, err := ioutil.ReadAll(jsonData)

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)

		return
	}

	var data model.Status

	json.Unmarshal(bytes, &data)

	waterStatus, windStatus := helper.GetWaterStatus(data.Status.Water), helper.GetWindStatus(data.Status.Wind)
	waterClass, winedClass := helper.GetWaterClass(data.Status.Water), helper.GetWindClass(data.Status.Wind)

	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"water":       data.Status.Water,
		"waterStatus": waterStatus,
		"waterClass":  waterClass,
		"wind":        data.Status.Wind,
		"windStatus":  windStatus,
		"windClass":   winedClass,
	})
}
