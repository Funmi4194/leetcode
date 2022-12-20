package main

import (
	"fmt"
	// "strings"
)
/*You are given  two string arrays, queries, and dictionary. All words
in each array ccomprise of lowercase English letters and have the same 
length. 
In one edit you can take a word from queries, and change any letter. Find 
all words from queries that , after a maximum of two edits, equal some 
word from dictionary. Return a list of all words from queries 
*/
func two_edits(queries []string, dictionary []string) []string {

	var str []string 
	/* str holds the list of words  equvalent to dictionary
	with or without two edits*/

	for _, query := range queries { // for every word in queries store it in query
		for _, dict := range dictionary { //for every word in dictionary store it in dict
			edit := 0 //edit is 0 at first instance
			for i:=0; i < len(query); i++{ //It gets a word from query to check if its in dict
				if query[i] != dict[i] { //If the word is not equal dict then edit 
					edit++ //edit 1 or 2 times
				}
			}
			if edit <= 2{ // if edit is less than 2 (after one or two equal)
				str = append(str, query) // add word to str using the keyword append
				continue // continue the loop until the last element in queries
			}
		}
	}
	
	return str
}
func main() {
	fmt.Println(two_edits([]string{"word", "note", "ants", "wood", "cays"}, []string{"joke", "wood","cays", "moat"}))
	 fmt.Println(two_edits([]string{"yes"}, []string{"not"}))
}
