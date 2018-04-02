package utilities

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/sync/errgroup"
)

type Ebook struct {
	URL   string `json:"url"`
	Title string `json:"title"`
	Image string `json:"image"`
}

type Ebooks struct {
	TotalPages  int     `json:"total_pages"`
	TotalEbooks int     `json:"total_ebooks"`
	List        []Ebook `json:"ebooks"`
}

type Player struct {
	FirstName   string `json:"first-name"`
	LastName    string `json:"last-name"`
	DOB         string `json:"DOB"`
	Position    string `json:"position"`
	Nationality string `json:"nationality"`
}
type Position struct {
	Name string `json:"name"`
}
type Team struct {
	Name        string   `json:"name"`
	StadiumName string   `json:"stadium-name"`
	ListPlayer  []Player `json:"players"`
}

type Teams struct {
	ListTeam []Team `json:"teams"`
}

func NewTeams() *Teams {
	return &Teams{}
}

func NewEbooks() *Ebooks {
	return &Ebooks{}
}

func (teams *Teams) GetTeamsByUrl(url string) error {

	doc, err := goquery.NewDocument(url)
	if err != nil {
		return err
	}
	linkUrls := []string{}
	doc.Find(".yui-t6 > #bd > #yui-main > .yui-b > .clearfix > .content-column > #page_competition_1_block_competition_tables_7-wrapper > .content > #page_competition_1_block_competition_tables_7 > #page_competition_1_block_competition_tables_7_block_competition_league_table_1 > form > table > tbody").Each(func(i int, table *goquery.Selection) {
		table.Find(".team_rank").Each(func(i int, tr *goquery.Selection) {
			linkUlr, exists := tr.Find(".large-link > a").Attr("href")
			if !exists {
				linkUlr = ""
			}
			linkUlr = "https://us.soccerway.com" + linkUlr
			linkUrls = append(linkUrls, linkUlr)
		})
	})
	//eg := errgroup.Group{}
	for _, uri := range linkUrls {
		if uri != "" {
			teams.GetTeamByUrl(uri)
		}
	}
	return nil
}

func (teams *Teams) GetTeamByUrl(uri string) error {
	doc, err := goquery.NewDocument(uri)
	if err != nil {
		return err
	}
	log.Println(doc.Text())
	clubName := doc.Find("#subheading h1").Text()
	stadiumName := doc.Find("#doc4 > #bd > #yui-main > .yui-b > .clearfix > .content-column > .block-horizontal-container > .second-element > #page_team_1_block_team_venue_4-wrapper > .content > #page_team_1_block_team_venue_4 > .fully-padded > .clearfix > .details > *:nth-child(2)").Text()
	team := Team{
		Name:        clubName,
		StadiumName: stadiumName,
	}
	linkPlayerUrls := []string{}

	doc.Find("#doc4 > #bd").Each(func(i int, table *goquery.Selection) {
		table.Find("#page_team_1_block_team_squad_8-wrapper > .content > #page_team_1_block_team_squad_8 > .squad-container > #page_team_1_block_team_squad_8-table > tbody > tr").Each(func(i int, tr *goquery.Selection) {
			linkUlr, exists := tr.Find("a").Attr("href")
			if !exists {
				linkUlr = ""
			}
			linkUlr = "https://us.soccerway.com" + linkUlr
			linkPlayerUrls = append(linkPlayerUrls, linkUlr)
		})
	})
	// eg := errgroup.Group{}
	// for _, uri := range linkPlayerUrls {
	// 	if uri != "" {
	// 		eg.Go(func() error {
	// 			err := team.GetPlayerByUrl(uri)
	// 			if err != nil {
	// 				return err
	// 			}
	// 			return nil
	// 		})
	// 	}
	// }
	for _, uri := range linkPlayerUrls {
		if uri != "" {
			team.GetPlayerByUrl(uri)
		}
	}
	teams.ListTeam = append(teams.ListTeam, team)
	return nil
}

func (team *Team) GetPlayerByUrl(url string) error {

	log.Println(url)

	doc, err := goquery.NewDocument("https://us.soccerway.com/national/england/premier-league/20172018/regular-season/r41547/")
	if err != nil {
		return err
	}
	// log.Println(doc.Text())
	// doc.Find("#doc4 > #bd > #yui-main > .yui-b").Each(func(i int, table *goquery.Selection) {
	// 	table.Find("#page_player_1_block_player_passport_3-wrapper > .content > #page_player_1_block_player_passport_3 > .fully-padded> .yui-gc > yui-u first > clearfix").Each(func(i int, div *goquery.Selection) {
	// 		fistName := div.Find("dl > *:nth-child(2)").Text()
	// 		lastName := div.Find("dl > *:nth-child(4)").Text()
	// 		nationality := div.Find("dl > *:nth-child(6)").Text()
	// 		dob := div.Find("dl > *:nth-child(8)").Text()
	// 		position := div.Find("dl > *:nth-child(16)").Text()
	// 		Player := Player{
	// 			FirstName:   fistName,
	// 			LastName:    lastName,
	// 			DOB:         dob,
	// 			Position:    position,
	// 			Nationality: nationality,
	// 		}
	// 		team.ListPlayer = append(team.ListPlayer, Player)
	// 	})

	// })
	doc.Find("*").Each(func(i int, table *goquery.Selection) {
		table.Find(".team_rank").Each(func(i int, tr *goquery.Selection) {
			linkUlr, exists := tr.Find(".large-link > a").Attr("href")
			if !exists {
				linkUlr = ""
			}
			linkUlr = "https://us.soccerway.com" + linkUlr
		})
	})
	return nil
}

func (ebooks *Ebooks) getEbooksByUrl(url string) error {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return err
	}

	doc.Find(".col-left ._2pin").Each(func(i int, s *goquery.Selection) {
		docTitle, exists := s.Find(".ellipsis a").Attr("title")
		if !exists {
			docTitle = ""
		}
		docLink, exists := s.Find(".ellipsis a").Attr("href")
		if !exists {
			docLink = "#"
		}
		docImg, exists := s.Find("a._3if7 img").Attr("src")
		if !exists {
			docImg = ""
		}
		Ebook := Ebook{
			URL:   docLink,
			Title: docTitle,
			Image: docImg,
		}
		ebooks.TotalEbooks++
		ebooks.List = append(ebooks.List, Ebook)
	})

	return nil
}

func (ebooks *Ebooks) GetTotalPages(url string) error {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return err
	}
	lastPageLink, _ := doc.Find("ul.pagination li:last-child a").Attr("href")
	if lastPageLink == "javascript:void();" {
		ebooks.TotalPages = 1
		return nil
	}
	split := strings.Split(lastPageLink, "?page=")
	totalPages, _ := strconv.Atoi(split[1])
	ebooks.TotalPages = totalPages
	return nil
}

func (ebooks *Ebooks) GetAllEbooks(currentUrl string) error {
	eg := errgroup.Group{}
	if ebooks.TotalPages > 0 {
		for i := 1; i <= ebooks.TotalPages; i++ {
			uri := fmt.Sprintf("%v?page=%v", currentUrl, i)
			eg.Go(func() error {
				err := ebooks.getEbooksByUrl(uri)
				if err != nil {
					return err
				}
				return nil
			})
		}
		if err := eg.Wait(); err != nil {
			return err
		}
	}
	return nil
}
