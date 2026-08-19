<template>
  <div class="flex-1 w-full flex flex-col bg-coffee-warmWhite text-coffee-charcoal min-h-screen">
    <!-- Header -->
    <div class="px-4 py-4 flex items-center gap-3 border-b border-coffee-cream bg-coffee-warmWhite/95 backdrop-blur-sm sticky top-0 z-10 select-none">
      <button
        @click="router.back()"
        class="grid w-9 h-9 flex-shrink-0 place-items-center rounded-full border border-coffee-latte/35 bg-coffee-cream/30 text-coffee-brown hover:border-coffee-brown/35 hover:bg-coffee-cream/70 hover:text-coffee-espresso focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-coffee-latte transition-colors"
        aria-label="返回"
      >
        <ArrowLeft class="w-[18px] h-[18px]" />
      </button>
      <div class="flex-1 min-w-0">
        <span class="text-[9px] uppercase tracking-[0.25em] font-semibold text-coffee-softGray">Weekly Review</span>
        <h2 class="font-serif text-2xl font-light text-coffee-espresso leading-none mt-1 truncate">本周回顾</h2>
      </div>
    </div>

    <!-- Content -->
    <div class="flex-1 overflow-y-auto px-6 py-5 space-y-6 pb-24 scrollbar-none">
      <!-- Loading State -->
      <div v-if="loading" class="flex items-center justify-center py-12">
        <Loader2 class="w-6 h-6 text-coffee-brown animate-spin" />
      </div>

      <!-- Empty State -->
      <div v-else-if="review && review.count === 0" class="text-center py-12 space-y-4">
        <div class="inline-grid w-16 h-16 place-items-center rounded-full bg-coffee-cream/60">
          <Coffee class="w-8 h-8 text-coffee-brown" />
        </div>
        <div>
          <h3 class="font-serif text-lg text-coffee-espresso">本周还没有记录</h3>
          <p class="text-sm text-coffee-softGray mt-2">记录你的第一杯咖啡，开始本周的旅程</p>
        </div>
        <button
          @click="router.push('/create')"
          class="px-6 py-2.5 bg-coffee-espresso text-coffee-warmWhite rounded-sm text-xs font-semibold tracking-wider hover:bg-coffee-brown transition-colors"
        >开始记录</button>
      </div>

      <!-- Review Content -->
      <div v-else-if="review" class="space-y-6">
        <!-- Week Info -->
        <div class="text-center">
          <p class="text-[10px] uppercase tracking-widest text-coffee-softGray">
            {{ review.week_start }} — {{ review.week_end }}
          </p>
          <div class="mt-2 text-4xl font-serif text-coffee-espresso font-light">{{ review.count }}</div>
          <p class="text-xs text-coffee-softGray">杯咖啡</p>
        </div>

        <!-- Favorite Type -->
        <div v-if="review.favorite_coffee_type" class="p-4 bg-coffee-cream/50 rounded-xl text-center">
          <p class="text-[9px] uppercase tracking-widest text-coffee-softGray mb-1">本周最爱</p>
          <p class="font-serif text-lg text-coffee-espresso">{{ getCoffeeTypeLabel(review.favorite_coffee_type) }}</p>
        </div>

        <!-- Coffee Types -->
        <div v-if="review.coffee_types && review.coffee_types.length > 0" class="space-y-3">
          <h3 class="text-[10px] uppercase tracking-widest font-semibold text-coffee-softGray">冲煮类型</h3>
          <div class="space-y-2">
            <div v-for="item in review.coffee_types" :key="item.coffee_type" class="flex items-center justify-between">
              <span class="text-sm text-coffee-charcoal">{{ getCoffeeTypeLabel(item.coffee_type) }}</span>
              <span class="text-xs text-coffee-softGray">{{ item.count }}杯</span>
            </div>
          </div>
        </div>

        <!-- Flavor Tags -->
        <div v-if="review.flavor_tags && review.flavor_tags.length > 0" class="space-y-3">
          <h3 class="text-[10px] uppercase tracking-widest font-semibold text-coffee-softGray">风味偏好</h3>
          <div class="flex flex-wrap gap-2">
            <span
              v-for="tag in review.flavor_tags"
              :key="tag.name"
              class="px-3 py-1 bg-coffee-cream/60 rounded-full text-xs text-coffee-espresso"
            >{{ tag.label }}</span>
          </div>
        </div>

        <!-- Mood Tags -->
        <div v-if="review.mood_tags && review.mood_tags.length > 0" class="space-y-3">
          <h3 class="text-[10px] uppercase tracking-widest font-semibold text-coffee-softGray">心情记录</h3>
          <div class="flex flex-wrap gap-2">
            <span
              v-for="tag in review.mood_tags"
              :key="tag.tag"
              class="px-3 py-1 bg-coffee-cream/60 rounded-full text-xs text-coffee-espresso"
            >{{ getMoodLabel(tag.tag) }}</span>
          </div>
        </div>

        <!-- Trend -->
        <div v-if="review.trend" class="p-4 bg-coffee-cream/30 rounded-xl">
          <p class="text-sm text-coffee-charcoal leading-relaxed">{{ review.trend }}</p>
        </div>

        <!-- Memory -->
        <div v-if="review.memory" class="p-4 bg-coffee-cream/50 rounded-xl text-center">
          <p class="text-sm text-coffee-espresso font-medium">{{ review.memory }}</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import request from '@/api/request'
import { ArrowLeft, Coffee, Loader2 } from 'lucide-vue-next'

const router = useRouter()

interface WeeklyReviewCoffeeType {
  coffee_type: string
  count: number
}

interface WeeklyReviewFlavorTag {
  name: string
  label: string
  count: number
}

interface WeeklyReviewLifestyleTag {
  tag: string
  count: number
}

interface WeeklyReviewFlavorProfile {
  acidity: number
  bitterness: number
  sweetness: number
  body: number
  aroma: number
  aftertaste: number
}

interface WeeklyReviewResponse {
  week: string
  week_start: string
  week_end: string
  count: number
  favorite_coffee_type: string
  coffee_types: WeeklyReviewCoffeeType[]
  flavor_tags: WeeklyReviewFlavorTag[]
  mood_tags: WeeklyReviewLifestyleTag[]
  scene_tags: WeeklyReviewLifestyleTag[]
  pairing_tags: WeeklyReviewLifestyleTag[]
  flavor_profile: WeeklyReviewFlavorProfile | null
  trend: string
  memory: string
}

const loading = ref(true)
const review = ref<WeeklyReviewResponse | null>(null)

const coffeeTypeLabels: Record<string, string> = {
  'Pour Over': '手冲',
  'Latte': '拿铁',
  'Americano': '美式',
  'Cold Brew': '冷萃',
  'Espresso': '浓缩',
  'Dirty': '脏咖啡',
  'Cappuccino': '卡布奇诺',
  'Flat White': '馥芮白'
}

const moodLabels: Record<string, string> = {
  'Calm': '平静',
  'Focused': '专注',
  'Tired': '疲惫',
  'Happy': '开心',
  'Rainy': '阴雨',
  'Slow': '慢活',
  'Productive': '高效'
}

function getCoffeeTypeLabel(type: string): string {
  return coffeeTypeLabels[type] || type
}

function getMoodLabel(mood: string): string {
  return moodLabels[mood] || mood
}

onMounted(async () => {
  try {
    const response = await request.get('/stats/weekly-review')
    review.value = response.data
  } catch (error) {
    console.error('Failed to load weekly review:', error)
  } finally {
    loading.value = false
  }
})
</script>
