package main

import (
	"encoding/json"
	_ "encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main() {

	go post()

	resp, err := http.Get("http://localhost:7000/courses")

	if err != nil {
		log.Fatalln(err)
		return
	}

	var result map[string]interface{}
	body, err := ioutil.ReadAll(resp.Body)

	// Unmarshal or Decode the JSON to the interface.
	json.Unmarshal([]byte(body), &result)

	if err != nil {
		log.Fatalln(err)
		return
	}

	sb := string(body)
	log.Printf(sb)


}


func addCourse(courseName string, rate int) url.Values{
	data := url.Values{
		"name":                    {courseName},
		"students":                {},
		"teachers":                {},
		"satisfaction_rating_avg": {string(rune(rate))},
	}
	return data
}

func post() {
	response, err := http.PostForm("http://localhost:7000/courses", addCourse("bdsa",10))

	if err != nil {
		log.Fatal(err)
	}

	var res map[string]interface{}

	json.NewDecoder(response.Body).Decode(&res)

}

