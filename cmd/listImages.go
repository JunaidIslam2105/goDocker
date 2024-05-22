package cmd

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
	"time"
)

var listImagesCmd = &cobra.Command{
	Use:   "listImages",
	Short: "Lists all the docker images on the system.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())

		if err != nil {
			panic(err)
		}
		images, err := cli.ImageList(ctx, types.ImageListOptions{})
		if err != nil {
			panic(err)
		}
		fmt.Printf("%-20s %-20s %-20s %-20s %-20s %-20s\n", "REPOSITORY", "TAG", "IMAGE ID", "CREATED", "SIZE", "CONTAINERS")
		for _, image := range images {
			for _, tag := range image.RepoTags {
				repo, tag := parseRepoTag(tag)
				sizeMB := strconv.Itoa(int(float64(image.Size/1024/1024))) + " MB"
				createdDays := strconv.Itoa(int(float64(time.Now().Unix()-image.Created)/86400)) + " Days ago"
				fmt.Printf("%-20s %-20s %-20s %-20s %-20s %-20d\n", repo, tag, image.ID[7:19], createdDays, sizeMB, image.Containers)
			}
		}
	},
}

func parseRepoTag(repoTag string) (repo string, tag string) {
	i := strings.LastIndex(repoTag, ":")
	if i == -1 {
		return repoTag, ""
	}
	return repoTag[:i], repoTag[i+1:]
}

func init() {
	rootCmd.AddCommand(listImagesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listImagesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listImagesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
