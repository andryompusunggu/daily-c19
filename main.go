package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	url := "http://ubsupdates.com/covid.ubs/ajax/isi_survey_baru.php?id=mnx7zkIb&ya=0&tidak=13&ket=&ip=125.161.212.29"
	method := "POST"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
