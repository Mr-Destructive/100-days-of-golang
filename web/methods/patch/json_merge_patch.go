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
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	Bio   string `json:"bio,omitempty"`
	Roles string `json:"roles,omitempt"`
}

func main() {
	apiURL := "https://dummy-json-patch.netlify.app/.netlify/functions/users/?id=2"

	userObj := User{
		Name:  "dummy",
		Roles: "dummy role",
	}

	var jsonPatchBody []byte
	jsonPatchBody, err := json.Marshal(userObj)
	if err != nil {
		panic(err)
	}
	fmt.Println("Request Body:", string(jsonPatchBody))

	//jsonPatchBody := `{
	//    "name": "dummy",
	//    "roles": "new dummy role"
	//}`

	req, err := http.NewRequest("PATCH", apiURL, strings.NewReader(string(jsonPatchBody)))
	req.Header.Set("Content-Type", "application/merge-patch+json")

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
	fmt.Println("bio:", user.Bio)
	fmt.Println("email:", user.Email)
	fmt.Println("roles:", user.Roles)
}
