<template>
  <div class="flex-1 w-full flex flex-col justify-between bg-coffee-warmWhite text-coffee-charcoal">
    
    <!-- Home Sticky Header -->
    <div class="px-6 py-5 flex justify-between items-end border-b border-coffee-cream/80 bg-coffee-warmWhite/80 backdrop-blur-md sticky top-0 z-20">
      <div>
        <span class="text-[9px] uppercase tracking-[0.25em] font-semibold text-coffee-softGray">{{ todayDate }}</span>
        <h2 class="font-serif text-2xl font-light text-coffee-espresso leading-none mt-1">今日风味志</h2>
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
    <div class="flex-1 overflow-y-auto px-6 py-5 space-y-6 pb-24 scrollbar-none">
      
      <!-- Hero Section: Dynamic banner for today's brew -->
      <div v-if="lastLog" class="block">
        <router-link 
          :to="'/coffee/' + lastLog.id" 
          class="block group relative overflow-hidden h-[190px] bg-coffee-espresso flex flex-col justify-between p-5 rounded-sm shadow-sm hover:scale-[0.99] transition-transform duration-300"
        >
          <img :src="lastLog.image_url" class="absolute inset-0 w-full h-full object-cover opacity-40 mix-blend-luminosity filter saturate-50 group-hover:scale-105 transition-all duration-700">
          
          <div class="z-10 flex justify-between items-start">
            <span class="text-[9px] uppercase tracking-[0.2em] font-semibold text-coffee-latte">今日风味首记</span>
            <span class="text-[9px] text-coffee-warmWhite border border-coffee-latte/40 px-2 py-0.5 rounded-full font-light font-serif italic uppercase">查看风账详情</span>
          </div>
          
          <div class="z-10 space-y-1">
            <h3 class="font-serif text-2xl text-coffee-warmWhite font-light truncate max-w-[280px]">{{ lastLog.coffee_name }}</h3>
            <div class="flex items-center gap-1.5 text-[10px] text-coffee-latte tracking-wider">
              <span class="font-serif italic">{{ lastLog.coffee_type }}</span>
              <span>•</span>
              <span class="truncate max-w-[180px]">{{ lastLog.shop_name.split(',')[0] }}</span>
            </div>
          </div>
        </router-link>
      </div>

      <!-- If no logs yet, show empty welcome card -->
      <div v-else>
        <router-link 
          to="/create" 
          class="block group relative overflow-hidden h-[190px] bg-coffee-espresso flex flex-col justify-between p-5 rounded-sm shadow-sm hover:scale-[0.99] transition-transform duration-300"
        >
          <div class="z-10 flex justify-between items-start">
            <span class="text-[9px] uppercase tracking-[0.2em] font-semibold text-coffee-latte">欢迎开启咖啡手账</span>
            <span class="text-[9px] text-coffee-warmWhite border border-coffee-latte/40 px-2 py-0.5 rounded-full font-light font-serif italic">开始记录</span>
          </div>
          <div class="z-10 text-left">
            <h3 class="font-serif text-2xl text-coffee-warmWhite font-light">记录你的第一杯美好</h3>
            <p class="text-[10px] text-coffee-latte tracking-wider mt-1">记录咖啡，也就是在记录你专属的生活格调。</p>
          </div>
        </router-link>
      </div>

      <!-- AI Tasting Summary Quote (Double border editorial style) -->
      <div class="p-5 bg-coffee-cream rounded-sm double-border text-center space-y-2.5">
        <span class="text-[9px] uppercase tracking-[0.25em] font-bold text-coffee-brown block">今日感官摘要</span>
        <p class="font-serif italic text-sm text-coffee-espresso leading-relaxed px-1">
          {{ lastLog ? `“今天你享用的是一杯${lastLog.coffee_name}${lastLog.coffee_type}。明亮高亢的香气像晨曦，口感与此时此刻😌 ${lastLog.mood} 的心境极佳契合。”` : '“咖啡的香气是属于清晨与安静午后的赞美诗。期待你写下今日的第一篇味觉手账，我将在此为你吟诵今日的风味总结。”' }}
        </p>
        <div class="text-[9px] text-coffee-softGray font-semibold tracking-wider uppercase">优雅平缓的生活律动</div>
      </div>

      <!-- Monthly Overview Grid -->
      <div class="grid grid-cols-3 gap-3 select-none">
        <div class="p-4 bg-coffee-cream/40 border border-coffee-cream/80 text-center rounded-sm">
          <div class="text-2xl font-serif text-coffee-espresso font-light">{{ store.monthBrews }}</div>
          <div class="text-[9px] uppercase tracking-wider text-coffee-softGray mt-1 font-semibold">本月冲煮</div>
        </div>
        <div class="p-4 bg-coffee-cream/40 border border-coffee-cream/80 text-center rounded-sm">
          <div class="text-2xl font-serif text-coffee-espresso font-light truncate px-1">{{ store.favoriteCoffeeType }}</div>
          <div class="text-[9px] uppercase tracking-wider text-coffee-softGray mt-1 font-semibold">最常喝</div>
        </div>
        <div class="p-4 bg-coffee-cream/40 border border-coffee-cream/80 text-center rounded-sm">
          <div class="text-2xl font-serif text-coffee-espresso font-light truncate px-1 uppercase">{{ store.favoriteFlavorTag }}</div>
          <div class="text-[9px] uppercase tracking-wider text-coffee-softGray mt-1 font-semibold">偏好风味</div>
        </div>
      </div>

      <!-- Recent Logs Section -->
      <div class="space-y-4">
        <div class="flex justify-between items-end border-b border-coffee-cream pb-1.5">
          <h3 class="font-serif text-lg font-light text-coffee-espresso">最近咖啡手账</h3>
          <router-link to="/timeline" class="text-[9px] uppercase tracking-widest text-coffee-brown font-semibold hover:text-coffee-espresso transition-colors">查看时间线</router-link>
        </div>
        
        <div class="space-y-3">
          <!-- Iterate over top 2 recent logs -->
          <router-link 
            v-for="log in recentLogs" 
            :key="log.id"
            :to="'/coffee/' + log.id" 
            class="block editorial-border p-3.5 bg-coffee-cream/20 flex gap-4 items-center rounded-sm hover:bg-coffee-cream/35 transition-colors"
          >
            <img :src="log.image_url" class="w-16 h-16 object-cover rounded-sm border border-coffee-cream flex-shrink-0">
            <div class="flex-1 min-w-0 space-y-1">
              <div class="flex justify-between items-center">
                <span class="text-[9px] uppercase tracking-widest font-semibold text-coffee-softGray">{{ log.coffee_type }}</span>
                <span class="text-[8px] font-mono text-coffee-softGray">{{ formatMonthDay(log.drink_date) }}</span>
              </div>
              <h4 class="font-serif text-base font-light text-coffee-espresso truncate italic leading-tight">{{ log.coffee_name }}</h4>
              <p class="text-[10px] text-coffee-brown font-light truncate leading-relaxed">{{ log.notes }}</p>
              <div class="flex gap-2 items-center text-[9px] text-coffee-softGray pt-0.5">
                <span class="px-1.5 py-0.5 bg-coffee-cream rounded-sm text-coffee-espresso">😌 {{ log.mood }}</span>
                <span class="truncate max-w-[120px]">at {{ log.shop_name.split(',')[0] }}</span>
              </div>
            </div>
          </router-link>
        </div>
      </div>

    </div>

    <!-- Sticky Bottom Navigation Bar -->
    <div class="h-16 border-t border-coffee-cream bg-coffee-warmWhite flex items-center justify-around px-6 z-30 select-none">
      <router-link to="/home" class="flex flex-col items-center gap-1 text-coffee-brown">
        <BookOpen class="w-5 h-5" />
        <span class="text-[10.5px] tracking-widest font-medium">咖啡日志</span>
      </router-link>
      <router-link to="/timeline" class="flex flex-col items-center gap-1 text-coffee-softGray hover:text-coffee-brown transition-colors">
        <Calendar class="w-5 h-5" />
        <span class="text-[10.5px] tracking-widest font-medium">时间线</span>
      </router-link>
      <router-link to="/stats" class="flex flex-col items-center gap-1 text-coffee-softGray hover:text-coffee-brown transition-colors">
        <BarChart3 class="w-5 h-5" />
        <span class="text-[10.5px] tracking-widest font-medium">咖迹</span>
      </router-link>
      <!-- Circular Quick add button -->
      <router-link to="/create" class="w-10 h-10 bg-coffee-espresso text-coffee-warmWhite rounded-full flex items-center justify-center -mt-8 shadow-md border-[4px] border-coffee-warmWhite hover:bg-coffee-brown transition-colors">
        <Plus class="w-5 h-5" />
      </router-link>
    </div>

  </div>
</template>

<script setup lang="ts">
import { computed, onMounted } from 'vue'
import { useCoffeeLogStore } from '@/stores/coffeeLog'
import { useAuthStore } from '@/stores/auth'
import { BookOpen, Calendar, BarChart3, Plus, Search } from 'lucide-vue-next'

const store = useCoffeeLogStore()
const authStore = useAuthStore()

// Fetch data on mount
onMounted(async () => {
  if (store.logs.length === 0) {
    await Promise.all([
      store.fetchLogs({ page: 1, page_size: 10 }),
      store.fetchStats()
    ])
  }
})

// Dynamic Dates
const todayDate = computed(() => {
  const date = new Date()
  const months = ["January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"]
  return `${months[date.getMonth()]} ${date.getDate()}, ${date.getFullYear()}`
})

// Logs Logic
const lastLog = computed(() => store.logs[0] || null)
const recentLogs = computed(() => store.logs.slice(0, 2))

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
