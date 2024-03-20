package main

import (
	"assignment3/routers"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Status struct {
	Water  int    `json:"water"`
	Wind   int    `json:"wind"`
	Status string `json:"status"`
}

func main() {
	const PORT = ":8088"

	fmt.Sprintln("server start at", PORT)
	go routers.StartServer().Run(PORT)

	ticker := time.NewTicker(15 * time.Second)
	defer ticker.Stop()

	go func() {
		for {
			select {
			case <-ticker.C:
				postStatus()
			}
		}
	}()

	select {}
}

func postStatus() {
	data := map[string]interface{}{
		"status": Status{
			Water:  rand.Intn(100),
			Wind:   rand.Intn(100),
			Status: "",
		},
	}

	reqJson, err := json.Marshal(data)
	if err != nil {
		log.Fatal(reqJson)
	}

	clinet := &http.Client{}

	req, err := http.NewRequest("POST", "http://localhost:8088/status", bytes.NewBuffer(reqJson))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := clinet.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}
