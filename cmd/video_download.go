package cmd

import (
	"errors"
	"fmt"
	"os"
	"os/exec"

	"github.com/ArchishmanSengupta/neo-cli/config"
	"github.com/ArchishmanSengupta/neo-cli/models"
	"github.com/kkdai/youtube/v2"
	"github.com/spf13/cobra"
)

var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download a video",
	Long: `Download command takes 2 arguments
					Argument 1. Video URL
					Argument 3. Video Folder Name where it will be saved`,

	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		videoUrl := args[0]
		videoFolderName := args[1]
		downloaderVideo(videoUrl, videoFolderName)
		saveInfoToDb(videoUrl, videoFolderName)
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)
}

func downloaderVideo(videoUrl, videoFolderName string) {
	//Check if the folder exists
	if _, err := os.Stat(videoFolderName); errors.Is(err, os.ErrNotExist) {
		fmt.Println(err)
		fmt.Println("The provided path does not exist")
		os.Mkdir(videoFolderName, 0755)
	}

	//command to download the video in the folder as in the github.com/kkdai/youtube/v2 documentation
	cmd := fmt.Sprintf(`youtubedr download -d ./ -o %s.mp4 %s`, videoFolderName, videoUrl)

	_, err := exec.Command(cmd).Output()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Video downloaded successfully at location: %s", videoFolderName)
}

func saveInfoToDb(videoUrl, videoFolderName string) {
	client := youtube.Client{}
	video, err := client.GetVideo(videoUrl)

	if err != nil {
		fmt.Println(err)
		return
	}

	metadata := models.Metadata{}

	metadata.Title = video.Title
	metadata.FolderName = videoFolderName
	metadata.Url = videoUrl

	err = metadata.Insert(videoUrl, videoFolderName, config.DbConn)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
