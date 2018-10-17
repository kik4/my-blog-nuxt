title: Nuxt.js でポートフォリオサイトを作成した記録

# 端書

これは Nuxt.js を使ってポートフォリオサイトを作ったのを GitHub の履歴を見返しながらメモ的に書き起こしたものです。
割と詳細を忘れていてきちんと作業日誌付けておけばよかったと後悔…。とはいえ詰まった部分をかなり思い出したのでこれを読んだ方が同じ轍を踏まない＆アドバイスをいただければ幸いです。
ちなみにスキルは React/Redux と Vue.js で業務経験あり。Nuxt.js や Firebase は初めてという具合です。

# モチベーション

一介の Web エンジニアとしてこの度は転職活動を始めようと思い立ったのですが、自分の場合はプライベートの開発物が全く無かったのです。Qiita などの細々とあるアウトプットや職務経歴書だけではアピールが足りません。というわけで手っ取り早くアピールできる制作物としてまずはポートフォリオサイトを作ることにしました。

# 最終的な成果物

https://kik4.work

# GitHub にリポジトリを作る

2018/09/18
https://github.com/kik4/my-portfolio-nuxt

# 開発ソフトの準備

エディタは VSCode。VCS は Git。js のモジュール管理は yarn。それぞれインストール済でした。

# Nuxt.js のインストール

