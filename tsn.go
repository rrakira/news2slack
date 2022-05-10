package main

import (
	"context"
	"fmt"
	"strconv"

	"cloud.google.com/go/translate"
	"github.com/PuerkitoBio/goquery"
	"github.com/slack-go/slack"
	"golang.org/x/text/language"
)

func main() {
	// 対象URLおよび取得する記事数を指定
	target_url := "https://tsn.ua/"
	num_articles := 3

	// 開始メッセージをSlackに投稿
	sendtoSlack("TSNによる最新のニュースを取得します。\n")

	// ページ情報取得
	doc, err := goquery.NewDocument(target_url)
	if err != nil {
		panic("htmlの取得に失敗しました")
	}

	// 各記事の情報を順に取得
	doc.Find("a.c-card__link").Each(func(i int, s *goquery.Selection) {
		// 指定した記事数に達したらメッセージを表示して終了
		if i >= num_articles {
			finished_message := strconv.Itoa(num_articles) + "件を表示しました。"
			sendtoSlack(finished_message + "全ニュースは下記リンクより確認。" + "\n" + target_url)
			panic("showed all")
		}

		// 要素を抽出
		article_title, _ := translateText("en", s.Text())
		href, _ := s.Attr("href")
		article_url := href

		// Slackに投稿
		sendtoSlack(article_title + "\n" + article_url)
	})
}

// Slackに投稿する関数
func sendtoSlack(text string) string {

	api := slack.New("YOUR_SLACK_API_HERE")

	_, _, err := api.PostMessage(
		"YOUR_SLACK_CHANNEL_LINK_HERE",
		slack.MsgOptionText(text, false),
		slack.MsgOptionAsUser(true),
	)
	if err != nil {
		fmt.Printf("%s\n", err)
		return ""
	}
	return ""
}

// Google Cloud Translateによる翻訳
func translateText(targetLanguage, text string) (string, error) {
	ctx := context.Background()

	lang, err := language.Parse(targetLanguage)
	if err != nil {
		return "", fmt.Errorf("language.Parse: %v", err)
	}

	client, err := translate.NewClient(ctx)
	if err != nil {
		return "", err
	}
	defer client.Close()

	resp, err := client.Translate(ctx, []string{text}, lang, nil)
	if err != nil {
		return "", fmt.Errorf("Translate: %v", err)
	}
	if len(resp) == 0 {
		return "", fmt.Errorf("Translate returned empty response to text: %s", text)
	}
	return resp[0].Text, nil
}
