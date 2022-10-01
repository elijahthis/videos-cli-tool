package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	// videos get subCommand
	getCmd := flag.NewFlagSet("get", flag.ExitOnError)

	// inputs for videos 'get' subcommand
	getAll := getCmd.Bool("all", false, "Get all videos")
	getId := getCmd.String("id", "", "Get YouTube video by ID")

	// videos add subCommand
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)

	// inputs for videos 'get' subcommand
	addID := addCmd.String("id", "", "YouTube Video ID")
	addTitle := addCmd.String("title", "", "YouTube Video title")
	addImageURL := addCmd.String("imageURL", "", "YouTube Video Image URL")
	addURL := addCmd.String("URL", "", "YouTube Video URL")
	addDesc := addCmd.String("desc", "", "YouTube Video description")

	// short validation
	if len(os.Args) < 2 {
		fmt.Println("expected 'get' or 'add' subcommands")
		os.Exit(1)

	}

	switch os.Args[1] {
	case "get":
		HandleGet(getCmd, getAll, getId)

	case "add":
		HandleAdd(addCmd, addID, addTitle, addImageURL, addURL, addDesc)

	default:
	}

}

func HandleGet(getCmd *flag.FlagSet, all *bool, id *string) {
	getCmd.Parse(os.Args[2:])

	if *all == false && *id == "" {
		fmt.Println("id is required or spacify --all to get all videos")
		getCmd.PrintDefaults()
		os.Exit(1)

	}
	if *all {
		//return all videos
		videos := getVideos()

		fmt.Printf("ID \t Title \t URL \t ImageURL \t Description \n")
		fmt.Printf("-------------\n")
		fmt.Printf("-------------\n")

		for _, video := range videos {
			fmt.Printf("%v \t %v \t %v \t %v \t %v \t", video.Id, video.Title, video.URL, video.ImageURL, video.Description)
		}

		return
	}

	if *id != "" {
		videos := getVideos()
		id := *id

		for _, video := range videos {
			if video.Id == id {
				fmt.Printf("%v \t %v \t %v \t %v \t %v \t", video.Id, video.Title, video.URL, video.ImageURL, video.Description)
			}
		}

	}

}

func ValidateVideo(addCmd *flag.FlagSet, id *string, title *string, imageURL *string, URL *string, desc *string) {
	addCmd.Parse(os.Args[2:])

	if *id == "" || *title == "" || *imageURL == "" || *URL == "" || *desc == "" {
		fmt.Println("All fields are required when adding a video")

		if *id == "" {
			fmt.Println("No ID")
		}
		if *title == "" {
			fmt.Println("No title")
		}
		if *imageURL == "" {
			fmt.Println("No imageURL")
		}
		if *URL == "" {
			fmt.Println("No URL")
		}
		if *desc == "" {
			fmt.Println("No desc")
		}

		addCmd.PrintDefaults()
		os.Exit(1)
	}
}

func HandleAdd(addCmd *flag.FlagSet, id *string, title *string, imageURL *string, URL *string, desc *string) {
	ValidateVideo(addCmd, id, title, imageURL, URL, desc)

	video := video{
		Id:          *id,
		Title:       *title,
		Description: *desc,
		ImageURL:    *imageURL,
		URL:         *URL,
	}

	videos := getVideos()
	videos = append(videos, video)

	saveVideos(videos)
}
