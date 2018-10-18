<template>
  <div class="container">
    <h1 class="page_title">ポートフォリオ</h1>
    <div class="margin_bottom">
      <h2>プロフィール</h2>
      <IndentBox>
        <InlineInfoBox>
          <InfoTitle>名前</InfoTitle>
          <InfoBody>kik4（きっくふぉー）[kíkfˈɔː]</InfoBody>
        </InlineInfoBox>
        <InlineInfoBox>
          <InfoTitle>職業</InfoTitle>
          <InfoBody>Webエンジニア（フルスタック）</InfoBody>
        </InlineInfoBox>
        <InlineInfoBox>
          <InfoTitle>業務経歴</InfoTitle>
          <InfoBody>ECサイト開発, 求職サイト開発</InfoBody>
        </InlineInfoBox>
        <InlineInfoBox>
          <InfoTitle>エディタ</InfoTitle>
          <InfoBody><TagList :list="['Visual Studio Code']" /></InfoBody>
        </InlineInfoBox>
        <InlineInfoBox>
          <InfoTitle>主な業務経験技術</InfoTitle>
          <InfoBody><TagList :list="['C#', 'JavaScript/TypeScript', 'ASP.NET Web Forms/MVC', 'PHP', 'Laravel', 'React+Redux', 'Vue.js', 'AWS', 'SQL Server', 'IIS', 'Jenkins']" /></InfoBody>
        </InlineInfoBox>
        <InlineInfoBox>
          <InfoTitle>プライベートで使う技術</InfoTitle>
          <InfoBody><TagList :list="['Go', 'GCP', 'DDD', 'PWA', 'SPA', 'サーバレス']" /></InfoBody>
        </InlineInfoBox>
      </IndentBox>
    </div>
    <div class="margin_bottom">
      <h2>業務経験</h2>
      <IndentBox class="margin_bottom">
        <InfoTitle>ECサイト開発</InfoTitle>
        <IndentBox>
          <InfoBody>担当工程</InfoBody>
          <IndentBox>
            <TagList :list="['要件定義', '設計', '実装', 'テスト', '運用/保守']" />
          </IndentBox>
          <InfoBody>技術</InfoBody>
          <IndentBox>
            <TagList :list="['C#', 'ASP.NET Web Forms/MVC', 'JavaScript', 'React+Redux', 'Vue.js', 'jQuery', 'AWS', 'SQL Server', 'IIS', 'Jenkins', 'GitBucket']" />
          </IndentBox>
          <InfoBody>概要</InfoBody>
          <IndentBox>
            <InfoBody>自社プロダクト</InfoBody>
            <InfoBody>フルスタックで何でも担当。比較的技術寄りの部分を解決。主な作業は以下の通り。</InfoBody>
            <InfoBody>
              <ul>
                <li>フロントエンド/バックエンド新規機能開発</li>
                <li>AWS EC2サーバ構築</li>
                <li>CI環境構築</li>
                <li>React/Vue.jsなどライブラリ技術選定</li>
                <li>DB設計</li>
                <li>ツール作成</li>
                <li>コードレビュー</li>
                <li>ASP.NET Web FormsからMVCへの順次移行提案・実施</li>
                <li>CSVデータ出力カスタムDSLの設計・実行システムの実装</li>
              </ul>
            </InfoBody>
          </IndentBox>
        </IndentBox>
      </IndentBox>
      <IndentBox class="margin_bottom">
        <InfoTitle>求職サイト開発</InfoTitle>
        <IndentBox>
          <InfoBody>担当工程</InfoBody>
          <IndentBox>
            <TagList :list="['設計', '実装', 'テスト', '運用/保守', 'デザイン']" />
          </IndentBox>
          <InfoBody>技術</InfoBody>
          <IndentBox>
            <TagList :list="['PHP', 'Laravel', 'JavaScript', 'React', 'jQuery', 'Apache', 'nginx', 'MariaDB', 'BootStrap', 'BitBucket']" />
          </IndentBox>
          <InfoBody>概要</InfoBody>
          <IndentBox>
            <InfoBody>自社プロダクト</InfoBody>
            <InfoBody>主に新規機能開発やバグ修正。一部デザインも担当。その他、vagrantを用いた環境構築など。</InfoBody>
          </IndentBox>
        </IndentBox>
      </IndentBox>
    </div>
    <div class="margin_bottom">
      <h2>個人の制作物</h2>
      <IndentBox class="margin_bottom">
        <InfoTitle>ポートフォリオサイト</InfoTitle>
        <IndentBox>
          <InfoBody>技術</InfoBody>
          <IndentBox>
            <TagList :list="['nuxt.js', 'Firebase Hosting', 'CircleCI', 'Github', 'Go', 'GAE', 'Cloud DataStore']" />
          </IndentBox>
          <InfoBody>詳細</InfoBody>
          <IndentBox>
            このサイトです。コードは公開中（<a href="https://github.com/kik4/my-portfolio-nuxt">Github</a>）。
          </IndentBox>
        </IndentBox>
      </IndentBox>
      <IndentBox class="margin_bottom">
        <InfoTitle>FlashANZAN (<a href="https://www.vector.co.jp/soft/winnt/edu/se480770.html">Vector</a>)</InfoTitle>
        <IndentBox>
          <InfoBody>詳細</InfoBody>
          <IndentBox>
            Windows用フラッシュ暗算練習ソフト。何年も前に作ったのでハンドルが違います。コードも残っていませんが、ちゃんと使えると思います。
          </IndentBox>
        </IndentBox>
      </IndentBox>
    </div>
  </div>
</template>

<script>
export default {
  head() {
    return {
      titleTemplate: "kik4",
      meta: [{ hid: "og:title", property: "og:title", content: "kik4" }],
    }
  },
}
</script>

<style scoped>
h2 {
  font-size: 1.3rem;
  padding-top: 3px;
  font-weight: bold;
}
.margin_bottom {
  margin-bottom: 12.5px;
}
</style>
