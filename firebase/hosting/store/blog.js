export const state = () => ({
  articles: [],
})

export const mutations = {
  setArticles(state, value) {
    state.articles = value
  },
}

export const actions = {
  async getArticles({ commit }) {
    const data = await this.$axios.$get("https://qiita.com/api/v2/users/kik4/items", { timeout: 5000 }).catch(() => {})
    if (data) {
      commit("setArticles", data)
    } else {
      commit("setArticles", null)
    }
  },
}
