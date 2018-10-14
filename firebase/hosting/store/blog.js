export const state = () => ({
  articles: null,
})

export const mutations = {
  setArticles(state, value) {
    state.articles = value
  },
}

export const actions = {
  async getArticles({ commit }) {
    const data = await this.$axios.$get(process.env.GAE_URL + "/qiita/items", { timeout: 5000 }).catch(() => {})
    if (data) {
      commit("setArticles", data.Value)
    } else {
      commit("setArticles", null)
    }
  },
}
