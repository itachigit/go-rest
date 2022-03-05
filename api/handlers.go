package api

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/itachigit/goREST/db/models"
)

func (svr *CowinServer) GetState(w http.ResponseWriter, r *http.Request) {
	req, err := http.NewRequest("GET", svr.config.Cowin.Url+"/v2/admin/location/states", http.NoBody)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.95 Safari/537.36")
	//req.Header.Set("Authorization", r.Header.Get("Authorization"))
	if err != nil {
		svr.logger.Error("error-creating-request: ", err)
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		svr.logger.Error("error-serving-request: ", err)
	}

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		svr.logger.Error("error-serving-request: ", res.Status)
	}
	states := []models.State{}
	jsonBody, err := io.ReadAll(res.Body)
	if err != nil {
		svr.logger.Error("error-reading-response: ", err)
	}
	jsonBody = []byte("[" + strings.Split(strings.Split(string(jsonBody), "[")[1], "]")[0] + "]")

	err = json.Unmarshal(jsonBody, &states)
	if err != nil {
		svr.logger.Error("error-unmarshaling-response: ", err)
	}

	db := svr.db
	db.Create(&states)
	w.Write(jsonBody)

}
