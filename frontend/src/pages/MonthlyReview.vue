<template>
  <div class="flex-1 w-full flex flex-col justify-between bg-coffee-warmWhite text-coffee-charcoal">
    
    <!-- Header -->
    <div class="px-6 py-5 flex justify-between items-end border-b border-coffee-cream bg-coffee-warmWhite sticky top-0 z-10 select-none">
      <div>
        <span class="text-[9px] uppercase tracking-[0.25em] font-semibold text-coffee-softGray">Monthly Review</span>
        <h2 class="font-serif text-2xl font-light text-coffee-espresso leading-none mt-1">月度咖啡回顾</h2>
      </div>
      <div class="flex items-center gap-2">
        <button @click="prevMonth" class="w-7 h-7 rounded-full border border-coffee-latte/50 flex items-center justify-center hover:bg-coffee-cream transition-colors">
          <ChevronLeft class="w-3.5 h-3.5 text-coffee-brown" />
        </button>
        <span class="text-[11px] font-serif text-coffee-espresso tracking-wider min-w-[80px] text-center">{{ displayMonth }}</span>
        <button @click="nextMonth" class="w-7 h-7 rounded-full border border-coffee-latte/50 flex items-center justify-center hover:bg-coffee-cream transition-colors" :class="isCurrentMonth ? 'opacity-30 pointer-events-none' : ''">
          <ChevronRight class="w-3.5 h-3.5 text-coffee-brown" />
        </button>
      </div>
    </div>

    <!-- Scrollable Body -->
    <div class="flex-1 overflow-y-auto px-6 py-5 space-y-6 pb-24 scrollbar-none">

      <!-- Empty State -->
      <div v-if="!review || review.count === 0" class="flex flex-col items-center justify-center py-16 space-y-4 select-none">
        <div class="text-5xl opacity-30">☕</div>
        <p class="text-sm text-coffee-softGray font-light italic text-center leading-relaxed">
          这个月的咖啡手账还是空白的。<br>新的月份，新的开始。
        </p>
        <router-link to="/create" class="px-5 py-2 text-[10px] uppercase tracking-widest font-semibold bg-coffee-espresso text-coffee-warmWhite rounded-sm hover:bg-coffee-brown transition-colors">
          记录第一杯
        </router-link>
      </div>

      <!-- Review Content -->
      <template v-if="review && review.count > 0">

        <!-- Hero: Month Count + Keywords -->
        <div class="p-6 rounded-sm border border-coffee-latte/40 select-none" style="background: linear-gradient(135deg, rgba(215,196,168,0.15) 0%, rgba(231,111,81,0.04) 100%);">
          <div class="text-center space-y-3">
            <span class="text-[9px] uppercase tracking-[0.25em] font-semibold text-coffee-softGray block">本月咖啡手账</span>
            <div class="text-6xl font-serif text-coffee-espresso font-light leading-none">{{ review.count }}</div>
            <span class="text-[10px] text-coffee-brown italic">cups of coffee this month</span>
          </div>
          <!-- Keywords -->
          <div v-if="review.keywords && review.keywords.length > 0" class="mt-5 flex flex-wrap justify-center gap-1.5">
            <span v-for="kw in review.keywords" :key="kw" class="px-3 py-1 bg-coffee-cream/70 border border-coffee-latte/30 rounded-full text-[10px] font-medium text-coffee-espresso tracking-wider">{{ kw }}</span>
          </div>
        </div>

        <!-- AI Editorial Summary -->
        <div v-if="monthlyReviewAI" class="space-y-3">
          <div class="flex items-center gap-2.5 select-none">
            <div class="h-px flex-1 max-w-[24px]" style="background: linear-gradient(to right, rgba(92,61,46,0.35), transparent);"></div>
            <h3 class="text-[10px] uppercase tracking-[0.2em] font-bold text-coffee-espresso">月度回顾 / Monthly Editorial</h3>
            <div class="h-px flex-1" style="background: linear-gradient(to left, rgba(92,61,46,0.35), transparent);"></div>
          </div>
          <div class="p-5 bg-coffee-cream/40 rounded-sm border border-coffee-latte/25 relative">
            <div class="text-[72px] leading-none font-serif text-coffee-latte/20 absolute top-2 left-3 select-none">&ldquo;</div>
            <p class="text-[15px] text-coffee-espresso leading-[1.8] font-light relative z-10 pl-2">{{ monthlyReviewAI }}</p>
            <div class="flex items-center justify-center gap-2 mt-4 select-none">
              <div class="h-px w-8 bg-coffee-latte/30"></div>
              <div class="w-1 h-1 rounded-full bg-coffee-latte/40"></div>
              <div class="h-px w-8 bg-coffee-latte/30"></div>
            </div>
          </div>
        </div>

        <!-- Coffee Type Breakdown -->
        <div class="space-y-3">
          <div class="flex items-center gap-2.5 select-none">
            <div class="h-px flex-1 max-w-[24px]" style="background: linear-gradient(to right, rgba(92,61,46,0.35), transparent);"></div>
            <h3 class="text-[10px] uppercase tracking-[0.2em] font-bold text-coffee-espresso">冲煮偏好 / Brew Style</h3>
            <div class="h-px flex-1" style="background: linear-gradient(to left, rgba(92,61,46,0.35), transparent);"></div>
          </div>
          <div v-if="review.coffee_types && review.coffee_types.length > 0" class="space-y-2.5">
            <div v-for="(item, idx) in review.coffee_types" :key="item.coffee_type">
              <div class="flex justify-between text-xs mb-1">
                <span>{{ coffeeTypeLabel(item.coffee_type) }}</span>
                <span class="font-semibold text-coffee-espresso">{{ item.count }} 杯 ({{ Math.round((item.count / review.count) * 100) }}%)</span>
              </div>
              <div class="w-full h-1.5 bg-coffee-cream rounded-full overflow-hidden select-none">
                <div class="h-full rounded-full transition-all duration-700"
                     :class="barColors[idx] || barColors[barColors.length - 1]"
                     :style="`width: ${Math.round((item.count / review.count) * 100)}%;`"></div>
              </div>
            </div>
          </div>
        </div>

        <!-- Flavor Tags -->
        <div v-if="review.flavor_tags && review.flavor_tags.length > 0" class="space-y-3">
          <div class="flex items-center gap-2.5 select-none">
            <div class="h-px flex-1 max-w-[24px]" style="background: linear-gradient(to right, rgba(92,61,46,0.35), transparent);"></div>
            <h3 class="text-[10px] uppercase tracking-[0.2em] font-bold text-coffee-espresso">风味图谱 / Flavor Map</h3>
            <div class="h-px flex-1" style="background: linear-gradient(to left, rgba(92,61,46,0.35), transparent);"></div>
          </div>
          <div class="flex flex-wrap gap-1.5">
            <span v-for="(tag, idx) in review.flavor_tags" :key="tag.name"
              class="px-3 py-1.5 rounded-full border tracking-wider"
              :class="idx < 2 ? 'bg-coffee-latte/15 border-coffee-latte/50 text-coffee-espresso font-medium' : 'bg-coffee-cream/60 border-coffee-latte/30 text-coffee-softGray font-light'"
            >{{ tag.label }} / {{ tag.name }} ({{ tag.count }})</span>
          </div>
        </div>

        <!-- Radar Chart (Monthly Flavor Profile) -->
        <div v-if="review.flavor_profile" class="space-y-3">
          <div class="flex items-center gap-2.5 select-none">
            <div class="h-px flex-1 max-w-[24px]" style="background: linear-gradient(to right, rgba(92,61,46,0.35), transparent);"></div>
            <h3 class="text-[10px] uppercase tracking-[0.2em] font-bold text-coffee-espresso">本月风味雷达 / Radar</h3>
            <div class="h-px flex-1" style="background: linear-gradient(to left, rgba(92,61,46,0.35), transparent);"></div>
          </div>
          <div class="w-[200px] h-[200px] mx-auto flex items-center justify-center p-2 bg-coffee-cream/40 rounded-full border border-coffee-latte/30">
            <FlavorRadarChart
              :values="monthlyFlavorValues"
              :size="190"
              :dimensions="['酸度', '苦感', '甜感', '醇厚', '香气', '余韵']"
              :dot-radius="3.0"
              :label-font-size="8.5"
            />
          </div>
        </div>

        <!-- Top Coffee Names -->
        <div v-if="review.coffee_names && review.coffee_names.length > 0" class="space-y-3">
          <div class="flex items-center gap-2.5 select-none">
            <div class="h-px flex-1 max-w-[24px]" style="background: linear-gradient(to right, rgba(92,61,46,0.35), transparent);"></div>
            <h3 class="text-[10px] uppercase tracking-[0.2em] font-bold text-coffee-espresso">常喝咖啡 / Top Picks</h3>
            <div class="h-px flex-1" style="background: linear-gradient(to left, rgba(92,61,46,0.35), transparent);"></div>
          </div>
          <div class="space-y-2">
            <div v-for="(item, idx) in review.coffee_names" :key="item.coffee_name"
              class="flex items-center gap-3 px-3 py-2.5 rounded-sm border"
              :class="idx === 0 ? 'bg-coffee-cream/50 border-coffee-latte/40' : 'bg-coffee-cream/20 border-coffee-latte/20'">
              <span class="text-[10px] font-serif text-coffee-softGray w-4">{{ idx + 1 }}</span>
              <span class="flex-1 text-xs font-medium text-coffee-espresso">{{ item.coffee_name }}</span>
              <span class="text-[10px] text-coffee-brown font-semibold">{{ item.count }}杯</span>
            </div>
          </div>
        </div>

        <!-- Lifestyle Tags Summary -->
        <div v-if="hasLifestyleTags" class="space-y-3">
          <div class="flex items-center gap-2.5 select-none">
            <div class="h-px flex-1 max-w-[24px]" style="background: linear-gradient(to right, rgba(92,61,46,0.35), transparent);"></div>
            <h3 class="text-[10px] uppercase tracking-[0.2em] font-bold text-coffee-espresso">生活标签 / Lifestyle</h3>
            <div class="h-px flex-1" style="background: linear-gradient(to left, rgba(92,61,46,0.35), transparent);"></div>
          </div>
          <div class="space-y-2">
            <!-- Mood Tags -->
            <div v-if="review.mood_tags && review.mood_tags.length > 0" class="flex flex-wrap gap-1.5">
              <span v-for="tag in review.mood_tags" :key="tag.tag" class="px-2.5 py-1 rounded-full bg-amber-50 border border-amber-200/50 text-[10px] text-amber-800 font-medium">{{ tag.tag }} ({{ tag.count }})</span>
            </div>
            <!-- Scene Tags -->
            <div v-if="review.scene_tags && review.scene_tags.length > 0" class="flex flex-wrap gap-1.5">
              <span v-for="tag in review.scene_tags" :key="tag.tag" class="px-2.5 py-1 rounded-full bg-sky-50 border border-sky-200/50 text-[10px] text-sky-800 font-medium">{{ tag.tag }} ({{ tag.count }})</span>
            </div>
            <!-- Pairing Tags -->
            <div v-if="review.pairing_tags && review.pairing_tags.length > 0" class="flex flex-wrap gap-1.5">
              <span v-for="tag in review.pairing_tags" :key="tag.tag" class="px-2.5 py-1 rounded-full bg-rose-50 border border-rose-200/50 text-[10px] text-rose-800 font-medium">{{ tag.tag }} ({{ tag.count }})</span>
            </div>
          </div>
        </div>

        <!-- Top Weekday -->
        <div v-if="review.top_weekday" class="p-4 bg-coffee-cream/30 rounded-sm border border-coffee-latte/25 text-center select-none">
          <span class="text-[9px] uppercase tracking-[0.2em] font-semibold text-coffee-softGray block">最常喝咖啡的日子</span>
          <div class="text-xl font-serif text-coffee-espresso mt-1">{{ weekdayLabel(review.top_weekday) }}</div>
          <span class="text-[10px] text-coffee-brown italic">the day you love coffee most</span>
        </div>

        <!-- Share CTA -->
        <button
          @click="shareReview"
          class="w-full py-3.5 text-[10px] uppercase tracking-widest font-semibold bg-coffee-espresso text-coffee-warmWhite hover:bg-coffee-brown transition-all rounded-sm flex items-center justify-center gap-2"
        >
          <Share2 class="w-3.5 h-3.5" />
          <span>分享月度回顾</span>
        </button>

      </template>
    </div>

    <!-- Sticky Bottom Navigation Bar -->
    <div class="relative h-16 border-t border-coffee-cream/60 bg-coffee-warmWhite flex items-center z-30 select-none">
      <router-link to="/home" class="flex-1 flex flex-col items-center gap-0.5 text-coffee-softGray hover:text-coffee-brown transition-colors">
        <BookOpen class="w-5 h-5" />
        <span class="text-[9.5px] tracking-widest font-medium">咖啡日志</span>
      </router-link>
      <router-link to="/timeline" class="flex-1 flex flex-col items-center gap-0.5 text-coffee-softGray hover:text-coffee-brown transition-colors">
        <Calendar class="w-5 h-5" />
        <span class="text-[9.5px] tracking-widest font-medium">时间线</span>
      </router-link>
      <div class="flex-1 flex flex-col items-center">
        <router-link to="/create" class="flex flex-col items-center gap-1 -translate-y-4 group">
          <div class="rounded-full flex items-center justify-center shadow-lg ring-4 ring-coffee-warmWhite transition-transform duration-200 group-hover:scale-105 group-active:scale-95"
               style="width:52px;height:52px;background:linear-gradient(145deg,#E76F51,#D4623E);">
            <Plus class="w-5 h-5 text-white" />
          </div>
        </router-link>
      </div>
      <router-link to="/stats" class="flex-1 flex flex-col items-center gap-0.5 text-coffee-softGray hover:text-coffee-brown transition-colors">
        <BarChart3 class="w-5 h-5" />
        <span class="text-[9.5px] tracking-widest font-medium">咖迹</span>
      </router-link>
      <router-link to="/profile" class="flex-1 flex flex-col items-center gap-0.5 text-coffee-softGray hover:text-coffee-brown transition-colors">
        <User class="w-5 h-5" />
        <span class="text-[9.5px] tracking-widest font-medium">我的</span>
      </router-link>
    </div>

  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, watch } from 'vue'
