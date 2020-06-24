package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type measurement struct {
	ID string `json:"id"`
	UserID string `json:"userID"`
	Concentration int `json:"concentration"`
	ConcentrationUnit string `json:"concentrationUnit"`
}

type response struct {
	limit int `json:"limit"`
	results []measurement `json:"results"`
	size int `json:"size"`
	start int `json:"start"`
}

func getMeasurements() []measurement {
	byteValue,_:= ioutil.ReadFile("data.json")

	var measurements []measurement
	json.Unmarshal(byteValue, &measurements)

	return measurements;
}

func handlePagination(offSet int, limit int)  response {
	measurements := getMeasurements()
	final := offSet + limit
	size := len(measurements)

	resOn := response{
		limit: limit,
		size: size,
		start: offSet,
	}

	if offSet > size {
		resOn.results = make([]measurement, 0)
		return resOn
	}

	if size < final {
		resOn.results = measurements[offSet:size]
		return resOn
	}

	resOn.results = measurements[offSet:final]
	return resOn
}

func main() {
	measurements := handlePagination(10, 10)

	for _, m := range measurements.results {
		fmt.Println(">> ", m.Concentration)
	}
}