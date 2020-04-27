# DomeSchedule

ドームのイベント情報を取得して任意のSlackのチャンネルに投稿するアプリ

## 事前準備

実行前に以下の事前準備が必要

### Slackアプリの追加

1. SlackAppを任意のWorkspaceに新規作成<https://api.slack.com/apps>
2. OAuth & PermissionsタブからScopesに"chat:write"を追加
3. 上部の"Install App To Workspace"をクリック
4. 表示された"Bot User OAuth Access Token"※1をコピーしておく
5. Slackの画面でイベント情報を通知したいチャンネルにアプリを追加

### 環境変数の追加

- "SLACK_BOT_TOKEN"
  - Slackのトークン
  - ※1でコピーしたもの
- "DOME_CHANNEL"
  - イベント情報を投稿するチャンネル
