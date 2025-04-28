<template>
  <div
    class="min-h-screen w-full flex items-start justify-center font-mono"
    style="background: linear-gradient(120deg, #232526 0%, #414345 100%); font-family: 'Roboto Mono', 'Menlo', monospace;"
  >
    <div class="w-full max-w-2xl mx-auto">
      <div class="flex flex-col items-center mt-6 mb-6 select-none cursor-pointer" @click="goBack">
        <span class="text-4xl text-yellow-400 leading-none">&#8592;</span>
        <div class="text-yellow-400 text-lg font-bold mt-1 tracking-wide">Volver</div>
      </div>
      <div class="bg-white rounded-none p-3 text-center border-b-2 border-gray-800 mb-2">
        <h1 class="text-2xl font-extrabold text-gray-900 tracking-wide">
          {{ ticker }}<span v-if="company && company.name"> - {{ company.name }}</span>
        </h1>
      </div>
      <div class="bg-black rounded-none shadow-none p-6 mt-6 border border-gray-500 border-b-2 border-b-gray-700 text-white">
        <h2 class="text-yellow-300 text-lg mb-3 font-bold tracking-wide">Company Info</h2>
        <LoaderBar v-if="loadingCompany" class="my-4" />
        <template v-else>
          <div v-if="company">
            <div class="mb-2 flex justify-center">
              <img v-if="company.logo" :src="company.logo" alt="Logo" class="w-20 h-20 object-contain rounded-lg border border-gray-300 bg-white" />
            </div>
            <div><span class="font-bold">Name:</span> {{ company.name }}</div>
            <div><span class="font-bold">Industry:</span> {{ company.finnhub_industry || company.finnhubIndustry }}</div>
            <div><span class="font-bold">Country:</span> {{ company.country }}</div>
            <div><span class="font-bold">IPO:</span> {{ company.ipo }}</div>
            <div><span class="font-bold">Website:</span> <a :href="company.web || company.weburl" target="_blank" class="text-yellow-300 underline">{{ company.web || company.weburl }}</a></div>
          </div>
          <div v-else class="text-yellow-300 text-center mt-2">No company info found.</div>
        </template>
      </div>
      <div class="bg-black rounded-none shadow-none p-6 mt-6 border border-gray-500 border-b-2 border-b-gray-700 text-white">
        <h2 class="text-yellow-300 text-lg mb-3 font-bold tracking-wide">Quote</h2>
        <LoaderBar v-if="loadingQuote" class="my-4" />
        <template v-else>
          <div v-if="quote && Array.isArray(quote) && quote.length">
            <div class="space-y-1">
              <div><span class="font-bold">Price:</span> <span :class="priceColor(quote[0].c, quote[0].pc)">{{ quote[0].c }}</span></div>
              <div><span class="font-bold">Change:</span> <span :class="quote[0].d > 0 ? 'text-green-500 font-bold' : quote[0].d < 0 ? 'text-red-500 font-bold' : ''">{{ formatChange(quote[0].d) }}</span></div>
              <div><span class="font-bold">% Change:</span> <span :class="quote[0].dp > 0 ? 'text-green-500 font-bold' : quote[0].dp < 0 ? 'text-red-500 font-bold' : ''">{{ formatChange(quote[0].dp, true) }}</span></div>
              <div><span class="font-bold">High:</span> {{ quote[0].h }}</div>
              <div><span class="font-bold">Low:</span> {{ quote[0].l }}</div>
              <div><span class="font-bold">Open:</span> {{ quote[0].o }}</div>
              <div><span class="font-bold">Prev Close:</span> {{ quote[0].pc }}</div>
            </div>
          </div>
          <div v-else-if="quote">
            <div class="space-y-1">
              <span class="font-bold">Price:</span> <span :class="priceColor(quote.c, quote.pc)">{{ quote.c }}</span><br />
              <span class="font-bold">Change:</span> <span :class="quote.d > 0 ? 'text-green-500 font-bold' : quote.d < 0 ? 'text-red-500 font-bold' : ''">{{ quote.d }}</span><br />
              <span class="font-bold">% Change:</span> <span :class="quote.dp > 0 ? 'text-green-500 font-bold' : quote.dp < 0 ? 'text-red-500 font-bold' : ''">{{ quote.dp }}%</span><br />
              <span class="font-bold">High:</span> {{ quote.h }}<br />
              <span class="font-bold">Low:</span> {{ quote.l }}<br />
              <span class="font-bold">Open:</span> {{ quote.o }}<br />
              <span class="font-bold">Prev Close:</span> {{ quote.pc }}
            </div>
          </div>
          <div v-else class="text-yellow-300 text-center mt-2">No quote info found.</div>
        </template>
      </div>
      <div class="bg-black rounded-none shadow-none p-6 mt-6 border border-gray-500 border-b-2 border-b-gray-700 text-white">
        <h2 class="text-yellow-300 text-lg mb-3 font-bold tracking-wide">Recomendaciones de analistas</h2>
        <RecommendationBar :loading-recommendations="loadingRecommendations" :recommendations="recommendationStore.recommendations" />
      </div>
      <div class="bg-black rounded-none shadow-none p-6 mt-6 border border-gray-500 border-b-2 border-b-gray-700 text-white">
        <h2 class="text-yellow-300 text-lg mb-3 font-bold tracking-wide">Noticias recientes</h2>
        <LoaderBar v-if="loadingNews" class="my-4" />
        <template v-else>
          <div v-if="news && news.length">
            <div v-for="n in news" :key="n.id" class="flex items-start mb-4 border-2 border-white rounded-none bg-gray-900 p-3 cursor-pointer hover:bg-gray-800 transition-colors">
              <img v-if="n.image" :src="n.image" alt="news" class="w-16 h-16 object-cover rounded-md mr-4 bg-gray-800" />
              <div class="flex-1">
                <a :href="n.url" class="text-yellow-300 font-bold underline text-base" target="_blank">{{ n.headline }}</a>
                <div class="text-gray-400 text-sm mb-1">{{ formatDate(n.datetime) }}</div>
                <div class="text-gray-200 text-sm mb-1">{{ n.summary }}</div>
                <div class="text-gray-500 text-xs">Fuente: {{ n.source }}</div>
              </div>
            </div>
          </div>
          <div v-else class="text-yellow-300 text-center mt-2">No hay noticias recientes para este ticker.</div>
        </template>
      </div>
    </div>
  </div>
