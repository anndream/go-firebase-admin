package admin_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	admin "github.com/acoshift/go-firebase-admin"
	"github.com/stretchr/testify/assert"
)

func TestSendToDevices(t *testing.T) {

	t.Run("send=success", func(t *testing.T) {
		// generate Response
		srv := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

			rw.WriteHeader(http.StatusOK)
			rw.Header().Set("Content-Type", "application/json")
			fmt.Fprint(rw, `{"success": 1,"failure": 0,"results": [{"message_id":"118218","registration_id": "abcd118218","error": ""}]}`)
		}))
		defer srv.Close()

		app := initApp()
		firFCM := app.FCM()
		firFCM.NewFcmSendEndpoint(srv.URL)

		assert.NotNil(t, app)
		assert.NotNil(t, firFCM)

		response, err := firFCM.SendToDevice(context.Background(), "mydevicetoken",
			admin.Message{Notification: admin.Notification{
				Title: "Hello go firebase admin",
				Body:  "My little Big Notification",
				Color: "#ffcc33"},
			})

		assert.Nil(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, 1, response.Success)
		assert.Equal(t, 0, response.Failure)
	})

	t.Run("send=failure", func(t *testing.T) {
		// generate Response
		srv := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

			rw.WriteHeader(http.StatusBadRequest)
		}))
		defer srv.Close()

		app := initApp()
		firFCM := app.FCM()
		firFCM.NewFcmSendEndpoint(srv.URL)

		assert.NotNil(t, app)
		assert.NotNil(t, firFCM)

		response, err := firFCM.SendToDevice(context.Background(), "mydevicetoken",
			admin.Message{Notification: admin.Notification{
				Title: "Hello go firebase admin",
				Body:  "My little Big Notification",
				Color: "#ffcc33"},
			})

		assert.NotNil(t, err)
		assert.Nil(t, response)

	})

}