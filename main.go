package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	rulesengine "github.com/cachefactory/go_rules_engine/internal/rulesengine"
)

func engine(w http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	engine, err := rulesengine.FromJson(string(body))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println("Processing rule set")
	engine.Run()

	jsonResponse, _ := json.Marshal(engine.JsonResponse())

	w.Header().Set("Content-Type", "application/json")

	w.Write(jsonResponse)
}

func main() {
	http.HandleFunc("/rules_engine", engine)
	http.ListenAndServe(":8090", nil)
	fmt.Println("Listening on port 8090")
}
