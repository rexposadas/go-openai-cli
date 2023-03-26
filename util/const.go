package util

import "github.com/sashabaranov/go-openai"

const (
	DefaultImagePrompt = "half human half machine, high detail, realistic light, unreal engine"
)

func DefaultImageSize() string {
	return openai.CreateImageSize256x256
}
