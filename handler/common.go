package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var bearerPrefix = "Bearer "

const (
	TypeErrMsg = "ERROR_MESSAGE"
	TypeMsg    = "MESSAGE"
	ERR_MSG    = "ERROR_MESSAGE"
	MSG        = "MESSAGE"
)

type ResStruct struct {
	Status   string `json:"status" example:"SUCCESS" example:"FAILED"`
	HTTPCode int    `json:"httpCode" example:"200" example:"500"`
	Message  string `json:"message" example:"pong" example:"could not connect to db"`
}

type Res500Struct struct {
	Status   string `json:"status" example:"FAILED"`
	HTTPCode int    `json:"httpCode" example:"500"`
	Message  string `json:"message" example:"could not connect to db"`
}

type Res400Struct struct {
	Status   string `json:"status" example:"FAILED"`
	HTTPCode int    `json:"httpCode" example:"400"`
	Message  string `json:"message" example:"Invalid param"`
}

type RequestData struct {
	Start time.Time
	w     http.ResponseWriter
	r     *http.Request
}

type RenderData struct {
	Data  interface{}
	Paths []string
}

type TemplateData struct {
	Data interface{}
}

func (t *TemplateData) SetConstants() {

}

func logAndGetContext(w http.ResponseWriter, r *http.Request) *RequestData {
	w.Header().Add("X-Content-Type-Options", "nosniff")
	w.Header().Add("X-Frame-Options", "DENY")

	start := time.Now()

	return &RequestData{
		Start: start,
		r:     r,
		w:     w,
	}
}

func jsonifyMessage(msg string, msgType string, httpCode int) ([]byte, int) {
	var data []byte
	var Obj struct {
		Status   string `json:"status"`
		HTTPCode int    `json:"httpCode"`
		Message  string `json:"message"`
	}
	Obj.Message = msg
	Obj.HTTPCode = httpCode
	switch msgType {
	case TypeErrMsg:
		Obj.Status = "FAILED"

	case TypeMsg:
		Obj.Status = "SUCCESS"
	}
	data, _ = json.Marshal(Obj)
	return data, httpCode
}

func writeJSONMessage(msg string, msgType string, httpCode int, rd *RequestData) int {
	d, code := jsonifyMessage(msg, msgType, httpCode)
	return writeJSONResponse(d, code, rd)
}

func writeJSONStruct(v interface{}, code int, rd *RequestData) int {
	fmt.Printf("writeJSONStruct req interface : %v", v)
	d, err := json.Marshal(v)
	if err != nil {
		return writeJSONMessage("Unable to marshal data. Err: "+err.Error(), TypeErrMsg, http.StatusInternalServerError, rd)
	}

	return writeJSONResponse(d, code, rd)
}

func writeJSONResponse(d []byte, code int, rd *RequestData) int {
	fmt.Print(rd.r, time.Since(rd.Start).Seconds(), code)
	fmt.Println("Status Code:", code, ", Response time:", time.Since(rd.Start))
	fmt.Println("Status Code:", code, ", Response time:", time.Since(rd.Start), " Response:", string(d))

	//rd.w.Header().Set("Access-Control-Allow-Origin", "*")
	rd.w.Header().Set("Content-Type", "application/json; charset=utf-8")
	rd.w.WriteHeader(code)
	rd.w.Write(d)
	return code
}

type ResStructErrorCode struct {
	StatusCode string `json:"status_code"`
	Message    string `json:"message"`
}

func respStruct(message string, statusCode string) *ResStructErrorCode {
	return &ResStructErrorCode{StatusCode: statusCode, Message: message}
}
