package cmd

import (
	"fmt"
	"os"
	"os/exec"

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

	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)
}

func downloaderVideo(videoUrl, videoFolderName string) {
	//Check if the folder exists
	if _, err := os.Stat(videoFolderName); os.IsNotExist(err) {
		os.Mkdir(videoFolderName, 0755)
	}

	command := fmt.Sprintf("neo -d %s .mp4 %s", videoFolderName, videoUrl)

	// Execute the command
	_, err := exec.Command("cmd", "/C", command).Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("Video downloaded successfully at location: %s", videoFolderName)
}

func saveInfoToDb(videoUrl, videoFolderName string) {

}
