package main

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

//PostMsg Slackにメッセージを送る
func (slack *SlackParameter) PostMsg(schedules []Schedule) error {
	// メッセージ用に結合
	msg := "【本日の京セラドームの予定】"
	for _, sche := range schedules {
		msg = msg + "\n" + sche.title
		msg = msg + "\n" + sche.dateInfo
	}

	//Slackメッセージ送信リクエスト作成
	jsonText := `{"channel":"` + slack.Channel + `","text":"` + msg + `"}`

	req, err := http.NewRequest(http.MethodPost, slack.URL, strings.NewReader(jsonText))
	if err != nil {
		return err
	}
	//ヘッダーの追加
	req.Header.Set("Content-type", "application/json;charset=utf-8")
	req.Header.Add("Authorization", "Bearer "+slack.Token)

	//リクエストの送信
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	//レスポンスの確認
	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return errors.New(string(body))
	}
	return nil
}
