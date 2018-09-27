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
      { hid: "description", name: "description", content: "kik4's page" }
    ],
    link: [
      { rel: "icon", type: "image/x-icon", href: "/favicon.ico" },
      { href: "https://fonts.googleapis.com/css?family=Noto+Sans+JP&amp;subset=japanese", rel: "stylesheet" }
    ]
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
          exclude: /(node_modules)/
        })
      }
    }
  },
  //
  css: ["@/assets/css/reboot.css"],
  generate: {
    fallback: true
  },
  plugins: [{ src: "~plugins/ga.js", ssr: false }]
}
