package arc

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/twlabs/personal-assistant/arc/logger"
)

const DEFAULT_PORT = "8090"

func Run() {
	logger.Init()

	startService()
}

func startService() {

	finalHandler := http.HandlerFunc(httpHandler)
	http.Handle("/", Authenticate(finalHandler))
	addr := ":" + port()
	logger.Fatal("ListenAndServe: ", http.ListenAndServe(addr, nil))
}

func port() string {
	if val, ok := os.LookupEnv("PORT"); ok {
		return val
	} else {
		return DEFAULT_PORT
	}
}

func planetRequestHandler(request ActionRequest) ActionResponse {
	request.init()

	actionExecutor, err := lookupActionExecutor(request.ActionName)

	if err != nil {
		return errorResponse(err)
	}

	return actionExecutor.Execute(request)
}

func errorResponse(err error) ActionResponse {
	logger.Error(err)
	response := SimpleResponse{Text: err.Error()}
	return response
}

func httpHandler(resp http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		resp.WriteHeader(http.StatusOK)
		writeBody(resp, map[string]interface{}{"hello": "planet world"})
		return
	}
	if request.Method != "POST" {
		resp.WriteHeader(http.StatusBadRequest)
	}
	decoder := json.NewDecoder(request.Body)
	var planetReq ActionRequest
	err := decoder.Decode(&planetReq)
	if err != nil {
		logger.Error(err)
		resp.WriteHeader(http.StatusBadRequest)
		writeBody(resp, err)
		return
	}

	writeBody(resp, planetRequestHandler(planetReq).ToPlanetResponse())
}

func writeBody(resp http.ResponseWriter, body interface{}) {
	resp.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(resp)
	encoder.Encode(body)
}
