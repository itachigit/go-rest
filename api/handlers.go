package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"go-rest/db/models"
)

func (svr *CowinServer) GetState(w http.ResponseWriter, r *http.Request) {
	req, err := http.NewRequest("GET", svr.config.Cowin.Url+"/entries", http.NoBody)
	//req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.95 Safari/537.36")
	//req.Header.Set("Authorization", r.Header.Get("Authorization"))
	logger := svr.GetLogger("get-state")
	if err != nil {
		logger.Error("error-creating-request: ", err)
	}
	logger.Info("entry")
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		logger.Error("error-serving-request: ", err)
	}

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		logger.Error("error-serving-request: ", res.Status)
	}
	entries := []models.Entry{}
	jsonBody, err := io.ReadAll(res.Body)
	if err != nil {
		logger.Error("error-reading-response: ", err)
	}
	jsonBody = []byte("[" + strings.Split(strings.Split(string(jsonBody), "[")[1], "]")[0] + "]")

	err = json.Unmarshal(jsonBody, &entries)
	if err != nil {
		logger.Error("error-unmarshaling-response: ", err)
	}
	db := svr.db
	count := 0
	for _, entry := range entries {
		if count > 15 {
			break
		}
		fmt.Print(count, " ")
		db.FirstOrCreate(&entry, models.Entry{API: entry.API})
		count++
	}

	logger.Info("states-created-in-db, returning-response")
	entries1 := []models.Entry{}
	if err = db.Find(&entries1).Error; err != nil {
		logger.Error("error-getting-entries-from-database: ", err)
	}
	jsonBody, err = json.Marshal(&entries1)
	fmt.Println(len(entries1))
	if err != nil {
		logger.Error("error-marshaling-entries: ", err)
	}
	w.Write(jsonBody)
	logger.Info("exit")
}
