package auth

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

func getApiKeyTestHeader(key, value string) http.Header {
	header := http.Header{}
	if key != "" {
		header.Add(key, value)
	}
	return header
}

func TestSplit(t *testing.T) {

	tests := map[string]struct {
		input string
		want  string
	}{
		"Missing Auth":     {input: "", want: "Error: no authorization header in http request"},
		"Basic":            {input: "Authorization: ApiKey absckdfper", want: "absckdfper"},
		"Wrong auth":       {input: "Authorization: Bearer absckdfper", want: "Error: wrong authorization header"},
		"Misstyped Auth 1": {input: "Authoration: ApiKey absckdfper", want: ""},
		"Misstyped Auth 2": {input: "Authorization: ApiKeyabsckdfper", want: ""},
	}

	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			var got = ""
			strs := strings.Split(testCase.input, ": ")
			log.Println("Checking that strings are passed correctly in the input")
			if len(strs) != 2 {
				got = "Error: no authorization header in http request"
				if !reflect.DeepEqual(testCase.want, got) {
					t.Fatalf("expected: %v, got: %v", testCase.want, got)
				}
				return
			} else if !strings.HasPrefix(strs[1], "ApiKey") {
				got = "Error: wrong authorization header"
				if !reflect.DeepEqual(testCase.want, got) {
					t.Fatalf("expected: %v, got: %v", testCase.want, got)
				}
				return
			}
			log.Println("Init header")
			header := getApiKeyTestHeader(strs[0], strs[1])
			log.Println("Call")
			got, err := GetAPIKey(header)
			if err != nil {
				fmt.Println(err)
			}
			if !reflect.DeepEqual(testCase.want, got) {
				t.Fatalf("expected: %v, got: %v", testCase.want, got)
			}
		})
	}
}
