<template>
  <div
    class="min-h-screen w-full flex items-start justify-center font-mono text-white"
    style="
      background: linear-gradient(120deg, #232526 0%, #414345 100%);
      font-family: 'Roboto Mono', 'Menlo', monospace;
    "
  >
    <div class="w-full max-w-5xl mx-auto">
      <div class="text-center mb-2">
        <h1 class="text-3xl font-extrabold text-white mt-5 mb-1 tracking-wide">
          StocksApp
        </h1>
        <span class="text-gray-400 text-base block mb-2"
          >by Juan Felipe Garcia</span
        >
      </div>
      <div class="flex w-full max-w-5xl mx-auto mb-2 justify-end">
        <input
          v-model="search"
          type="text"
          placeholder="Buscar ticker..."
          class="bg-black border border-gray-500 text-yellow-300 font-mono text-lg px-4 py-1 rounded-none outline-none focus:border-yellow-300 focus:text-yellow-200 w-52 transition-all"
        />
      </div>
      <div class="w-full max-w-5xl mx-auto bg-black/80 overflow-x-auto">
        <LoaderBar v-if="isLoading" />
        <div v-else>
          <div
            v-if="errorQuotes"
            class="text-yellow-300 text-center my-6 text-xl"
          >
            Sin conexión
          </div>
          <table
            v-else
            class="w-full border-collapse bg-gray-950/80 text-white text-base font-mono"
            style="font-family: 'Roboto Mono', 'Menlo', monospace"
          >
            <thead>
              <tr>
                <th
                  v-for="col in columns"
                  :key="col.key"
                  @click="sortBy(col.key)"
                  :class="[
                    'bg-black font-extrabold px-3 py-2 h-8 min-h-[32px] border border-gray-700 cursor-pointer select-none transition-colors align-middle',
                    col.key === 'ticker' ? 'text-yellow-300' : '',
                    (col.key === 'valcambio' || col.key === 'porcencambio') &&
                    sortKey === col.key
                      ? sortAsc
                        ? 'text-red-400'
                        : 'text-green-400'
                      : '',
                  ]"
                >
                  {{ col.label }}
                  <span v-if="sortKey === col.key">{{
                    sortAsc ? "▲" : "▼"
                  }}</span>
                </th>
              </tr>
            </thead>
            <tbody class="bg-black">
              <tr
                v-for="(row, index) in filteredAndSortedQuotes"
                :key="row.ticker"
                :class="[
                  index % 2 === 0 ? 'bg-gray-900' : 'bg-[#232526]',
                  'hover:bg-gradient-to-r hover:from-gray-700 hover:to-gray-900 cursor-pointer transition-colors group',
                ]"
                @mouseover="hovered = row.ticker"
                @mouseleave="hovered = null"
                @click="openTicker(row.ticker)"
              >
                <td
                  :class="[
                    'text-yellow-300 font-extrabold bg-inherit px-3 py-2 h-8 min-h-[32px] transition-colors align-middle',
                    hovered === row.ticker
                      ? 'bg-gradient-to-r from-yellow-700 to-yellow-900'
                      : '',
                    'group-hover:bg-gradient-to-r group-hover:from-yellow-700 group-hover:to-yellow-900',
                  ]"
                >
                  {{ row.ticker }}
                </td>
                <td class="px-3 py-2 h-8 min-h-[32px] align-middle">
                  {{ row.valor }}
                </td>
                <td
                  :class="[
                    row.valcambio > 0
                      ? 'text-green-400'
                      : row.valcambio < 0
                      ? 'text-red-400'
                      : 'text-white',
                    'px-3 py-2 h-8 min-h-[32px] align-middle',
                  ]"
                >
                  {{ formatChange(row.valcambio) }}
                </td>
                <td
                  :class="[
                    row.porcencambio > 0
                      ? 'text-green-400'
                      : row.porcencambio < 0
                      ? 'text-red-400'
                      : 'text-white',
                    'px-3 py-2 h-8 min-h-[32px] align-middle',
                  ]"
                >
                  {{ formatChange(row.porcencambio, true) }}
                </td>
                <td
                  class="text-gray-400 text-sm px-3 py-2 h-8 min-h-[32px] align-middle"
                >
                  {{ row.hora }}
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </div>
</template>



<script setup>
import { ref, computed, onMounted } from "vue";
import { useRouter } from "vue-router";
import LoaderBar from "./LoaderBar.vue";
import { useQuotesStore } from "../stores/quotes";
import { storeToRefs } from "pinia";

function formatChange(n, percent = false) {
  if (n > 0) return `+${n}${percent ? "%" : ""}`;
  if (n < 0) return `${n}${percent ? "%" : ""}`;
  return `0${percent ? "%" : ""}`;
}

const quotesStore = useQuotesStore();
const { quotes } = storeToRefs(quotesStore);
const search = ref("");
const sortKey = ref("ticker");
const sortAsc = ref(true);
const hovered = ref(null);

const columns = [
  { key: "ticker", label: "Ticker" },
  { key: "valor", label: "Valor" },
  { key: "valcambio", label: "Cambio" },
  { key: "porcencambio", label: "% Cambio" },
  { key: "hora", label: "Hora" },
];

const router = useRouter();
function openTicker(ticker) {
  router.push({ path: `/${ticker}` });
}

const isLoading = ref(false);
const errorQuotes = ref(false);

async function fetchData() {
  // TTL de 1 minuto
  const now = Date.now();
  const ttl = 60 * 1000;
  if (
    quotes.value.length > 0 &&
    quotesStore.lastFetched &&
    now - quotesStore.lastFetched < ttl
  ) {
    return;
  }
  isLoading.value = true;
  errorQuotes.value = false;
  try {
    const data = await fetch("/api/quotes").then((r) => r.json());
    const hora = new Date().toLocaleTimeString([], {
      hour: "2-digit",
      minute: "2-digit",
    });
    const mapped = data.map((q) => ({
      ticker: q.ticker,
      valor: q.c,
      valcambio: q.d,
      porcencambio: q.dp,
      hora,
    }));
    quotesStore.setQuotes(mapped);
  } catch (e) {
    errorQuotes.value = true;
  } finally {
    isLoading.value = false;
  }
}

function sortBy(key) {
  if (sortKey.value === key) {
    sortAsc.value = !sortAsc.value;
  } else {
    sortKey.value = key;
    sortAsc.value = true;
  }
}

const filteredAndSortedQuotes = computed(() => {
  let filtered = quotes.value;
  if (search.value) {
    filtered = filtered.filter((q) =>
      q.ticker.toLowerCase().includes(search.value.toLowerCase())
    );
  }
  const sorted = [...filtered].sort((a, b) => {
    if (a[sortKey.value] < b[sortKey.value]) return sortAsc.value ? -1 : 1;
    if (a[sortKey.value] > b[sortKey.value]) return sortAsc.value ? 1 : -1;
    return 0;
  });
  return sorted;
});

onMounted(fetchData);
</script>



