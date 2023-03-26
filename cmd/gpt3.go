package cmd

import (
	"context"
	"fmt"
	"github.com/rexposadas/go-openai-cli/util/require"
	"github.com/sashabaranov/go-openai"

	"github.com/spf13/cobra"
)

// gpt3Cmd represents the gpt3 command
var gpt3Cmd = &cobra.Command{
	Use:   "gpt-3",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		token := require.APIToken()

		c := openai.NewClient(token)
		ctx := context.Background()

		req := openai.CompletionRequest{
			Model:     openai.GPT3Ada,
			MaxTokens: 5,
			Prompt:    "some prompt",
		}
		resp, err := c.CreateCompletion(ctx, req)
		if err != nil {
			fmt.Printf("Completion error: %v\n", err)
			return
		}
		fmt.Println(resp.Choices[0].Text)
	},
}

func init() {
	rootCmd.AddCommand(gpt3Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// gpt3Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// gpt3Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
