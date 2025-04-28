<template>
  <div class="bloomberg-table">
    <div class="main-content">
      <div class="back-btn" @click="goBack">
        <span class="arrow">&#8592;</span>
        <div class="volver">Volver</div>
      </div>
      <div class="header-box">
        <h1 class="title">{{ ticker }}<span v-if="company && company.name"> - {{ company.name }}</span></h1>
      </div>
      <div v-if="loading" class="loading"><span class="loading-yellow">Cargando...</span></div>
      <div v-else>
        <div class="company-card block-card">
          <h2 class="block-title">Company Info</h2>
          <div v-if="company">
            <div style="margin-bottom:10px;">
              <img v-if="company.logo" :src="company.logo" alt="Logo" style="width:80px;height:80px;object-fit:contain;border-radius:10px;border:1px solid #ddd;background:#fff;" />
            </div>
            <div><b>Name:</b> {{ company.name }}</div>
            <div><b>Industry:</b> {{ company.finnhub_industry || company.finnhubIndustry }}</div>
            <div><b>Country:</b> {{ company.country }}</div>
            <div><b>IPO:</b> {{ company.ipo }}</div>
            <div><b>Website:</b> <a :href="company.web || company.weburl" target="_blank">{{ company.web || company.weburl }}</a></div>
          </div>
          <div v-else>No company info found.</div>
        </div>
        <div class="quote-card block-card">
          <h2 class="block-title">Quote</h2>
          <div v-if="quote && Array.isArray(quote) && quote.length">
            <div>
              <div><b>Price:</b> <span :class="priceColor(quote[0].c, quote[0].pc)">{{ quote[0].c }}</span></div>
              <div><b>Change:</b> <span :class="quote[0].d > 0 ? 'pos' : quote[0].d < 0 ? 'neg' : ''">{{ formatChange(quote[0].d) }}</span></div>
              <div><b>% Change:</b> <span :class="quote[0].dp > 0 ? 'pos' : quote[0].dp < 0 ? 'neg' : ''">{{ formatChange(quote[0].dp, true) }}</span></div>
              <div><b>High:</b> {{ quote[0].h }}</div>
              <div><b>Low:</b> {{ quote[0].l }}</div>
              <div><b>Open:</b> {{ quote[0].o }}</div>
              <div><b>Prev Close:</b> {{ quote[0].pc }}</div>
            </div>
          </div>
          <div v-else-if="quote">
            <div><b>Price:</b> <span :class="priceColor(quote.c, quote.pc)">{{ quote.c }}</span></div>
            <div><b>Change:</b> <span :class="quote.d > 0 ? 'pos' : quote.d < 0 ? 'neg' : ''">{{ quote.d }}</span></div>
            <div><b>% Change:</b> <span :class="quote.dp > 0 ? 'pos' : quote.dp < 0 ? 'neg' : ''">{{ quote.dp }}%</span></div>
            <div><b>High:</b> {{ quote.h }}</div>
            <div><b>Low:</b> {{ quote.l }}</div>
            <div><b>Open:</b> {{ quote.o }}</div>
            <div><b>Prev Close:</b> {{ quote.pc }}</div>
          </div>
          <div v-else>No quote info found.</div>
        </div>
        <div class="recommend-card block-card">
          <h2 class="block-title">Recomendaciones de analistas</h2>
          <div v-if="recommendationStore.recommendations && recommendationStore.recommendations.length">
            <RecommendationBar />
          </div>
          <div v-else class="no-news">No hay recomendaciones recientes para este ticker.</div>
        </div>
        <div class="news-card block-card">
          <h2 class="block-title">Noticias recientes</h2>
          <div v-if="news && news.length">
            <div v-for="n in news" :key="n.id" class="news-item">
              <img v-if="n.image" :src="n.image" alt="news" class="news-img" />
              <div class="news-content">
                <a :href="n.url" class="news-headline" target="_blank">{{ n.headline }}</a>
                <div class="news-date">{{ formatDate(n.datetime) }}</div>
                <div class="news-summary">{{ n.summary }}</div>
                <div class="news-source">Fuente: {{ n.source }}</div>
              </div>
            </div>
          </div>
          <div v-else class="no-news">No hay noticias recientes para este ticker.</div>
        </div>
      </div>
    </div>
  </div>
</template>
<script setup>


import { useRouter } from 'vue-router'
const router = useRouter()
function goBack() {
  router.push('/')
}

import { ref, onMounted } from 'vue'
import { computed } from 'vue'

function formatChange(n, percent = false) {
  if (n > 0) return `+${n}${percent ? '%' : ''}`;
  if (n < 0) return `${n}${percent ? '%' : ''}`;
  return `0${percent ? '%' : ''}`;
}
import RecommendationBar from './RecommendationBar.vue'
import { useRoute } from 'vue-router'
import { useRecommendationStore } from '../stores/recommendation'

const route = useRoute()
const ticker = route.params.ticker
const company = ref(null)
const quote = ref(null)
const news = ref([])
const loading = ref(true)
const recommendationStore = useRecommendationStore()

