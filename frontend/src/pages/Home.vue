<template>
  <div class="flex-1 w-full flex flex-col justify-between bg-coffee-warmWhite text-coffee-charcoal">
    
    <!-- Home Sticky Header -->
    <div class="px-6 py-5 flex justify-between items-end border-b border-coffee-cream/80 bg-coffee-warmWhite/80 backdrop-blur-md sticky top-0 z-20">
      <div>
        <span class="text-[9px] uppercase tracking-[0.25em] font-semibold text-coffee-softGray">{{ todayDate }}</span>
        <h2 class="font-serif text-2xl font-light text-coffee-espresso leading-none mt-1">{{ greeting }}</h2>
      </div>
      <div class="flex items-center gap-4">
        <router-link to="/timeline" class="text-coffee-espresso hover:text-coffee-brown transition-colors">
          <Search class="w-5 h-5" />
        </router-link>
        <router-link to="/profile" class="block">
          <img :src="userAvatar" alt="Avatar" class="w-8 h-8 rounded-full object-cover border border-coffee-latte hover:opacity-85 transition-opacity">
        </router-link>
      </div>
    </div>

    <!-- Scrollable Body Content -->
    <div class="flex-1 overflow-y-auto px-6 py-5 space-y-5 pb-24 scrollbar-none">
      
      <!-- Today's Record Status Card -->
      <div v-if="store.todayLog" class="block">
        <router-link 
          :to="'/coffee/' + store.todayLog.id" 
          class="block group relative overflow-hidden bg-coffee-espresso flex flex-col justify-between p-5 rounded-sm shadow-sm hover:scale-[0.99] transition-transform duration-300"
          style="min-height: 180px;"
        >
          <img :src="store.todayLog.image_url" class="absolute inset-0 w-full h-full object-cover opacity-35 mix-blend-luminosity filter saturate-50 group-hover:scale-105 transition-all duration-700">
          
          <!-- Paper grain overlay -->
          <div class="absolute inset-0 opacity-[0.04]" style="background-image: radial-gradient(rgba(255,242,219,0.6) 1px, transparent 0); background-size: 18px 18px;"></div>

          <div class="z-10 flex justify-between items-start">
            <div class="flex items-center gap-2">
              <span class="w-1.5 h-1.5 rounded-full bg-green-400 animate-pulse"></span>
              <span class="text-[9px] uppercase tracking-[0.2em] font-semibold text-coffee-latte">今日已记录</span>
            </div>
            <span class="text-[9px] text-coffee-warmWhite border border-coffee-latte/40 px-2 py-0.5 rounded-full font-light font-serif italic uppercase">查看详情</span>
          </div>
          
          <div class="z-10 space-y-1.5">
            <h3 class="font-serif text-2xl text-coffee-warmWhite font-light truncate max-w-[280px]">{{ store.todayLog.coffee_name }}</h3>
            <div class="flex items-center gap-1.5 text-[10px] text-coffee-latte tracking-wider">
              <span class="font-serif italic">{{ store.todayLog.coffee_type }}</span>
              <span>·</span>
              <span class="truncate max-w-[180px]">{{ store.todayLog.shop_name.split(',')[0] }}</span>
              <span>·</span>
              <span class="inline-flex items-center gap-1">
                <AppIcon :name="moodIconName(store.todayLog.mood)" :size="11" />
                {{ moodLabel(store.todayLog.mood) }}
              </span>
            </div>
          </div>
        </router-link>
      </div>

      <!-- Hero: Last log when no today's record -->
      <div v-else-if="lastLog" class="block">
        <router-link 
          :to="'/coffee/' + lastLog.id" 
          class="block group relative overflow-hidden bg-coffee-espresso flex flex-col justify-between p-5 rounded-sm shadow-sm hover:scale-[0.99] transition-transform duration-300"
          style="min-height: 160px;"
        >
          <img :src="lastLog.image_url" class="absolute inset-0 w-full h-full object-cover opacity-35 mix-blend-luminosity filter saturate-50 group-hover:scale-105 transition-all duration-700">
          
          <div class="absolute inset-0 opacity-[0.04]" style="background-image: radial-gradient(rgba(255,242,219,0.6) 1px, transparent 0); background-size: 18px 18px;"></div>

          <div class="z-10 flex justify-between items-start">
            <span class="text-[9px] uppercase tracking-[0.2em] font-semibold text-coffee-latte">最近风味记录</span>
            <span class="text-[9px] text-coffee-warmWhite border border-coffee-latte/40 px-2 py-0.5 rounded-full font-light font-serif italic uppercase">查看详情</span>
          </div>
          
          <div class="z-10 space-y-1.5">
            <h3 class="font-serif text-2xl text-coffee-warmWhite font-light truncate max-w-[280px]">{{ lastLog.coffee_name }}</h3>
            <div class="flex items-center gap-1.5 text-[10px] text-coffee-latte tracking-wider">
              <span class="font-serif italic">{{ lastLog.coffee_type }}</span>
              <span>·</span>
              <span class="truncate max-w-[180px]">{{ lastLog.shop_name.split(',')[0] }}</span>
            </div>
          </div>
        </router-link>

        <!-- Quick record prompt -->
        <router-link 
          to="/create"
          class="block mt-2.5 group relative overflow-hidden p-4 rounded-sm border border-dashed border-coffee-latte/40 bg-coffee-cream/20 hover:bg-coffee-cream/35 transition-colors"
        >
          <div class="flex items-center justify-between">
            <div class="flex items-center gap-2.5">
              <div class="w-8 h-8 rounded-full flex items-center justify-center" style="background: linear-gradient(145deg, #E76F51, #D4623E);">
                <Plus class="w-4 h-4 text-white" />
              </div>
              <div>
                <span class="text-xs font-serif text-coffee-espresso font-light">记录今日咖啡</span>
                <span class="block text-[9px] text-coffee-softGray tracking-wider">今天还未记录，点击开始</span>
              </div>
            </div>
            <span class="text-[9px] text-coffee-latte font-serif italic">Quick Log →</span>
          </div>
        </router-link>
      </div>

      <!-- Empty welcome card when no logs at all -->
      <div v-else>
        <router-link 
          to="/create" 
          class="block group relative overflow-hidden h-[200px] flex flex-col justify-between p-5 rounded-2xl shadow-md hover:scale-[0.99] transition-transform duration-300"
          style="background: linear-gradient(145deg, #5C3D2E 0%, #3A2418 55%, #6B4A38 100%);"
        >
          <!-- Paper grain overlay -->
          <div class="absolute inset-0 opacity-[0.05]" style="background-image: radial-gradient(rgba(255,242,219,0.6) 1px, transparent 0); background-size: 18px 18px;"></div>

          <!-- Subtle side gradient sheen -->
          <div class="absolute inset-y-0 right-0 w-1/2 pointer-events-none" style="background: linear-gradient(to left, rgba(231,111,81,0.15), transparent);"></div>

          <!-- Large watermark letter -->
          <div class="absolute right-3 bottom-0 font-serif leading-none select-none pointer-events-none translate-y-3 text-coffee-warmWhite/[0.04]" style="font-size: 140px;">C</div>

          <!-- Corner bracket decorations -->
          <div class="absolute top-3.5 left-3.5 w-4 h-4 border-t border-l border-coffee-latte/25"></div>
          <div class="absolute top-3.5 right-3.5 w-4 h-4 border-t border-r border-coffee-latte/25"></div>
          <div class="absolute bottom-3.5 left-3.5 w-4 h-4 border-b border-l border-coffee-latte/25"></div>
          <div class="absolute bottom-3.5 right-3.5 w-4 h-4 border-b border-r border-coffee-latte/25"></div>

          <!-- Header row -->
          <div class="relative z-10 flex justify-between items-start">
            <span class="text-[9px] uppercase tracking-[0.22em] font-semibold text-coffee-latte/80">欢迎开启咖啡手账</span>
            <span class="text-[9px] text-coffee-warmWhite/90 border border-coffee-latte/30 px-2.5 py-0.5 font-light font-serif italic tracking-wider">开始记录</span>
          </div>

          <!-- Hairline divider -->
          <div class="relative z-10 flex items-center gap-2">
            <div class="w-5 h-px bg-coffee-latte/30"></div>
            <div class="flex-1 h-px" style="background: linear-gradient(to right, rgba(255,242,219,0.2), transparent);"></div>
          </div>

          <!-- Bottom text -->
          <div class="relative z-10 text-left space-y-2">
            <h3 class="font-serif text-[1.6rem] text-coffee-warmWhite font-light leading-snug tracking-wide">记录你的第一杯美好</h3>
            <div class="flex items-center gap-2">
              <div class="w-5 h-px bg-coffee-latte/40 flex-shrink-0"></div>
              <p class="text-[10px] text-coffee-latte/75 tracking-widest font-light">记录咖啡，也就是在记录你专属的生活格调。</p>
            </div>
          </div>
        </router-link>
      </div>

      <!-- AI Lifestyle Quote (Double border editorial style) -->
      <div class="relative overflow-hidden rounded-sm" style="background: linear-gradient(145deg, #E76F51 0%, #D4623E 55%, #E87D60 100%); border: 4px double rgba(255,242,219,0.35);">
        <!-- Decorative large quotation mark -->
        <div class="absolute top-0 left-3 font-serif leading-none select-none pointer-events-none text-coffee-warmWhite/[0.18]" style="font-size: 96px; line-height: 1;">&ldquo;</div>
        <!-- Subtle bottom-right closing quote -->
        <div class="absolute bottom-0 right-3 font-serif leading-none select-none pointer-events-none text-coffee-warmWhite/[0.10] rotate-180" style="font-size: 64px; line-height: 1;">&ldquo;</div>

        <div class="relative z-10 px-5 pt-5 pb-4 text-center space-y-3">
          <!-- Title with flanking lines -->
          <div class="flex items-center justify-center gap-2.5">
            <div class="h-px flex-1 max-w-[36px]" style="background: linear-gradient(to left, rgba(255,242,219,0.55), transparent);"></div>
            <span class="text-[9px] uppercase tracking-[0.28em] font-bold text-coffee-warmWhite">生活手账摘要</span>
            <div class="h-px flex-1 max-w-[36px]" style="background: linear-gradient(to right, rgba(255,242,219,0.55), transparent);"></div>
          </div>

          <!-- Quote text from AI -->
          <p class="font-serif italic text-[0.82rem] text-coffee-warmWhite leading-[1.75] px-2">
            {{ displayQuote }}
          </p>

          <!-- Attribution with ornament -->
          <div class="flex items-center justify-center gap-2 pt-0.5">
            <div class="w-4 h-px bg-coffee-warmWhite/30"></div>
            <span class="text-[8.5px] text-coffee-warmWhite/70 font-semibold tracking-[0.2em] uppercase">{{ quoteAttribution }}</span>
            <div class="w-4 h-px bg-coffee-warmWhite/30"></div>
          </div>
        </div>
      </div>

      <!-- Monthly Dashboard Grid -->
      <div class="space-y-3 select-none">
        <div class="flex items-center gap-2">
          <div class="w-4 h-px bg-coffee-softGray/40"></div>
          <span class="text-[9px] uppercase tracking-[0.22em] font-semibold text-coffee-softGray">本月咖啡仪表</span>
          <div class="flex-1 h-px bg-coffee-cream"></div>
        </div>

        <div class="grid grid-cols-2 gap-3">
          <!-- Month brews count -->
          <div class="p-4 bg-coffee-cream/40 border border-coffee-cream/80 text-center rounded-sm">
            <div class="text-3xl font-serif font-light italic text-coffee-espresso">{{ store.monthBrews }}</div>
            <div class="text-[9px] uppercase tracking-wider text-coffee-softGray mt-1 font-semibold">本月冲煮</div>
          </div>
          <!-- Favorite coffee type -->
          <div class="p-4 bg-coffee-cream/40 border border-coffee-cream/80 text-center rounded-sm">
            <div class="text-3xl font-serif font-light italic text-coffee-espresso truncate px-1">{{ coffeeTypeShortLabel(store.favoriteCoffeeType) }}</div>
            <div class="text-[9px] uppercase tracking-wider text-coffee-softGray mt-1 font-semibold">最常喝</div>
          </div>
        </div>

        <!-- Recent Flavor Tags -->
        <div v-if="store.recentFlavorTags.length > 0" class="p-4 bg-coffee-cream/30 border border-coffee-cream/60 rounded-sm">
          <div class="text-[9px] uppercase tracking-wider text-coffee-softGray font-semibold mb-2.5">偏好风味图谱</div>
          <div class="flex flex-wrap gap-2">
            <span 
              v-for="(tag, idx) in store.recentFlavorTags" 
              :key="tag.name"
              class="px-3 py-1.5 rounded-full text-[11px] font-medium tracking-wide border"
              :class="idx === 0 ? 'bg-coffee-latte/15 text-coffee-espresso border-coffee-latte/40 font-semibold' : 'bg-coffee-cream/50 text-coffee-brown border-coffee-cream/80 font-light'"
            >{{ tag.label }} <span class="text-[9px] opacity-60">{{ tag.count }}杯</span></span>
          </div>
        </div>
      </div>

      <!-- Recent Logs Section -->
      <div class="space-y-4">
        <div class="flex justify-between items-end border-b border-coffee-cream pb-1.5">
          <h3 class="font-serif text-lg font-light text-coffee-espresso">最近咖啡手账</h3>
          <router-link to="/timeline" class="text-[9px] uppercase tracking-widest text-coffee-brown font-semibold hover:text-coffee-espresso transition-colors">查看时间线</router-link>
        </div>
        
        <div class="space-y-3">
          <!-- Iterate over top 3 recent logs -->
          <router-link 
            v-for="log in recentLogs" 
            :key="log.id"
            :to="'/coffee/' + log.id" 
            class="block editorial-border p-3.5 bg-coffee-cream/20 flex gap-4 items-center rounded-sm hover:bg-coffee-cream/35 transition-colors"
          >
            <img :src="log.image_url" class="w-16 h-16 object-cover rounded-sm border border-coffee-cream flex-shrink-0">
            <div class="flex-1 min-w-0 space-y-1">
              <div class="flex justify-between items-center">
                <span class="text-[9px] uppercase tracking-widest font-semibold text-coffee-softGray">{{ coffeeTypeLabel(log.coffee_type) }}</span>
                <span class="text-[8px] font-mono text-coffee-softGray">{{ formatMonthDay(log.drink_date) }}</span>
              </div>
              <h4 class="font-serif text-base font-light text-coffee-espresso truncate italic leading-tight">{{ log.coffee_name }}</h4>
              <p class="text-[10px] text-coffee-brown font-light truncate leading-relaxed">{{ log.notes }}</p>
              <div class="flex gap-2 items-center text-[9px] text-coffee-softGray pt-0.5">
                <span class="inline-flex items-center gap-1 px-1.5 py-0.5 bg-coffee-cream rounded-sm text-coffee-espresso">
                  <AppIcon :name="moodIconName(log.mood)" :size="11" />
                  {{ moodLabel(log.mood) }}
                </span>
                <span class="truncate max-w-[120px]">在 {{ log.shop_name.split(',')[0] }}</span>
              </div>
            </div>
          </router-link>
        </div>
      </div>

    </div>

    <!-- Sticky Bottom Navigation Bar -->
    <div class="relative h-16 border-t border-coffee-cream/60 bg-coffee-warmWhite flex items-center z-30 select-none">

      <!-- Left: 咖啡日志 -->
      <router-link to="/home" class="flex-1 flex flex-col items-center gap-0.5 text-coffee-brown">
        <BookOpen class="w-5 h-5" />
        <span class="text-[9.5px] tracking-widest font-medium">咖啡日志</span>
      </router-link>

      <!-- Left: 时间线 -->
      <router-link to="/timeline" class="flex-1 flex flex-col items-center gap-0.5 text-coffee-softGray hover:text-coffee-brown transition-colors">
        <Calendar class="w-5 h-5" />
        <span class="text-[9.5px] tracking-widest font-medium">时间线</span>
      </router-link>

      <!-- Center FAB: 记录 (elevated above bar) -->
      <div class="flex-1 flex flex-col items-center">
        <router-link
          to="/create"
          class="flex flex-col items-center gap-1 -translate-y-4 group"
        >
          <div class="w-13 h-13 rounded-full flex items-center justify-center shadow-lg ring-4 ring-coffee-warmWhite transition-transform duration-200 group-hover:scale-105 group-active:scale-95"
               style="width: 52px; height: 52px; background: linear-gradient(145deg, #E76F51, #D4623E);">
            <Plus class="w-5 h-5 text-white" />
          </div>
        </router-link>
      </div>

      <!-- Right: 咖迹 -->
      <router-link to="/stats" class="flex-1 flex flex-col items-center gap-0.5 text-coffee-softGray hover:text-coffee-brown transition-colors">
        <BarChart3 class="w-5 h-5" />
        <span class="text-[9.5px] tracking-widest font-medium">咖迹</span>
      </router-link>

      <!-- Right: 个人中心 -->
      <router-link to="/profile" class="flex-1 flex flex-col items-center gap-0.5 text-coffee-softGray hover:text-coffee-brown transition-colors">
        <User class="w-5 h-5" />
        <span class="text-[9.5px] tracking-widest font-medium">我的</span>
      </router-link>

    </div>

  </div>
