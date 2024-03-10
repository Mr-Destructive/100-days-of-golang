package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*
GET
https://dummyjson.com/products/1

{"id":1,"title":"iPhone 9","description":"An apple mobile which is nothing like apple","price":549,"discountPercentage":12.96,"rating":4.69,"stock":94,"brand":"Apple","category":"smartphones","thumbnail":"https://cdn.dummyjson.com/product-images/1/thumbnail.jpg","images":["https://cdn.dummyjson.com/product-images/1/1.jpg","https://cdn.dummyjson.com/product-images/1/2.jpg","https://cdn.dummyjson.com/product-images/1/3.jpg","https://cdn.dummyjson.com/product-images/1/4.jpg","https://cdn.dummyjson.com/product-images/1/thumbnail.jpg"]}
*/

// we define a struct for the json data
// we can only include the fields that we want
type Product struct {
	ID                 int      `json:"id"`
	Title              string   `json:"title"`
	Description        string   `json:"description"`
	Price              float64  `json:"price"`
	DiscountPercentage float64  `json:"discountPercentage"`
	Rating             float64  `json:"rating"`
	Stock              int      `json:"stock"`
	Brand              string   `json:"brand"`
	Category           string   `json:"category"`
	Thumbnail          string   `json:"thumbnail,omitempty"`
	Images             []string `json:"-"`
	/*
	   // we can skip the fields that we don't want
	      Thumbnail          string  `json:"thumbnail"`
	      Images             []string `json:"images"`
	   // or include omitempty flag
	   // the omitempty flag will skip the field if it is empty
	       Thumbnail          string  `json:"thumbnail,omitempty"`
	       Images             []string `json:"images,omitempty"`
	*/
}

func main() {
	reqURL := "https://dummyjson.com/products/1"
	resp, err := http.Get(reqURL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var data Product
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&data)
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
	//we can now access the data with the struct
	fmt.Println(data.Title)
}
