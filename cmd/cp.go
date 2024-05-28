package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"path/filepath"
)

// cpCmd represents the cp command
var cpCmd = &cobra.Command{
	Use:   "cp",
	Short: "Copy files/folders between a container and the local filesystem",
	Long:  `This command will copy files/folders between a container and the local filesystem.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Println("Error: Source path and container ID must be provided")
			os.Exit(1)
		}

		srcPath := args[0]
		containerID := args[1]

		// Use the docker cp command to copy the file
		cpCmd := exec.Command("docker", "cp", srcPath, fmt.Sprintf("%s:/", containerID))

		// Run the command and capture the output
		_, err := cpCmd.CombinedOutput()
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			os.Exit(1)
		}

		fmt.Printf("Successfully copied %s to container %s\n", filepath.Base(srcPath), containerID)

	},
}

func init() {
	rootCmd.AddCommand(cpCmd)
}
