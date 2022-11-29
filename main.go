package main

import (
	"encoding/json"
	"fmt"
	"log"
)

const data = `[
	{
	 "Title": "moby dick",
	 "Price": 10,
	 "Released": 118281600
	},
	{
	 "Title": "odyssey",
	 "Price": 15,
	 "Released": 733622400
	},
	{
	 "Title": "hobbit",
	 "Price": 25,
	 "Released": -62135596800
	}
   ]`

func main() {

	// var pocket money = 10

	// fmt.Println("I have ", pocket)
	// l := list{
	// 	{Title: "moby dick", Price: 10, Released: toTimestamp(118281600)},
	// 	{Title: "odyssey", Price: 15, Released: toTimestamp("733622400")},
	// 	{Title: "hobbit", Price: 25},
	// }

	// l.discount(.5)

	// sort.Sort(l)

	// sort.Sort(byReleaseDate(l))
	// sort.Sort(sort.Reverse(byReleaseDate(l)))

	// fmt.Print(l)

	// data, err := json.MarshalIndent(l, "", " ")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(string(data))

	var l list
	err := json.Unmarshal([]byte(data), &l)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(l)
}
