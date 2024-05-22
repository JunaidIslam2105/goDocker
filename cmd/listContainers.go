package cmd

import (
	"context"
	"fmt"
	containertypes "github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/spf13/cobra"
)

// listContainersCmd represents the listContainers command
var listContainersCmd = &cobra.Command{
	Use:   "listContainers",
	Short: "List all running containers",
	Long:  `This command will list all the running containers in the same format as docker ps.`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
		if err != nil {
			panic(err)
		}
		defer cli.Close()

		containers, err := cli.ContainerList(ctx, containertypes.ListOptions{All: true})
		if err != nil {
			panic(err)
		}

		fmt.Printf("%-12s %-20s %-15s %-15s\n", "CONTAINER ID", "IMAGE", "COMMAND", "STATUS")
		for _, container := range containers {
			fmt.Printf("%-12.12s %-20.20s %-15.15s %-15s\n", container.ID, container.Image, container.Command, container.State)
		}
	},
}

func init() {
	rootCmd.AddCommand(listContainersCmd)
}
