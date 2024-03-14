package main

import (
	"fmt"
	"log"
	"os"
	"github.com/karashiiro/bingode"
	"github.com/xivapi/godestone/v2"
)

func checkFile () error {
	_, err := os.Stat("players.txt") //if the fle doens't exist, it'll store the error in err
	if os.IsExist(err) {
		err := os.Remove("players.txt")
		if err != nil {
			fmt.Println(err)
		}
		} else {
			f, err := os.Create("players.txt")
			if err != nil {
				fmt.Println("Error creating file")
		}
		defer f.Close()
	}
	return err
}


func main() {
	checkFile()
	s := godestone.NewScraper(bingode.New(), godestone.EN)

	for member := range s.FetchFreeCompanyMembers("9229142273877457819") {
		if member.Error != nil {
			log.Fatalln(member.Error)
		}
		//log.Println(member.Name)
		f, err := os.OpenFile("players.txt", os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("Error opening file")
			return
		}

		newName := member.Name
		_, err = fmt.Fprintln(f, newName)
		if err != nil {
			fmt.Println("Failed to write names")
			f.Close()
			return
		}
		err = f.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	fmt.Println("File written to")

}