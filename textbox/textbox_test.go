package textbox

import (
	"bytes"
	"encoding/json"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"log"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("Text Analyze with invalid param", func() {

	textbox := []byte(`{"status":false}`)
	requestBody := new(bytes.Buffer)
	encodeErr := json.NewEncoder(requestBody).Encode(textbox)
	if encodeErr != nil {
		log.Fatal(encodeErr)
	}

	request, err := http.NewRequest("POST", "/textAnalyze", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(TextAnalyze)
	handler.ServeHTTP(recorder, request)

	Describe("analyze text", func() {
		Context("analyze", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Text Analyze", func() {

	textbox := MachineBox{Text: "I can't wait for @machineboxio to release Textbox; it provides natural language processing and a whole host of other #useful things."}
	requestBody := new(bytes.Buffer)
	encodeErr := json.NewEncoder(requestBody).Encode(textbox)
	if encodeErr != nil {
		log.Fatal(encodeErr)
	}

	request, err := http.NewRequest("POST", "/textAnalyze", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(TextAnalyze)
	handler.ServeHTTP(recorder, request)

	Describe("analyze text", func() {
		Context("analyze", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})
