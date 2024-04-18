package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func BankService(ctx *gin.Context, responsePtr interface{}) (interface{}, error) {
	endpointURL := viper.GetString("BANK_URL")

	requestBody, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", endpointURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	// req.Header.Set("bearer", "paseto token")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, responsePtr)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error en la solicitud: %v", resp.Status)
	}

	//for the moment is returning null,
	return nil, nil
}