とりあえず[ドキュメント通りに](https://ja.nuxtjs.org/guide/installation)テンプレートを入れましょう。

## Vue CLI のインストール

```bash
yarn add @vue/cli @vue/cli-init
```

## テンプレートのインストール

```bash
vue init nuxt-community/starter-template <project-name>
```

## 依存関係のインストール

```bash
cd <project-name>
yarn install
```

## 試しに起動

```bash
yarn dev
```

http://localhost:3000 にアクセスして動作確認しました。

# 開発ツールの準備

[ESLint と Prettier](https://ja.nuxtjs.org/guide/development-tools/#eslint)の導入の仕方が書いてあったので実施。準備ばかりですが、開発のしやすさの向上は何より重要です。

```bash
yarn add babel-eslint eslint eslint-config-prettier eslint-loader eslint-plugin-vue eslint-plugin-prettier prettier
```

VSCode に ESLint のプラグインを入れ、 .eslintrc.js ファイルを追加して有効にします。
prettier プラグインを入れたり .prettierrc ファイルを追加したりしましたが、ESLint 側から設定するんですね。最終的に落ち着いた設定はこちら。

```js
module.exports = {
  root: true,
  env: {
    browser: true,
    node: true
  },
  parserOptions: {
    parser: "babel-eslint"
  },
  extends: [
    // https://github.com/vuejs/eslint-plugin-vue#priority-a-essential-error-prevention
    // consider switching to `plugin:vue/strongly-recommended` or `plugin:vue/recommended` for stricter rules.
    "plugin:vue/recommended",
    "plugin:prettier/recommended"
  ],
  // required to lint *.vue files
  plugins: ["vue"],
  // add your custom rules here
  rules: {
    "prettier/prettier": [
      "error",
      {
        semi: false,
        printWidth: 120,
        trailingComma: "all"
      }
    ]
  }
};
```

VSCode のプラグインで保存時に動作するように設定を追加。

```json
{
  "eslint.validate": [
    "javascript",
    {
      "language": "vue",
      "autoFix": true
    }
  ],
  "eslint.autoFixOnSave": true
}
```

nuxt.config.json もおすすめ通りに設定を追加。
以下の通り、package.json にも設定を追加しました。

```json
"scripts": {
    ...
  "lint": "eslint --ext .js,.vue --ignore-path .gitignore .",
  "lintfix": "eslint --fix --ext .js,.vue --ignore-path .gitignore .",
  "precommit": : "npm run lint",
}
```

# 自分でページを作成

- components/AppLogo.vue の削除
- reboot.css を導入
- nuxt.config.js にページ情報を入力
- layouts/default.vue と pages/index.vue の編集
- favicon.ico の用意

などなどやってようやく最初のページができました。

# 紆余曲折

9/19 - 24。
firebase function を導入しようとして一旦やめたり、テンプレートをいじったり、とりあえずボタンを押したら数字が増えるページを作ったりしていました。a タグで他のページへのリンク書いたら普通に画面遷移して、nuxt-link 使わなきゃ SPA にならないじゃんとかなったり。

# Firebase Hosting にホスト

[ガイド](https://ja.nuxtjs.org/guide/routing/#firebase-%E3%83%9B%E3%82%B9%E3%83%86%E3%82%A3%E3%83%B3%E3%82%B0%E5%90%91%E3%81%91%E3%81%AE%E5%AE%9F%E8%A3%85)があるのでその通りにしました。
404 ページ用に layouts/error.vue も追加。色々試行錯誤しつつ、デプロイまで達成。

```bash
yarn add firebase
firebase deploy
```

# CircleCI の導入

9/25。いちいち手作業でデプロイなんてやっていられなくて、GitHub にプッシュしたら自動でデプロイまで行ってほしいわけです。
node のバージョンが足りないとかコマンド足りないとか環境変数とか試行錯誤。
使い方はいくらでも出てくるのではしょりますがこの時点ではこれで落ち着きました。

https://github.com/kik4/my-portfolio-nuxt/commit/3387f30b53d318f6ccb573d06d077ec7fa536d66

この時点で数時間しか触っていない牛歩具合。

# ドメインの取得

お名前.com でドメインを取得。Firebase の画面の指示通りに設定をするもアクセスできない…のでそのまま放置。翌日にはアクセスできました。
SSL 対応は Firebase なのでこの時点で完了。便利ですね。

# 紆余曲折

9/25 - 27。
サイトの中身を拡充させたり、[google analytics 対応](https://ja.nuxtjs.org/faq/google-analytics/#google-%E3%82%A2%E3%83%8A%E3%83%AA%E3%83%86%E3%82%A3%E3%82%AF%E3%82%B9%E3%82%92%E4%BD%BF%E3%81%86%E3%81%AB%E3%81%AF-)したり、ディレクトリ構成を試行錯誤したり。Edge じゃ適用されないスタイルを使っているのを見つけてまあいいやと諦めたり。

# blog ページの追加

9/28。
blog という名前ですが、Qiita の記事を Qiita API で取得して一覧表示させるようにしました。
使用する API は[こちら](https://qiita.com/api/v2/docs#get-apiv2usersuser_iditems)。
ページを開く度に axios で取得します。Qiita API にはレート制限があって「認証していない状態では IP アドレスごとに 1 時間に 60 回まで」なので、実はこれだと割と簡単に上限に達してしまいます。ここの改善は後々。

https://github.com/kik4/my-portfolio-nuxt/commit/902d1e5d7db6fb9412459f3fcd712e43cf92e212

# OGP タグの設定

OGP タグの設定。少し後になって page 側でプロパティを変更しようとしてうまく行っていないことが発覚します。
[ここに書いてある通り](https://ja.nuxtjs.org/api/pages-head)なのですが、hid を指定しないと更新できないのでした。

# コンポーネントの分割

10/1。
nuxt.js でページの内容を拡充させてきたのですが、これまでコンポーネントを切っていません。せっかくなのでコンポーネントに分けて構造やスタイルを隠蔽してしまいました。

https://github.com/kik4/my-portfolio-nuxt/commit/fffa539b3e18cd226f473dac9019e49fe38d9f62

# css scoped

10/2。
ページで設定したスタイルがページ遷移後も残っていることが発覚。ページの css に scoped を付けたら解消されました。

# コンポーネントを自動でグローバルに登録

10/3。
Vue は色々便利ではあるんですが、コンポーネントを import してさらにそれを components として登録しないといけないのが面倒です。
よく使うコンポーネントはグローバルに自動で登録してしまいました。

https://github.com/kik4/my-portfolio-nuxt/commit/0e5372d81b585a1d10fbe5c756acbfa9554bc672

# sitemap の追加

SEO 的に大事な sitemap を追加します。sitemap は[モジュールが提供されている](https://github.com/nuxt-community/sitemap-module)のでそれを使用。

https://github.com/kik4/my-portfolio-nuxt/commit/91079ecd59c78daed5c99584fa41f63927387815

Google Seach Console にも登録します。

# Qiita の記事情報をキャッシュ

毎回 Qiita API を叩いて記事情報を取りに行っていられないので、一度取ったら Vuex の store に保存するようにしました。
Vuex の使い方は Redux の経験からなんとなくで把握。
この実装でもページリロードされると store が消えるのでレートに引っかかる可能性はあります。

https://github.com/kik4/my-portfolio-nuxt/commit/c6a69b9acc95b869bdc22295dcd7dee4b29787ad

ページで[fetch メソッド](https://ja.nuxtjs.org/api/pages-fetch/)を使うと読み込むまでページ遷移を待ってくれます。瞬時に切り替わると読み込んでから内容が差し込まれて画面がガクガクと動作してしまうのですが、これで綺麗に動いてくれます。

# GAE/Go と DataStore で Qiita の記事情報を保存

10/6。
GAE を使って Qiita の記事情報を Cloud DataStore に キャッシュするようにします。一時間毎にキャッシュを取得して、ポートフォリオサイトの方ではそのキャッシュを自作の API を通して参照します。
Firebase の DB といえば FireStore があるんですが、アクセスルールとか cron が無いとかで使うのに手間取りそうだったのでこの時点はやめました。リアルタイム性も要らないですしね。
GAE（と Go）を使ったことがあるので手っ取り早いという理由で技術選定。

https://github.com/kik4/my-portfolio-nuxt/commit/9cf1c02806e1eb242d523793b7f9bc80a0e7b1bf

DataStore を使うのは初めてでした。json をそのまま格納できなくて、色々調べることに。

- 配列をエンティティの基底にできない
- 配列のオブジェクトに配列は入れられない
- long string を入れるには index を外す

などの問題をどうにか乗り越えます。

```
dev_appserver.py .
```

が動かなくて.のところはフルパスで app.yaml を指定したりしました。

# GAE の CircleCI 対応

GAE の方のデプロイもいちいちコンソールからやっていられないので CircleCI に任せます。

https://github.com/kik4/my-portfolio-nuxt/commit/94b4b36c27ca61de9356e037b97b10c4d8354cf4

上記からも config.yml に色々修正を重ねました。
さらに GCP の方で App Engine Admin API(appengine.googleapis.com) を有効にして上げる必要がありました。一度コンソールからやればいけるのかもしれませんが、面倒だったので GCP のコンソールから直接有効化。（Marketplace にありました）

# GAE のコードのパッケージ分け

main.go 一つに収まる程度のコード量しかなかったのですがせっかくなので分割しておきます。
GOPATH の中で作業していなかったので物理フォルダを移動させたり、app.yaml のサブディレクトリにパッケージがあるとエラーになるのを忘れていてしばらく詰まったりしていました。

# 内容の整理

blog ページを articles ページに変更したり、コンテンツを整理しました。

# 終わりに

まだまだ開発を続けていきますが、ここまでで一旦は形になりました。
途中 PWA にしようとして思った通りに動かないので諦めました。次のステップはこれかもしれません。
ここまで色々ありましたが、終わってみれば爆裂に簡単に作れたので Nuxt.js 便利ですね！
