package utils

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
)

// SendToGPTはOpenAI GPTにメッセージを送信します。
// 送信したメッセージの応答とエラーを返します。
func SendToGPT(token string, messages []openai.ChatCompletionMessage) (string, error) {
	// トークンを使用してOpenAIクライアントを作成
	client := openai.NewClient(token)
	// チャットコンプリーションリクエストを作成
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    openai.GPT3Dot5Turbo,
			Messages: messages,
		},
	)

	// エラーチェック
	if err != nil {
		// エラーメッセージを出力
		fmt.Printf("ChatCompletion error: %v\n", err)
		return "", err
	}

	// レスポンスからメッセージコンテンツを取り出して返す
	return resp.Choices[0].Message.Content, nil
}
