package cmd

import (
	"context"
	"errors"
	"fmt"
	"github.com/rexposadas/go-openai-cli/util/require"
	"github.com/sashabaranov/go-openai"
	"io"

	"github.com/spf13/cobra"
)

var ChatGPTCmd = &cobra.Command{
	Use:   "ChatGPT",
	Short: "Run ChatGPT completion",
	Long: `
Example: 
	open-ai-cli ChatGPT

`,
	Run: func(cmd *cobra.Command, args []string) {

		token := require.APIToken()

		c := openai.NewClient(token)
		ctx := context.Background()

		req := openai.ChatCompletionRequest{
			Model:     openai.GPT3Dot5Turbo,
			MaxTokens: 20,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "Lorem ipsum",
				},
			},
			Stream: true,
		}
		stream, err := c.CreateChatCompletionStream(ctx, req)
		if err != nil {
			fmt.Printf("ChatCompletionStream error: %v\n", err)
			return
		}
		defer stream.Close()

		fmt.Printf("Stream response: ")
		for {
			response, err := stream.Recv()
			if errors.Is(err, io.EOF) {
				fmt.Println("\nStream finished")
				return
			}

			if err != nil {
				fmt.Printf("\nStream error: %v\n", err)
				return
			}

			fmt.Printf(response.Choices[0].Delta.Content)
		}
		fmt.Println("ChatGPT called")
	},
}

func init() {
	rootCmd.AddCommand(ChatGPTCmd)
}
