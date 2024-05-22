package cmd

import (
	"context"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"io"
	"os"

	"github.com/spf13/cobra"
)

var logsCmd = &cobra.Command{
	Use:   "logs",
	Short: "Print out logs of a given container ID",
	Long:  `This command will print out the logs of a given container ID.`,
	Run: func(cmd *cobra.Command, args []string) {
		containerId := args[0]
		ctx := context.Background()
		cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
		if err != nil {
			panic(err)
		}
		defer cli.Close()

		options := container.LogsOptions{ShowStdout: true}

		out, err := cli.ContainerLogs(ctx, containerId, options)
		if err != nil {
			panic(err)
		}

		io.Copy(os.Stdout, out)
	},
}

func init() {
	rootCmd.AddCommand(logsCmd)
}
