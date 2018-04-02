package main

import (
	"log"

	"./src/utilities"
)

func checkError(err error) {
	if err != nil {
		log.Println(err)
	}
}

func main() {

	currentUrl := "https://us.soccerway.com/national/england/premier-league/20172018/regular-season/r41547/"
	// ebooks := utilities.NewEbooks()
	// err := ebooks.GetTotalPages(currentUrl)
	// checkError(err)
	// err = ebooks.GetAllEbooks(currentUrl)
	// checkError(err)
	// ebooksJson, err := json.Marshal(ebooks)
	// checkError(err)
	// err = ioutil.WriteFile("output.json", ebooksJson, 0644)
	// checkError(err)
	teams := utilities.NewTeams()
	err := teams.GetTeamsByUrl(currentUrl)
	checkError(err)
}
