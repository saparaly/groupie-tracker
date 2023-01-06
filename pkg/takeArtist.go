package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func TakeArtist(id int, s string) Artisttest {
	resp, err := http.Get(s)
	if err != nil {
		fmt.Println("No response from request")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var artist []Artisttest

	json.Unmarshal(body, &artist)

	defer resp.Body.Close()
	artist[20].Image = "https://www.vagalume.com.br/mamonas-assassinas/images/mamonas-assassinas.jpg"
	// fmt.Println(artist[id])

	return artist[id-1]
}

func DatesConcert(id int) Relations {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/relation/" + strconv.Itoa(id))
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var concert Relations
	json.Unmarshal(body, &concert)
	return concert
}
