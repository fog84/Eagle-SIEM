package main

import (
	"fmt"
	"io/ioutil" // c'est déprécié fraudra changer
	"os" // maintenant on utilise os
	"strings"
	"strconv"
	"encoding/json"
	"net/http"
	"time"
)

const TOKEN_PATH = "/etc/eagle_agent/conf/token"
const FILES_TO_MONITOR = "/etc/eagle_agent/conf/files_to_monitor.lst"
const LAST_LINE_READ_SAVES_PATH = "/etc/eagle_agent/save_last_lines/last_lines_"
const INDEXER_URL = "http://localhost:8000/events"

func handleError(err error) {
	if err != nil{
		fmt.Println(err)
	}
}

func getToken() string {
	data, err := ioutil.ReadFile(TOKEN_PATH)
	handleError(err)
	return string(data)
}

func getFilesToMonitor() []string {
	data, err := ioutil.ReadFile(FILES_TO_MONITOR)
	handleError(err)
	return strings.Split(string(data), "\n")
}

func readNewLines(filesToMonitor []string) []string{
	var newLinesFromAllFilesToMonitor []string

	for _, file := range filesToMonitor {
		if file == "" {
			continue
		}
		
		lastLineReadSavePath := LAST_LINE_READ_SAVES_PATH + strings.ReplaceAll(file, "/", "_")

		data, err := ioutil.ReadFile(lastLineReadSavePath)

		var lastLineRead int
		if err != nil {
			lastLineRead = 0
		} else {
			lastLineRead, err = strconv.Atoi(string(data))
			handleError(err)
		}
		lines, err := ioutil.ReadFile(file)
		handleError(err)
		newLines := strings.Split(string(lines), "\n")[lastLineRead:]
		newLinesFromAllFilesToMonitor = append(newLinesFromAllFilesToMonitor, newLines...)


		lastLineRead_current := lastLineRead + len(newLines)
		if lastLineRead == 0 {
			f, err := os.Create(lastLineReadSavePath)
			handleError(err)
			defer f.Close()
			f.WriteString(strconv.Itoa(lastLineRead_current))
		} else {
			f, err := os.OpenFile(lastLineReadSavePath, os.O_WRONLY|os.O_TRUNC, 0644)
			handleError(err)
			f.WriteString(strconv.Itoa(lastLineRead_current))
		}
	}
	return newLinesFromAllFilesToMonitor
}

func sendNewLinesToDB(newLinesFromAllFilesToMonitor []string, token string) {
	mapD := map[string]interface{}{
		"newLines": newLinesFromAllFilesToMonitor,
	}
	jsonData, _ := json.Marshal(mapD)

    client := &http.Client{
        Timeout: time.Second * 10,
    }

    method := "POST"
    payload := strings.NewReader(string(jsonData))

    req, err := http.NewRequest(method, INDEXER_URL, payload)
	handleError(err)
	req.Header.Add("Content-Type", "application/json")
    req.Header.Add("Authorization", token)

	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
}

func main() {
	token := getToken()

	fmt.Println(token)

	for {
		filesToMonitor := getFilesToMonitor()

		newLinesFromAllFilesToMonitor := readNewLines(filesToMonitor)
		fmt.Println(newLinesFromAllFilesToMonitor)
	
		if len(newLinesFromAllFilesToMonitor) != 0 {
			sendNewLinesToDB(newLinesFromAllFilesToMonitor, token)
		}

		time.Sleep(10 * time.Second)
	}
}