import { useCoffeeLogStore } from '@/stores/coffeeLog'
import FlavorRadarChart from '@/components/charts/FlavorRadarChart.vue'
import { coffeeTypeLabel } from '@/constants/coffee'
import { BookOpen, Calendar, BarChart3, Plus, User, Share2, ChevronLeft, ChevronRight } from 'lucide-vue-next'

const store = useCoffeeLogStore()

const barColors = ['bg-coffee-brown', 'bg-coffee-brown/60', 'bg-coffee-brown/35', 'bg-coffee-brown/20']

// Month navigation
const now = new Date()
const currentYear = now.getFullYear()
const currentMonth = now.getMonth() + 1
const selectedYear = ref(currentYear)
const selectedMonth = ref(currentMonth)

const isCurrentMonth = computed(() => selectedYear.value === currentYear && selectedMonth.value === currentMonth)

const monthParam = computed(() => {
  return `${selectedYear.value}-${String(selectedMonth.value).padStart(2, '0')}`
})

const displayMonth = computed(() => {
  return `${selectedYear.value}.${String(selectedMonth.value).padStart(2, '0')}`
})

function prevMonth() {
  if (selectedMonth.value === 1) {
    selectedMonth.value = 12
    selectedYear.value--
  } else {
    selectedMonth.value--
  }
}

