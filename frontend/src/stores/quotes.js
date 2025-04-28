import { defineStore } from 'pinia'

export const useQuotesStore = defineStore('quotes', {
  state: () => ({
    quotes: [],
    lastFetched: null // opcional: para cache TTL
  }),
  actions: {
    setQuotes(data) {
      this.quotes = data
      this.lastFetched = Date.now()
    },
    clearQuotes() {
      this.quotes = []
      this.lastFetched = null
    }
  }
})
