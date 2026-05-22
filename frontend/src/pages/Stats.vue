<template>
  <div class="flex-1 w-full flex flex-col justify-between bg-coffee-warmWhite text-coffee-charcoal">
    
    <!-- Header -->
    <div class="px-6 py-5 flex justify-between items-end border-b border-coffee-cream bg-coffee-warmWhite sticky top-0 z-10 select-none">
      <div>
        <span class="text-[9px] uppercase tracking-[0.25em] font-semibold text-coffee-softGray">Sensory Metrics</span>
        <h2 class="font-serif text-2xl font-light text-coffee-espresso leading-none mt-1">我的风味足迹</h2>
      </div>
      <img src="https://images.unsplash.com/photo-1534528741775-53994a69daeb?auto=format&fit=crop&q=80&w=120" alt="Avatar" class="w-8 h-8 rounded-full object-cover border border-coffee-latte select-none">
    </div>

    <!-- Scrollable Body Contents -->
    <div class="flex-1 overflow-y-auto px-6 py-5 space-y-6 pb-24 scrollbar-none">
      
      <!-- Big Numbers Grid -->
      <div class="grid grid-cols-2 gap-4 select-none">
        <div class="p-5 bg-coffee-cream rounded-2xl double-border text-center space-y-1">
          <span class="text-[9px] uppercase tracking-wider text-coffee-softGray">总饮用手账</span>
          <div class="text-4xl font-serif text-coffee-espresso font-light">{{ store.totalBrews }}</div>
          <span class="text-[9px] text-coffee-brown italic">Quiet mornings</span>
        </div>
        <div class="p-5 bg-coffee-cream rounded-2xl double-border text-center space-y-1">
          <span class="text-[9px] uppercase tracking-wider text-coffee-softGray">本月饮用杯数</span>
          <div class="text-4xl font-serif text-coffee-espresso font-light">{{ store.monthBrews }}</div>
          <span class="text-[9px] text-coffee-brown italic">May logs</span>
        </div>
      </div>

      <!-- Radar Average Footprint -->
      <div class="space-y-4">
        <div class="text-center select-none">
          <span class="text-[9px] uppercase tracking-[0.2em] font-bold text-coffee-brown block">多维综合风味雷达</span>
          <span class="text-[10px] text-coffee-softGray">基于你所有冲煮日记计算的平均感官足迹</span>
        </div>
        
        <!-- Center-aligned FlavorRadarChart component -->
        <div class="w-[200px] h-[200px] mx-auto flex items-center justify-center p-2 bg-coffee-cream/40 rounded-full border border-coffee-latte/30">
          <FlavorRadarChart 
            :values="store.averageSensoryValues"
            :size="190"
            :dimensions="['酸度', '苦感', '甜感', '醇厚', '香气', '余韵']"
            :dot-radius="3.0"
            :label-font-size="8.5"
          />
        </div>
      </div>

      <!-- Style Breakdown Metrics -->
      <div class="space-y-4">
        <h3 class="font-serif text-lg font-light text-coffee-espresso border-b border-coffee-cream pb-1 select-none">多维度冲煮偏好</h3>
        
        <!-- Brew styles proportion bar (real data) -->
        <div class="space-y-3">
          <div class="text-[10px] uppercase tracking-widest font-semibold text-coffee-softGray mb-1 select-none">最爱冲煮风格</div>
          <div v-if="brewTypeStats.length > 0" class="space-y-2.5">
            <div v-for="(item, idx) in brewTypeStats" :key="item.type">
              <div class="flex justify-between text-xs mb-1">
                <span>{{ coffeeTypeLabel(item.type) }}</span>
                <span class="font-semibold text-coffee-espresso">{{ item.count }} brews ({{ item.pct }}%)</span>
              </div>
              <div class="w-full h-1.5 bg-coffee-cream rounded-full overflow-hidden select-none">
                <div class="h-full rounded-full transition-all duration-700"
                     :class="barColors[idx] || barColors[barColors.length - 1]"
                     :style="`width: ${item.pct}%;`"></div>
              </div>
            </div>
          </div>
          <p v-else class="text-[10px] text-coffee-softGray italic">暂无冲煮记录</p>
        </div>

        <!-- Tag clouds proportion (real data) -->
        <div class="space-y-2 select-none">
          <div class="text-[10px] uppercase tracking-widest font-semibold text-coffee-softGray mb-1">高频出现的风味特征</div>
          <div v-if="flavorTagStats.length > 0" class="flex flex-wrap gap-1.5 text-xs font-medium">
            <span
              v-for="(tag, idx) in flavorTagStats" :key="tag.name"
              class="px-3 py-1 rounded-full border border-coffee-latte/50"
              :class="idx < 2 ? 'bg-coffee-latte/15 text-coffee-espresso' : 'bg-coffee-cream/60 text-coffee-softGray font-light'"
            >{{ tag.label }} / {{ tag.name }} ({{ tag.count }}杯)</span>
          </div>
          <p v-else class="text-[10px] text-coffee-softGray italic">暂无风味标签记录</p>
        </div>
      </div>

      <!-- ---- Monthly Review Entry ---- -->
      <router-link to="/monthly-review" class="block p-5 rounded-sm border border-coffee-latte/40 hover:border-coffee-brown/40 transition-colors select-none" style="background: linear-gradient(135deg, rgba(215,196,168,0.15) 0%, rgba(231,111,81,0.04) 100%);">
        <div class="flex items-center justify-between">
          <div class="flex items-center gap-3">
            <div class="text-2xl leading-none">📅</div>
            <div>
              <h4 class="font-serif text-sm font-semibold text-coffee-espresso leading-tight">月度咖啡回顾</h4>
              <span class="text-[10px] text-coffee-brown font-serif italic">Monthly Review / 回顾你的咖啡月</span>
            </div>
          </div>
          <ChevronRight class="w-4 h-4 text-coffee-softGray" />
        </div>
      </router-link>

      <!-- ---- Coffee Personality Module ---- -->
      <div v-if="store.personalities.length > 0" class="space-y-4">
        <div class="flex items-center gap-2.5 select-none">
          <div class="h-px flex-1 max-w-[24px]" style="background: linear-gradient(to right, rgba(92,61,46,0.35), transparent);"></div>
          <h3 class="text-[10px] uppercase tracking-[0.2em] font-bold text-coffee-espresso">你的咖啡人格 / Coffee Personality</h3>
          <div class="h-px flex-1" style="background: linear-gradient(to left, rgba(92,61,46,0.35), transparent);"></div>
        </div>

        <!-- Primary Personality Card -->
        <div class="p-5 rounded-sm border border-coffee-latte/40" style="background: linear-gradient(135deg, rgba(215,196,168,0.2) 0%, rgba(231,111,81,0.06) 100%);">
          <div class="flex items-start gap-4">
            <div class="text-4xl leading-none select-none">{{ store.personalities[0].icon }}</div>
            <div class="flex-1 space-y-1.5">
              <div>
                <h4 class="font-serif text-lg font-semibold text-coffee-espresso leading-tight">{{ store.personalities[0].title }}</h4>
                <span class="text-[10px] text-coffee-brown font-serif italic">{{ store.personalities[0].subtitle }}</span>
              </div>
              <p class="text-xs text-coffee-brown leading-relaxed font-light">{{ store.personalities[0].description }}</p>
            </div>
          </div>
        </div>

        <!-- Secondary Personality Tags -->
        <div v-if="store.personalities.length > 1" class="flex flex-wrap gap-2">
          <div v-for="p in store.personalities.slice(1)" :key="p.slug" class="flex items-center gap-2 px-3 py-2 bg-coffee-cream/50 border border-coffee-latte/30 rounded-sm">
            <span class="text-lg leading-none select-none">{{ p.icon }}</span>
            <div>
              <div class="text-[11px] font-semibold text-coffee-espresso leading-none">{{ p.title }}</div>
              <div class="text-[9px] text-coffee-softGray font-serif italic">{{ p.subtitle }}</div>
            </div>
          </div>
        </div>

        <!-- Share Personality CTA -->
        <button
          @click="sharePersonality"
          class="w-full py-3 text-[10px] uppercase tracking-widest font-semibold bg-coffee-espresso text-coffee-warmWhite hover:bg-coffee-brown transition-all rounded-sm flex items-center justify-center gap-2"
        >
          <Share2 class="w-3.5 h-3.5" />
          <span>分享我的咖啡人格</span>
        </button>
      </div>

      <!-- ---- AI Coffee Profile ---- -->
      <div v-if="store.coffeeProfileAI" class="space-y-3">
        <div class="flex items-center gap-2.5 select-none">
          <div class="h-px flex-1 max-w-[24px]" style="background: linear-gradient(to right, rgba(92,61,46,0.35), transparent);"></div>
          <h3 class="text-[10px] uppercase tracking-[0.2em] font-bold text-coffee-espresso flex items-center gap-1.5"><Sparkles class="w-3 h-3 text-amber-500" />AI 咖啡画像 / Profile</h3>
          <div class="h-px flex-1" style="background: linear-gradient(to left, rgba(92,61,46,0.35), transparent);"></div>
        </div>
        <div class="p-4 bg-coffee-cream/30 border border-coffee-latte/30 rounded-sm">
          <p class="text-[13px] text-coffee-brown leading-[1.75] font-light font-serif italic pl-2 border-l-2 border-coffee-latte/30">{{ store.coffeeProfileAI }}</p>
        </div>
      </div>

      <!-- ---- AI Preference Insight ---- -->
      <div v-if="store.preferenceInsight" class="space-y-3">
        <div class="flex items-center gap-2.5 select-none">
          <div class="h-px flex-1 max-w-[24px]" style="background: linear-gradient(to right, rgba(92,61,46,0.35), transparent);"></div>
          <h3 class="text-[10px] uppercase tracking-[0.2em] font-bold text-coffee-espresso flex items-center gap-1.5"><Sparkles class="w-3 h-3 text-amber-500" />偏好洞察 / Insight</h3>
          <div class="h-px flex-1" style="background: linear-gradient(to left, rgba(92,61,46,0.35), transparent);"></div>
        </div>
        <div class="p-4 bg-coffee-cream/30 border border-coffee-latte/30 rounded-sm">
          <p class="text-[13px] text-coffee-brown leading-[1.75] font-light font-serif italic pl-2 border-l-2 border-coffee-latte/30">{{ store.preferenceInsight }}</p>
        </div>
      </div>

      <!-- ---- AI Status Indicator ---- -->
      <div v-if="store.aiStatus" class="flex items-center justify-between px-4 py-3 bg-coffee-cream/20 border border-coffee-cream rounded-sm select-none">
        <div class="flex items-center gap-2">
          <div class="w-2 h-2 rounded-full" :class="store.aiStatus.enabled ? 'bg-emerald-400' : 'bg-coffee-softGray'"></div>
          <span class="text-[10px] text-coffee-softGray">AI 引擎</span>
        </div>
        <span class="text-[10px] font-mono" :class="store.aiStatus.enabled ? 'text-emerald-600' : 'text-coffee-softGray'">{{ store.aiStatus.enabled ? store.aiStatus.model : 'Mock 模式' }}</span>
      </div>

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
      <router-link to="/stats" class="flex-1 flex flex-col items-center gap-0.5 text-coffee-brown">
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
import { onMounted, computed } from 'vue'
import { useCoffeeLogStore } from '@/stores/coffeeLog'
import { getFlavorTags } from '@/api/flavorTag'
import FlavorRadarChart from '@/components/charts/FlavorRadarChart.vue'
import { coffeeTypeLabel } from '@/constants/coffee'
import { BookOpen, Calendar, BarChart3, Plus, User, Share2, ChevronRight, Sparkles } from 'lucide-vue-next'

