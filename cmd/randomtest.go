package cmd

import (
	"bufio"
	"context"
	"fmt"
	"github.com/rexposadas/go-openai-cli/util/require"
	"io"
	"log"
	"os"
	"strings"

	api "github.com/sashabaranov/go-openai"
	"github.com/spf13/cobra"
)

var randomtestCmd = &cobra.Command{
	Use:   "randomtest",
	Short: "run a random test",
	Long: `
Used for debugging and testing the go-openai library.

`,
	Run: func(cmd *cobra.Command, args []string) {
		token := require.APIToken()
		msg := "some random message"

		result := createImage(token, msg)
		log.Println(result)
	},
}

func createImage(key, msg string) error {

	config := api.DefaultConfig(key)
	client := api.NewClientWithConfig(config)
	ctx := context.Background()

	req := api.ImageRequest{
		ResponseFormat: api.CreateImageResponseFormatB64JSON,
	}
	req.Prompt = msg
	resp, err := client.CreateImage(ctx, req)
	if err != nil {
		return err
	}

	data := resp.Data[0].B64JSON
	if err := ioCopyURL(client, ctx, data); err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Created)

	return nil
}

func ioCopyURL(client *api.Client, ctx context.Context, content string) error {
	file, err := os.Create("copied.PNG")
	if err != nil {
		return fmt.Errorf("failed to create file %s", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	io.Copy(writer, strings.NewReader(content))

	writer.Flush()

	req := api.ImageEditRequest{
		Image:  file,
		Mask:   file,
		Prompt: "There is a turtle in the pool",
		N:      3,
		Size:   "1024x1024",
	}
	_, err = client.CreateEditImage(ctx, req)
	if err != nil {
		return err
	}
	return nil
}

func fetch(key, msg string) string {
	c := api.NewClient(key)
	ctx := context.Background()

	req := api.CompletionRequest{
		Model:     api.GPT3Ada,
		MaxTokens: 5,
		Prompt:    msg,
	}
	resp, err := c.CreateCompletion(ctx, req)
	if err != nil {
		fmt.Printf("Completion error: %v\n", err)
		return ""
	}
	fmt.Println(resp.Choices[0].Text)
	return resp.Choices[0].Text
}

func init() {
	rootCmd.AddCommand(randomtestCmd)
}
