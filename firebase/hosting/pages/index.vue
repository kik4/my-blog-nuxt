<template>
  <div class="container">
    <h1 class="page_title">ポートフォリオ</h1>
    <div class="margin_bottom">
      <h2>プロフィール</h2>
      <IndentBox>
        <InlineInfoBox>
          <InfoTitle>名前</InfoTitle>
          <InfoBody>kik4</InfoBody>
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
          <InfoTitle>業務経験技術</InfoTitle>
          <InfoBody><TagList :list="['C#', 'JavaScript/TypeScript', 'ASP.NET Web Forms/MVC', 'Laravel', 'React+Redux', 'Vue.js', 'AWS', 'SQL Server', 'IIS', 'Jenkins']" /></InfoBody>
        </InlineInfoBox>
      </IndentBox>
    </div>
    <div class="margin_bottom">
      <h2>活動</h2>
      <IndentBox>
        <nuxt-link to="/blog">記事一覧</nuxt-link>
      </IndentBox>
    </div>
    <div class="margin_bottom">
      <h2>外部リンク</h2>
      <IndentBox>
        <InfoBody><a href="https://twitter.com/_kik4_">Twitter (@_kik4_)</a></InfoBody>
        <InfoBody><a href="https://qiita.com/kik4">Qiita</a></InfoBody>
        <InfoBody><a href="https://github.com/kik4">Github</a></InfoBody>
        <InfoBody><a href="https://kik4.hatenablog.com/">ブログ</a></InfoBody>
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
          <InfoBody>概要</InfoBody>
          <IndentBox>
            <InfoBody>自社プロダクト</InfoBody>
            <InfoBody>新規機能開発やらサーバ構築やらCI構築やらフロントエンドライブラリ技術選定・導入やらDB設計やら提案・実施。</InfoBody>
          </IndentBox>
        </IndentBox>
      </IndentBox>
      <IndentBox class="margin_bottom">
        <InfoTitle>求職サイト開発</InfoTitle>
        <IndentBox>
          <InfoBody>担当工程</InfoBody>
          <IndentBox>
            <TagList :list="['実装', 'テスト', '運用/保守', 'デザイン']" />
          </IndentBox>
          <InfoBody>概要</InfoBody>
          <IndentBox>
            <InfoBody>自社プロダクト</InfoBody>
            <InfoBody>主に新規機能開発やバグ修正。一部デザインも担当。</InfoBody>
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
            <TagList :list="['nuxt.js', 'Firebase Hosting', 'CircleCI', 'Github']" />
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
import TagList from "@/components/TagList.vue"
import IndentBox from "@/components/IndentBox.vue"
import InfoBody from "@/components/InfoBody.vue"
import InfoTitle from "@/components/InfoTitle.vue"
import InlineInfoBox from "@/components/InlineInfoBox.vue"
export default {
  components: { TagList, IndentBox, InfoBody, InfoTitle, InlineInfoBox },
  head() {
    return {
      titleTemplate: "kik4",
      meta: [{ hid: "og:title", property: "og:title", content: "kik4" }],
    }
  },
}
</script>

<style>
h2 {
  font-size: 1.3rem;
  padding-top: 3px;
  font-weight: bold;
}
.margin_bottom {
  margin-bottom: 12.5px;
}
</style>

