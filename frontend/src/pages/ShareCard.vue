<template>
  <div v-if="log" class="flex-1 w-full flex flex-col justify-between bg-black/65 backdrop-blur-sm p-6 relative">
    
    <!-- Top Bar Close button -->
    <div class="flex justify-between items-center text-white mb-2 select-none z-10 relative">
      <router-link 
        :to="'/coffee/' + log.id" 
        class="w-9 h-9 rounded-full bg-white/20 flex items-center justify-center hover:bg-white/35 transition-colors"
      >
        <X class="w-4 h-4" />
      </router-link>
      <span class="text-[10px] uppercase tracking-[0.25em] font-bold">分享海报生成</span>
      <div class="w-9 h-9"></div> <!-- spacer -->
    </div>

    <!-- Aspect Ratio Switches Row -->
    <div class="flex justify-center gap-2 bg-neutral-900/60 p-1.5 rounded-full border border-neutral-800 backdrop-blur-md w-full max-w-[290px] mx-auto text-[9px] text-neutral-400 font-bold uppercase tracking-wider mb-2 select-none z-10 relative">
      <button 
        @click="ratio = '1:1'"
        class="flex-1 py-1.5 rounded-full transition-all"
        :class="ratio === '1:1' ? 'bg-white/10 text-white' : 'hover:text-white'"
      >
        1:1 方形
      </button>
      <button 
        @click="ratio = '3:4'"
        class="flex-1 py-1.5 rounded-full transition-all"
        :class="ratio === '3:4' ? 'bg-white/10 text-white' : 'hover:text-white'"
      >
        3:4 封面
      </button>
      <button 
        @click="ratio = '9:16'"
        class="flex-1 py-1.5 rounded-full transition-all"
        :class="ratio === '9:16' ? 'bg-white/10 text-white' : 'hover:text-white'"
      >
        9:16 朋友圈
      </button>
    </div>

    <!-- Centered Card Viewport with elegant responsive scaling -->
    <div class="flex-1 flex items-center justify-center py-4 z-10">
      
      <!-- The Card DOM element that html2canvas will target -->
      <div 
        id="card-element"
        class="card-paper p-5 double-border text-coffee-charcoal transition-all duration-300 ease-out flex flex-col justify-between overflow-hidden"
        :style="cardDimensions"
      >
        <!-- Card Watermark Header -->
        <div class="flex justify-between items-center text-[8px] tracking-[0.3em] font-bold text-coffee-softGray pb-2 border-b border-coffee-cream select-none">
          <span>MY COFFEE LOG</span>
          <span>CHRONICLE OF FLAVOR</span>
        </div>

        <!-- Content Stack -->
        <div class="flex-1 flex flex-col justify-center space-y-3.5 my-3">
          <!-- Portrait photo inside card -->
          <div 
            class="overflow-hidden rounded-xl relative border border-coffee-cream/40 bg-neutral-100 transition-all duration-300 flex-shrink-0 select-none"
            :class="ratio === '1:1' ? 'h-16' : (ratio === '3:4' ? 'h-28' : 'h-40')"
          >
            <img :src="log.image_url" class="w-full h-full object-cover">
            <span class="absolute bottom-2.5 left-2.5 bg-coffee-espresso/85 backdrop-blur-sm text-coffee-warmWhite px-1.5 py-0.5 text-[7px] tracking-[0.2em] uppercase font-serif font-bold">
              {{ log.coffee_type }}
            </span>
          </div>

          <!-- Typography Details -->
          <div class="text-center select-none">
            <h4 class="font-serif text-lg italic font-light text-coffee-espresso leading-none break-words">{{ log.coffee_name }}</h4>
            <p class="text-[8px] uppercase tracking-wider text-coffee-brown mt-0.5 font-semibold truncate">{{ log.shop_name.split(',')[0] }}</p>
          </div>

          <!-- AI Sensory summary shortened nicely -->
          <p 
            class="font-serif italic text-coffee-espresso leading-relaxed text-center px-1"
            :class="ratio === '1:1' ? 'text-[9.5px] line-clamp-2' : 'text-[11px] line-clamp-3'"
          >
            "{{ log.ai_summary.slice(0, ratio === '1:1' ? 45 : 85) }}..."
          </p>

          <!-- SVG Radar component with scale response -->
          <div class="flex justify-center select-none">
            <FlavorRadarChart 
              :values="[log.acidity, log.bitterness, log.sweetness, log.body, log.aroma, log.aftertaste]"
              :size="ratio === '1:1' ? 48 : 64"
              :dimensions="['Acid', 'Bit', 'Sweet', 'Body', 'Aroma', 'After']"
              :label-font-size="5"
              :dot-radius="1.5"
            />
          </div>
        </div>

        <!-- Footer watermark stamp -->
        <div class="flex justify-between items-center pt-2 border-t border-coffee-cream text-[7.5px] text-coffee-softGray font-bold tracking-widest uppercase select-none">
          <span>{{ formatFullDate(log.drink_date) }}</span>
          <span>MOOD: 😌 {{ log.mood }}</span>
        </div>

      </div>

    </div>

    <!-- Download CTA button -->
    <div class="space-y-4">
      <button 
        @click="exportCardImage" 
        class="w-full py-4 bg-coffee-cream text-coffee-espresso hover:bg-coffee-latte transition-all duration-300 rounded-xl text-xs font-semibold tracking-[0.25em] uppercase shadow-md flex items-center justify-center gap-2"
        :disabled="isExporting"
      >
        <template v-if="isExporting">
          <div class="w-4 h-4 border-2 border-coffee-espresso border-t-transparent rounded-full animate-spin"></div>
          <span>生成高清PNG海报中...</span>
        </template>
        <template v-else>
          <Download class="w-4 h-4" />
          <span>下载高清无损海报</span>
        </template>
      </button>
      <p class="text-[9px] text-neutral-400 text-center tracking-wider font-light uppercase select-none">
        直接将底层 HTML DOM 节点超采样渲染导出为高清相册卡片
      </p>
    </div>

  </div>

  <div v-else class="flex-1 flex flex-col justify-center items-center p-8 text-center text-coffee-softGray space-y-4">
    <Coffee class="w-12 h-12 stroke-[1px]" />
    <p class="font-serif text-sm uppercase tracking-widest">Entry Not Found</p>
    <router-link to="/home" class="px-4 py-2 bg-coffee-espresso text-coffee-warmWhite text-xs tracking-wider rounded-xl font-semibold uppercase">Go Back</router-link>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { useCoffeeLogStore } from '@/stores/coffeeLog'