async function fetchDetails() {
  loading.value = true
  try {
    const [companyRes, quoteRes] = await Promise.all([
      fetch(`/api/companyinfo/${ticker}`),
      fetch(`/api/quotes/${ticker}`)
    ])
    company.value = await companyRes.json()
    quote.value = await quoteRes.json()
    // Noticias
    const today = new Date();
    const to = today.toISOString().slice(0,10);
    const fromDate = new Date(today.getTime() - 30*24*60*60*1000); // 30 días atrás
    const from = fromDate.toISOString().slice(0,10);
    const token = import.meta.env.VITE_FINNHUB_TOKEN;
    const newsRes = await fetch(`https://finnhub.io/api/v1/company-news?symbol=${ticker}&from=${from}&to=${to}&token=${token}`);
    news.value = await newsRes.json();
    // Recomendaciones
    const recRes = await fetch(`https://finnhub.io/api/v1/stock/recommendation?symbol=${ticker}&token=${token}`);
    const recData = await recRes.json();
    recommendationStore.setRecommendations(recData);
  } catch (e) {
    company.value = null
    quote.value = null
    news.value = []
  }
  loading.value = false
}

onMounted(fetchDetails)
function priceColor(price, prevClose) {
  if (price > prevClose) return 'pos';
  if (price < prevClose) return 'neg';
  return '';
}

function formatDate(ts) {
  if (!ts) return '';
  const d = new Date(ts * 1000);
  return d.toLocaleDateString('es-ES', { year: 'numeric', month: 'short', day: 'numeric' });
}

</script>

<style scoped>
.bloomberg-table {
  min-height: 100vh;
  width: 100vw;
  background: #181818;
  margin: 0;
  padding: 0;
  box-sizing: border-box;
  display: flex;
  align-items: flex-start;
  justify-content: center;
}
.main-content {
  width: 100vw;
  max-width: 600px;
  margin: 0 auto;
}
.header-box {
  background: #fff;
  border-radius: 0;
  padding: 10px;
  text-align: center;
  place-items: center;
  border-bottom: 2px solid #222;
}
.title {
  font-size: 2.2rem;
  font-family: 'Menlo', 'Roboto Mono', monospace;
  font-weight: 800;
  color: #222;
  letter-spacing: 0.03em;
}
.subtitle {
  font-size: 0.95rem;
  color: #aaa;
  font-family: 'Menlo', 'Roboto Mono', monospace;
  margin-bottom: 10px;
}
.loading {
  color: #222;
  font-size: 1.1rem;
  text-align: center;
  margin-top: 24px;
  margin-bottom: 24px;
}
.loading-yellow {
  color: #ffe600;
  font-weight: bold;
}
.block-card {
  background: #000;
  border-radius: 0;
  box-shadow: none;
  padding: 22px 20px 18px 20px;
  margin: 22px 0 0 0;
  color: #fff;
  font-family: 'Menlo', 'Roboto Mono', monospace;
  border: 1.5px solid #888;
  border-bottom: 2px solid #333;
}
.block-title {
  color: #ffe600;
  font-size: 1.15rem;
  margin-bottom: 12px;
  letter-spacing: 0.02em;
}
.info-block {
  margin-bottom: 28px;
}
.info-block h2 {
  font-size: 1.15rem;
  color: #ffe600;
  margin-bottom: 10px;
  letter-spacing: 0.02em;
}
.pos {
  color: #00c853;
  font-weight: bold;
}
.neg {
  color: #d50000;
  font-weight: bold;
}
a {
  color: #ffe600;
  text-decoration: underline;
}
.back-btn {
  display: flex;
  flex-direction: column;
  align-items: center;
  cursor: pointer;
  margin-top: 18px;
  margin-bottom: 18px;
  user-select: none;
}
.arrow {
  font-size: 2.5rem;
  color: #ffd600;
  line-height: 1;
}
.volver {
  color: #ffd600;
  font-size: 1.1rem;
  font-family: 'Menlo', 'Roboto Mono', monospace;
  font-weight: bold;
  margin-top: 2px;
  letter-spacing: 0.03em;
}

.news-block {
  background: #000;
  border-radius: 16px;
  box-shadow: 0 2px 8px rgba(0,0,0,0.13);
  margin: 0 auto 28px auto;
  padding: 18px 18px 10px 18px;
  max-width: 600px;
  color: #fff;
}
.news-block h2 {
  color: #ffe600;
  font-size: 1.18rem;
  margin-bottom: 14px;
}
.news-item {
  display: flex;
  align-items: flex-start;
  margin-bottom: 18px;
  border: 2.5px solid #fff;
  border-radius: 0;
  background: #111;
  padding: 12px 10px 12px 10px;
}
.news-img {
  width: 60px;
  height: 60px;
  object-fit: cover;
  border-radius: 8px;
  margin-right: 16px;
  background: #222;
}
.news-content {
  flex: 1;
}
.news-headline {
  color: #ffe600;
  font-weight: bold;
  font-size: 1.02rem;
  text-decoration: underline;
}
.news-date {
  color: #bbb;
  font-size: 0.92rem;
  margin-bottom: 4px;
}
.news-summary {
  color: #eee;
  font-size: 0.99rem;
  margin-bottom: 3px;
}
.news-source {
  color: #888;
  font-size: 0.89rem;
}
.no-news {
  color: #ffd600;
  margin: 16px 0 0 0;
  text-align: center;
  font-size: 1.05rem;
}

</style>
<style scoped>
.bloomberg-table {
  font-family: 'Roboto Mono', 'Menlo', monospace;
  background: linear-gradient(120deg, #232526 0%, #414345 100%);
}

</style>

