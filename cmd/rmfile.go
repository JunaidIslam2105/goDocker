package cmd

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/spf13/cobra"
	"io"
	"os"
)

var rmfileCmd = &cobra.Command{
	Use:   "rmfile",
	Short: "Remove a file from a container",
	Long:  `This command will remove a file from a container given the container ID and the file path.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Println("Error: Container ID and file path must be provided")
			os.Exit(1)
		}

		containerID := args[0]
		filePath := args[1]

		ctx := context.Background()
		cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
		if err != nil {
			panic(err)
		}
		defer cli.Close()

		execConfig := types.ExecConfig{
			AttachStdout: true,
			AttachStderr: true,
			Cmd:          []string{"rm", filePath},
		}

		response, err := cli.ContainerExecCreate(ctx, containerID, execConfig)
		if err != nil {
			panic(err)
		}

		attachOptions := types.ExecStartCheck{
			Detach: false,
			Tty:    false,
		}

		hijackedResponse, err := cli.ContainerExecAttach(ctx, response.ID, attachOptions)
		if err != nil {
			panic(err)
		}
		defer hijackedResponse.Close()

		_, _ = io.Copy(os.Stdout, hijackedResponse.Reader)
	},
}

func init() {
	rootCmd.AddCommand(rmfileCmd)
}