import FlavorRadarChart from '@/components/charts/FlavorRadarChart.vue'
import { X, Download, Coffee } from 'lucide-vue-next'
import html2canvas from 'html2canvas'

const props = defineProps<{
  id: string
}>()

const store = useCoffeeLogStore()
const log = computed(() => store.getLogById(parseInt(props.id)))

const ratio = ref('3:4')
const isExporting = ref(false)

// Dimensions calculations for the tactile card
const cardDimensions = computed(() => {
  if (ratio.value === '1:1') {
    return { width: '290px', height: '290px' }
  } else if (ratio.value === '3:4') {
    return { width: '290px', height: '386px' }
  } else {
    return { width: '260px', height: '462px' }
  }
})

// html2canvas export rendering mechanism! (High precision!)
const exportCardImage = async () => {
  const cardElement = document.getElementById('card-element')
  if (!cardElement) return
  
  isExporting.value = true
  
  try {
    // Wait for brief moments to allow any font rendering to settle
    await new Promise(resolve => setTimeout(resolve, 300))
    
    // Call html2canvas with optimal high-res settings
    const canvas = await html2canvas(cardElement, {
      scale: 3, // Upscale 3x to get high-definition retina rendering (for social media posts!)
      useCORS: true, // Handle images cross origin safely
      backgroundColor: '#FFF2DB', // Ensure warm background is baked in
      logging: false
    })
    
    // Convert to PNG URI
    const dataUrl = canvas.toDataURL('image/png')
    
    // Trigger download in local browser environment
    const link = document.createElement('a')
    link.download = `MCL-Share-${props.id}-${ratio.value.replace(':', '_')}.png`
    link.href = dataUrl
    document.body.appendChild(link)
    link.click()
    document.body.removeChild(link)
    
    alert('📷 咖啡卡片导出成功！已将高精度小红书海报卡片保存至相册/下载列表中。')
  } catch (err) {
    console.error('Canvas export failed:', err)
    alert('导出海报失败，可能是部分图片资源跨域。真实环境下我们将通过上传后端中转来妥善解决 CORS 跨域问题。')
  } finally {
    isExporting.value = false
  }
}

// Helpers
const formatFullDate = (dateStr: string) => {
  const date = new Date(dateStr)
  const months = ["JAN", "FEB", "MAR", "APR", "MAY", "JUN", "JUL", "AUG", "SEP", "OCT", "NOV", "DEC"]
  return `${date.getDate()} ${months[date.getMonth()]} ${date.getFullYear()}`
}
</script>

<style scoped>
.card-paper {
  background-color: #FFF2DB;
  box-shadow: 0 25px 50px -12px rgba(44, 26, 14, 0.45);
}
</style>
