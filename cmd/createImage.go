package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

// createImageCmd represents the createImage command
var createImageCmd = &cobra.Command{
	Use:   "createImage",
	Short: "Create a new image from a Dockerfile.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Println("Please provide a name for the image and the path to the Dockerfile.")
			return
		}

		imageName := args[0]
		dockerfilePath := args[1]

		dockerBuildCmd := exec.Command("docker", "build", "-t", imageName, dockerfilePath)

		output, err := dockerBuildCmd.CombinedOutput()
		if err != nil {
			fmt.Printf("Error while creating Docker image: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Output: %s\n", output)
	},
}

func init() {
	rootCmd.AddCommand(createImageCmd)
}
