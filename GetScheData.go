package main

import (
	//"fmt"
	"fmt"
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
func GetSchedule() []Schedule {
	year := time.Now().Format("2006")
	month := time.Now().Format("1")
	url := "https://www.kyoceradome-osaka.jp/events/?yearId=" + year + "&monthId= " + month
	schedules := []Schedule{}
	nowDate := time.Now()

	doc, err := goquery.NewDocument(url)
	if err != nil {
		panic(err)
	}
	//classが"event-box"のみ抽出
	doc.Find(".event-box").Each(func(i int, s *goquery.Selection) {
		//
		topArea := s.Find(".top")
		dateText := topArea.Find(".date.sp").Text()
		dateText = checkRegexp(`\d{4}年\d{2}月\d{2}日`, dateText)
		//dateText = ""
		if dateText == "" {
			fmt.Errorf("can't find date")
		}
		//本日予定か確認
		if nowDate.Format("2006年01月02日") == dateText {
			title := topArea.Find("a").Text()

			btmArea := s.Find(".btm")
			dateInfo := btmArea.Find(".date").Text()
			//fmt.Println(strings.Replace(date.Text(), "\n", "", -1))

			//スケジュールの取得
			var schedule Schedule
			schedule.title = title
			schedule.dateInfo = strings.Replace(dateInfo, "\n", "", -1)
			schedules = append(schedules, schedule)
		}

		//fmt.Println(dateText)
		//fmt.Println(date)
		//fmt.Println(title.Text())

	})

	if len(schedules) <= 0 {
		var noSche Schedule
		noSche.title = "本日" + nowDate.Format("1/2") + "のイベント予定はありません"
		noSche.dateInfo = ""
		schedules = append(schedules, noSche)
	}

	return schedules
}

func checkRegexp(reg, str string) string {
	return regexp.MustCompile(reg).FindString(str)
}