</template>

<script setup lang="ts">
import { computed, onMounted } from 'vue'
import { useCoffeeLogStore } from '@/stores/coffeeLog'
import { useAuthStore } from '@/stores/auth'
import { coffeeTypeShortLabel, coffeeTypeLabel, moodLabel, moodIconName } from '@/constants/coffee'
import AppIcon from '@/components/AppIcon.vue'
import { BookOpen, Calendar, BarChart3, Plus, Search, User } from 'lucide-vue-next'

const store = useCoffeeLogStore()
const authStore = useAuthStore()

// Fetch data on mount
onMounted(async () => {
  await Promise.all([
    store.fetchStats(),
    store.logs.length === 0
      ? store.fetchLogs({ page: 1, page_size: 10 })
      : Promise.resolve()
  ])
  await store.fetchLifestyleQuote()
})

// Dynamic Dates
const todayDate = computed(() => {
  const date = new Date()
  const months = ["January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"]
  return `${months[date.getMonth()]} ${date.getDate()}, ${date.getFullYear()}`
})

// Greeting based on time of day
const greeting = computed(() => {
  const hour = new Date().getHours()
  if (hour < 6) return '夜深了，来杯咖啡'
  if (hour < 12) return '早安，今日风味志'
  if (hour < 18) return '午后风味志'
  return '晚安，今日风味志'
})

