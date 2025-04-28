import { defineStore } from 'pinia'

export const useRecommendationStore = defineStore('recommendation', {
  state: () => ({
    recommendations: []
  }),
  actions: {
    setRecommendations(recs) {
      this.recommendations = recs
    },
    clearRecommendations() {
      this.recommendations = []
    }
  }
})
