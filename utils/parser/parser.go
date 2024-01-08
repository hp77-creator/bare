package parser

import (
	"bare/utils/osutil"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const rawUrl = "https://raw.githubusercontent.com/"
const rawApiUrl = "https://api.github.com/repos/"

type Bare struct {
	Author       string
	BareName     string
	Version      string
	BarePath     string
	Variants     []string // Template name -> description (to be asked in prompt)
	Placeholders map[string][]string
	License      string
}

var BareObj Bare

func Parser(filePath string) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// var bareObj Bare;
	err = json.Unmarshal(data, &BareObj)
	if err != nil {
		log.Fatal(err)
	}

	if BareObj.BareName == "" {
		log.Fatal("Bare name not present")
	}
}

func GetRecipe(user string, repo string, branch string) {
	if user == "bare-cli" {
		reqUrl := rawUrl + user + "/" + repo + "/" + branch + "/recipe.json"
		log.Println("Info: " + reqUrl)
		resp, err := http.Get(reqUrl)
		if err != nil {
			log.Fatal(err.Error())
		}
		log.Println("Response: " + resp.Status)
		respCode, _ := strconv.Atoi(resp.Status[0:3])
		if respCode >= 400 {
			log.Fatal("recipe.json doesn't exist for the repo")
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		//Convert the body to type string
		err = json.Unmarshal(body, &BareObj)
		if err != nil {
			log.Fatal(err)
		}

		BareObj.BarePath = osutil.GetBarePath()
	} else {
		reqUrl := rawApiUrl + user + "/" + repo
		log.Println("Info: " + reqUrl)
		resp, err := http.Get(reqUrl)
		if err != nil {
			log.Fatal(err)
		}
		respCode, _ := strconv.Atoi(resp.Status[0:3])
		if respCode >= 400 {
			log.Println("Response Code: " + resp.Status)
			log.Fatal("Request didn't go through")
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		//Convert the body to type string
		err = json.Unmarshal(body, &BareObj)
		if err != nil {
			log.Fatal(err)
		}

		BareObj.BarePath = osutil.GetBarePath()
	}
}

// Return user, repo, branch
func ParseGithubRepo(bareName string) (string, string, string) {

	userRepo := strings.Split(bareName, "/")

	if len(userRepo) <= 1 {
		fmt.Println("Invalid Bare path")
		os.Exit(-1)
	}

	user := userRepo[0]
	repo := userRepo[1]

	repoBranch := strings.Split(repo, "#")
	branch := "main"
	if len(repoBranch) > 1 {
		branch = repoBranch[1]
	}

	repo = repoBranch[0]

	return user, repo, branch
}