function nextMonth() {
  if (isCurrentMonth.value) return
  if (selectedMonth.value === 12) {
    selectedMonth.value = 1
    selectedYear.value++
  } else {
    selectedMonth.value++
  }
}

const review = computed(() => store.monthlyReview)
const monthlyReviewAI = computed(() => store.monthlyReviewAI)

const monthlyFlavorValues = computed(() => {
  if (review.value?.flavor_profile) {
    const p = review.value.flavor_profile
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

const hasLifestyleTags = computed(() => {
  return (review.value?.mood_tags && review.value.mood_tags.length > 0) ||
    (review.value?.scene_tags && review.value.scene_tags.length > 0) ||
    (review.value?.pairing_tags && review.value.pairing_tags.length > 0)
})

const weekdayNamesMap: Record<number, string> = {
  1: '周日', 2: '周一', 3: '周二', 4: '周三',
  5: '周四', 6: '周五', 7: '周六'
}

function weekdayLabel(weekday: number) {
  return weekdayNamesMap[weekday] || ''
}

function shareReview() {
  if (!review.value) return
  const lines: string[] = []
  lines.push(`☕ ${displayMonth.value} 月度咖啡回顾`)
  lines.push(`本月 ${review.value.count} 杯`)
  if (review.value.favorite_coffee_type) {
    lines.push(`最常喝：${coffeeTypeLabel(review.value.favorite_coffee_type)}`)
  }
  if (review.value.keywords && review.value.keywords.length > 0) {
    lines.push(`关键词：${review.value.keywords.join(' · ')}`)
  }
  if (monthlyReviewAI.value) {
    lines.push(monthlyReviewAI.value)
  }
  lines.push('')
  lines.push('— MY COFFEE LOG')
  const text = lines.join('\n')
  if (navigator.share) {
    navigator.share({ title: 'My Coffee Monthly Review', text }).catch(() => {})
  } else {
    navigator.clipboard.writeText(text).then(() => {
      alert('月度回顾已复制到剪贴板！')
    }).catch(() => {})
  }
}

async function loadData() {
  await store.fetchMonthlyReview(monthParam.value)
}

onMounted(loadData)

watch(monthParam, loadData)
</script>

<style scoped>
/* Monthly Review page specific scoped styles */
</style>
