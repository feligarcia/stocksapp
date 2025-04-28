<template>
  <div class="relative w-full my-5">
    <LoaderBar v-if="loadingRecommendations" class="my-6" />
    <template v-else-if="latest && latest.period">
      <!-- Flecha -->
      <!-- Flecha -->
      <!-- Flecha por fuera, apuntando la barra -->
      <div class="absolute z-20 transition-all duration-300" :style="{ ...arrowStyle, top: '-5px' }">
        <span class="text-yellow-300 text-2xl select-none drop-shadow">▼</span>
      </div>
      <!-- Etiquetas extremos -->
      <div class="flex justify-between items-center font-bold text-yellow-300 mb-1 text-base tracking-wide">
        <span class="ml-1">Sell</span>
        <span class="mr-1">Buy</span>
      </div>
      <!-- Barra con gradiente y marcas -->
      <div class="relative w-full rounded bg-gray-900" :style="{ background: gradient, height: '6px', margin: '0 0 10px 0' }">
        <!-- Marcas extremos -->
        <div class="absolute left-0 top-[-6px] w-[2px] h-4 bg-yellow-300 rounded z-10"></div>
        <div class="absolute right-0 top-[-6px] w-[2px] h-4 bg-yellow-300 rounded z-10"></div>
        <!-- Marcas hold (a la izquierda y derecha del centro) -->
        <div class="absolute top-[-6px] w-[2px] h-4 bg-yellow-300 rounded z-10" :style="holdTickStyle('left')"></div>
        <div class="absolute top-[-6px] w-[2px] h-4 bg-yellow-300 rounded z-10" :style="holdTickStyle('right')"></div>
      </div>
      <span v-if="latest.period" class="block text-gray-400 text-xs mt-2">(Fecha recomendación: {{ latest.period }})</span>
    </template>
    <template v-else>
      <div class="text-yellow-300 text-center mt-2">Sin conexión</div>
    </template>
  </div>
</template>

<script setup>
import LoaderBar from './LoaderBar.vue'
import { computed } from 'vue'
import { storeToRefs } from 'pinia'
import { useRecommendationStore } from '../stores/recommendation'
const recommendationStore = useRecommendationStore()
const { recommendations } = storeToRefs(recommendationStore)

const props = defineProps({
  loadingRecommendations: {
    type: Boolean,
    required: true
  }
})

const categories = [
  { key: 'strongSell', label: 'Strong Sell', color: '#e53935' }, // rojo intenso
  { key: 'sell', label: 'Sell', color: '#fb8c00' }, // naranja
  { key: 'hold', label: 'Hold', color: '#ffe600' }, // amarillo
  { key: 'buy', label: 'Buy', color: '#b2d900' }, // verde amarillento
  { key: 'strongBuy', label: 'Strong Buy', color: '#00bfae' } // verde intenso
]

// Calcula los porcentajes de cada categoría para el gradiente
const percents = computed(() => {
  if (!latest.value) return [0,0,0,0,0]
  const total = categories.reduce((acc, cat) => acc + (latest.value[cat.key] || 0), 0)
  if (!total) return [0,0,0,0,0]
  return categories.map(cat => ((latest.value[cat.key] || 0) / total) * 100)
})

// Construir el gradiente lineal
const gradient = computed(() => {
  if (!percents.value.length) return '#e53935'
  let stops = []
  let acc = 0
  for (let i = 0; i < categories.length; i++) {
    const start = acc
    acc += percents.value[i]
    // Si el porcentaje es 0, no agregues el color
    if (percents.value[i] === 0) continue
    stops.push(`${categories[i].color} ${start.toFixed(2)}%`)
    stops.push(`${categories[i].color} ${(acc).toFixed(2)}%`)
  }
  return `linear-gradient(to right, ${stops.join(', ')})`
})

// Posición de las marcas del hold (inicio y fin del segmento hold)
function holdTickStyle(side) {
  // Buscar el porcentaje acumulado hasta el segmento hold
  const idxHold = categories.findIndex(cat => cat.key === 'hold')
  let leftPercent = 0
  for (let i = 0; i < idxHold; i++) {
    leftPercent += percents.value[i] || 0
  }
  let rightPercent = leftPercent + (percents.value[idxHold] || 0)
  return {
    left: side === 'left' ? `${leftPercent}%` : `${rightPercent}%`
  }
}

// Ordena por fecha descendente y toma el más reciente
const latest = computed(() => {
  if (!recommendations.value || recommendations.value.length === 0) return {}
  const sorted = [...recommendations.value].sort((a, b) => new Date(b.period) - new Date(a.period))
  return sorted[0]
})

// Promedio ponderado
const weightedAvg = computed(() => {
  if (!latest.value) return 0;
  const total =
    1 * (latest.value.strongSell || 0) +
    2 * (latest.value.sell || 0) +
    3 * (latest.value.hold || 0) +
    4 * (latest.value.buy || 0) +
    5 * (latest.value.strongBuy || 0);
  const count =
    (latest.value.strongSell || 0) +
    (latest.value.sell || 0) +
    (latest.value.hold || 0) +
    (latest.value.buy || 0) +
    (latest.value.strongBuy || 0);
  return count ? total / count : 0;
});

// Calcula la posición de la flecha en %
const arrowStyle = computed(() => {
  // 1 = extremo izq, 5 = extremo der
  const percent = ((weightedAvg.value - 1) / 4) * 100;
  return {
    left: `calc(${percent}% - 10px)`
  };
});

function segmentStyle(key) {
  const cat = categories.find(c => c.key === key)
  return {
    background: cat.color
  }
}
</script>
