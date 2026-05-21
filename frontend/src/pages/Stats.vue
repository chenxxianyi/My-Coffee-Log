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
        <div class="p-5 bg-coffee-cream rounded-sm double-border text-center space-y-1">
          <span class="text-[9px] uppercase tracking-wider text-coffee-softGray">总饮用手账</span>
          <div class="text-4xl font-serif text-coffee-espresso font-light">{{ store.totalBrews }}</div>
          <span class="text-[9px] text-coffee-brown italic">Quiet mornings</span>
        </div>
        <div class="p-5 bg-coffee-cream rounded-sm double-border text-center space-y-1">
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
        
        <!-- Brew styles proportion bar -->
        <div class="space-y-3">
          <div class="text-[10px] uppercase tracking-widest font-semibold text-coffee-softGray mb-1 select-none">最爱冲煮风格</div>
          <div class="space-y-2.5">
            <div>
              <div class="flex justify-between text-xs mb-1">
                <span>Pour Over / 手冲</span>
                <span class="font-semibold text-coffee-espresso">12 brews (66%)</span>
              </div>
              <div class="w-full h-1.5 bg-coffee-cream rounded-full overflow-hidden select-none">
                <div class="h-full bg-coffee-brown rounded-full" style="width: 66%;"></div>
              </div>
            </div>
            
            <div>
              <div class="flex justify-between text-xs mb-1">
                <span>Latte / 拿铁</span>
                <span class="font-semibold text-coffee-espresso">4 brews (22%)</span>
              </div>
              <div class="w-full h-1.5 bg-coffee-cream rounded-full overflow-hidden select-none">
                <div class="h-full bg-coffee-brown/60 rounded-full" style="width: 22%;"></div>
              </div>
            </div>

            <div>
              <div class="flex justify-between text-xs mb-1">
                <span>Americano / 美式</span>
                <span class="font-semibold text-coffee-espresso">2 brews (12%)</span>
              </div>
              <div class="w-full h-1.5 bg-coffee-cream rounded-full overflow-hidden select-none">
                <div class="h-full bg-coffee-brown/30 rounded-full" style="width: 12%;"></div>
              </div>
            </div>
          </div>
        </div>

        <!-- Tag clouds proportion -->
        <div class="space-y-2 select-none">
          <div class="text-[10px] uppercase tracking-widest font-semibold text-coffee-softGray mb-1">高频出现的风味特征</div>
          <div class="flex flex-wrap gap-1.5 text-xs text-coffee-espresso font-medium">
            <span class="px-3 py-1 bg-coffee-cream/60 border border-coffee-latte/50 rounded-full">柑橘 / citrus (14杯)</span>
            <span class="px-3 py-1 bg-coffee-cream/60 border border-coffee-latte/50 rounded-full">花香 / floral (11杯)</span>
            <span class="px-3 py-1 bg-coffee-cream/60 border border-coffee-latte/50 rounded-full font-light text-coffee-softGray">坚果 / nutty (9杯)</span>
            <span class="px-3 py-1 bg-coffee-cream/60 border border-coffee-latte/50 rounded-full font-light text-coffee-softGray">焦糖 / caramel (6杯)</span>
          </div>
        </div>
      </div>

    </div>

    <!-- Floating Navigation Bar -->
    <div class="h-16 border-t border-coffee-cream bg-coffee-warmWhite flex items-center justify-around px-6 z-30 select-none">
      <router-link to="/home" class="flex flex-col items-center gap-1 text-coffee-softGray hover:text-coffee-brown transition-colors">
        <BookOpen class="w-5 h-5" />
        <span class="text-[10.5px] tracking-widest font-medium">咖啡日志</span>
      </router-link>
      <router-link to="/timeline" class="flex flex-col items-center gap-1 text-coffee-softGray hover:text-coffee-brown transition-colors">
        <Calendar class="w-5 h-5" />
        <span class="text-[10.5px] tracking-widest font-medium">时间线</span>
      </router-link>
      <router-link to="/stats" class="flex flex-col items-center gap-1 text-coffee-brown">
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
import { onMounted } from 'vue'
import { useCoffeeLogStore } from '@/stores/coffeeLog'
import FlavorRadarChart from '@/components/charts/FlavorRadarChart.vue'
import { BookOpen, Calendar, BarChart3, Plus } from 'lucide-vue-next'

const store = useCoffeeLogStore()

onMounted(async () => {
  await store.fetchStats()
})
</script>

<style scoped>
/* Stats page specific scoped styles */
</style>
