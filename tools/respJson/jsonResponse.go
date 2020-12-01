package respJson

import (
	"encoding/json"
	"fmt"
	"github.com/vivaldy22/mekar-regis-client/model"
	"net/http"
)

func WriteJSON(success bool, code int, msg string, data interface{}, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	if success {
		resp := model.JSONResponse{
			Success: success,
			Message: msg,
			Data:    data,
		}
		jsonResp, _ := json.Marshal(resp)
		w.Write(jsonResp)
	} else {
		fail := model.JSONFailResponse{
			Success:   success,
			Message:   fmt.Sprintf("%v, %v", msg, err),
			ErrorCode: code,
			Data:      nil,
		}
		jsonResp, _ := json.Marshal(fail)
		w.Write(jsonResp)
	}
}
