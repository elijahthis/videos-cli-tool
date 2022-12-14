package main

import (
	"encoding/json"
	"io/ioutil"
	// "encoding/json"
)

type video struct {
	Id, Title, Description, ImageURL, URL string
}

func getVideos() (videos []video) {
	fileBytes, err := ioutil.ReadFile("./videos.json")

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(fileBytes, &videos)

	if err != nil {
		panic(err)
	}

	return videos

}

func saveVideos(videos []video) {
	videoBytes, err := json.Marshal(videos)

	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("./videos.json", videoBytes, 0644)

	if err != nil {
		panic(err)
	}
}
