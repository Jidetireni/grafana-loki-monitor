package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestTickHandler(t *testing.T) {

	router := gin.Default()
	SetRoute(router)
	w := httptest.NewRecorder()

	// Prepare request

	payload := RequestBody{
		ChannelID: "019527af-f248-7e03-b2fc-bed0265814a7",
		ReturnURL: "https://ping.telex.im/v1/webhooks/019527af-f248-7e03-b2fc-bed0265814a7",
		Settings: []Setting{
			{"Loki Server URL", "text", true, "http://100.27.210.53:3100"},
			{"Loki Query", "text", true, "{job=\"varlogs\"}"},
			{"interval", "text", true, "2 * * * *"},
		},
	}

	JsonPayload, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "/tick", bytes.NewBuffer(JsonPayload))
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	// Validate immediate response
	assert.Equal(t, http.StatusOK, w.Code)

	fmt.Println("Response Body:", w.Body.String())
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, "019527af-f248-7e03-b2fc-bed0265814a7", response["channel_id"])
	// Wait briefly to allow async goroutine to execute
	time.Sleep(500 * time.Millisecond)
}
