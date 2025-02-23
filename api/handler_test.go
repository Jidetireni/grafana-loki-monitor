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
		ChannelID: "01952e92-8ab0-7c08-9df4-dbaa1f4d6c9d",
		ReturnURL: "https://ping.telex.im/v1/webhooks/01952e92-8ab0-7c08-9df4-dbaa1f4d6c9d",
		Settings: []Setting{
			{"Loki Server URL", "text", true, "http://100.27.210.53:3100"},
			{"Loki Query", "text", true, "{job=\"varlogs\"}"},
			{"interval", "text", true, "*/2 * * * *"},
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

	assert.Equal(t, "Loki integration", response["event_name"])
	// Wait briefly to allow async goroutine to execute
	time.Sleep(500 * time.Millisecond)
}
