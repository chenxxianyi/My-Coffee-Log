<template>
  <div class="flex-1 w-full flex flex-col justify-between bg-coffee-warmWhite text-coffee-charcoal">
    
    <!-- Header -->
    <div class="px-6 py-4 border-b border-coffee-cream flex justify-between items-center bg-coffee-warmWhite sticky top-0 z-10 select-none">
      <button @click="router.push('/home')" class="text-coffee-brown hover:text-coffee-espresso">
        <X class="w-5 h-5" />
      </button>
      <span class="font-serif text-lg font-light text-coffee-espresso uppercase tracking-wider">记录今日咖啡</span>
      <span class="text-[10px] tracking-wider text-coffee-softGray font-semibold uppercase">步骤 {{ step }}/3</span>
    </div>

    <!-- Progress Indicator Bar -->
    <div class="h-1 bg-coffee-cream w-full flex select-none">
      <div class="h-full bg-coffee-brown transition-all duration-300" :style="{ width: (step * 33.3) + '%' }"></div>
    </div>

    <!-- Form Body -->
    <div class="flex-1 overflow-y-auto px-6 py-5">
      
      <!-- STEP 1: Basic Information -->
      <div v-if="step === 1" class="space-y-6">
        <div class="space-y-2">
          <label class="text-[10px] uppercase tracking-[0.2em] font-semibold text-coffee-espresso block">1. 咖啡/豆子名称 *</label>
          <input 
            type="text" 
            v-model="form.coffee_name" 
            placeholder="例如: 埃塞俄比亚 耶加雪菲" 
            class="w-full p-3 bg-coffee-cream/40 border border-coffee-latte/60 focus:border-coffee-brown focus:outline-none rounded-sm font-serif text-sm transition-colors"
          >
        </div>

        <div class="space-y-2">
          <label class="text-[10px] uppercase tracking-[0.2em] font-semibold text-coffee-espresso block select-none">2. 冲煮咖啡类型 *</label>
          <div class="grid grid-cols-2 gap-2">
            <button 
              v-for="t in typePresets" 
              :key="t.val"
              @click="form.coffee_type = t.val"
              type="button"
              class="p-3 border rounded-sm text-center text-xs font-serif font-light transition-all duration-200"
              :class="form.coffee_type === t.val 
                ? 'bg-coffee-cream border-coffee-brown ring-1 ring-coffee-brown text-coffee-espresso' 
                : 'bg-coffee-cream/30 border-coffee-latte/50 hover:border-coffee-brown text-coffee-espresso'"
            >
              {{ t.label }}
            </button>
          </div>
        </div>

        <div class="space-y-2">
          <label class="text-[10px] uppercase tracking-[0.2em] font-semibold text-coffee-espresso block select-none">3. 手账封面图 (点击预设或上传本地照片)</label>
          <div class="grid grid-cols-4 gap-2">
            <div 
              v-for="(imgUrl, idx) in store.DEFAULT_PHOTOS" 
              :key="idx"
              @click="form.image_url = imgUrl"
              class="aspect-square relative cursor-pointer overflow-hidden rounded-sm border transition-all"
              :class="form.image_url === imgUrl ? 'border-2 border-coffee-brown scale-[1.02]' : 'border-transparent opacity-80 hover:opacity-100'"
            >
              <img :src="imgUrl" class="w-full h-full object-cover">
              <div v-if="form.image_url === imgUrl" class="absolute inset-0 bg-coffee-espresso/20 flex items-center justify-center text-white">
                <Check class="w-4 h-4" />
              </div>
            </div>
          </div>

          <!-- Upload Local Photo Component inside Create Journal -->
          <div 
            @click="triggerFileSelect"
            class="mt-3 p-3 border border-dashed border-coffee-latte/60 hover:border-coffee-brown bg-coffee-cream/15 rounded-sm cursor-pointer transition-colors flex items-center justify-center gap-2 select-none"
          >
            <template v-if="isUploading">
              <div class="w-4 h-4 border-2 border-coffee-espresso border-t-transparent rounded-full animate-spin"></div>
              <span class="text-xs text-coffee-espresso">正在上传相片...</span>
            </template>
            <template v-else-if="isLocalUploaded">
              <div class="w-6 h-6 rounded-full overflow-hidden border border-coffee-espresso">
                <img :src="form.image_url" class="w-full h-full object-cover">
              </div>
              <span class="text-xs text-green-700 font-medium">本地咖啡照片上传并设为封面！</span>
            </template>
            <template v-else>
              <Plus class="w-4 h-4 text-coffee-softGray" />
              <span class="text-xs text-coffee-espresso font-medium">使用手机拍摄/本地相片作为手账封面</span>
            </template>
          </div>
          <!-- Hidden Native File Input -->
          <input 
            type="file" 
            ref="fileInput" 
            accept="image/*" 
            @change="handleFileChange" 
            class="hidden"
          >
        </div>
      </div>

      <!-- STEP 2: Sensory Sliders with live SVG radar rendering -->
      <div v-if="step === 2" class="space-y-6">
        <div class="flex justify-between items-center mb-2 select-none">
          <span class="text-[10px] uppercase tracking-[0.2em] font-semibold text-coffee-espresso">感官风味指纹 / 6维风味参数</span>
          <span class="text-[10px] text-coffee-softGray italic">0 (无感) - 5 (浓郁)</span>
        </div>

        <!-- Horizontal layout: sliders on left, live responsive radar chart on right -->
        <div class="flex gap-4 items-center">
          <div class="flex-1 space-y-4">
            <!-- Render 6 dynamic sliders -->
            <div v-for="s in sliderSpecs" :key="s.key" class="space-y-1">
              <div class="flex justify-between text-xs text-coffee-espresso">
                <span>{{ s.label }}</span>
                <span class="font-semibold font-mono">{{ form[s.key] }}</span>
              </div>
              <input 
                type="range" 
                min="0" 
                max="5" 
                step="1" 
                v-model.number="form[s.key]" 
                class="w-full h-1 bg-coffee-cream rounded-lg appearance-none cursor-pointer accent-coffee-brown"
              >
            </div>
          </div>

          <!-- Dynamic SVG Radar Chart Component (Instant reactive updates!) -->
          <div class="w-[130px] h-[130px] flex-shrink-0 bg-coffee-cream/40 rounded-full flex items-center justify-center p-1 border border-coffee-latte/40 select-none">
            <FlavorRadarChart 
              :values="[form.acidity, form.bitterness, form.sweetness, form.body, form.aroma, form.aftertaste]"
              :size="120"
              :show-labels="false"
              :dot-radius="2.0"
            />
          </div>
        </div>
      </div>

      <!-- STEP 3: Flavor tags, Mood, Spot & Notes -->
      <div v-if="step === 3" class="space-y-5">
        <div class="space-y-2">
          <label class="text-[10px] uppercase tracking-[0.2em] font-semibold text-coffee-espresso block font-medium select-none">1. 风味特征标签 (点击多选)</label>
          <div class="flex flex-wrap gap-1.5">
            <button
              v-for="tag in tagPresets"
              :key="tag.name"
              @click="toggleTag(tag.name)"
              type="button"
              class="px-3 py-1 text-[11px] border rounded-full transition-all duration-150"
              :class="form.flavor_tags.includes(tag.name)
                ? 'bg-coffee-espresso text-coffee-warmWhite border-coffee-espresso'
                : 'bg-coffee-cream/40 text-coffee-espresso border-coffee-latte/50 hover:border-coffee-brown'"
            >
              {{ tag.label }}
            </button>
          </div>
        </div>

        <div class="space-y-2">
          <label class="text-[10px] uppercase tracking-[0.2em] font-semibold text-coffee-espresso block font-medium select-none">2. 此时此地心情 / Mood</label>
          <div class="grid grid-cols-4 gap-2">
            <button
              v-for="m in moodPresets"
              :key="m.val"
              @click="form.mood = m.val"
              type="button"
              class="p-2 border text-center rounded-sm text-xs font-serif transition-all"
              :class="form.mood === m.val
                ? 'border-coffee-brown bg-coffee-cream text-coffee-espresso font-semibold'
                : 'border-coffee-latte/50 text-coffee-softGray hover:border-coffee-brown'"
            >
              {{ m.label }}
            </button>
          </div>
        </div>

        <div class="space-y-2">
          <label class="text-[10px] uppercase tracking-[0.2em] font-semibold text-coffee-espresso block">3. 咖啡出品馆 / Shop & Spot</label>
          <input 
            type="text" 
            v-model="form.shop_name" 
            placeholder="例如: Blue Bottle, 上海" 
            class="w-full p-2.5 bg-coffee-cream/40 border border-coffee-latte/60 focus:border-coffee-brown focus:outline-none rounded-sm text-xs"
          >
        </div>

        <div class="space-y-2">
          <label class="text-[10px] uppercase tracking-[0.2em] font-semibold text-coffee-espresso block">4. 味觉手账备忘 / Notes</label>
          <textarea 
            v-model="form.notes"
            rows="3" 
            placeholder="写下属于这杯咖啡、此时此刻的心情和香气备注..." 
            class="w-full p-2.5 bg-coffee-cream/40 border border-coffee-latte/60 focus:border-coffee-brown focus:outline-none rounded-sm text-xs font-serif"
          ></textarea>
        </div>
      </div>

    </div>

    <!-- Bottom Controls -->
    <div class="p-6 border-t border-coffee-cream flex gap-3 bg-coffee-warmWhite sticky bottom-0 z-10 select-none">
      <button 
        v-if="step > 1" 
        @click="step--" 
        class="flex-1 py-3 text-xs uppercase tracking-wider font-semibold bg-coffee-cream text-coffee-espresso hover:bg-coffee-latte transition-all rounded-sm"
      >
        上一步
      </button>
      <button 
        @click="handleNext" 
        class="flex-1 py-3 text-xs uppercase tracking-wider font-semibold bg-coffee-espresso text-coffee-warmWhite hover:bg-coffee-brown transition-all rounded-sm flex items-center justify-center gap-1.5"
        :disabled="isSubmitting"
      >
        <template v-if="isSubmitting">
          <div class="w-4 h-4 border-2 border-white border-t-transparent rounded-full animate-spin"></div>
          <span>正在智能撰写感官日志...</span>
        </template>
        <template v-else>
          <span>{{ step === 3 ? 'AI 总结并保存手账' : '下一步' }}</span>
        </template>
      </button>
    </div>

  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'
