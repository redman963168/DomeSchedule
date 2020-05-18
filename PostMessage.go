package main

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

//PostMsg Slackにメッセージを送る
func PostMsg(schedules []Schedule) error {
	// メッセージ用に結合
	msg := "【本日の京セラドームの予定】"
	for _, sche := range schedules {
		msg = msg + "\n" + sche.title
		msg = msg + "\n" + sche.dateInfo
	}

	//Slackメッセージ送信リクエスト作成
	slackURL := "https://slack.com/api/chat.postMessage"
	values := url.Values{}
	values.Set("token", slackAPIToken)
	values.Add("channel", slackChannel)
	values.Add("text", msg)

	req, err := http.NewRequest(http.MethodPost, slackURL, strings.NewReader(values.Encode()))
	if err != nil {
		return err
	}
	//ヘッダーの追加
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	//リクエストの送信
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	//レスポンスの確認
	if resp.StatusCode != 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		return errors.New(string(body))
	}
	return nil
}
