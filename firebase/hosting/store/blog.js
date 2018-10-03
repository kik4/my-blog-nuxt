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
    const response = await this.$axios
      .$get("https://qiita.com/api/v2/users/kik4/items", { timeout: 5000 })
      .catch(() => {})
    if (response && response.data) {
      commit("setArticles", response.data)
    } else {
      commit("setArticles", null)
    }
  },
}
