package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const url = "https://myroadsafety.rsa.ie/api/v1/Availability/slots/1647b57c-5043-ef11-af89-005056b9b50c/null/89d8d682-8ad4-e911-a2d7-005056827428/0fed074d-c2d6-e811-a2c0-005056823b22/e5bbe47a-3f94-e911-a2be-0050568fd8e0/a7e0e690-6c0f-d07f-1d48-d88138a5055e"
const bearerToken = ""

type Response struct {
	Slots  []interface{} `json:"slots"`
	Dates  []interface{} `json:"dates"`
	Months []interface{} `json:"months"`
}

func checkAvailability() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Authorization", "Bearer "+bearerToken)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Referer", "https://myroadsafety.rsa.ie/portal/booking/new/e5bbe47a-3f94-e911-a2be-0050568fd8e0/d2dc5f8c-2506-ea11-a2c3-0050568fd8e0")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error making request: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response: %v", err)
	}

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatalf("Error unmarshalling response: %v", err)
	}

	timeNow := time.Now().Format(time.TimeOnly)
	if len(response.Slots) > 0 || len(response.Dates) > 0 {
		fmt.Printf("[%s] Available slots: %v\n", timeNow, response.Slots)
		fmt.Printf("[%s] Available dates: %v\n", timeNow, response.Dates)
	} else {
		fmt.Printf("[%s] No slots or dates available\n", timeNow)
	}
}

func main() {
	for {
		checkAvailability()
		time.Sleep(1 * time.Minute)
	}
}
