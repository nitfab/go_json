package main

import (
	"fmt"
	"encoding/json"	

	"github.com/alphadose/haxmap"
)

func main() {
	// initialize map with key type `int` and value type `string`
	mep := haxmap.New[string, interface{}]()

	// set a value (overwrites existing value if present)
	mep.Set("name", "Nitin Sharma")
	mep.Set("age", 22)
	mep.Set("job", "Full Stack Developer")

	val, ok := mep.Get("name")
	if ok {
		println(val)
		println(ok)
	}

	output, err := json.Marshal(val)
	if err != nil {
		fmt.Println("Error")
	}
	fmt.Println(string(output))
}