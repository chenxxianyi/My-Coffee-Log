import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import * as coffeeLogApi from '@/api/coffeeLog'
import * as statsApi from '@/api/stats'
import { getLocalDateString } from '@/utils/date'

export interface CoffeeLog {
  id: number
  coffee_name: string
  coffee_type: string
  shop_name: string
  location?: string
  image_url: string
  drink_date: string
  mood: string
  notes: string
  acidity: number
  bitterness: number
  sweetness: number
  body: number
  aroma: number
  aftertaste: number
  ai_summary: string
  flavor_tags: string[]
}

export interface NewCoffeeLog {
  coffee_name: string
  coffee_type: string
  shop_name: string
  location?: string
  image_url: string
  mood: string
  notes: string
  acidity: number
  bitterness: number
  sweetness: number
  body: number
  aroma: number
  aftertaste: number
  flavor_tags: string[]
}

export const useCoffeeLogStore = defineStore('coffeeLog', () => {
  
  const DEFAULT_PHOTOS = [
    "https://images.unsplash.com/photo-1514432324607-a09d9b4aefdd?auto=format&fit=crop&q=80&w=600",
    "https://images.unsplash.com/photo-1541167760496-1628856ab772?auto=format&fit=crop&q=80&w=600",
    "https://images.unsplash.com/photo-1495474472287-4d71bcdd2085?auto=format&fit=crop&q=80&w=600",
    "https://images.unsplash.com/photo-1507133750040-4a8f57021571?auto=format&fit=crop&q=80&w=600"
  ]

  const logs = ref<CoffeeLog[]>([])
  const isLoading = ref(false)

  // Stats from API
  const statsOverview = ref<statsApi.StatsOverview | null>(null)
  const flavorProfile = ref<statsApi.FlavorProfile | null>(null)

  // Getters
  const getLogById = computed(() => {
    return (id: number) => logs.value.find((log: CoffeeLog) => log.id === id)
  })

  const totalBrews = computed(() => statsOverview.value?.total_count || 0)
  const monthBrews = computed(() => statsOverview.value?.month_count || 0)
  const favoriteCoffeeType = computed(() => statsOverview.value?.favorite_coffee_type || 'Pour Over')
  const favoriteFlavorTag = computed(() => statsOverview.value?.favorite_flavor_tag || 'citrus')
  const recentFlavorTags = computed(() => statsOverview.value?.recent_flavor_tags || [])

  // Lifestyle quote from AI
  const lifestyleQuote = ref('')

  // Today's log check
  const todayLog = computed(() => {
    const today = getLocalDateString()
    return logs.value.find((log: CoffeeLog) => log.drink_date === today) || null
  })

  const averageSensoryValues = computed(() => {
    if (flavorProfile.value) {
      const p = flavorProfile.value
      return [
        Number(p.acidity.toFixed(1)),
        Number(p.bitterness.toFixed(1)),
        Number(p.sweetness.toFixed(1)),
        Number(p.body.toFixed(1)),
        Number(p.aroma.toFixed(1)),
        Number(p.aftertaste.toFixed(1))
      ]
    }
    return [0, 0, 0, 0, 0, 0]
  })

  // Actions
  async function fetchLogs(params?: { page?: number; page_size?: number; month?: string; coffee_type?: string }) {
    isLoading.value = true
    try {
      const res = await coffeeLogApi.getCoffeeLogs(params)
      logs.value = res.list
      return res
    } finally {
      isLoading.value = false
    }
  }

  async function fetchLogById(id: number) {
    const log = await coffeeLogApi.getCoffeeLogById(id)
    // Update in local cache if exists, otherwise add
    const idx = logs.value.findIndex(l => l.id === id)
    if (idx >= 0) {
      logs.value[idx] = log
    } else {
      logs.value.unshift(log)
    }
    return log
  }

  async function addLog(newLog: NewCoffeeLog): Promise<CoffeeLog> {
    const created = await coffeeLogApi.createCoffeeLog({
      ...newLog,
      drink_date: getLocalDateString()
    })
    logs.value.unshift(created)
    return created
  }

  async function deleteLog(id: number) {
    await coffeeLogApi.deleteCoffeeLog(id)
    logs.value = logs.value.filter((log: CoffeeLog) => log.id !== id)
  }

  async function fetchStats() {
    try {
      const [overview, profile] = await Promise.all([
        statsApi.getStatsOverview(),
        statsApi.getFlavorProfile()
      ])
      statsOverview.value = overview
      flavorProfile.value = profile
    } catch (e) {
      console.error('Failed to fetch stats:', e)
    }
  }

  async function fetchLifestyleQuote() {
    try {
      const res = await statsApi.getLifestyleQuote()
      lifestyleQuote.value = res.quote
    } catch (e) {
      console.error('Failed to fetch lifestyle quote:', e)
      lifestyleQuote.value = ''
    }
  }

  return {
    logs,
    isLoading,
    DEFAULT_PHOTOS,
    getLogById,
    totalBrews,
    monthBrews,
    favoriteCoffeeType,
    favoriteFlavorTag,
    recentFlavorTags,
    averageSensoryValues,
    lifestyleQuote,
    todayLog,
    fetchLogs,
    fetchLogById,
    addLog,
    deleteLog,
    fetchStats,
    fetchLifestyleQuote
  }
})
