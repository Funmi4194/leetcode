package main

import (
	"log"
	"fmt"
	"strings"
)




func twoEditWords(queries []string, dictionary []string) []string {

var str []string // hold
var word string	// hold each of the word in convertToLower in queries
var word2 string
var newWord2 string // neword2 hold the value of second edit
var newWord string // neword hold the value of first edit
var convertTolower []string
    

	
for i := 0; i <len(queries); i++{ 
	convertTolower = append(convertTolower, strings.ToLower(queries[i])) //incase any word in queries is upper case convert to lowercase
	//store value in convertFlower

	fmt.Print("Enter a word in queries: ")// this segmet of code is not really neccessary but i just added
	_, err := fmt.Scanln(&word)
	if err != nil{
	log.Fatal(err)
	}
	word = strings.ToLower(word)

	if word == convertTolower[i]{
		convertTolower[i] = word
	}
	for j := 0; j < len(dictionary); j++{ // check if word is in dict then append to str
		if word == dictionary[j]{
			str = append(str, word)
			}
			continue
	   }
    }

    for k := 0; k < len(dictionary); k++{
		
		fmt.Print("Let's edit twice to check if and item in queries will be equals dictionary")
		fmt.Println("")
		fmt.Println("FIRST EDIT")
		fmt.Print("Enter a word to edit: ")
		_, err := fmt.Scanln(&word2)
	    if err != nil{
	    log.Fatal(err)
	    }
		
		if word2 != dictionary[k]{
		fmt.Print("Enter the letter you want to replace in word: ")

		var firstReplace string
		_, err := fmt.Scanln(&firstReplace)
		if err != nil{
			log.Fatal(err)
		}

		fmt.Print("Enter the replaced letter: ")
		var secondReplace string
		_, err = fmt.Scanln(&secondReplace)
		if err != nil{
			log.Fatal(err)
		   }
		newWord = strings.Replace(word2, firstReplace, secondReplace, 1)
		
		if newWord == dictionary[k]{
			str = append(str, word2)
		}else{
			fmt.Println("")
			fmt.Println("SECOND EDIT")


		fmt.Print("Enter the letter you want to replace in word: ")
		var replace string
		_, err := fmt.Scanln(&replace)
		if err != nil{
			log.Fatal(err)
		   }

		fmt.Print("Enter the replaced letter: ")
		var replace2 string
		_, err = fmt.Scanln(&replace2)
		if err != nil{
			log.Fatal(err)
		   }
		newWord2 = strings.Replace(newWord, replace, replace2, 1)
		if newWord2 == dictionary[k]{
			str = append(str, word2)
		}else{
			fmt.Printf("It will take more than 2 edits for %q\n", word2)

		}
		}
	}else{
		
		fmt.Printf("%q is in dictionary", word2)
		str = append(str, word2)
		
	}

}	
return str

}



func main() {
 fmt.Println(twoEditWords([]string{"word", "note", "ants", "wood", "car"}, []string{"wood", "joke", "moat", "car"}))
// fmt.Println(twoEditWords([]string{"yes"}, []string{"not"}))
}