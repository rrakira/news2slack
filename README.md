# 概要
ニュースサイトからトップニュースのタイトルとURLを取得し、タイトルを翻訳した上でSlackに投稿するプログラム。
近年の国際情勢を踏まえて、ロシアとウクライナにおいて各国主要メディアがどのような報道をしているかを知るために作成しました。

# 使い方
使用前準備
1. Slackのアプリ設定
2. 投稿先のSlackワークスペースAPIとチャネルリンクを取得し、`rt.go`および`tsn.go`の`sendtoSlack`関数内該当箇所に貼り付け（APIやリンクの取得方法は[こちら](https://api.slack.com/apps)）
3. Google Cloud Translateの設定（設定方法は[こちら](https://cloud.google.com/translate/docs/setup?hl=ja)）
4. 翻訳先言語を指定（デフォルトでは英語）
5. 取得する記事の数を指定（デフォルトでは5）

ニュース取得先ウェブサイト
- `rt.go`を実行すると「Russia Today」（ロシアメディア）からニュースを取得することができます。
- `tsn.go`を実行すると「TSN」（ウクライナメディア）からニュースを取得することができます。

![scraping](https://user-images.githubusercontent.com/103403806/167711625-04d00114-097b-4a26-85d1-cfd24103ffd4.gif)
