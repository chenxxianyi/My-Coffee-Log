<template>
  <div class="flex-1 w-full flex flex-col justify-between bg-coffee-warmWhite text-coffee-charcoal">
    
    <!-- Header -->
    <div class="px-6 py-4 border-b border-coffee-cream flex justify-between items-center bg-coffee-warmWhite sticky top-0 z-10 select-none">
      <button @click="router.push('/home')" class="text-coffee-brown hover:text-coffee-espresso">
        <X class="w-5 h-5" />
      </button>
      <!-- Mode Toggle -->
      <div class="flex items-center gap-1 bg-coffee-cream/60 rounded-full p-0.5">
        <button 
          @click="logMode = 'quick'" 
          type="button"
          class="px-3 py-1 rounded-full text-[10px] tracking-wider font-semibold transition-all duration-200"
          :class="logMode === 'quick' ? 'bg-coffee-espresso text-coffee-warmWhite shadow-sm' : 'text-coffee-softGray hover:text-coffee-espresso'"
        >快速记录</button>
        <button 
          @click="logMode = 'detailed'" 
          type="button"
          class="px-3 py-1 rounded-full text-[10px] tracking-wider font-semibold transition-all duration-200"
          :class="logMode === 'detailed' ? 'bg-coffee-espresso text-coffee-warmWhite shadow-sm' : 'text-coffee-softGray hover:text-coffee-espresso'"
        >精细记录</button>
      </div>
      <span v-if="logMode === 'detailed'" class="text-[10px] tracking-wider text-coffee-softGray font-semibold uppercase">步骤 {{ step }}/3</span>
      <span v-else class="text-[10px] tracking-wider text-coffee-softGray font-semibold uppercase">快速记录</span>
    </div>

    <!-- Progress Indicator Bar (only in detailed mode) -->
    <div v-if="logMode === 'detailed'" class="h-1 bg-coffee-cream w-full flex select-none">
      <div class="h-full bg-coffee-brown transition-all duration-300" :style="{ width: (step * 33.3) + '%' }"></div>
    </div>

    <!-- Form Body -->
    <div class="flex-1 overflow-y-auto px-6 py-5">

      <!-- ==================== QUICK LOG MODE ==================== -->
      <div v-if="logMode === 'quick'" class="space-y-6">
        <!-- Photo Selection -->
        <div class="space-y-2">
          <label class="text-[10px] uppercase tracking-[0.2em] font-semibold text-coffee-espresso block select-none">封面图</label>
          <div class="grid grid-cols-4 gap-2">
            <div 
              v-for="(imgUrl, idx) in store.DEFAULT_PHOTOS" 
              :key="idx"
              @click="quickForm.image_url = imgUrl"
              class="aspect-square relative cursor-pointer overflow-hidden rounded-sm border transition-all"
              :class="quickForm.image_url === imgUrl ? 'border-2 border-coffee-brown scale-[1.02]' : 'border-transparent opacity-80 hover:opacity-100'"
            >
              <img :src="imgUrl" class="w-full h-full object-cover">
              <div v-if="quickForm.image_url === imgUrl" class="absolute inset-0 bg-coffee-espresso/20 flex items-center justify-center text-white">
                <Check class="w-4 h-4" />
              </div>
            </div>
          </div>

          <!-- Upload Local Photo -->
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
                <img :src="quickForm.image_url" class="w-full h-full object-cover">
              </div>
              <span class="text-xs text-green-700 font-medium">本地咖啡照片上传并设为封面！</span>
            </template>
            <template v-else>
              <Plus class="w-4 h-4 text-coffee-softGray" />
              <span class="text-xs text-coffee-espresso font-medium">使用手机拍摄/本地相片作为手账封面</span>
            </template>
          </div>
          <input 
            type="file" 
            ref="fileInput" 
            accept="image/*" 
            @change="handleFileChange" 
            class="hidden"
          >
        </div>

        <!-- Coffee Type Selection -->
        <div class="space-y-2">
          <label class="text-[10px] uppercase tracking-[0.2em] font-semibold text-coffee-espresso block select-none">咖啡类型</label>
          <div class="grid grid-cols-3 gap-2">
            <button 
              v-for="t in typePresets" 
              :key="t.val"
              @click="quickForm.coffee_type = t.val"
              type="button"
              class="p-2.5 border rounded-sm text-center text-xs font-serif font-light transition-all duration-200"
              :class="quickForm.coffee_type === t.val 
                ? 'bg-coffee-cream border-coffee-brown ring-1 ring-coffee-brown text-coffee-espresso' 
                : 'bg-coffee-cream/30 border-coffee-latte/50 hover:border-coffee-brown text-coffee-espresso'"
            >
              {{ t.label }}
            </button>
          </div>
        </div>

        <!-- Mood Selection -->
        <div class="space-y-2">
          <label class="text-[10px] uppercase tracking-[0.2em] font-semibold text-coffee-espresso block select-none">此时心情</label>
          <div class="grid grid-cols-4 gap-2">
            <button
              v-for="m in moodPresets"
              :key="m.val"
              @click="quickForm.mood = m.val"
              type="button"
              class="p-2.5 border rounded-sm text-xs font-serif transition-all flex flex-col items-center gap-1"
              :class="quickForm.mood === m.val
                ? 'border-coffee-brown bg-coffee-cream text-coffee-espresso font-semibold'
                : 'border-coffee-latte/50 text-coffee-softGray hover:border-coffee-brown'"
            >
              <component :is="m.icon" class="w-3.5 h-3.5 flex-shrink-0" />
              <span>{{ m.label }}</span>
            </button>
          </div>
        </div>

        <!-- Flavor Impression Quick Selector -->
        <div class="space-y-3">
          <div class="flex justify-between items-center">
            <label class="text-[10px] uppercase tracking-[0.2em] font-semibold text-coffee-espresso select-none">风味印象 <span class="text-coffee-softGray font-normal normal-case">(选填)</span></label>
            <button
              v-if="quickForm.flavor_preset"
              @click="quickForm.flavor_preset = ''"
              type="button"
              class="text-[9px] text-coffee-softGray hover:text-coffee-brown tracking-wider transition-colors"
            >清除选择</button>
          </div>

          <div class="grid grid-cols-2 gap-2">
            <button
              v-for="f in flavorPresets"
              :key="f.val"
              @click="quickForm.flavor_preset = quickForm.flavor_preset === f.val ? '' : f.val"
              type="button"
              class="p-3 border rounded-sm text-left transition-all duration-200 flex items-center gap-3"
              :class="quickForm.flavor_preset === f.val
                ? 'border-coffee-brown bg-coffee-cream text-coffee-espresso'
                : 'border-coffee-latte/50 text-coffee-softGray hover:border-coffee-brown'"
            >
              <component :is="f.icon" class="w-4 h-4 flex-shrink-0" />
              <div>
                <div class="text-xs font-semibold font-serif leading-none">{{ f.label }}</div>
                <div class="text-[10px] mt-1 opacity-70">{{ f.desc }}</div>
              </div>
            </button>
          </div>

          <!-- Live Radar Preview -->
          <Transition name="fade">
            <div v-if="quickForm.flavor_preset" class="flex items-center gap-4 p-3 bg-coffee-cream/30 border border-coffee-latte/40 rounded-sm">
              <div class="w-[72px] h-[72px] flex-shrink-0">
                <FlavorRadarChart
                  :values="quickRadarValues"
                  :size="72"
                  :show-labels="false"
                  :dot-radius="1.5"
                />
              </div>
              <div class="flex-1 space-y-1">
                <div class="text-[9px] uppercase tracking-[0.2em] font-semibold text-coffee-espresso">
                  {{ flavorPresets.find(f => f.val === quickForm.flavor_preset)?.label }} 风味图谱
                </div>
                <div class="grid grid-cols-2 gap-x-3 gap-y-0.5 text-[9px] text-coffee-softGray">
                  <span>酸度 <span class="font-mono text-coffee-espresso">{{ quickRadarValues[0] }}</span></span>
                  <span>苦感 <span class="font-mono text-coffee-espresso">{{ quickRadarValues[1] }}</span></span>
                  <span>甜感 <span class="font-mono text-coffee-espresso">{{ quickRadarValues[2] }}</span></span>
                  <span>醇厚 <span class="font-mono text-coffee-espresso">{{ quickRadarValues[3] }}</span></span>
                  <span>香气 <span class="font-mono text-coffee-espresso">{{ quickRadarValues[4] }}</span></span>
                  <span>余韵 <span class="font-mono text-coffee-espresso">{{ quickRadarValues[5] }}</span></span>
                </div>
              </div>
            </div>
          </Transition>
        </div>

        <!-- Optional: Coffee Name (quick fill) -->
        <div class="space-y-2">
          <label class="text-[10px] uppercase tracking-[0.2em] font-semibold text-coffee-espresso block">咖啡名称 <span class="text-coffee-softGray font-normal normal-case">(选填)</span></label>
          <input 
            type="text" 
            v-model="quickForm.coffee_name" 
            placeholder="不填则自动生成" 
            class="w-full p-3 bg-coffee-cream/40 border border-coffee-latte/60 focus:border-coffee-brown focus:outline-none rounded-sm font-serif text-sm transition-colors"
          >
        </div>
      </div>

      <!-- ==================== DETAILED LOG MODE ==================== -->

      <!-- STEP 1: Basic Information -->
      <div v-if="logMode === 'detailed' && step === 1" class="space-y-6">
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

        <!-- 4. Diary Writing Section -->
        <div class="space-y-2">
          <div class="flex items-center gap-2">
            <label class="text-[10px] uppercase tracking-[0.2em] font-semibold text-coffee-espresso select-none">4. 手账日记 <span class="text-coffee-softGray font-normal normal-case">(选填)</span></label>
          </div>
          <div class="relative">
            <textarea 
              v-model="form.notes"
              rows="5"
              placeholder="今天在哪里？除了咖啡还有什么让你印象深刻？就用这段文字，把这一刻永久留下来……"
              class="w-full p-4 bg-coffee-cream/30 border border-coffee-latte/40 focus:border-coffee-brown focus:outline-none rounded-sm text-sm font-serif leading-relaxed resize-none transition-colors placeholder:text-coffee-softGray/60 placeholder:italic"
            ></textarea>
            <div class="absolute bottom-3 right-3 text-[9px] text-coffee-softGray/50 font-mono select-none">{{ form.notes.length }} 字</div>
          </div>
          <p class="text-[9px] text-coffee-softGray/70 italic select-none">AI 将在保存时参考这段文字，生成更有温度的感官评语。</p>
        </div>
      </div>

      <!-- STEP 2: Sensory Sliders with live SVG radar rendering -->
      <div v-if="logMode === 'detailed' && step === 2" class="space-y-6">
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
      <div v-if="logMode === 'detailed' && step === 3" class="space-y-5">
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
              class="p-2 border rounded-sm text-xs font-serif transition-all flex flex-col items-center gap-1"
              :class="form.mood === m.val
                ? 'border-coffee-brown bg-coffee-cream text-coffee-espresso font-semibold'
                : 'border-coffee-latte/50 text-coffee-softGray hover:border-coffee-brown'"
            >
              <component :is="m.icon" class="w-3.5 h-3.5 flex-shrink-0" />
              <span>{{ m.label }}</span>
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

      </div>

    </div>

    <!-- Bottom Controls -->
    <div class="p-6 border-t border-coffee-cream flex gap-3 bg-coffee-warmWhite sticky bottom-0 z-10 select-none">
      <!-- Detailed mode: show back button -->
      <button 
        v-if="logMode === 'detailed' && step > 1" 
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
          <span v-if="logMode === 'quick'">一键保存</span>
          <span v-else>{{ step === 3 ? 'AI 总结并保存手账' : '下一步' }}</span>
        </template>
      </button>
    </div>

  </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed } from 'vue'
