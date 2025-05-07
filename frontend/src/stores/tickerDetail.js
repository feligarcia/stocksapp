import { defineStore } from 'pinia'

export const useTickerDetailStore = defineStore('tickerDetail', {
  state: () => ({
    companyInfo: {}, 
    quotes: {},      
    recommendations: {}, 
    news: {} 
  }),
  actions: {
    setCompanyInfo(ticker, data) {
      this.companyInfo[ticker] = { data, lastFetched: Date.now() }
    },
    setQuote(ticker, data) {
      this.quotes[ticker] = { data, lastFetched: Date.now() }
    },
    setRecommendations(ticker, data) {
      this.recommendations[ticker] = { data, lastFetched: Date.now() }
    },
    clearCompanyInfo(ticker) {
      delete this.companyInfo[ticker]
    },
    clearQuote(ticker) {
      delete this.quotes[ticker]
    },
    clearRecommendations(ticker) {
      delete this.recommendations[ticker]
    },
    setNews(ticker, data) {
      this.news[ticker] = { data, lastFetched: Date.now() }
    },
    clearNews(ticker) {
      delete this.news[ticker]
    }
  }
})
