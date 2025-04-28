import { createRouter, createWebHistory } from 'vue-router'
import BloombergTable from './components/BloombergTable.vue'
import TickerDetail from './components/TickerDetail.vue'

const routes = [
  { path: '/', component: BloombergTable },
  { path: '/:ticker', component: TickerDetail, props: true },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router
