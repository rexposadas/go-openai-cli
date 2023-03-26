package cmd

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"github.com/rexposadas/go-openai-cli/util"
	"github.com/rexposadas/go-openai-cli/util/require"
	"github.com/sashabaranov/go-openai"
	"image/png"
	"os"

	"github.com/spf13/cobra"
)

var dalle2Cmd = &cobra.Command{
	Use:   "dalle2",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		token := require.APIToken()
		c := openai.NewClient(token)
		ctx := context.Background()

		// Sample image by link
		reqUrl := openai.ImageRequest{
			Prompt:         util.DefaultImagePrompt,
			Size:           util.DefaultImageSize(),
			ResponseFormat: openai.CreateImageResponseFormatURL,
			N:              1,
		}

		respUrl, err := c.CreateImage(ctx, reqUrl)
		if err != nil {
			fmt.Printf("Image creation error: %v\n", err)
			return
		}
		fmt.Println(respUrl.Data[0].URL)

		// Example image as base64
		reqBase64 := openai.ImageRequest{
			Prompt:         util.DefaultImagePrompt,
			Size:           util.DefaultImageSize(),
			ResponseFormat: openai.CreateImageResponseFormatB64JSON,
			N:              1,
		}

		respBase64, err := c.CreateImage(ctx, reqBase64)
		if err != nil {
			fmt.Printf("Image creation error: %v\n", err)
			return
		}

		imgBytes, err := base64.StdEncoding.DecodeString(respBase64.Data[0].B64JSON)
		if err != nil {
			fmt.Printf("Base64 decode error: %v\n", err)
			return
		}

		r := bytes.NewReader(imgBytes)
		imgData, err := png.Decode(r)
		if err != nil {
			fmt.Printf("PNG decode error: %v\n", err)
			return
		}

		file, err := os.Create("example.png")
		if err != nil {
			fmt.Printf("File creation error: %v\n", err)
			return
		}
		defer file.Close()

		if err := png.Encode(file, imgData); err != nil {
			fmt.Printf("PNG encode error: %v\n", err)
			return
		}

		fmt.Println("The image was saved as example.png")
	},
}

func init() {
	rootCmd.AddCommand(dalle2Cmd)
}
