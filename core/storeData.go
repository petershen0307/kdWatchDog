package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// for google script transfer json structure
type gsData struct {
	ID      string
	DailyKD []KDResult
}

// SaveDataToGoogleSheet send the data to google sheet
func SaveDataToGoogleSheet(kdInfo []KDResult, stockID string, tabName pricePeriod) error {
	googleScriptURL := fmt.Sprintf("https://script.google.com/macros/s/AKfycbwpX0l_OVCz9jV5JpKoBexfyk_8zzgtCRCaySL8hOlarjWaTbox/exec?tab=%s", tabName)
	timeoutRequest := http.Client{Timeout: time.Minute * 1}
	url := fmt.Sprintf(googleScriptURL)
	sendData := gsData{
		ID:      stockID,
		DailyKD: kdInfo[len(kdInfo)-1:],
	}
	encodeData, err := json.Marshal(sendData)
	if err != nil {
		return err
	}
	response, err := timeoutRequest.Post(url, "application/json", bytes.NewBuffer(encodeData))
	if err != nil {
		return err
	}
	defer response.Body.Close()
	if http.StatusOK != response.StatusCode {
		return fmt.Errorf("Got http error code %v", response.StatusCode)
	}
	return nil
}
