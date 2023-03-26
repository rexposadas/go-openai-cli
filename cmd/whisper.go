package cmd

import (
	"context"
	"fmt"

	"github.com/rexposadas/go-openai-cli/util/require"
	"github.com/sashabaranov/go-openai"
	"github.com/spf13/cobra"
)

var whisperCmd = &cobra.Command{
	Use:   "whisper",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		token := require.APIToken()
		filePath := require.File(file)

		c := openai.NewClient(token)
		ctx := context.Background()

		req := openai.AudioRequest{
			Model: openai.Whisper1,
			//FilePath: "tmp/sample.mp3",
			FilePath: filePath,
		}
		resp, err := c.CreateTranscription(ctx, req)
		if err != nil {
			fmt.Printf("transcription error: %v", err)
			return
		}
		fmt.Println(resp.Text)
	},
}

func init() {
	rootCmd.AddCommand(whisperCmd)
}
