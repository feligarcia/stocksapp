<template>
  <div class="bloomberg-table">
    <div class="main-content">
      <div class="header">
        <h1 class="title">StocksApp</h1>
        <span class="subtitle">by Juan Felipe Garcia</span>
      </div>
      <div class="search-bar">
        <input
          v-model="search"
          type="text"
          placeholder="Search ticker..."
          class="search-input"
          @input="onSearch"
        />
      </div>
      <div class="table-container">
        <table class="bb-table">
          <thead>
            <tr>
              <th v-for="col in columns" :key="col.key" @click="sortBy(col.key)" :class="['sortable', col.key==='ticker' ? 'ticker' : '', sortKey===col.key ? (sortAsc ? 'sort-asc' : 'sort-desc') : '']">
                {{ col.label }}
                <span v-if="sortKey===col.key">{{ sortAsc ? '▲' : '▼' }}</span>
              </th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="row in filteredAndSortedQuotes" :key="row.ticker"
                :class="[$index % 2 === 0 ? 'even' : 'odd', 'bb-row']"
                @mouseover="hovered = row.ticker"
                @mouseleave="hovered = null"
                @click="openTicker(row.ticker)"
                :style="hovered === row.ticker ? 'background:#383838;cursor:pointer;' : ''"
            >
              <td class="ticker">{{ row.ticker }}</td>
              <td>{{ row.valor }}</td>
              <td :class="row.valcambio > 0 ? 'pos' : row.valcambio < 0 ? 'neg' : ''">{{ formatChange(row.valcambio) }}</td>
              <td :class="row.porcencambio > 0 ? 'pos' : row.porcencambio < 0 ? 'neg' : ''">{{ formatChange(row.porcencambio, true) }}</td>
              <td class="time">{{ row.hora }}</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<style scoped>
.bloomberg-table {
  min-height: 100vh;
  width: 100vw;
  background: #000;
  margin: 0;
  padding: 0;
  box-sizing: border-box;
  display: flex;
  align-items: flex-start;
  justify-content: center;
  font-family: 'Menlo', 'Roboto Mono', monospace;
  color: #fff;
}
.main-content {
  width: 100vw;
  max-width: 1100px;
  margin: 0 auto;
}
.header {
  text-align: center;
  margin-bottom: 6px;
}
.title {
  font-size: 2.2rem;
  font-weight: 800;
  color: #fff;
  margin: 20px 0 2px 0;
  letter-spacing: 0.03em;
}
.subtitle {
  font-size: 0.95rem;
  color: #aaa;
  margin-bottom: 10px;
  display: block;
}
.table-container {
  width: 100vw;
  max-width: 1100px;
  margin: 0 auto;
  background: #000;
  border: none;
  box-shadow: none;
  overflow-x: auto;
}
.bb-table {
  width: 100%;
  border-collapse: collapse;
  background: #000;
  font-size: 1rem;
}
.bb-table th, .bb-table td {
  border: 1px solid #444;
  padding: 2px 10px;
  text-align: left;
  background: transparent;
  color: #fff;
  font-weight: 600;
}
.bb-table th {
  background: #3a3a3a;
  color: #fff;
  font-size: 1.05em;
  font-weight: 800;
  letter-spacing: 0.02em;
}
.bb-table th.sortable {
  cursor: pointer;
  transition: background 0.12s;
}
.bb-table th.sortable:hover {
  background: #464646;
}

.bb-table th.ticker, .bb-table td.ticker {
  color: #ffe600;
  font-weight: 800;
  font-size: 1.06em;
}
.bb-table td.ticker {
  background: #222;
}
.bb-table tr.even {
  background: #191919;
}
.bb-table tr.odd {
  background: #232323;
}
.bb-row:hover, .bb-row:focus {
  background: #383838 !important;
  outline: 1.5px solid #ffe600;
  cursor: pointer;
}

.bb-table td.pos {
  color: #00ff00;
  font-weight: 400;
}
.bb-table td.neg {
  color: #ff3333;
  font-weight: 400;
}
.bb-table td.time {
  color: #aaa;
  font-size: 0.9em;
}
.search-bar {
  width: 100%;
  max-width: 1100px;
  margin: 0 auto 10px auto;
  display: flex;
  justify-content: flex-end;
}
.search-input {
  background: #000;
  border: 1.5px solid #888;
  color: #ffe600;
  font-size: 1.1em;
  padding: 4px 14px;
  border-radius: 0;
  outline: none;
  box-shadow: none;
  transition: border 0.15s;
  width: 210px;
}
.search-input:focus {
  border-color: #ffe600;
  color: #fff200;
}
</style>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'

function formatChange(n, percent = false) {
  if (n > 0) return `+${n}${percent ? '%' : ''}`;
  if (n < 0) return `${n}${percent ? '%' : ''}`;
  return `0${percent ? '%' : ''}`;
}

const quotes = ref([])
const search = ref('')
const sortKey = ref('ticker')
const sortAsc = ref(true)
const hovered = ref(null)

const columns = [
  { key: 'ticker', label: 'Ticker' },
  { key: 'valor', label: 'Valor' },
  { key: 'valcambio', label: 'Cambio' },
  { key: 'porcencambio', label: '% Cambio' },
  { key: 'hora', label: 'Hora' },
]

const router = useRouter()
function openTicker(ticker) {
  router.push({ path: `/${ticker}` })
}

async function fetchData() {
  const data = await fetch('/api/quotes').then(r => r.json())
  const now = new Date().toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' })
  quotes.value = data.map(q => ({
    ticker: q.ticker,
    valor: q.c,
    valcambio: q.d,
    porcencambio: q.dp,
    hora: now
  }))
}

function sortBy(key) {
  if (sortKey.value === key) {
    sortAsc.value = !sortAsc.value
  } else {
    sortKey.value = key
    sortAsc.value = true
  }
}

const filteredAndSortedQuotes = computed(() => {
  let filtered = quotes.value
  if (search.value) {
    filtered = filtered.filter(q => q.ticker.toLowerCase().includes(search.value.toLowerCase()))
  }
  const sorted = [...filtered].sort((a, b) => {
    if (a[sortKey.value] < b[sortKey.value]) return sortAsc.value ? -1 : 1
    if (a[sortKey.value] > b[sortKey.value]) return sortAsc.value ? 1 : -1
    return 0
  })
  return sorted
})

function onSearch() {
  // No-op, v-model triggers computed
}

onMounted(fetchData)
</script>

<style scoped>
.bloomberg-table {
  font-family: 'Roboto Mono', 'Menlo', monospace;
  background: #181818;
}
</style>


<style scoped>
.bloomberg-table {
  font-family: 'Roboto Mono', 'Menlo', monospace;
  background: linear-gradient(120deg, #232526 0%, #414345 100%);
}
th, td {
  border-bottom: 1px solid #222;
}
</style>
