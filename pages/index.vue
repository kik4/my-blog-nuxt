<template>
  <div class="container">
    <h1 class="page_title">ポートフォリオ</h1>
    <div>
      <h2>プロフィール</h2>
      <ul class="no_list profile_list">
        <li>
          <dl>
            <dt>名前</dt>
            <dd>kik4</dd>
          </dl>
        </li>
        <li>
          <dl>
            <dt>職業</dt>
            <dd>Webエンジニア</dd>
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
            <dt>主なスキル</dt>
            <dd>
              <ul 
                class="tags" 
                style="padding: 0;">
                <li>C#</li>
                <li>ASP.NET</li>
                <li>React+Redux</li>
                <li>Vue.js</li>
                <li>JavaScript/TypeScript</li>
                <li>Laravel</li>
                <li>AWS</li>
                <li>SQL Server</li>
                <li>IIS</li>
                <li>golang</li>
              </ul>
            </dd>
          </dl>
        </li>
      </ul>
    </div>
    <div>
      <h2>リンク</h2>
      <ul class="no_list">
        <li><a href="https://twitter.com/_kik4_">Twitter (@_kik4_)</a></li>
        <li><a href="https://qiita.com/kik4">Qiita</a></li>
        <li><a href="https://github.com/kik4">Github</a></li>
        <li><a href="https://kik4.hatenablog.com/">ブログ</a></li>
      </ul>
    </div>
    <div>
      <h2>個人の制作物</h2>
      <ul class="no_list portfolio_list">
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
      <ul class="portfolio_list no_list">
        <li>FlashANZAN (<a href="https://www.vector.co.jp/soft/winnt/edu/se480770.html">Vector</a>)
        </li>
        <ul>
          <li>
            詳細
            <ul><li>フラッシュ暗算練習ソフト。何年も前に作ったのでハンドルも違いますしコードも残っていませんが、ちゃんと使えると思います。</li></ul>
          </li>
        </ul>
      </ul>
    </div>
  </div>
</template>

<script>
export default {
  head: {
    titleTemplate: "kik4"
  }
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
  padding-right: 5px;
}
.profile_list dt:after {
  padding-right: 5px;
  content: ":";
}

.no_list,
.no_list ul {
  list-style: none;
  padding-left: 20px;
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

