package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type User struct {
	ID    int    `json:"id,omitempty"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Roles string `json:"Roles"`
}

func main() {
	apiURL := "https://dummy-json-patch.netlify.app/.netlify/functions/users/?id=2"

	jsonPatchBody := `[
        {
            "op": "replace",
            "path": "/name",
            "value": "some dummy name"
        }
    ]`

	req, err := http.NewRequest("PATCH", apiURL, strings.NewReader(jsonPatchBody))
	req.Header.Set("Content-Type", "application/json-patch+json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var user User

	resBody, err := io.ReadAll(resp.Body)
	fmt.Println(string(resBody))
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(resBody, &user)
	if err != nil {
		panic(err)
	}

	fmt.Println("Updated/Patched User", user)
	fmt.Println("id:", user.ID)
	fmt.Println("name:", user.Name)
	fmt.Println("email:", user.Email)
	fmt.Println("roles:", user.Roles)

}
