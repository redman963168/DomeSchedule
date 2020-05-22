package main

import (
	"errors"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

//Schedule タイトル、日付
type Schedule struct {
	title    string
	dateInfo string
}

//GetSchedule スケジュールの取得
func (dome *DomeEventPage) GetSchedule() ([]Schedule, error) {
	nowDate := time.Now()
	year := nowDate.Format("2006")
	month := nowDate.Format("1")
	url := dome.url
	url = strings.Replace(url, "%YEAR%", year, -1)
	url = strings.Replace(url, "%MONTH%", month, -1)
	schedules := []Schedule{}
	ExistedDate := true

	doc, err := goquery.NewDocument(url)
	if err != nil {
		return nil, err
	}
	//classが"event-box"のみ抽出
	doc.Find(".event-box").EachWithBreak(func(i int, s *goquery.Selection) bool {
		topArea := s.Find(".top")
		dateText := topArea.Find(".date.sp").Text()
		dateText = checkRegexp(`\d{4}年\d{2}月\d{2}日`, dateText)
		if dateText == "" {
			ExistedDate = false
			return false
		}
		//本日予定か確認
		if nowDate.Format("2006年01月02日") == dateText {
			title := topArea.Find("h2").Text()

			btmArea := s.Find(".btm")
			dateInfo := btmArea.Find(".date").Text()

			//スケジュールの取得
			var schedule Schedule
			schedule.title = title
			schedule.dateInfo = strings.Replace(dateInfo, "\n", "", -1)
			schedules = append(schedules, schedule)
		}
		return true
	})
	if !ExistedDate {
		return nil, errors.New("日付要素をを見つけられませんでした。")
	}
	if len(schedules) <= 0 {
		var noSche Schedule
		noSche.title = "本日" + nowDate.Format("1/2") + "のイベント予定はありません"
		noSche.dateInfo = ""
		schedules = append(schedules, noSche)
	}

	return schedules, nil
}

func checkRegexp(reg, str string) string {
	return regexp.MustCompile(reg).FindString(str)
}