const store = useCoffeeLogStore()

const barColors = ['bg-coffee-brown', 'bg-coffee-brown/60', 'bg-coffee-brown/35', 'bg-coffee-brown/20']

const brewTypeStats = computed(() => {
  const total = store.logs.length
  if (total === 0) return []
  const counts: Record<string, number> = {}
  store.logs.forEach(log => {
    counts[log.coffee_type] = (counts[log.coffee_type] || 0) + 1
  })
  return Object.entries(counts)
    .sort((a, b) => b[1] - a[1])
    .slice(0, 4)
    .map(([type, count]) => ({
      type,
      count,
      pct: Math.round((count / total) * 100)
    }))
})

const flavorTagStats = computed(() => {
  const counts: Record<string, number> = {}
  store.logs.forEach(log => {
    ;(log.flavor_tags || []).forEach(tag => {
      counts[tag] = (counts[tag] || 0) + 1
    })
  })
  const labelMap: Record<string, string> = {}
  getFlavorTags().forEach(t => { labelMap[t.name] = t.label })
  return Object.entries(counts)
    .sort((a, b) => b[1] - a[1])
    .slice(0, 6)
    .map(([name, count]) => ({ name, label: labelMap[name] || name, count }))
})

const sharePersonality = () => {
  if (store.personalities.length === 0) return
  const primary = store.personalities[0]
  const text = `${primary.icon} ${primary.title} — ${primary.subtitle}\n${primary.description}\n\n— MY COFFEE LOG`
  if (navigator.share) {
    navigator.share({ title: 'My Coffee Personality', text }).catch(() => {})
  } else {
    navigator.clipboard.writeText(text).then(() => {
      alert('咖啡人格已复制到剪贴板！')
    }).catch(() => {})
  }
}

onMounted(async () => {
  await store.fetchStats()
  if (store.logs.length === 0) {
    await store.fetchLogs({ page: 1, page_size: 100 })
  }
  store.fetchPersonality()
  store.fetchAIStatus()
  store.fetchCoffeeProfile()
  store.fetchPreferenceInsight()
})
</script>

<style scoped>
/* Stats page specific scoped styles */
</style>
