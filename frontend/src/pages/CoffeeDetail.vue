<template>
  <div v-if="log" class="flex-1 w-full flex flex-col justify-between bg-coffee-warmWhite text-coffee-charcoal">
    
    <!-- Header transparent gradient overlay -->
    <div class="absolute top-0 inset-x-0 h-16 flex justify-between items-center px-6 z-20 bg-gradient-to-b from-black/55 to-transparent select-none">
      <button @click="router.push('/home')" class="w-9 h-9 rounded-full bg-white/20 backdrop-blur-md text-white flex items-center justify-center hover:bg-white/35 transition-colors">
        <ArrowLeft class="w-4 h-4" />
      </button>
      <button @click="router.push(`/share/${log.id}`)" class="w-9 h-9 rounded-full bg-white/20 backdrop-blur-md text-white flex items-center justify-center hover:bg-white/35 transition-colors">
        <Share2 class="w-4 h-4" />
      </button>
    </div>

    <!-- Scrollable Detail Contents -->
    <div class="flex-1 overflow-y-auto pb-12 scrollbar-none">
      
      <!-- Cover Photo with Badge -->
      <div class="w-full h-72 relative bg-coffee-espresso">
        <img :src="log.image_url" class="w-full h-full object-cover">
        <div class="absolute inset-0 bg-gradient-to-t from-coffee-warmWhite via-transparent to-transparent"></div>
        <span class="absolute bottom-5 left-6 bg-coffee-brown text-coffee-warmWhite px-3 py-1 rounded-sm text-[10px] font-serif uppercase tracking-[0.25em] leading-none font-bold select-none">
          {{ log.coffee_type }}
        </span>
      </div>

      <!-- Typography Body Contents -->
      <div class="px-6 space-y-6 -mt-2 relative z-10">
        
        <!-- Header details -->
        <div class="space-y-1.5">
          <span class="text-[9px] uppercase tracking-[0.25em] font-semibold text-coffee-softGray select-none">{{ fullDate }}</span>
          <h1 class="font-serif text-[38px] font-light italic text-coffee-espresso leading-none break-words">{{ log.coffee_name }}</h1>
          <div class="flex gap-2 items-center text-xs text-coffee-brown font-medium select-none">
            <MapPin class="w-3.5 h-3.5" />
            <span>{{ log.shop_name }}</span>
          </div>
        </div>

        <div class="h-[1px] bg-coffee-cream"></div>

        <!-- AI Sensory prose Double border -->
        <div class="p-5 bg-coffee-cream rounded-sm double-border space-y-2">
          <div class="flex justify-between items-center select-none">
            <span class="text-[9px] uppercase tracking-[0.25em] font-bold text-coffee-brown">AI 感官评语</span>
            <span class="text-[10px] text-coffee-espresso italic font-semibold">😌 {{ log.mood }}</span>
          </div>
          <p class="font-serif italic text-[13.5px] text-coffee-espresso leading-relaxed">
            {{ log.ai_summary }}
          </p>
        </div>

        <!-- SVG Radar Component Grid -->
        <div class="space-y-3.5">
          <h3 class="text-xs uppercase tracking-[0.2em] font-bold text-coffee-espresso select-none">风味足迹雷达 / Sensory Radar</h3>
          
          <div class="grid grid-cols-5 gap-4 items-center">
            <!-- Live Radar Chart Component -->
            <div class="col-span-2 aspect-square flex items-center justify-center p-1 bg-coffee-cream/40 rounded-full border border-coffee-latte/30">
              <FlavorRadarChart 
                :values="[log.acidity, log.bitterness, log.sweetness, log.body, log.aroma, log.aftertaste]"
                :size="110"
              />
            </div>
            <!-- Score lists -->
            <div class="col-span-3 text-[11px] text-coffee-brown space-y-1.5 font-medium pl-3 border-l border-coffee-cream select-none">
              <div class="flex justify-between"><span>Acidity / 酸度</span><span class="font-mono text-coffee-espresso">{{ log.acidity }}.0 / 5</span></div>
              <div class="flex justify-between"><span>Bitterness / 苦感</span><span class="font-mono text-coffee-espresso">{{ log.bitterness }}.0 / 5</span></div>
              <div class="flex justify-between"><span>Sweetness / 甜感</span><span class="font-mono text-coffee-espresso">{{ log.sweetness }}.0 / 5</span></div>
              <div class="flex justify-between"><span>Body / 醇厚度</span><span class="font-mono text-coffee-espresso">{{ log.body }}.0 / 5</span></div>
              <div class="flex justify-between"><span>Aroma / 香气</span><span class="font-mono text-coffee-espresso">{{ log.aroma }}.0 / 5</span></div>
              <div class="flex justify-between"><span>Aftertaste / 余韵</span><span class="font-mono text-coffee-espresso">{{ log.aftertaste }}.0 / 5</span></div>
            </div>
          </div>
        </div>

        <!-- Flavor Tags -->
        <div class="space-y-2.5">
          <h3 class="text-xs uppercase tracking-[0.2em] font-bold text-coffee-espresso select-none">感官风味标签</h3>
          <div class="flex flex-wrap gap-2 select-none">
            <span 
              v-for="tag in log.flavor_tags" 
              :key="tag" 
              class="px-3 py-1 text-xs bg-coffee-cream text-coffee-espresso border border-coffee-latte/45 rounded-full font-mono uppercase"
            >
              ★ {{ tag }}
            </span>
          </div>
        </div>

        <!-- Taste notes diary -->
        <div class="space-y-2">
          <h3 class="text-xs uppercase tracking-[0.2em] font-bold text-coffee-espresso select-none">感官味觉日记 / Diary</h3>
          <p class="text-xs text-coffee-brown leading-relaxed font-light font-serif">
            "{{ log.notes }}"
          </p>
        </div>

        <!-- Danger Actions Row -->
        <div class="grid grid-cols-2 gap-3 pt-4 select-none">
          <button 
            @click="handleDelete" 
            class="py-3 text-[10px] uppercase tracking-widest font-semibold border border-red-200 text-red-700 hover:bg-red-50 transition-all rounded-sm"
          >
            删除记录
          </button>
          <router-link 
            :to="'/share/' + log.id" 
            class="py-3 text-[10px] uppercase tracking-widest font-semibold bg-coffee-espresso text-coffee-warmWhite hover:bg-coffee-brown transition-all rounded-sm flex items-center justify-center gap-1.5"
          >
            <Share2 class="w-3.5 h-3.5" />
            <span>生成分享海报</span>
          </router-link>
        </div>

      </div>
    </div>

  </div>

  <!-- If not found -->
  <div v-else class="flex-1 flex flex-col justify-center items-center p-8 text-center text-coffee-softGray space-y-4">
    <Coffee class="w-12 h-12 stroke-[1px]" />
    <p class="font-serif text-sm uppercase tracking-widest">Entry Not Found</p>
    <router-link to="/home" class="px-4 py-2 bg-coffee-espresso text-coffee-warmWhite text-xs tracking-wider rounded-sm font-semibold uppercase">Go Back</router-link>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useCoffeeLogStore } from '@/stores/coffeeLog'
import FlavorRadarChart from '@/components/charts/FlavorRadarChart.vue'
import { ArrowLeft, Share2, MapPin, Coffee } from 'lucide-vue-next'

const props = defineProps<{
  id: string
}>()

const router = useRouter()
const store = useCoffeeLogStore()

const log = computed(() => store.getLogById(parseInt(props.id)))

// Fetch from API if not in local cache
onMounted(async () => {
  if (!log.value) {
    try {
      await store.fetchLogById(parseInt(props.id))
    } catch {
      // log not found, will show empty state
    }
  }
})

const fullDate = computed(() => {
  if (!log.value) return ''
  const date = new Date(log.value.drink_date)
  const days = ["Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"]
  const months = ["January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"]
  return `${days[date.getDay()]}, ${months[date.getMonth()]} ${date.getDate()}, ${date.getFullYear()}`
})

const handleDelete = async () => {
  if (confirm('确认删除本条手账日志吗？该操作无法撤销。')) {
    try {
      await store.deleteLog(parseInt(props.id))
      router.push('/home')
    } catch (e: any) {
      alert(e.message || '删除失败')
    }
  }
}
</script>

<style scoped>
/* Detail specific scoped styles */
</style>
