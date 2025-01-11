package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Document struct {
	ID   int         `json:"id"`
	Data interface{} `json:"data"`
}

type JSONPatch struct {
	Op    string `json:"op"`
	Path  string `json:"path"`
	From  string `json:"from,omitempty"`
	Value string `json:"value,omitempty"`
}

func main() {
	baseURL := "https://dummy-json-patch.netlify.app/.netlify/functions"
	docId := 4
	apiURL := fmt.Sprintf("%s/documents/?id=%d", baseURL, docId)

	jsonPatchData := []JSONPatch{
		{
			Op:    "test",
			Path:  "/new_name",
			Value: "dummy",
		},
		{
			Op:    "add",
			Path:  "/user/id",
			Value: "123",
		},
	}
	jsonPatchBytes, err := json.Marshal(jsonPatchData)
	if err != nil {
		panic(err)
	}
	jsonPatchBody := string(jsonPatchBytes)
	/*
			jsonPatchBody := `[
		        {
		            "op": "test",
		            "path": "/new_name",
		            "value": "dummy"
		        },
		        {
		            "op": "add",
		            "path": "/user/id",
		            "value": 123
		        }
		    ]`
	*/

	req, err := http.NewRequest("PATCH", apiURL, strings.NewReader(jsonPatchBody))
	req.Header.Set("Content-Type", "application/json-patch+json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var updatedDoc Document

	resBody, err := io.ReadAll(resp.Body)
	fmt.Println(string(resBody))
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(resBody, &updatedDoc)
	if err != nil {
		panic(err)
	}

	fmt.Println("Updated/Patched Document", updatedDoc)
	fmt.Println("id:", updatedDoc.ID)
	fmt.Println("document:", updatedDoc.Data)

}
