// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/newrelic/newrelic-telemetry-sdk-go/telemetry"
)

type Operands struct {
	OperandOne float32 `json:"operandOne,string"`
	OperandTwo float32 `json:"operandTwo,string"`
}

func add(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var operands Operands
	json.NewDecoder(r.Body).Decode(&operands)
	fmt.Println(fmt.Sprintf("%s%f%s%f", "Adding ", operands.OperandOne, " to ", operands.OperandTwo))
	json.NewEncoder(w).Encode(operands.OperandOne + operands.OperandTwo)
}

func main() {
	h, err := telemetry.NewHarvester(telemetry.ConfigAPIKey(os.Getenv("NEW_RELIC_INSIGHTS_INSERT_API_KEY")))
	if err != nil {
		fmt.Println(err)
	}

	// Record Gauge, Count, and Summary metrics using RecordMetric. These
	// metrics are not aggregated.  This is useful for exporting metrics
	// recorded by another system.
	h.RecordMetric(telemetry.Gauge{
		Timestamp: time.Now(),
		Value:     1,
		Name:      "myMetric",
		Attributes: map[string]interface{}{
			"color": "purple",
		},
	})

	// Record spans using RecordSpan.
	h.RecordSpan(telemetry.Span{
		ID:          "12345",
		TraceID:     "67890",
		Name:        "purple-span",
		Timestamp:   time.Now(),
		Duration:    time.Second,
		ServiceName: "DistributedCalculatorGoAdderOpenTelemetry",
		Attributes: map[string]interface{}{
			"color": "purple",
		},
	})

	router := mux.NewRouter()

	router.HandleFunc("/add", add).Methods("POST", "OPTIONS")
	log.Fatal(http.ListenAndServe(":6000", router))
}
