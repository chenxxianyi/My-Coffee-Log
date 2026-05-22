<template>
  <div v-if="log" class="flex-1 w-full flex flex-col justify-between bg-black/65 backdrop-blur-sm p-6 relative">
    
    <!-- Top Bar -->
    <div class="flex justify-between items-center text-white mb-2 select-none z-10 relative">
      <router-link 
        :to="'/coffee/' + log.id" 
        class="w-9 h-9 rounded-full bg-white/20 flex items-center justify-center hover:bg-white/35 transition-colors"
      >
        <X class="w-4 h-4" />
      </router-link>
      <span class="text-[10px] uppercase tracking-[0.25em] font-bold">分享海报生成</span>
      <div class="w-9 h-9"></div>
    </div>

    <!-- Template Switcher -->
    <div class="flex justify-center gap-1.5 bg-neutral-900/60 p-1 rounded-full border border-neutral-800 backdrop-blur-md w-full max-w-[320px] mx-auto text-[8px] text-neutral-400 font-bold uppercase tracking-wider mb-1.5 select-none z-10 relative">
      <button 
        v-for="t in templates" :key="t.id"
        @click="cardTemplate = t.id"
        class="flex-1 py-1.5 rounded-full transition-all"
        :class="cardTemplate === t.id ? 'bg-white/10 text-white' : 'hover:text-white'"
      >
        {{ t.label }}
      </button>
    </div>

    <!-- Aspect Ratio Switcher -->
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

    <!-- Card Viewport -->
    <div class="flex-1 flex items-center justify-center py-3 z-10">
      
      <!-- ====== MINIMAL TEMPLATE ====== -->
      <div 
        v-if="cardTemplate === 'minimal'"
        id="card-element"
        class="card-minimal transition-all duration-300 ease-out flex flex-col justify-between overflow-hidden"
        :style="cardDimensions"
      >
        <!-- Brand watermark -->
        <div class="flex justify-between items-center text-[7px] tracking-[0.3em] font-bold text-coffee-softGray/60 pb-1.5 border-b border-coffee-cream/50 select-none">
          <span>MY COFFEE LOG</span>
          <span>MINIMAL</span>
        </div>

        <div class="flex-1 flex flex-col justify-center space-y-3 my-2.5">
          <!-- Photo -->
          <div 
            class="overflow-hidden rounded-lg relative bg-neutral-100 flex-shrink-0 select-none"
            :class="ratio === '1:1' ? 'h-14' : (ratio === '3:4' ? 'h-24' : 'h-36')"
          >
            <img :src="log.image_url" class="w-full h-full object-cover">
          </div>
          <!-- Name + Shop -->
          <div class="text-center select-none">
            <h4 class="font-serif text-base italic font-light text-coffee-espresso leading-tight break-words">{{ log.coffee_name }}</h4>
            <p class="text-[7px] uppercase tracking-wider text-coffee-brown/70 mt-0.5 font-medium truncate">{{ log.shop_name.split(',')[0] }}</p>
          </div>
          <!-- AI summary -->
          <p class="font-serif italic text-coffee-espresso/80 leading-relaxed text-center px-2"
            :class="ratio === '1:1' ? 'text-[8.5px] line-clamp-2' : 'text-[10px] line-clamp-3'"
          >
            "{{ log.ai_summary.slice(0, ratio === '1:1' ? 40 : 80) }}..."
          </p>
        </div>

        <!-- Footer -->
        <div class="flex justify-between items-center pt-1.5 border-t border-coffee-cream/50 text-[6.5px] text-coffee-softGray/70 font-bold tracking-widest uppercase select-none">
          <span>{{ formatFullDate(log.drink_date) }}</span>
          <span>{{ log.coffee_type }}</span>
        </div>
      </div>

      <!-- ====== MAGAZINE TEMPLATE ====== -->
      <div 
        v-if="cardTemplate === 'magazine'"
        id="card-element"
        class="card-magazine double-border transition-all duration-300 ease-out flex flex-col justify-between overflow-hidden"
        :style="cardDimensions"
      >
        <!-- Magazine Header -->
        <div class="flex justify-between items-center text-[7px] tracking-[0.3em] font-bold text-coffee-softGray pb-1.5 border-b border-coffee-cream select-none">
          <span>MY COFFEE LOG</span>
          <span>CHRONICLE OF FLAVOR</span>
        </div>

        <div class="flex-1 flex flex-col justify-center space-y-3 my-2.5">
          <!-- Photo with type badge -->
          <div 
            class="overflow-hidden rounded-xl relative border border-coffee-cream/40 bg-neutral-100 flex-shrink-0 select-none"
            :class="ratio === '1:1' ? 'h-16' : (ratio === '3:4' ? 'h-28' : 'h-40')"
          >
            <img :src="log.image_url" class="w-full h-full object-cover">
            <span class="absolute bottom-2 left-2.5 bg-coffee-espresso/85 backdrop-blur-sm text-coffee-warmWhite px-1.5 py-0.5 text-[6.5px] tracking-[0.2em] uppercase font-serif font-bold">
              {{ log.coffee_type }}
            </span>
          </div>
          <!-- Name + Shop -->
          <div class="text-center select-none">
            <h4 class="font-serif text-lg italic font-light text-coffee-espresso leading-none break-words">{{ log.coffee_name }}</h4>
            <p class="text-[7.5px] uppercase tracking-wider text-coffee-brown mt-0.5 font-semibold truncate">{{ log.shop_name.split(',')[0] }}</p>
          </div>
          <!-- AI summary -->
          <p class="font-serif italic text-coffee-espresso leading-relaxed text-center px-1"
            :class="ratio === '1:1' ? 'text-[9px] line-clamp-2' : 'text-[10.5px] line-clamp-3'"
          >
            "{{ log.ai_summary.slice(0, ratio === '1:1' ? 45 : 85) }}..."
          </p>
          <!-- Radar -->
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

        <!-- Footer -->
        <div class="flex justify-between items-center pt-1.5 border-t border-coffee-cream text-[7px] text-coffee-softGray font-bold tracking-widest uppercase select-none">
          <span>{{ formatFullDate(log.drink_date) }}</span>
          <span>MOOD: {{ moodEmoji(log.mood) }} {{ log.mood }}</span>
        </div>
      </div>

      <!-- ====== CAFE RECEIPT TEMPLATE ====== -->
      <div 
        v-if="cardTemplate === 'receipt'"
        id="card-element"
        class="card-receipt transition-all duration-300 ease-out flex flex-col justify-between overflow-hidden"
        :style="cardDimensions"
      >
        <!-- Receipt Header -->
        <div class="text-center pb-1.5 border-b-2 border-dashed border-coffee-espresso/25 select-none">
          <div class="text-[9px] tracking-[0.35em] font-bold text-coffee-espresso">MY COFFEE LOG</div>
          <div class="text-[6px] tracking-[0.2em] text-coffee-softGray mt-0.5">— CAFE RECEIPT —</div>
        </div>

        <div class="flex-1 flex flex-col justify-center space-y-2 my-2">
          <!-- Photo strip -->
          <div 
            class="overflow-hidden bg-neutral-100 flex-shrink-0 select-none"
            :class="ratio === '1:1' ? 'h-12' : (ratio === '3:4' ? 'h-20' : 'h-32')"
          >
            <img :src="log.image_url" class="w-full h-full object-cover">
          </div>
          <!-- Receipt line items -->
          <div class="space-y-1 text-[8px] text-coffee-espresso font-mono select-none">
            <div class="flex justify-between border-b border-dotted border-coffee-cream/60 pb-0.5">
              <span>COFFEE</span>
              <span class="font-serif italic text-[10px]">{{ log.coffee_name }}</span>
            </div>
            <div class="flex justify-between border-b border-dotted border-coffee-cream/60 pb-0.5">
              <span>TYPE</span>
              <span>{{ log.coffee_type }}</span>
            </div>
            <div class="flex justify-between border-b border-dotted border-coffee-cream/60 pb-0.5">
              <span>SHOP</span>
              <span class="truncate ml-2">{{ log.shop_name.split(',')[0] }}</span>
            </div>
            <div class="flex justify-between border-b border-dotted border-coffee-cream/60 pb-0.5">
              <span>MOOD</span>
              <span>{{ moodEmoji(log.mood) }} {{ log.mood }}</span>
            </div>
          </div>
          <!-- Flavor scores as receipt items -->
          <div class="grid grid-cols-3 gap-x-2 gap-y-0.5 text-[7px] font-mono text-coffee-brown select-none">
            <div class="flex justify-between"><span>Acd</span><span>{{ log.acidity }}/5</span></div>
            <div class="flex justify-between"><span>Bit</span><span>{{ log.bitterness }}/5</span></div>
            <div class="flex justify-between"><span>Swt</span><span>{{ log.sweetness }}/5</span></div>
            <div class="flex justify-between"><span>Bod</span><span>{{ log.body }}/5</span></div>
            <div class="flex justify-between"><span>Arm</span><span>{{ log.aroma }}/5</span></div>
            <div class="flex justify-between"><span>Aft</span><span>{{ log.aftertaste }}/5</span></div>
          </div>
          <!-- AI summary -->
          <p class="font-serif italic text-[9px] text-coffee-espresso/80 leading-relaxed text-center px-1 line-clamp-2">
            "{{ log.ai_summary.slice(0, 60) }}..."
          </p>
        </div>

        <!-- Receipt Footer -->
        <div class="pt-1.5 border-t-2 border-dashed border-coffee-espresso/25 select-none">
          <div class="flex justify-between items-center text-[6.5px] text-coffee-softGray font-bold tracking-widest uppercase">
            <span>{{ formatFullDate(log.drink_date) }}</span>
            <span>MY COFFEE LOG</span>
          </div>
          <!-- Barcode-style decoration -->
          <div class="flex justify-center gap-px mt-1">
            <div v-for="i in 30" :key="i" class="bg-coffee-espresso/20" :style="{ width: (i % 3 === 0 ? '2px' : '1px'), height: '8px' }"></div>
          </div>
        </div>
      </div>

    </div>

    <!-- Download CTA -->
    <div class="space-y-3">
      <!-- AI Share Copy -->
      <button @click="generateAIShareCopy" :disabled="isGeneratingCopy" class="w-full py-3.5 text-[10px] uppercase tracking-widest font-semibold bg-white/10 text-white hover:bg-white/20 transition-all rounded-sm flex items-center justify-center gap-2 border border-white/20">
        <template v-if="isGeneratingCopy">
          <span class="animate-pulse">AI 文案生成中...</span>
        </template>
        <template v-else>
          <Sparkles class="w-4 h-4" />
          <span>AI 生成分享文案</span>
        </template>
      </button>
      <div v-if="aiCopy" class="p-3 bg-white/5 border border-white/10 rounded-sm space-y-2">
        <p class="text-[11px] text-white/80 font-serif italic leading-relaxed">{{ aiCopy }}</p>
        <button @click="copyAIText" class="text-[9px] text-white/50 hover:text-white/80 transition-colors uppercase tracking-wider font-semibold">复制文案</button>
      </div>
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
      <p class="text-[8px] text-neutral-400 text-center tracking-wider font-light uppercase select-none">
        {{ cardTemplate === 'minimal' ? '极简风格' : cardTemplate === 'magazine' ? '杂志风格' : '小票风格' }} · {{ ratio }} · 超采样3x渲染
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
import { X, Download, Coffee, Sparkles } from 'lucide-vue-next'
import html2canvas from 'html2canvas'
import * as statsApi from '@/api/stats'

