package app

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type measurement struct {
	ID                string `json:"id"`
	UserID            string `json:"userID"`
	Concentration     int    `json:"concentration"`
	ConcentrationUnit string `json:"concentrationUnit"`
}

type response struct {
	Limit   int           `json:"limit"`
	Results []measurement `json:"results"`
	Size    int           `json:"size"`
	Start   int           `json:"start"`
}

func getMeasurements() []measurement {
	contents, err := ioutil.ReadFile("app/data.json")

	if err != nil {
		log.Fatal(err)
	}

	var measurements []measurement
	json.Unmarshal(contents, &measurements)

	return measurements
}

func handlePagination(offSet int, limit int) response {
	measurements := getMeasurements()
	final := offSet + limit
	size := len(measurements)

	resOn := response{
		Limit: limit,
		Size:  size,
		Start: offSet,
	}

	if offSet > size {
		resOn.Results = make([]measurement, 0)
		return resOn
	}

	if size < final {
		resOn.Results = measurements[offSet:size]
		return resOn
	}

	resOn.Results = measurements[offSet:final]
	return resOn
}

// GetMeasurementsJSON get JSON info about Measurements
func GetMeasurementsJSON(offSet int, limit int) []byte {
	measurements := handlePagination(offSet, limit)
	json, _ := json.Marshal(measurements)
	return json
}
