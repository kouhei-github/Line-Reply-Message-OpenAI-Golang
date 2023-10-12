package firebase_facade

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"fmt"
	"google.golang.org/api/option"
)

type MyFireStore struct {
	App *firestore.Client
	Ctx context.Context
}

// ConfigFireBase 関数は Firebase へのコネクションを初期化し、 Firestore のクライアントを返します
func ConfigFireBase(ctx context.Context) (*MyFireStore, error) {
	// Firebaseの認証情報をjsonファイルからロードする
	sa := option.WithCredentialsFile("./zeus-firebase-admin.json")
	// プロジェクトIDを設定する
	conf := &firebase.Config{ProjectID: "zeus-line-talk"}
	// Firebaseアプリを初期化する
	app, err := firebase.NewApp(ctx, conf, sa)

	// Firebaseアプリの初期化時のエラーハンドリング
	if err != nil {
		fmt.Println("接続エラー。")
		return &MyFireStore{}, err
	}

	// Authクライアントの取得を試みる
	_, err = app.Firestore(ctx)
	// Authクライアント取得時のエラーハンドリング
	if err != nil {
		fmt.Println("error getting Auth client: \n", err)
		return &MyFireStore{}, err
	}

	// Firestoreクライアントの取得を試みる
	client, err := app.Firestore(ctx)
	// Firestoreクライアント取得時のエラーハンドリング
	if err != nil {
		return &MyFireStore{}, err
	}

	// Firestoreクライアントとコンテキストを持つ新しいMyFireStoreストラクチャを返す
	return &MyFireStore{App: client, Ctx: ctx}, nil
}