const props = defineProps<{
  id: string
}>()

const store = useCoffeeLogStore()
const log = computed(() => store.getLogById(parseInt(props.id)))

const ratio = ref('3:4')
const cardTemplate = ref('magazine')
const isExporting = ref(false)
const aiCopy = ref('')
const isGeneratingCopy = ref(false)

const templates = [
  { id: 'minimal', label: '极简' },
  { id: 'magazine', label: '杂志' },
  { id: 'receipt', label: '小票' },
]

const cardDimensions = computed(() => {
  if (ratio.value === '1:1') {
    return { width: '290px', height: '290px' }
  } else if (ratio.value === '3:4') {
    return { width: '290px', height: '386px' }
  } else {
    return { width: '260px', height: '462px' }
  }
})

const moodEmoji = (mood: string) => {
  const map: Record<string, string> = { Calm: '😌', Energetic: '⚡', Reflective: '💭', Tired: '🥱' }
  return map[mood] ?? '😌'
}

const exportCardImage = async () => {
  const cardElement = document.getElementById('card-element')
  if (!cardElement) return
  
  isExporting.value = true
  
  try {
    await new Promise(resolve => setTimeout(resolve, 300))
    
    const bgColor = cardTemplate.value === 'receipt' ? '#FFFEF7' : '#FFF2DB'
    
    const canvas = await html2canvas(cardElement, {
      scale: 3,
      useCORS: true,
      backgroundColor: bgColor,
      logging: false
    })
    
    const dataUrl = canvas.toDataURL('image/png')
    
    const link = document.createElement('a')
    link.download = `MCL-${cardTemplate.value}-${props.id}-${ratio.value.replace(':', '_')}.png`
    link.href = dataUrl
    document.body.appendChild(link)
    link.click()
    document.body.removeChild(link)
    
    alert('📷 咖啡卡片导出成功！已将高精度海报保存至相册/下载列表中。')
  } catch (err) {
    console.error('Canvas export failed:', err)
    alert('导出海报失败，可能是部分图片资源跨域。')
  } finally {
    isExporting.value = false
  }
}

