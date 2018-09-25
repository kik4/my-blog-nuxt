<template>
  <div class="container">
    <h1 class="page_title">ポートフォリオ</h1>
    <div>
      <h2>プロフィール</h2>
      <ul class="profile">
        <li><b>名前:</b>kik4</li>
        <li><b>職業:</b>Webエンジニア</li>
        <li><b>業務経歴:</b>ECサイト開発, 求職サイト開発</li>
        <li><b>技術:</b>C#, ASP.NET, React+Redux, Vue.js, TypeScript, Laravel, AWS, SQL Server, IIS, golang</li>
      </ul>
    </div>
    <div>
      <h2>リンク</h2>
      <ul class="profile">
        <li><a href="https://twitter.com/_kik4_">Twitter (@_kik4_)</a></li>
        <li><a href="https://qiita.com/kik4">Qiita</a></li>
        <li><a href="https://github.com/kik4">Github</a></li>
        <li><a href="https://kik4.hatenablog.com/">ブログ</a></li>
      </ul>
    </div>
    <div>
      <p><nuxt-link to="counter">カウンタへ</nuxt-link></p>
    </div>
    <div>
      Deployed by Ciecle CI
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
.profile b {
  padding-right: 5px;
}
li {
  padding: 2px 0;
}
</style>

