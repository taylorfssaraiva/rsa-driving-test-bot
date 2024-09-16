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
const bearerToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRheWxvcmZzc2FyYWl2YUBnbWFpbC5jb20iLCJ1bmlxdWVfbmFtZSI6IlRBWUxPUiIsImZhbWlseV9uYW1lIjoiRk9OU0VDQSBTQVJBSVZBIiwic3ViIjoiOGViZmU2NmItNzU0Mi1lZjExLWFmODktMDA1MDU2YjliNTBjIiwicHBzbiI6IjE5NjIzOTZQQSIsIjJmYWF1dGgiOiJ0cnVlIiwibXlnb3YiOiJ0cnVlIiwibXlnb3Z0b2tlbiI6ImV5SmhiR2NpT2lKU1V6STFOaUlzSW10cFpDSTZJbk5wWjI1cGJtZHJaWGt1YlhsbmIzWnBaQzUyTVNJc0luUjVjQ0k2SWtwWFZDSjkuZXlKbGVIQWlPakUzTWpFeU1EazNOVE1zSW01aVppSTZNVGN5TVRJd056azFNeXdpZG1WeUlqb2lNUzR3SWl3aWFYTnpJam9pYUhSMGNITTZMeTloWTJOdmRXNTBMbTE1WjI5MmFXUXVhV1V2WlRFNU4yVmhPV1V0TURKbE5TMDBZMkkyTFRrMllqSXROVFUzTVdOa05qUTFOelUwTDNZeUxqQXZJaXdpYzNWaUlqb2lNV3BaVFZFdlRrZEpWVEZqV2pCdldUZHFWMnMyZVUxSFdFRkNjR3RGY2xNdlRVbG1ia0UyV1dGRFdUMGlMQ0poZFdRaU9pSmtOV1k1WmpobVppMWxPVEExTFRSa09ERXRPVGc0WWkxaFlXSTFNV1l6T0RjeU9UTWlMQ0pwWVhRaU9qRTNNakV5TURjNU5UTXNJbUYxZEdoZmRHbHRaU0k2TVRjeU1USXdOemsxTXl3aVpXMWhhV3dpT2lKMFlYbHNiM0ptYzNOaGNtRnBkbUZBWjIxaGFXd3VZMjl0SWl3aWIybGtJam9pT0RkalpUTXpZek10TlRNMk55MDBaV0ZpTFRnM05Ua3RPV0UwTjJFMk1qQXhPRGsySWl3aVVIVmliR2xqVTJWeWRtbGpaVTUxYldKbGNpSTZJakU1TmpJek9UWlFRU0lzSWtKcGNuUm9SR0YwWlNJNklqQXpMekEzTHpFNU9UUWlMQ0pNWVhOMFNtOTFjbTVsZVNJNklreHZaMmx1SWl3aVoybDJaVzVPWVcxbElqb2lWR0Y1Ykc5eUlpd2ljM1Z5Ym1GdFpTSTZJa1p2Ym5ObFkyRWdVMkZ5WVdsMllTSXNJbTF2WW1sc1pTSTZJak0xTXpBNE16ZzNOVFkwTVRZaUxDSkVVMUJQYm14cGJtVk1aWFpsYkNJNklqSWlMQ0pFVTFCUGJteHBibVZNWlhabGJGTjBZWFJwWXlJNklqSWlMQ0pEZFhOMGIyMWxja2xrSWpvaU9URXpPRFV5TUNJc0lrRmpZMlZ3ZEdWa1VISnBkbUZqZVZSbGNtMXpJanAwY25WbExDSkJZMk5sY0hSbFpGQnlhWFpoWTNsVVpYSnRjMVpsY25OcGIyNU9kVzFpWlhJaU9pSTNJaXdpVTAxVE1rWkJSVzVoWW14bFpDSTZabUZzYzJVc0lrRmpZMlZ3ZEdWa1VISnBkbUZqZVZSbGNtMXpSR0YwWlZScGJXVWlPakUyTlRnNE16STJPVGtzSW5SeWRYTjBSbkpoYldWM2IzSnJVRzlzYVdONUlqb2lRakpEWHpGQlgzTnBaMjVwYmkxV05TMU1TVlpGSWl3aVEyOXljbVZzWVhScGIyNUpaQ0k2SW1VMllqYzFaak15TFROak5UY3RORGd3WmkwNE1ERTJMVFpsWXpoaFpqRmhOekl3WVNKOS5sZGM5d1UxaThMTy1LYTZJR0g4c1FOWnJfdXBHdC1wMUpJN1ZfWWM4Y0p0aXVZUHpJb2ZMNnZMdG5RYVBYdG1seldDMnNWRDlsTlgwVWtTSGljMlhvQWVsYjJZWkt6YjFsc2pEWjlDNmpUSWh2R0xwSGhDRDYtRUxKcjJpY3VTc3R3cHRJSWVoUEhaMFA1clBuZGNwdUhyVkxfTW1QNzBtd09tem12QVN6WHd6VzhaeGdxNHU2S3dPbDY2TjJBTzBpSm9JZWhEMjVXN1BfMTZjTjk5OEdRWHZGNmtsN1BXUFVEeUkyMzh0OE5LcURRcjkyb3RRYnBVZm5nQnZSaXJSOUtqc1JlbUZVcGstU2dxMVByeDBGQi11MkYwdFl4TDZSM2ZMcVRzeUl4LUU0MmRxV1oyNlctQ2Q1bTlVdXVjNm9iSXZVc2ZrNEs3aC1KU1dhZXFGNXciLCJuYmYiOjE3MjEyMDc5NTQsImV4cCI6MTcyMTIxNTE1NCwiaWF0IjoxNzIxMjA3OTU0LCJpc3MiOiJteXJvYWRzYWZldHkucnNhLmllIiwiYXVkIjoibXlyb2Fkc2FmZXR5LnJzYS5pZSJ9.oBJBotT9a2wTa_JE-6x6YCckJrRZqIZu2FZm3UWF2ew"

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