// Logs Logic
const lastLog = computed(() => store.logs[0] || null)
const recentLogs = computed(() => store.logs.slice(0, 3))

// AI Lifestyle Quote display
const displayQuote = computed(() => {
  if (store.lifestyleQuote) {
    return `\u201c${store.lifestyleQuote}\u201d`
  }
  if (store.logs.length > 0) {
    return '\u201c每一杯咖啡，都是生活赠予的温柔时刻。\u201d'
  }
  return '\u201c咖啡的香气是属于清晨与安静午后的赞美诗。期待你写下今日的第一篇味觉手账。\u201d'
})

const quoteAttribution = computed(() => {
  if (store.monthBrews >= 10) return '醇厚的生活律动'
  if (store.monthBrews >= 3) return '温柔的日常节奏'
  return '期待你的风味故事'
})

// Avatar Logic
const userAvatar = computed(() => authStore.user?.avatar_url || 'https://images.unsplash.com/photo-1534528741775-53994a69daeb?auto=format&fit=crop&q=80&w=120')

// Helpers
const formatMonthDay = (dateStr: string) => {
  const date = new Date(dateStr)
  const months = ["Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"]
  return `${months[date.getMonth()]} ${date.getDate()}`
}

</script>

<style scoped>
/* Scoped custom styling if any */
</style>
