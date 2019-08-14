package textbox

import (
	"encoding/json"
	result "github.com/heaptracetechnology/machinebox-textbox/result"
	"github.com/machinebox/sdk-go/textbox"
	"net/http"
	"strings"
)

//MachineBox struct
type MachineBox struct {
	Text    string `json:"text,omitempty"`
	Address string `json:"address,omitempty"`
}

//TextAnalyze MachineBox-Textbox
func TextAnalyze(responseWriter http.ResponseWriter, request *http.Request) {

	decoder := json.NewDecoder(request.Body)
	var param MachineBox
	decodeErr := decoder.Decode(&param)
	if decodeErr != nil {
		result.WriteErrorResponseString(responseWriter, decodeErr.Error())
		return
	}

	client := textbox.New(param.Address)

	text := strings.NewReader(param.Text)

	textResult, resultErr := client.Check(text)
	if resultErr != nil {
		result.WriteErrorResponseString(responseWriter, resultErr.Error())
		return
	}

	bytes, _ := json.Marshal(textResult)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}
