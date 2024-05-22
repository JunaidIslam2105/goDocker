package cmd

import (
	"context"
	"fmt"
	containertypes "github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"os"

	"github.com/spf13/cobra"
)

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop one or more containers",
	Long:  `This command will stop one or more containers given their IDs.`,
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
			noWaitTimeout := 0
			if err := cli.ContainerStop(ctx, containerID, containertypes.StopOptions{Timeout: &noWaitTimeout}); err != nil {
				panic(err)
			}
			fmt.Println("Container stopped: ", containerID)
		}
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
