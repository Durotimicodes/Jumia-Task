package test

//
//import (
//	"context"
//	"encoding/json"
//	"errors"
//	"github.com/Durotimicodes/jumia-phone-number-task/database/mocks"
//	"github.com/Durotimicodes/jumia-phone-number-task/handlers"
//	"github.com/Durotimicodes/jumia-phone-number-task/models"
//	"github.com/Durotimicodes/jumia-phone-number-task/routers"
//	"github.com/golang/mock/gomock"
//	"github.com/stretchr/testify/assert"
//	"net/http"
//	"net/http/httptest"
//	"strings"
//	"testing"
//)
//
//func TestGetMobileNumbers(t *testing.T) {
//	ctrl := gomock.NewController(t)
//
//	defer ctrl.Finish()
//
//	mockDB := mocks.NewMockDBInterface(ctrl)
//
//	h := &handlers.Handler{
//		Repository: mockDB,
//	}
//
//	route, _ := routers.SetUpRouters(h)
//
//	countryCode := "258"
//	stateOne := "OK"
//	Mobile := "(258) 291244811"
//
//
//	configNum := models.ConfigureMobileNumber{
//		Country:      "Cameroon",
//		State:        stateOne,
//		CountryCode:  countryCode,
//		MobileNumber: Mobile,
//	}
//
//	bodyJSON, err := json.Marshal(configNum)
//	if err != nil {
//		t.Fail()
//	}
//
//	t.Run("Testing For Error", func(t *testing.T) {
//		c := context.Context()
//		mockDB.EXPECT().GetMobileNumbers(c , countryCode, stateOne).Return(nil, errors.New("Error Exist"))
//		rw := httptest.NewRecorder()
//		req, _ := http.NewRequest(http.MethodGet, "/api/v1/user/mobile/"+countryCode, strings.NewReader(string(bodyJSON)))
//		route.ServeHTTP(rw, req)
//		assert.Equal(t, http.StatusBadRequest, rw.Code)
//		assert.Contains(t, rw.Body.String(), "Bad request")
//
//	})
//
//}
