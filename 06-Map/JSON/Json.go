package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	res := make(map[string]interface{})
	res["code"] = 200
	res["message"] = "success"
	res["data"] = map[string]interface{}{
		"username": "Tom",
		"age":      31,
		"hobby":    []string{"reading,running"},
	}
	fmt.Println("map data:", res)

	jsons, errs := json.Marshal(res)
	if errs != nil {
		fmt.Println("json marshal error:", errs)
	}
	fmt.Println("--- map to json")
	fmt.Println(string(jsons))

	res2 := make(map[string]interface{})
	errs = json.Unmarshal(jsons, &res2)
	if errs != nil {
		fmt.Println("json unmarshal error:", errs)
	}
	fmt.Println("--- json to map ---")
	fmt.Println("map data :", res2)

	person := map[int]string{
		1: "tom",
		2: "jack",
		3: "John",
	}
	fmt.Println("map data:", person)

	delete(person, 2)
	fmt.Println("data :", person)

	person[3] = "LaLa"
	fmt.Println("map data:", person)
}
