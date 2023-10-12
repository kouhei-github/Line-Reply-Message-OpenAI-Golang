package line

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/sashabaranov/go-openai"
	"net-http/myapp/utils"
	"net-http/myapp/utils/firebase_facade"
	"net/http"
	"os"
	"time"
)

// MessageHandler 関数はLINEからのメッセージを処理します
func MessageHandler(w http.ResponseWriter, r *http.Request) {

	// LINEチャネルのシークレットとアクセストークンを用いて新たなbotを作成します
	bot, err := linebot.New(
		os.Getenv("LINE_CHANNEL_SECRET"),
		os.Getenv("LINE_CHANNEL_ACCESS_TOKEN"),
	)
	var message map[string]string
	// リクエストの解析を試みます
	events, err := bot.ParseRequest(r)
	// もし解析に失敗した場合はエラーメッセージとともにステータス500を返します
	if err != nil {
		message = map[string]string{"status": "500", "message": err.Error() + "=> 解析に失敗しました"}
		fmt.Println(message)
		json.NewEncoder(w).Encode(message)
		return
	}

	// botからのメッセージのユーザーIDを取得します
	userLineId := events[0].Source.UserID
	var userInputMessage string
	// メッセージイベントが来たときの処理
	if events[0].Type == linebot.EventTypeMessage {
		// テキストメッセージが来たときの処理
		switch userInput := events[0].Message.(type) {
		case *linebot.TextMessage:
			userInputMessage = userInput.Text
		}
	}

	// Firebaseの設定を読み込みます
	ctx := context.Background()
	fireStore, err := firebase_facade.ConfigFireBase(ctx)
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(err.Error())
	}

	// LINEのプロファイル情報を取得します
	profileRes, err := bot.GetProfile(userLineId).Do()
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(err.Error())
	}

	// Firebaseへのメッセージを作成します
	fireBaseMessage := firebase_facade.Message{
		Text:      userInputMessage,
		CreatedAt: time.Now(),
		ChatUser: firebase_facade.User{
			Name: profileRes.DisplayName,
			Id:   userLineId, // TODO lineのユーザーID
		},
	}
	// メッセージをFirebaseに追加します
	err = fireStore.AddDoc(userLineId, fireBaseMessage)
	// TODO firebaseからデータを全て取得
	talks := fireStore.GetDocs(userLineId)
	var messages []openai.ChatCompletionMessage
	// Firebaseから取得した全てのトークをOpenAIのメッセージフォーマットに変換します
	for _, talk := range talks {
		messageRole := openai.ChatMessageRoleUser
		if talk.Name == "system" {
			messageRole = openai.ChatMessageRoleAssistant
		}
		chat := openai.ChatCompletionMessage{
			Role:    messageRole,
			Content: talk.Text,
		}
		messages = append(messages, chat)
	}

	// TODO openaiでの会話の生成
	answer, err := utils.SendToGPT(os.Getenv("OPENAI_API_KEY"), messages)
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(err.Error())
	}

	// OpenAIからの応答をFirebaseに保存します
	fireBaseMessage = firebase_facade.Message{
		Text:      answer,
		CreatedAt: time.Now(),
		ChatUser: firebase_facade.User{
			Name: "system",
		},
	}
	err = fireStore.AddDoc(userLineId, fireBaseMessage)

	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(err.Error())
	}

	// ユーザーへOpenAIからの応答をLINEを通して返信します
	textMsg := linebot.NewTextMessage(answer)
	// 返信送信時のエラーハンドリング
	res, err := bot.PushMessage(userLineId, textMsg).Do()
	if err != nil {
		message = map[string]string{"status": "500", "message": err.Error()}
		fmt.Println(message)
		json.NewEncoder(w).Encode(message)
		return
	}

	// 成功のレスポンスメッセージを出力します
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(map[string]string{"status": "500", "message": "成功しました: " + res.RequestID})
}
