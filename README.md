# App Review Slack Bot

## Overview

ユーザが投稿したスマートフォンアプリの新着レビューを
Slackに通知するバッチプログラム.  

*連携可能サービス 2019/05/04時点*

- レビュー収集先
  - [App Store CustomerReview RSS](https://www.apple.com/jp/rss/)

- 通知先
  - [Slack Incoming WebHooks integration](https://slack.com/apps/A0F7XDUAZ-incoming-webhooks)

## Requirements

- golang : version 1.11.x ~
- [dep](https://github.com/golang/dep) ( goの依存関係管理ツール)

## How to Use

### 環境変数の設定

#### レビュー最新収集日時ログファイルの作成
`cp files/.updated.log files/updated.log`

#### .env ファイルの作成
`cp .env.example .env`

#### 環境変数のセット

##### Slackに表示されるアプリ名
```
SLACK_BOT_NAME=AppReviewer
```

##### アイコン名（Slackのアイコン）
```
SLACK_BOT_ICON=:speech_balloon:
```

##### Slack Webhook URL
Slack側での設定方法やWebhookURLの取得方法については[こちら](https://slack.com/apps/A0F7XDUAZ-incoming-webhooks)を参照.  
```
SLACK_WEBHOOK_URL=https://hooks.slack.com/services/XXXX
```

##### Slack の通知チャンネル
```
SLACK_TARGET_CHANNEL=random
```

##### iOSアプリのRSSフィードURL
下記の`XXXX`の部分にはiOSアプリのアプリIDが入る.  
```
TARGET_IOS_APP_ID=XXXX
```

##### レビュー最新収集日時ログファイルパス
基本的にデフォルトのままでOK.  
```
UPDATED_LOG_PATH=./files/updated.log
```

### バッチの実行

```
# 依存関係の解決
dep ensure 

# ビルド
go build

# 実行
./app-review-slackbot
```

## License
This Application is open-sourced software licensed under the [MIT](https://opensource.org/licenses/MIT)

## Author

[shirobrak](https://github.com/shirobrak)