</template>
<script setup>
import LoaderBar from './LoaderBar.vue'
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
const router = useRouter()
function goBack() {
  router.push('/')
}

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
const loadingCompany = ref(true)
const loadingQuote = ref(true)
const loadingNews = ref(true)
const loading = ref(true)
const loadingRecommendations = ref(true)
const recommendationStore = useRecommendationStore()

async function fetchDetails() {
  loading.value = true
  loadingCompany.value = true
  loadingQuote.value = true
  loadingNews.value = true
  try {
    // Company
    const companyRes = await fetch(`/api/companyinfo/${ticker}`)
    company.value = await companyRes.json()
    loadingCompany.value = false
    // Quote
    const quoteRes = await fetch(`/api/quotes/${ticker}`)
    quote.value = await quoteRes.json()
    loadingQuote.value = false
    // Noticias
    const today = new Date();
    const to = today.toISOString().slice(0,10);
    const fromDate = new Date(today.getTime() - 30*24*60*60*1000); // 30 días atrás
    const from = fromDate.toISOString().slice(0,10);
    const token = import.meta.env.VITE_FINNHUB_TOKEN;
    const newsRes = await fetch(`https://finnhub.io/api/v1/company-news?symbol=${ticker}&from=${from}&to=${to}&token=${token}`);
    news.value = await newsRes.json();
    loadingNews.value = false
    // Recomendaciones
    const recRes = await fetch(`https://finnhub.io/api/v1/stock/recommendation?symbol=${ticker}&token=${token}`);
    const recData = await recRes.json();
    recommendationStore.setRecommendations(recData);
    loadingRecommendations.value = false;
  } catch (e) {
    company.value = null
    quote.value = null
    news.value = []
    loadingCompany.value = false
    loadingQuote.value = false
    loadingNews.value = false
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