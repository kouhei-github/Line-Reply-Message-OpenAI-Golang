package firebase_facade

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id   string
	Name string
}

type Message struct {
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
	ChatUser  User      `json:"chat_user"`
}

// AddDoc 関数は "users" コレクションに新しいドキュメントを追加します
func (store *MyFireStore) AddDoc(target string, insertDocs Message) error {
	// "users" コレクション内の指定されたドキュメントIDのドキュメントに対して Set オペレーションを行い、 insertDocs の内容を設定します
	// ドキュメントが既に存在する場合、その内容は insertDocs によって置き換えられます
	u, err := uuid.NewRandom()
	if err != nil {
		return nil
	}
	_, err = store.App.Collection(
		"users").Doc(target).Collection(
		"messages").Doc(u.String()).Set(store.Ctx, insertDocs)

	// Set オペレーション中にエラーが発生した場合、そのエラーを返します
	if err != nil {
		return err
	}

	// エラーがなければ nil を返します
	return nil
}
