package cmd

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
	"github.com/spf13/cobra"
	"io"
	"os"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Runs a container using the specified image.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		imageName := args[0]
		ctx := context.Background()
		cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
		if err != nil {
			panic(err)
		}
		defer cli.Close()

		out, err := cli.ImagePull(ctx, imageName, image.PullOptions{})
		if err != nil {
			panic(err)
		}
		defer out.Close()
		io.Copy(os.Stdout, out)

		resp, err := cli.ContainerCreate(ctx, &container.Config{
			Image: imageName,
		}, nil, nil, nil, "")
		if err != nil {
			panic(err)
		}

		if err := cli.ContainerStart(ctx, resp.ID, container.StartOptions{}); err != nil {
			panic(err)
		}

		fmt.Println(resp.ID)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
