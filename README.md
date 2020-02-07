# graduation-checker

卒業要件をチェックしたい！

## インストール

`$ go get -u github.com/nu50218/graduation-checker/cmd/graduation-checker`

## 使い方

1. [これ](https://gist.github.com/nu50218/bdbec2e4bc39a053a00ec9f72b4394d7)で成績のjsonファイルをゲットします。

2. `$ graduation-checker -target=[TARGET] credit.json`

こんな感じで出力されます。

```bash
[✗] 合計128単位 [条件] 合計で128単位取得すること。
[✔] 基礎セミナー
[✗] 英語 [条件] 「言語文化」として英語6単位を含むこと。
[✔] 第二外国語
[✔] 健康・スポーツ科学
[✗] 文系基礎・教養科目 [条件] 「文系基礎科目」及び「文系教養科目」6単位を含むこと。
[✔] 理系教養科目
[✔] 全学教養科目
[✗] 専門基礎科目 [条件] 30単位以上取得すること。
```

### TARGET

TARGETの部分はチェックする専攻を指定します。以下に対応しています。

例: `$ graduation-checker -target=suuri credit.json`

| 専攻        | TARGET      |
| --------- | ----------- |
| 自然情報学科    | `sizen`     |
| 数理情報系     | `suuri`     |
| 複雑システム系専攻 | `fukuzatsu` |
