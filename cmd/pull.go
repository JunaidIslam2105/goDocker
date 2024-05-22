package cmd

import (
	"context"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"

	"io"
	"os"

	"github.com/spf13/cobra"
)

// pullCmd represents the pull command
var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pulls the latest image of the specified name from dockerhub.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		imgName := args[0]
		ctx := context.Background()
		cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
		if err != nil {
			panic(err)
		}
		defer cli.Close()

		out, err := cli.ImagePull(ctx, imgName, image.PullOptions{})
		if err != nil {
			panic(err)
		}

		defer out.Close()

		io.Copy(os.Stdout, out)
	},
}

func init() {
	rootCmd.AddCommand(pullCmd)
}
