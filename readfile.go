package main

import (
	"fmt"
	"log"
	"readfile/datafile"
)

func main() {
	lines, err := datafile.GetStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	votes := make(map[string]int)
	for _, line := range lines {
		votes[line]++
	}
	for name, vote := range votes {
		fmt.Printf("Votes for %s: %d\n", name, vote)
	}
	delete(votes, "Amber Graham")
	fmt.Println(votes)
}
