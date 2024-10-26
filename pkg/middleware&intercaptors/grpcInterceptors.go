package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
)

func main() {
	var URL string
	var port int
	var a, b int

	fmt.Scan(&URL, &port, &a, &b)

	var nums []int

	url := fmt.Sprintf("%s:%d?a=%d&b=%d", URL, port, a, b)

	res, _ := http.Get(url)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(body, nums)

	sort.Ints(nums)

	for i := len(nums) - 1; i > 0; i-- {
		fmt.Println(nums[i])
	}
}
