package imgur

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const apiURL = "https://api.imgur.com/3/image"

// UploadImage upload image to imgur
func UploadImage(clientID string, data []byte) (string, error) {
	formData := url.Values{
		"image": []string{base64.StdEncoding.EncodeToString(data)},
		"type":  []string{"base64"},
	}
	request, _ := http.NewRequest("POST", apiURL, strings.NewReader(formData.Encode()))
	request.Header.Add("Authorization", "Client-ID "+clientID)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Printf("upload image with error=%v", err)
		return "", err
	}
	defer response.Body.Close()
	respData, err := ioutil.ReadAll(response.Body)
	jsonMap := map[string]interface{}{}
	json.Unmarshal(respData, &jsonMap)
	imageLink := jsonMap["data"].(map[string]interface{})["link"].(string)
	return imageLink, nil
}
