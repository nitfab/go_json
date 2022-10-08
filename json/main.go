package main

import (
	"encoding/json"
	"fmt"
)

// Want to convert a data in json from the data(it can be anything array, key-value pair)is simply called an encoding or marhsal.

type course struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"price"`
	Platform string   `json:"platform"`
	Password string   `json:"-"`              //- will not show the password
	Tags     []string `json:"tags,omitempty"` //If there is a nil value don't show it
}

func main() {
	fmt.Println("Welcome to JSON")
	EncodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"ReactJS", 299, "Udemy", "ABCDEF", []string{"web-dev", "js"}},
		{"MERN", 299, "Udemy", "ABCDEF", []string{"mern", "js"}},
		{"Angular", 299, "Udemy", "ABCDEF", nil},
	}

	finalJson, err := json.Marshal(lcoCourses)
	// finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")

	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)

}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "MERN",
		"price": 299,
		"platform": "Udemy",
		"tags": ["mern","js"]
	}
	`)

	var lcocourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lcocourse)
		fmt.Printf("%#v\n", lcocourse)
	} else {
		fmt.Println("JSON was not valid")
	}

	//Some cases where you just want to add data to key value
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v, Value is %v and Type is %T\n", k, v, v)
	}
}