import { useRouter } from 'vue-router'
import { useCoffeeLogStore, NewCoffeeLog } from '@/stores/coffeeLog'
import FlavorRadarChart from '@/components/charts/FlavorRadarChart.vue'
import request from '@/api/request'
import { X, Check, Plus, Smile, Zap, Moon, CloudRain, Sun, Leaf, Heart, Flame } from 'lucide-vue-next'

const router = useRouter()
const store = useCoffeeLogStore()

// Mode: 'quick' = single page, 'detailed' = 3-step wizard
const logMode = ref<'quick' | 'detailed'>('quick')
const step = ref(1)
const isSubmitting = ref(false)

// File upload states
const fileInput = ref<HTMLInputElement | null>(null)
const isUploading = ref(false)
const isLocalUploaded = ref(false)

// Quick Log Form — minimal fields, defaults for the rest
const quickForm = reactive({
  image_url: store.DEFAULT_PHOTOS[1],
  coffee_type: 'Pour Over',
  mood: 'Calm',
  coffee_name: '',
  flavor_preset: ''
})

// Detailed Log Form — full 3-step wizard
const form = reactive<NewCoffeeLog>({
  coffee_name: '',
  coffee_type: 'Pour Over',
  image_url: store.DEFAULT_PHOTOS[1],
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

// Auto-generate coffee name for quick log
const generateQuickName = (type: string) => {
  const typeMap: Record<string, string> = {
    'Pour Over': '手冲咖啡',
    'Latte': '拿铁',
    'Americano': '美式咖啡',
    'Cold Brew': '冷萃咖啡',
    'Espresso': '浓缩咖啡',
    'Dirty': 'Dirty 咖啡'
  }
  return typeMap[type] || '咖啡'
}

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
const flavorPresets = [
  {
    val: 'bright',
    label: '清新明亮',
    desc: '高酸 · 轻盈',
    icon: Sun,
    values: { acidity: 5, bitterness: 1, sweetness: 3, body: 2, aroma: 4, aftertaste: 3 }
  },
  {
    val: 'floral',
    label: '花果芬芳',
    desc: '果香 · 花香',
    icon: Leaf,
    values: { acidity: 4, bitterness: 1, sweetness: 4, body: 2, aroma: 5, aftertaste: 3 }
  },
  {
    val: 'smooth',
    label: '甜美柔滑',
    desc: '低酸 · 甜润',
    icon: Heart,
    values: { acidity: 2, bitterness: 1, sweetness: 5, body: 3, aroma: 3, aftertaste: 4 }
  },
  {
    val: 'bold',
    label: '浓郁醇厚',
    desc: '厚重 · 回甘',
    icon: Flame,
    values: { acidity: 1, bitterness: 4, sweetness: 2, body: 5, aroma: 4, aftertaste: 5 }
  }
]

const quickRadarValues = computed(() => {
  const preset = flavorPresets.find(f => f.val === quickForm.flavor_preset)
  if (preset) {
    const v = preset.values
    return [v.acidity, v.bitterness, v.sweetness, v.body, v.aroma, v.aftertaste]
  }
  return [3, 2, 3, 3, 3, 3]
})

const moodPresets = [
  { val: 'Calm', label: '平静', icon: Smile },
  { val: 'Energetic', label: '愉悦', icon: Zap },
  { val: 'Reflective', label: '沉浸', icon: Moon },
  { val: 'Tired', label: '疲惫', icon: CloudRain }
]

const toggleTag = (tag: string) => {
  if (form.flavor_tags.includes(tag)) {
    form.flavor_tags = form.flavor_tags.filter((t: string) => t !== tag)
  } else {
    form.flavor_tags.push(tag)
  }
}

const handleNext = async () => {
  // Quick Log: save immediately
  if (logMode.value === 'quick') {
    isSubmitting.value = true
    const quickLog: NewCoffeeLog = {
      coffee_name: quickForm.coffee_name.trim() || generateQuickName(quickForm.coffee_type),
      coffee_type: quickForm.coffee_type,
      image_url: quickForm.image_url,
      mood: quickForm.mood,
      shop_name: 'Local Coffee Spot',
      notes: '一杯温润安静的手账记录。',
      ...(flavorPresets.find(f => f.val === quickForm.flavor_preset)?.values ?? { acidity: 3, bitterness: 2, sweetness: 3, body: 3, aroma: 3, aftertaste: 3 }),
      flavor_tags: []
    }
    try {
      const created = await store.addLog(quickLog)
      isSubmitting.value = false
      router.push(`/coffee/${created.id}`)
    } catch (e: any) {
      isSubmitting.value = false
      alert(e.message || '保存失败，请稍后重试')
    }
    return
  }

  // Detailed Log: step-by-step wizard
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
    isSubmitting.value = true
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
      // Set image on the active form based on current mode
      if (logMode.value === 'quick') {
        quickForm.image_url = res.url
      } else {
        form.image_url = res.url
      }
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
.fade-enter-active, .fade-leave-active { transition: opacity 0.2s ease, transform 0.2s ease; }
.fade-enter-from, .fade-leave-to { opacity: 0; transform: translateY(-4px); }

/* Range slider styling */
input[type="range"]::-webkit-slider-thumb {
  border-radius: 50%;
  width: 12px;
  height: 12px;
  background: #7A5638;
}
</style>
