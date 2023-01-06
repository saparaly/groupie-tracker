package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func DisplayCards(s string) Artists {
	resp, err := http.Get(s)
	if err != nil {
		fmt.Println("No response from request")

		// return "err1"
	}

	// fmt.Println(resp)

	body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(body))

	var res Artists
	json.Unmarshal(body, &res)

	return res
}
