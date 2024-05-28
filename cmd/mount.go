package cmd

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/spf13/cobra"
	"os"
)

var mountCmd = &cobra.Command{
	Use:   "mount",
	Short: "Mount a directory from the host to the container",
	Long:  `This command will mount a directory from the host to the container.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 3 {
			fmt.Println("Error: Image ID, source directory, and target directory must be provided")
			os.Exit(1)
		}

		imageID := args[0]
		srcDir := args[1]
		targetDir := args[2]

		ctx := context.Background()
		cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
		if err != nil {
			panic(err)
		}
		defer cli.Close()

		hostConfig := &container.HostConfig{
			Binds: []string{fmt.Sprintf("%s:%s", srcDir, targetDir)},
		}

		resp, err := cli.ContainerCreate(ctx, &container.Config{
			Image: imageID,
		}, hostConfig, nil, nil, "")
		if err != nil {
			panic(err)
		}

		if err := cli.ContainerStart(ctx, resp.ID, container.StartOptions{}); err != nil {
			panic(err)
		}

		fmt.Printf("Successfully mounted %s to %s in container %s\n", srcDir, targetDir, resp.ID)
	},
}

func init() {
	rootCmd.AddCommand(mountCmd)
}
