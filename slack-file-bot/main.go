package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)


func main(){
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-4560896334343-4659667105523-g8PlGM5vrwNsMMDE05mBFRum")
	os.Setenv("CHANNEL_ID", "C04GX9HUWP5")

	api:=slack.New(os.Getenv("SLACK_BOT_TOKEN"))

	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"pic.jpg"}


	for i := 0; i < len(fileArr); i++{
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File: fileArr[i],
		}
		file, err := api.UploadFile(params)
		if err != nil{
			fmt.Printf("&s\n", err)
			return
		}
		fmt.Printf("Name: %s, URL:%s\n", file.Name, file.URL)
	}
}
