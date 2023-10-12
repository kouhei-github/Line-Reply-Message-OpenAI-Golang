package firebase_facade

import (
	"cloud.google.com/go/firestore"
	"errors"
	"google.golang.org/api/iterator"
	"log"
	"time"
)

// Talk はチャットメッセージの内容と作成日時を保持する構造体です
type Talk struct {
	Name      string
	Text      string
	CreatedAt time.Time
}

// GetDocs 関数は "users" コレクションから指定したユーザーのドキュメントを取得します
// また、3日以前のメッセージは削除します
func (store *MyFireStore) GetDocs(target string) []Talk {
	// アプリケーションのユーザーコレクションに対するイテレータを作成
	iter := store.App.Collection("users").Doc(target).Collection("messages").OrderBy("CreatedAt", firestore.Desc).Documents(store.Ctx)
	var result []Talk

	for {
		// ドキュメントを順番に取り出す
		doc, err := iter.Next()
		if errors.Is(err, iterator.Done) {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}

		talk := castArrayString(doc.Data())                   // ドキュメントのデータをTalkに変換
		fiveMinutesBefore := time.Now().Add(-5 * time.Minute) // 5分前

		compare := time.Date(talk.CreatedAt.Year(), talk.CreatedAt.Month(), talk.CreatedAt.Day(), 0, 0, 0, 0, talk.CreatedAt.Location())
		fiveMinutesBefore = time.Date(fiveMinutesBefore.Year(), fiveMinutesBefore.Month(), fiveMinutesBefore.Day(), 0, 0, 0, 0, fiveMinutesBefore.Location())
		// メッセージの作成日が5分前であれば削除
		if compare.Before(fiveMinutesBefore) {
			store.App.Collection("users").Doc(target).Collection("messages").Doc(doc.Ref.ID).Delete(store.Ctx)
		} else {
			// 結果のスライスに追加
			result = append(result, talk)
		}
	}
	return result
}

// castArrayString 関数は map を Talk に変換します
func castArrayString(userMessage map[string]interface{}) Talk {
	var talk Talk
	// ドキュメントのデータを整形し、Talk 構造体に格納
	for key, value := range userMessage {
		switch key {
		case "Text":
			talk.Text = value.(string)
		case "ChatUser":
			result, ok := value.(map[string]interface{})
			if ok {
				name, exists := result["Name"]
				if exists {
					talk.Name = name.(string)
				}
			}
		case "CreatedAt":
			createdAt, ok := value.(time.Time)
			if ok {
				talk.CreatedAt = createdAt
			}
		}

	}
	return talk
}
