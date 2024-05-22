package cmd

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"os"

	"github.com/spf13/cobra"
)

// rmiCmd represents the rmi command
var rmiCmd = &cobra.Command{
	Use:   "rmi",
	Short: "Remove one or more images",
	Long:  `This command will remove one or more images given their IDs.`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
		if err != nil {
			panic(err)
		}
		defer cli.Close()

		if len(args) < 1 {
			fmt.Println("Error: Image ID not provided")
			os.Exit(1)
		}

		for _, imageID := range args {
			_, err := cli.ImageRemove(ctx, imageID, types.ImageRemoveOptions{})
			if err != nil {
				fmt.Printf("Error removing image %s: %v\n", imageID, err)
				continue
			}
			fmt.Printf("Image removed: %s\n", imageID)
		}
	},
}

func init() {
	rootCmd.AddCommand(rmiCmd)
}
