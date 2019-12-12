package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/bgoldovsky/counter-api/internal/models"
	"github.com/bgoldovsky/counter-api/internal/services/mock"
	"github.com/golang/mock/gomock"
)

func TestRequest(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	service := mock.NewMockCounter(ctrl)

	exp := models.State{Total: 150, LastUpdate: time.Now().UTC()}

	service.EXPECT().Inc().Return(nil)
	service.EXPECT().State().Return(exp)

	server := NewServer(service)
	handler := http.HandlerFunc(server.ServeHTTP)

	r, err := http.NewRequest(http.MethodGet, "/", strings.NewReader(""))
	if err != nil {
		fmt.Printf("new request error: %v", err)
	}

	w := httptest.NewRecorder()
	handler.ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Fatalf("wrong code. expected %d, got %d\n", http.StatusOK, w.Code)
	}

	act := models.State{}
	err = json.Unmarshal(w.Body.Bytes(), act)
	if err != nil {
		fmt.Printf("unmarshaling error: %v", err)
	}

	if reflect.DeepEqual(exp, act) {
		t.Errorf("wrong result. extected %v, got %v\n", exp, w.Body.String())
	}
}
