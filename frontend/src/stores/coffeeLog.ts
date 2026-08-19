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
  mood_tags: string[]
  scene_tags: string[]
  pairing_tags: string[]
  flavor_tags: string[]
  bean_id?: number | null
  bean?: CoffeeBeanInfo | null
  brew_ratio: string
  water_temp: string
  grind_size: string
  // Data quality fields (v2)
  record_mode: string
  coffee_name_source: string
  notes_source: string
  shop_source: string
  sensory_recorded: boolean
  source_log_id: number | null
  is_test_data: boolean
}

export interface CoffeeBeanInfo {
  id: number
  name: string
  origin: string
  processing_method: string
  roast_level: string
  roaster: string
  image_url: string
}

export interface NewCoffeeLog {
  coffee_name: string
  coffee_type: string
  shop_name: string
  location?: string
  image_url: string
  mood: string
  notes: string
  generate_ai?: boolean
  acidity: number
  bitterness: number
  sweetness: number
  body: number
  aroma: number
  aftertaste: number
  flavor_tags: string[]
  mood_tags?: string[]
  scene_tags?: string[]
  pairing_tags?: string[]
  bean_id?: number | null
  bean_name?: string
  brew_ratio?: string
  water_temp?: string
  grind_size?: string
  // Data quality fields (v2)
  record_mode?: string
  sensory_recorded?: boolean
  source_log_id?: number | null
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
  const statsMeta = ref<statsApi.StatsMeta | null>(null)
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

  // Coffee Personality
  const personalities = ref<statsApi.PersonalityTag[]>([])

  // Monthly Review
  const monthlyReview = ref<statsApi.MonthlyReviewData | null>(null)
  const monthlyReviewAI = ref('')

  // AI Status
  const aiStatus = ref<statsApi.AIStatus | null>(null)

  // AI Coffee Profile
  const coffeeProfileAI = ref('')

  // AI Preference Insight
  const preferenceInsight = ref('')

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

  function removeLogFromCache(id: number) {
    logs.value = logs.value.filter((log: CoffeeLog) => log.id !== id)
  }

  async function deleteLog(id: number, options?: { removeFromCache?: boolean }) {
    await coffeeLogApi.deleteCoffeeLog(id)
    if (options?.removeFromCache !== false) {
      removeLogFromCache(id)
    }
  }

  async function fetchStats() {
    try {
      const [overviewResp, profile] = await Promise.all([
        statsApi.getStatsOverview(),
        statsApi.getFlavorProfile()
      ])
      // Handle meta-wrapped response
      if (overviewResp && typeof overviewResp === 'object' && 'data' in overviewResp && 'meta' in overviewResp) {
        statsOverview.value = (overviewResp as statsApi.StatsResponseWrapper<statsApi.StatsOverview>).data
        statsMeta.value = (overviewResp as statsApi.StatsResponseWrapper<statsApi.StatsOverview>).meta
      } else {
        statsOverview.value = overviewResp as unknown as statsApi.StatsOverview
        statsMeta.value = null
      }
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

  async function fetchPersonality() {
    try {
      const res = await statsApi.getPersonality()
      personalities.value = res.personalities || []
    } catch (e) {
      console.error('Failed to fetch personality:', e)
      personalities.value = []
    }
  }

  async function fetchMonthlyReview(month?: string) {
    try {
      const [review, aiRes] = await Promise.all([
        statsApi.getMonthlyReview(month),
        statsApi.getMonthlyReviewAI(month)
      ])
      monthlyReview.value = review
      monthlyReviewAI.value = aiRes.summary || ''
    } catch (e) {
      console.error('Failed to fetch monthly review:', e)
    }
  }

  async function fetchAIStatus() {
    try {
      aiStatus.value = await statsApi.getAIStatus()
    } catch (e) {
      console.error('Failed to fetch AI status:', e)
    }
  }

  async function fetchCoffeeProfile() {
    try {
      const res = await statsApi.generateCoffeeProfile()
      coffeeProfileAI.value = res.profile || ''
    } catch (e) {
      console.error('Failed to fetch coffee profile:', e)
      coffeeProfileAI.value = ''
    }
  }

  async function fetchPreferenceInsight() {
    try {
      const res = await statsApi.generatePreferenceInsight()
      preferenceInsight.value = res.insight || ''
    } catch (e) {
      console.error('Failed to fetch preference insight:', e)
      preferenceInsight.value = ''
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
    personalities,
    monthlyReview,
    monthlyReviewAI,
    aiStatus,
    coffeeProfileAI,
    preferenceInsight,
    statsMeta,
    fetchLogs,
    fetchLogById,
    addLog,
    deleteLog,
    removeLogFromCache,
    fetchStats,
    fetchLifestyleQuote,
    fetchPersonality,
    fetchMonthlyReview,
    fetchAIStatus,
    fetchCoffeeProfile,
    fetchPreferenceInsight
  }
})