import { useCoffeeLogStore, NewCoffeeLog } from '@/stores/coffeeLog'
import FlavorRadarChart from '@/components/charts/FlavorRadarChart.vue'
import request from '@/api/request'
import { X, Check, Plus } from 'lucide-vue-next'

const router = useRouter()
const store = useCoffeeLogStore()

const step = ref(1)
const isSubmitting = ref(false)

// File upload states
const fileInput = ref<HTMLInputElement | null>(null)
const isUploading = ref(false)
const isLocalUploaded = ref(false)

// State Form Data
const form = reactive<NewCoffeeLog>({
  coffee_name: '',
  coffee_type: 'Pour Over',
  image_url: store.DEFAULT_PHOTOS[1], // default pour over image
  acidity: 4,
  bitterness: 1,
  sweetness: 3,
  body: 2,
  aroma: 5,
  aftertaste: 4,
  flavor_tags: ['citrus', 'floral'],
  mood: 'Calm',
  shop_name: '',
  notes: ''
})

// Presets Specs
const typePresets = [
  { val: 'Pour Over', label: 'Pour Over / 手冲' },
  { val: 'Latte', label: 'Latte / 拿铁' },
  { val: 'Americano', label: 'Americano / 美式' },
  { val: 'Cold Brew', label: 'Cold Brew / 冷萃' },
  { val: 'Espresso', label: 'Espresso / 浓缩' },
  { val: 'Dirty', label: 'Dirty / 脏咖啡' }
]

