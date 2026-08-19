<template>
  <div class="flex-1 w-full flex flex-col bg-coffee-warmWhite text-coffee-charcoal min-h-screen">

    <!-- Header -->
    <div class="px-5 py-3.5 border-b border-coffee-cream flex justify-between items-center bg-coffee-warmWhite/95 backdrop-blur-sm sticky top-0 z-10 select-none">
      <button
        @click="router.push('/home')"
        class="grid w-9 h-9 place-items-center -ml-2 rounded-full text-coffee-brown hover:bg-coffee-cream/60 hover:text-coffee-espresso transition-colors"
        aria-label="返回首页"
      >
        <Home class="w-5 h-5" />
      </button>
      <div class="text-center leading-none">
        <h1 class="font-serif text-[17px] font-semibold tracking-wide text-coffee-espresso">记录成功</h1>
      </div>
      <div class="w-9"></div>
    </div>

    <!-- Success Content -->
    <div class="flex-1 overflow-y-auto px-6 py-8">

      <!-- Success Icon & Title -->
      <div class="text-center mb-8 animate-fade-in">
        <div class="inline-grid w-16 h-16 place-items-center rounded-full bg-green-50 border border-green-200 mb-4">
          <Check class="w-8 h-8 text-green-600" />
        </div>
        <h2 class="font-serif text-xl font-semibold text-coffee-espresso mb-2">
          {{ progress?.is_first_record ? '这是你的第一杯咖啡 ☕' : '这杯已经被收进你的咖啡手账' }}
        </h2>
        <p class="text-sm text-coffee-softGray">
          {{ progress?.is_first_record ? '记录之旅从此开始' : getMonthText() }}
        </p>
      </div>

      <!-- Record Summary Card -->
      <div v-if="log" class="rounded-lg border border-coffee-cream bg-white/60 p-5 mb-6 shadow-sm">
        <div class="flex gap-4">
          <!-- Cover Image -->
          <div class="w-20 h-20 rounded-md overflow-hidden flex-shrink-0 bg-coffee-cream">
            <img
              v-if="log.image_url"
              :src="log.image_url"
              class="w-full h-full object-cover"
              :alt="log.coffee_name"
            >
            <div v-else class="w-full h-full grid place-items-center">
              <Coffee class="w-8 h-8 text-coffee-latte" />
            </div>
          </div>
          <!-- Info -->
          <div class="flex-1 min-w-0">
            <h3 class="font-serif text-base font-semibold text-coffee-espresso truncate">{{ log.coffee_name || getTypeName(log.coffee_type) }}</h3>
            <div class="mt-1 flex items-center gap-2 text-xs text-coffee-softGray">
              <span>{{ getTypeName(log.coffee_type) }}</span>
              <span v-if="log.mood" class="flex items-center gap-1">
                <span class="w-1 h-1 rounded-full bg-coffee-latte"></span>
                {{ getMoodName(log.mood) }}
              </span>
            </div>
            <div class="mt-2 text-xs text-coffee-softGray/70">{{ log.drink_date }}</div>
          </div>
        </div>
      </div>

      <!-- Insight Progress -->
      <div v-if="progress && progress.next_insight_name" class="rounded-lg border border-coffee-cream bg-coffee-cream/20 p-4 mb-6">
        <div class="flex items-center gap-2 mb-2">
          <Sparkles class="w-4 h-4 text-coffee-brown" />
          <span class="text-xs font-semibold text-coffee-espresso tracking-wide">下一项洞察</span>
        </div>
        <p class="text-sm text-coffee-espresso">
          再记录 <span class="font-semibold text-coffee-brown">{{ progress.next_insight_delta }}</span> 杯，即可生成「{{ progress.next_insight_name }}」
        </p>
        <div class="mt-3 h-1.5 bg-coffee-cream rounded-full overflow-hidden">
          <div
            class="h-full bg-coffee-brown rounded-full transition-all duration-500"
            :style="{ width: progressWidth + '%' }"
          ></div>
        </div>
      </div>

      <!-- Fun Stats -->
      <div v-if="progress" class="grid grid-cols-2 gap-3 mb-8">
        <div class="rounded-lg border border-coffee-cream bg-white/60 p-4 text-center">
          <div class="text-2xl font-serif font-semibold text-coffee-espresso">{{ progress.month_cup_count }}</div>
          <div class="mt-1 text-xs text-coffee-softGray">本月杯数</div>
        </div>
        <div class="rounded-lg border border-coffee-cream bg-white/60 p-4 text-center">
          <div class="text-2xl font-serif font-semibold text-coffee-espresso">{{ progress.total_count }}</div>
          <div class="mt-1 text-xs text-coffee-softGray">总计杯数</div>
        </div>
      </div>

    </div>

    <!-- Bottom Actions -->
    <div class="border-t border-coffee-cream bg-coffee-warmWhite/95 backdrop-blur-sm sticky bottom-0 z-10 select-none p-5 space-y-3">
      <button
        v-if="log"
        @click="router.push(`/coffee/${log.id}`)"
        class="w-full py-3 text-xs uppercase tracking-wider font-semibold bg-coffee-cream text-coffee-espresso hover:bg-coffee-latte transition-all rounded-sm"
      >
        查看详情
      </button>
      <div class="flex gap-3">
        <button
          v-if="log"
          @click="handleRebrew"
          class="flex-1 py-3 text-xs uppercase tracking-wider font-semibold border border-coffee-brown text-coffee-brown hover:bg-coffee-cream/40 transition-all rounded-sm"
        >
          再次冲煮
        </button>
        <button
          @click="router.push('/home')"
          class="flex-1 py-3 text-xs uppercase tracking-wider font-semibold bg-coffee-espresso text-coffee-warmWhite hover:bg-coffee-brown transition-all rounded-sm"
        >
          返回首页
        </button>
      </div>
    </div>

  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useCoffeeLogStore } from '@/stores/coffeeLog'
