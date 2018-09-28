<template>
  <div class="container">
    <h1 class="page_title">ポートフォリオ</h1>
    <div>
      <h2>プロフィール</h2>
      <ul class="profile_list no_list_dot">
        <li>
          <dl>
            <dt>名前</dt>
            <dd>kik4</dd>
          </dl>
        </li>
        <li>
          <dl>
            <dt>職業</dt>
            <dd>Webエンジニア（フルスタック）</dd>
          </dl>
        </li>
        <li>
          <dl>
            <dt>業務経歴</dt>
            <dd>ECサイト開発, 求職サイト開発</dd>
          </dl>
        </li>
        <li>
          <dl>
            <dt>エディタ</dt>
            <dd>
              <ul class="tags no_list_padding">
                <li>Visual Studio Code</li>
              </ul>
            </dd>
          </dl>
        </li>
        <li>
          <dl>
            <dt>業務経験技術</dt>
            <dd>
              <ul class="tags no_list_padding">
                <li>C#</li>
                <li>JavaScript/TypeScript</li>
                <li>ASP.NET Web Forms/MVC</li>
                <li>Laravel</li>
                <li>React+Redux</li>
                <li>Vue.js</li>
                <li>AWS</li>
                <li>SQL Server</li>
                <li>IIS</li>
                <li>Jenkins</li>
              </ul>
            </dd>
          </dl>
        </li>
      </ul>
    </div>
    <div>
      <h2>活動</h2>
      <ul class="no_list_dot">
        <li><a href="https://twitter.com/_kik4_">Twitter (@_kik4_)</a></li>
        <li><a href="https://qiita.com/kik4">Qiita</a></li>
        <li><a href="https://github.com/kik4">Github</a></li>
        <li><a href="https://kik4.hatenablog.com/">ブログ</a></li>
      </ul>
    </div>
    <div>
      <h2>業務経験</h2>
      <ul class="portfolio_list no_list_dot">
        <li>ECサイト開発</li>
        <ul>
          <li>
            担当工程
            <ul class="tags">
              <li>要件定義</li>
              <li>設計</li>
              <li>実装</li>
              <li>テスト</li>
              <li>運用/保守</li>
            </ul>
          </li>
          <li>
            概要
            <ul>
              <li>自社プロダクト</li>
              <li>新規機能開発やらサーバ構築やらCI構築やらフロントエンドライブラリ技術選定・導入やらDB設計やら何でもやりました。</li>
            </ul>
          </li>
        </ul>
      </ul>
      <ul class="portfolio_list no_list_dot">
        <li>求職サイト開発</li>
        <ul>
          <li>
            担当工程
            <ul class="tags">
              <li>実装</li>
              <li>テスト</li>
              <li>運用/保守</li>
              <li>デザイン</li>
            </ul>
          </li>
          <li>
            概要
            <ul>
              <li>自社プロダクト</li>
              <li>主に新規機能開発やバグ修正。一部デザインも担当。</li>
            </ul>
          </li>
        </ul>
      </ul>
    </div>
    <div>
      <h2>個人の制作物</h2>
      <ul class="portfolio_list no_list_dot">
        <li>ポートフォリオサイト</li>
        <ul>
          <li>
            技術
            <ul class="tags">
              <li>nuxt.js</li>
              <li>Firebase Hosting</li>
              <li>CircleCI</li>
              <li>Github</li>
            </ul>
          </li>
          <li>
            詳細
            <ul>
              <li>このサイトです。コードは公開中（<a href="https://github.com/kik4/my-portfolio-nuxt">Github</a>）。</li>
            </ul>
          </li>
        </ul>
      </ul>
      <ul class="portfolio_list no_list_dot">
        <li>FlashANZAN (<a href="https://www.vector.co.jp/soft/winnt/edu/se480770.html">Vector</a>)
        </li>
        <ul>
          <li>
            詳細
            <ul><li>Windows用フラッシュ暗算練習ソフト。何年も前に作ったのでハンドルが違います。コードも残っていませんが、ちゃんと使えると思います。</li></ul>
          </li>
        </ul>
      </ul>
    </div>
  </div>
</template>

<script>
export default {
  head: {
    titleTemplate: "kik4",
  },
}
</script>

<style>
h2 {
  font-size: 1.3rem;
  padding-top: 3px;
}

.profile_list dl {
  display: flex;
  flex-wrap: wrap;
  margin: 0;
}
.profile_list dt {
  padding-right: 10px;
}

.no_list_dot,
.no_list_dot ul {
  list-style: none;
}

ul {
  padding-left: 12px;
}
@media (min-width: 600px) {
  ul {
    padding-left: 20px;
  }
}
.no_list_padding {
  padding: 0;
}

.portfolio_list li {
  font-weight: normal;
}
.portfolio_list > li {
  font-weight: bold;
}

.tags {
  display: flex;
  flex-wrap: wrap;
}
.tags > li {
  list-style: none;
  margin-right: 2px;
  white-space: nowrap;
  border: 1px solid #cccccc;
  border-radius: 5px;
  padding: 0 5px;
  background-color: #fff9f9f9;
}
</style>

