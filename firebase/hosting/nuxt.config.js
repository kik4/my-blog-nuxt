module.exports = {
  /*
  ** Headers of the page
  */
  head: {
    title: "kik4's page",
    titleTemplate: "%s | kik4",
    meta: [
      { charset: "utf-8" },
      { name: "viewport", content: "width=device-width, initial-scale=1" },
      { hid: "description", name: "description", content: "kik4's page: The engineer's portfolio, blog and others." },
      { property: "og:title", content: "kik4" },
      { property: "og:type", content: "website" },
      { property: "og:image", content: (process.env.BASE_URL || "http://localhost:3000") + "/icon-512x512.png" },
      { property: "og:site_name", content: "kik4" },
      { property: "og:description", content: "kik4" },
      { name: "twitter:card", content: "summary" },
    ],
    link: [
      { rel: "icon", type: "image/x-icon", href: "/favicon.ico" },
      { rel: "icon", size: "192x192", href: "/icon-192x192.png" },
      { rel: "apple-touch-icon", href: "/apple-touch-icon.png" },
      { rel: "manifest", href: "/manifest.json" },
      { rel: "stylesheet", href: "https://fonts.googleapis.com/css?family=Noto+Sans+JP&amp;subset=japanese" },
    ],
  },
  /*
  ** Customize the progress bar color
  */
  loading: { color: "#3B8070" },
  /*
  ** Build configuration
  */
  build: {
    /*
    ** Run ESLint on save
    */
    extend(config, { isDev, isClient }) {
      if (isDev && isClient) {
        config.module.rules.push({
          enforce: "pre",
          test: /\.(js|vue)$/,
          loader: "eslint-loader",
          exclude: /(node_modules)/,
        })
      }
    },
  },
  //
  css: ["@/assets/css/reboot.css"],
  generate: {
    fallback: true,
  },
  plugins: [{ src: "~plugins/ga.js", ssr: false }],
  env: {
    baseUrl: process.env.BASE_URL || "http://localhost:3000",
  },
}
