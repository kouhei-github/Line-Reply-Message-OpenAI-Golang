# Line × OpenAi API × Golang
質問やお悩み相談が来た際に、GPTエンジンを使用して、<br>
完璧な返答をするLINE公式アカウント用のバックエンドアプリを作成しました。

![385544301_628796422781454_6186934858729003709_n](https://github.com/kouhei-github/Line-Reply-Message-OpenAI-Golang/assets/49782052/2625b4fc-f7da-4d05-865f-297760a2d866)


---

## 1. 動かすために必要なこと

### 1.1 .envファイルの作成
```shell
cp .env.sample .env
```

・Lineでアカウントを作成して、チャンネルシークレットとアクセストークをセットする<br>
・OpenAIのAPIキーをセットする
```text
LINE_CHANNEL_SECRET=
LINE_CHANNEL_ACCESS_TOKEN=

OPENAI_API_KEY=
```

---

### 1.2 Lineのデベロッパー画面でwebhookを指定する
localhostだと動かないので、<br>
local出たてたwebサーバーを**ngrok**で外部からアクセスできるようにする必要がある

![スクリーンショット 2023-10-12 10 16 22](https://github.com/kouhei-github/Line-Reply-Message-OpenAI-Golang/assets/49782052/ddcf8ba8-7ff5-43f6-adcb-87cb4fcdc125)


---

## 2. 起動・停止方法
### 2.1 起動方法
imageの作成
```shell
docker compose build
```

imageからコンテナの起動
```shell
docker compose up -d
```

---

### 2.2 停止方法
```shell
docker compose down
```

---


## 3. 動作確認
下記から公式アカウントを追加して動作確認してください。

![スクリーンショット 2023-10-12 9 22 15](https://github.com/kouhei-github/Line-Reply-Message-OpenAI-Golang/assets/49782052/6352a4b0-a852-4936-9035-f175804cc8e3)

---