const sliderSpecs = [
  { key: 'acidity', label: '酸度 Acidity' },
  { key: 'bitterness', label: '苦感 Bitterness' },
  { key: 'sweetness', label: '甜感 Sweetness' },
  { key: 'body', label: '醇厚度 Body' },
  { key: 'aroma', label: '香气 Aroma' },
  { key: 'aftertaste', label: '余韵 Aftertaste' }
] as const

const tagPresets = [
  { name: 'floral', label: '花香' },
  { name: 'citrus', label: '柑橘' },
  { name: 'berry', label: '莓果' },
  { name: 'nutty', label: '坚果' },
  { name: 'chocolate', label: '巧克力' },
  { name: 'caramel', label: '焦糖' },
  { name: 'creamy', label: '奶油' },
  { name: 'winey', label: '酒香' }
]
const moodPresets = [
  { val: 'Calm', label: '😌 Calm' },
  { val: 'Energetic', label: '⚡ Joy' },
  { val: 'Reflective', label: '💭 Mood' },
  { val: 'Tired', label: '🥱 Tired' }
]

const toggleTag = (tag: string) => {
  if (form.flavor_tags.includes(tag)) {
    form.flavor_tags = form.flavor_tags.filter((t: string) => t !== tag)
  } else {
    form.flavor_tags.push(tag)
  }
}

const handleNext = async () => {
  if (step.value === 1) {
    if (!form.coffee_name.trim()) {
      alert('请输入咖啡名称哦')
      return
    }
    step.value = 2
  } 
  else if (step.value === 2) {
    step.value = 3
  } 
  else if (step.value === 3) {
    // Save record with Pinia
    isSubmitting.value = true
    
    // Create actual save parameters
    const savedLog = {
      ...form,
      shop_name: form.shop_name.trim() || 'Local Coffee Spot',
      notes: form.notes.trim() || '一杯温润安静的手账记录。'
    }

    try {
      const created = await store.addLog(savedLog)
      isSubmitting.value = false
      router.push(`/coffee/${created.id}`)
    } catch (e: any) {
      isSubmitting.value = false
      alert(e.message || '保存失败，请稍后重试')
    }
  }
}

// Local File Upload Handlers
const triggerFileSelect = () => {
  if (fileInput.value) {
    fileInput.value.click()
  }
}

const handleFileChange = async (event: Event) => {
  const target = event.target as HTMLInputElement
  const files = target.files
  if (!files || files.length === 0) return

  const file = files[0]
  if (file.size > 5 * 1024 * 1024) {
    alert('照片大小不能超过 5MB 喔！')
    return
  }

  isUploading.value = true
  const formData = new FormData()
  formData.append('file', file)

  try {
    const res = await request.post('/upload', formData, {
      headers: {
        'Content-Type': 'multipart/form-data'
      }
    }) as any

    if (res && res.url) {
      form.image_url = res.url
      isLocalUploaded.value = true
    } else {
      throw new Error('未获取到返回的图片地址')
    }
  } catch (e: any) {
    alert(e.message || '相片上传失败，请重试')
  } finally {
    isUploading.value = false
    target.value = '' // Clear input
  }
}
</script>

<style scoped>
/* Range slider styling */
input[type="range"]::-webkit-slider-thumb {
  border-radius: 50%;
  width: 12px;
  height: 12px;
  background: #7A5638;
}
</style>