const formatFullDate = (dateStr: string) => {
  const date = new Date(dateStr)
  const months = ["JAN", "FEB", "MAR", "APR", "MAY", "JUN", "JUL", "AUG", "SEP", "OCT", "NOV", "DEC"]
  return `${date.getDate()} ${months[date.getMonth()]} ${date.getFullYear()}`
}

async function generateAIShareCopy() {
  if (!log.value || isGeneratingCopy.value) return
  isGeneratingCopy.value = true
  try {
    const res = await statsApi.generateShareCopy({
      coffee_name: log.value.coffee_name,
      coffee_type: log.value.coffee_type,
      shop_name: log.value.shop_name,
      mood: log.value.mood,
      notes: log.value.notes
    })
    aiCopy.value = res.copy || ''
  } catch (e) {
    console.error('Failed to generate AI share copy:', e)
    aiCopy.value = ''
  } finally {
    isGeneratingCopy.value = false
  }
}

function copyAIText() {
  if (!aiCopy.value) return
  navigator.clipboard.writeText(aiCopy.value).then(() => {
    alert('AI 文案已复制到剪贴板！')
  }).catch(() => {})
}
</script>

<style scoped>
.card-minimal {
  background-color: #FFF2DB;
  padding: 16px;
  border-radius: 4px;
  box-shadow: 0 25px 50px -12px rgba(44, 26, 14, 0.35);
}

.card-magazine {
  background-color: #FFF2DB;
  padding: 20px;
  box-shadow: 0 25px 50px -12px rgba(44, 26, 14, 0.45);
}

.card-receipt {
  background-color: #FFFEF7;
  padding: 16px;
  border-radius: 2px;
  box-shadow: 0 25px 50px -12px rgba(44, 26, 14, 0.35);
  border: 1px solid rgba(92, 61, 46, 0.1);
}
</style>
