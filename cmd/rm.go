package cmd

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"os"

	"github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove one or more containers",
	Long:  `This command will remove one or more containers given their IDs.`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
		if err != nil {
			panic(err)
		}
		defer cli.Close()

		if len(args) < 1 {
			fmt.Println("Error: No container IDs provided")
			os.Exit(1)
		}

		for _, containerID := range args {
			if err := cli.ContainerRemove(ctx, containerID, container.RemoveOptions{}); err != nil {
				fmt.Printf("Error removing container %s: %v\n", containerID, err)
				continue
			}
			fmt.Println("Container removed: ", containerID)
		}
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)
}