import { getRecordProgress, type RecordProgressResponse } from '@/api/stats'
import { Home, Check, Coffee, Sparkles } from 'lucide-vue-next'

const router = useRouter()
const route = useRoute()
const store = useCoffeeLogStore()

const logId = computed(() => Number(route.params.id))
const log = computed(() => store.getLogById(logId.value))
const progress = ref<RecordProgressResponse | null>(null)

const progressWidth = computed(() => {
  if (!progress.value || !progress.value.next_insight_name) return 100
  const total = progress.value.total_count + progress.value.next_insight_delta
  if (total === 0) return 0
  return Math.min(100, (progress.value.total_count / total) * 100)
})

const typeMap: Record<string, string> = {
  'Pour Over': '手冲', 'Latte': '拿铁', 'Americano': '美式',
  'Cold Brew': '冷萃', 'Espresso': '浓缩', 'Dirty': '脏咖啡',
  'Cappuccino': '卡布奇诺', 'Flat White': '馥芮白'
}

const moodMap: Record<string, string> = {
  'Calm': '平静', 'Energetic': '愉悦', 'Reflective': '沉浸', 'Tired': '疲惫'
}

function getTypeName(type: string) {
  return typeMap[type] || type
}

function getMoodName(mood: string) {
  return moodMap[mood] || mood
}

function getMonthText() {
  const cupCount = progress.value?.month_cup_count || 0
  return `这是你本月记录的第 ${cupCount} 杯`
}

function handleRebrew() {
  if (log.value) {
    router.push(`/create?from_log_id=${log.value.id}`)
  }
}

onMounted(async () => {
  if (logId.value && log.value) {
    try {
      progress.value = await getRecordProgress(logId.value)
    } catch (e) {
      console.error('Failed to fetch record progress:', e)
    }
  }
})
</script>

<style scoped>
@keyframes fadeIn {
  from { opacity: 0; transform: translateY(8px); }
  to { opacity: 1; transform: translateY(0); }
}
.animate-fade-in {
  animation: fadeIn 0.4s ease-out;
}
</style>
