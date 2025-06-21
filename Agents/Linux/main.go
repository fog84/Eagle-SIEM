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

const TOKEN_PATH = "conf/token"
const FILES_TO_MONITOR = "conf/files_to_monitor.lst"
const LAST_LINE_READ_SAVES_PATH = "save_last_lines/last_lines_"
const INDEXER_URL = "http://localhost:8080"

func getToken() string {
	data, err := ioutil.ReadFile(TOKEN_PATH)
	if err != nil {
		fmt.Println(err)
	}
	return string(data)
}

func getFilesToMonitor() []string {
	data, _ := ioutil.ReadFile(FILES_TO_MONITOR)
	return strings.Split(string(data), "\n")
}

func readNewLines(filesToMonitor []string) []string{
	var newLinesFromAllFilesToMonitor []string

	for _, file := range filesToMonitor {
		lastLineReadSavePath := LAST_LINE_READ_SAVES_PATH + strings.ReplaceAll(file, "/", "_")
		
		data, err := ioutil.ReadFile(lastLineReadSavePath)

		var lastLineRead int
		if err != nil {
			lastLineRead = 0
		} else {
			lastLineRead, _ = strconv.Atoi(string(data))
		}
		lines, _ := ioutil.ReadFile(file)
		newLines := strings.Split(string(lines), "\n")[lastLineRead:]
		newLinesFromAllFilesToMonitor = append(newLinesFromAllFilesToMonitor, newLines...)


		lastLineRead_current := lastLineRead + len(newLines)
		if lastLineRead == 0 {
			f, _ := os.Create(lastLineReadSavePath)
			defer f.Close()
			f.WriteString(strconv.Itoa(lastLineRead_current))
		} else {
			f, _ := os.OpenFile(lastLineReadSavePath, os.O_WRONLY|os.O_TRUNC, 0644)
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

    req, _ := http.NewRequest(method, INDEXER_URL, payload)
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
	filesToMonitor := getFilesToMonitor()

	newLinesFromAllFilesToMonitor := readNewLines(filesToMonitor)

	if len(newLinesFromAllFilesToMonitor) != 0 {
		sendNewLinesToDB(newLinesFromAllFilesToMonitor, token)
	}
}