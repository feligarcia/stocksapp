<template>
  <div class="recommend-bar-container" style="width:100%;">
    <!-- Flecha -->
    <div class="recommend-arrow" :style="arrowStyle">
      <span class="arrow">▼</span>
    </div>
    <!-- Etiquetas extremos -->
    <div class="recommend-bar-labels">
      <span class="label-sell">Sell</span>
      <span class="label-buy">Buy</span>
    </div>
    <!-- Barra con gradiente y marcas -->
    <div class="recommend-bar-bg" :style="{ background: gradient, height: '6px', position: 'relative', borderRadius: '3px', margin: '0 0 10px 0' }">
      <!-- Marcas extremos -->
      <div class="tick tick-left" />
      <div class="tick tick-right" />
      <!-- Marcas hold (a la izquierda y derecha del centro) -->
      <div class="tick tick-hold-left" :style="holdTickStyle('left')" />
      <div class="tick tick-hold-right" :style="holdTickStyle('right')" />
    </div>
    <span v-if="latest.period" class="rec-period-date">(Fecha recomendación: {{ latest.period }})</span>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { storeToRefs } from 'pinia'
import { useRecommendationStore } from '../stores/recommendation'
const recommendationStore = useRecommendationStore()
const { recommendations } = storeToRefs(recommendationStore)

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

<style scoped>
.recommend-bar-container {
  width: 100%;
  position: relative;
  margin: 18px 0 10px 0;
}
.recommend-bar-labels {
  display: flex;
  justify-content: space-between;
  align-items: center;
  font-weight: bold;
  color: #ffe600;
  margin-bottom: 2px;
  font-size: 1.1em;
  letter-spacing: 1px;
}
.label-sell {
  margin-left: 2px;
}
.label-buy {
  margin-right: 2px;
}
.recommend-bar-bg {
  width: 100%;
  box-sizing: border-box;
  position: relative;
  border: none;
  box-shadow: none;
  outline: none;
  background-clip: padding-box;
  border-radius: 3px;
  overflow: visible;
  background: #232526;
  display: flex;
  flex-direction: row;
  height: 32px;
}
.recommend-arrow {
  position: absolute;
  left: 0;
  top: 16px; /* Justo debajo de la barra de 6px + margen */
  z-index: 2;
  font-size: 18px;
  color: #ffe600;
  transition: left 0.3s;
  pointer-events: none;
}
.tick {
  position: absolute;
  width: 2px;
  height: 18px;
  background: #ffe600;
  top: -6px;
  z-index: 2;
  border-radius: 1px;
}
.tick-left {
  left: 0;
}
.tick-right {
  right: 0;
}
.tick-hold-left {
  left: 48%; /* Ajuste fino para el lado izquierdo de hold */
}
.tick-hold-right {
  left: 52%; /* Ajuste fino para el lado derecho de hold */
}



.recommend-segment {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 0.95em;
  color: #181818;
  font-weight: 700;
  border-right: 1.5px solid #232526;
}
.recommend-segment:last-child {
border-right: none;
}
.recommend-label {
z-index: 2;
}
.recommend-arrow {
position: absolute;
top: -20px;
transition: left 0.3s;
z-index: 3;
font-size: 1.7em;
color: #ffe600;
}
.arrow {
font-size: 1.5em;
}
.recommend-summary {
margin-top: 10px;
color: #fff;
font-size: 1em;
display: flex;
flex-wrap: wrap;
gap: 18px;
align-items: center;
  flex-wrap: wrap;
  gap: 18px;

}
.rec-count {
  color: #ffe600;
}
.rec-period {
  color: #aaa;
  font-size: 0.95em;
}
.rec-period-date {
  color: #aaa;
  font-size: 0.85em;
  align-self: right;
}
</style>
