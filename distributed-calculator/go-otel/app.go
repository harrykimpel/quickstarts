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

	"github.com/gorilla/mux"
	"github.com/newrelic/go-agent/v3/integrations/nrgorilla"
	newrelic "github.com/newrelic/go-agent/v3/newrelic"
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
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("distributed-calculator-go-adder"),
		newrelic.ConfigLicense("<NEW-RELIC-LICENSE-KEY>"),
		newrelic.ConfigDebugLogger(os.Stdout),
		func(config *newrelic.Config) {
			config.DatastoreTracer.SlowQuery.Enabled = true
			config.DatastoreTracer.SlowQuery.Threshold = 1
			config.DistributedTracer.Enabled = true
		},
	)
	if nil != err {
		fmt.Println(err)
		os.Exit(1)
	}

	router := mux.NewRouter()
	router.Use(nrgorilla.Middleware(app))

	router.HandleFunc(newrelic.WrapHandleFunc(app, "/add", add))
	log.Fatal(http.ListenAndServe(":6000", router))
}
