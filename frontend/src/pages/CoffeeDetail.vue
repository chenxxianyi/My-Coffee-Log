<template>
  <div v-if="log" class="flex-1 w-full flex flex-col justify-between bg-coffee-warmWhite text-coffee-charcoal">

    <!-- Celebration Banner (shown after creation) -->
    <Transition name="celebration">
      <div v-if="showCelebration" class="fixed top-0 inset-x-0 z-50 flex justify-center pointer-events-none">
        <div class="mt-4 px-5 py-2.5 bg-coffee-espresso/95 backdrop-blur-md rounded-sm shadow-lg flex items-center gap-2.5 pointer-events-auto">
          <div class="w-5 h-5 rounded-full bg-green-500/20 flex items-center justify-center">
            <Check class="w-3 h-3 text-green-400" />
          </div>
          <span class="text-[11px] text-coffee-warmWhite font-semibold tracking-wider uppercase">手账记录成功</span>
          <div class="w-px h-3 bg-coffee-latte/30"></div>
          <span class="text-[10px] text-coffee-latte font-serif italic">本月第 {{ store.monthBrews }} 杯</span>
        </div>
      </div>
    </Transition>

    <!-- Celebration Particles -->
    <Transition name="fade">
      <div v-if="showCelebration" class="fixed inset-0 z-40 pointer-events-none overflow-hidden">
        <div v-for="i in 8" :key="i" class="celebration-particle" :style="particleStyle(i)"></div>
      </div>
    </Transition>

    <!-- Floating Share FAB -->
    <div class="fixed bottom-8 right-5 z-30">
      <button 
        @click="router.push(`/share/${log.id}`)" 
        class="w-12 h-12 rounded-full flex items-center justify-center shadow-xl ring-4 ring-coffee-warmWhite transition-all duration-200 hover:scale-110 active:scale-95"
        style="background: linear-gradient(145deg, #E76F51, #D4623E);"
      >
        <Share2 class="w-5 h-5 text-white" />
      </button>
      <span class="block text-center text-[8px] text-coffee-softGray mt-1 tracking-wider font-semibold select-none">分享</span>
    </div>

    <!-- Header transparent gradient overlay -->
    <div class="absolute top-0 inset-x-0 h-16 flex justify-between items-center px-6 z-20 bg-gradient-to-b from-black/55 to-transparent select-none">
      <button @click="router.push('/home')" class="w-9 h-9 rounded-full bg-white/20 backdrop-blur-md text-white flex items-center justify-center hover:bg-white/35 transition-colors">
        <ArrowLeft class="w-4 h-4" />
      </button>
      <div class="flex items-center gap-2">
        <button @click="router.push(`/share/${log.id}`)" class="w-9 h-9 rounded-full bg-white/20 backdrop-blur-md text-white flex items-center justify-center hover:bg-white/35 transition-colors">
          <Share2 class="w-4 h-4" />
        </button>
      </div>
    </div>

    <!-- Scrollable Detail Contents -->
    <div class="flex-1 overflow-y-auto pb-16 scrollbar-none">
      
      <!-- ============ HERO COVER PHOTO ============ -->
      <div class="w-full h-[340px] relative bg-coffee-espresso overflow-hidden" :class="justCreated ? 'animate-fade-in' : ''">
        <img :src="log.image_url" class="w-full h-full object-cover scale-105">
        <!-- Multi-layer gradient for magazine depth -->
        <div class="absolute inset-0 bg-gradient-to-t from-coffee-warmWhite via-coffee-warmWhite/20 to-transparent"></div>
        <div class="absolute inset-0 bg-gradient-to-b from-black/40 via-transparent to-transparent"></div>
        <!-- Paper grain texture -->
        <div class="absolute inset-0 opacity-[0.04]" style="background-image: radial-gradient(rgba(255,242,219,0.6) 1px, transparent 0); background-size: 18px 18px;"></div>
        <!-- Coffee type badge (bilingual, editorial) -->
        <div class="absolute bottom-6 left-6 flex items-center gap-2.5 select-none">
          <span class="bg-coffee-espresso/90 backdrop-blur-sm text-coffee-warmWhite px-3.5 py-1.5 rounded-sm text-[10px] font-serif uppercase tracking-[0.25em] leading-none font-bold">
            {{ log.coffee_type }}
          </span>
          <span class="text-[10px] text-coffee-warmWhite/70 font-serif italic">/ {{ coffeeTypeShortLabel(log.coffee_type) }}</span>
        </div>
        <!-- Corner bracket decorations on cover -->
        <div class="absolute top-14 left-5 w-5 h-5 border-t border-l border-white/20"></div>
        <div class="absolute top-14 right-5 w-5 h-5 border-t border-r border-white/20"></div>
      </div>

      <!-- ============ MAGAZINE BODY CONTENTS ============ -->
      <div class="px-6 space-y-7 -mt-3 relative z-10">
        
        <!-- ---- Editorial Title Block ---- -->
        <div class="space-y-3" :class="justCreated ? 'animate-slide-up' : ''">
          <!-- Date with flanking lines -->
          <div class="flex items-center gap-2.5 select-none">
            <div class="h-px flex-1 max-w-[28px]" style="background: linear-gradient(to right, rgba(92,61,46,0.4), transparent);"></div>
            <span class="text-[9px] uppercase tracking-[0.25em] font-semibold text-coffee-softGray">{{ fullDate }}</span>
            <div class="h-px flex-1 max-w-[28px]" style="background: linear-gradient(to left, rgba(92,61,46,0.4), transparent);"></div>
          </div>
          <!-- Magazine-style large serif title -->
          <h1 class="font-serif text-[42px] font-light italic text-coffee-espresso leading-[1.05] break-words tracking-wide">{{ log.coffee_name }}</h1>
          <!-- Meta info strip: location + mood -->
          <div class="flex items-center gap-3 text-xs select-none">
            <div class="flex items-center gap-1.5 text-coffee-brown font-medium">
              <MapPin class="w-3 h-3" />
              <span>{{ log.shop_name }}</span>
            </div>
            <div class="w-px h-3 bg-coffee-cream"></div>
            <span class="text-coffee-espresso italic font-semibold">{{ moodLabel(log.mood) }}</span>
          </div>
        </div>

        <!-- Hairline separator -->
        <div class="flex items-center gap-3 select-none">
          <div class="w-3 h-px bg-coffee-softGray/40"></div>
          <div class="flex-1 h-px bg-coffee-cream"></div>
          <div class="w-3 h-px bg-coffee-softGray/40"></div>
        </div>

        <!-- ---- Monthly Progress Card (shown after creation) ---- -->
        <Transition name="slide-fade">
          <div v-if="justCreated" class="p-4 rounded-sm border border-coffee-latte/40" style="background: linear-gradient(135deg, rgba(215,196,168,0.25) 0%, rgba(231,111,81,0.08) 100%);">
            <div class="flex justify-between items-center mb-2.5">
              <span class="text-[9px] uppercase tracking-[0.2em] font-bold text-coffee-brown">本月咖啡进度</span>
              <span class="text-[10px] font-serif italic text-coffee-espresso">{{ store.monthBrews }} / {{ nextMilestone }}</span>
            </div>
            <div class="w-full h-1.5 bg-coffee-cream/60 rounded-full overflow-hidden">
              <div 
                class="h-full rounded-full transition-all duration-700 ease-out"
                :style="{ width: progressPercent + '%', background: 'linear-gradient(90deg, #D7C4A8, #E76F51)' }"
              ></div>
            </div>
            <div class="flex justify-between mt-1.5 text-[8px] text-coffee-softGray select-none">
              <span v-for="m in milestones" :key="m" :class="store.monthBrews >= m ? 'text-coffee-brown font-semibold' : ''">{{ m }}杯</span>
            </div>
            <p class="text-[10px] text-coffee-brown font-serif italic mt-2">
              {{ milestoneMessage }}
            </p>
          </div>
        </Transition>

        <!-- ---- AI Editorial Prose (Magazine Feature) ---- -->
        <div class="relative p-6 bg-coffee-cream/60 rounded-sm double-border space-y-3" :class="justCreated ? 'animate-slide-up-delay-1' : ''">
          <!-- Decorative large opening quote -->
          <div class="absolute top-2 left-4 font-serif leading-none select-none pointer-events-none text-coffee-espresso/[0.12]" style="font-size: 72px; line-height: 1;">&ldquo;</div>
          <!-- Section header -->
          <div class="flex justify-between items-center select-none relative z-10">
            <span class="text-[9px] uppercase tracking-[0.25em] font-bold text-coffee-brown">AI 感官评语 / Editorial</span>
            <span class="text-[10px] text-coffee-espresso italic font-semibold">{{ moodLabel(log.mood) }}</span>
          </div>
          <!-- AI prose text -->
          <p class="font-serif italic text-[15px] text-coffee-espresso leading-[1.8] relative z-10">
            {{ log.ai_summary }}
          </p>
          <!-- Closing ornament -->
          <div class="flex items-center justify-center gap-2 pt-1 select-none">
            <div class="w-4 h-px bg-coffee-brown/25"></div>
            <div class="w-1 h-1 rounded-full bg-coffee-brown/30"></div>
            <div class="w-4 h-px bg-coffee-brown/25"></div>
          </div>
        </div>

        <!-- ---- Sensory Radar (Centered, Larger) ---- -->
        <div class="space-y-4" :class="justCreated ? 'animate-slide-up-delay-2' : ''">
          <div class="flex items-center gap-2.5 select-none">
            <div class="h-px flex-1 max-w-[24px]" style="background: linear-gradient(to right, rgba(92,61,46,0.35), transparent);"></div>
            <h3 class="text-[10px] uppercase tracking-[0.2em] font-bold text-coffee-espresso">风味足迹雷达 / Sensory Radar</h3>
            <div class="h-px flex-1" style="background: linear-gradient(to left, rgba(92,61,46,0.35), transparent);"></div>
          </div>
          
          <!-- Centered radar with circular frame -->
          <div class="flex justify-center">
            <div class="w-[170px] h-[170px] flex items-center justify-center p-2 bg-coffee-cream/30 rounded-full border border-coffee-latte/25">
              <FlavorRadarChart 
                :values="[log.acidity, log.bitterness, log.sweetness, log.body, log.aroma, log.aftertaste]"
                :size="150"
                :dimensions="['Acid', 'Bitter', 'Sweet', 'Body', 'Aroma', 'After']"
                :label-font-size="8"
                :dot-radius="2.5"
              />
            </div>
          </div>
          <!-- Score list below radar -->
          <div class="grid grid-cols-3 gap-x-4 gap-y-2 text-[10.5px] text-coffee-brown font-medium select-none">
            <div class="flex justify-between items-center"><span>Acidity / 酸度</span><span class="font-mono text-coffee-espresso font-semibold">{{ log.acidity }}</span></div>
            <div class="flex justify-between items-center"><span>Bitter / 苦感</span><span class="font-mono text-coffee-espresso font-semibold">{{ log.bitterness }}</span></div>
            <div class="flex justify-between items-center"><span>Sweet / 甜感</span><span class="font-mono text-coffee-espresso font-semibold">{{ log.sweetness }}</span></div>
            <div class="flex justify-between items-center"><span>Body / 醇厚</span><span class="font-mono text-coffee-espresso font-semibold">{{ log.body }}</span></div>
            <div class="flex justify-between items-center"><span>Aroma / 香气</span><span class="font-mono text-coffee-espresso font-semibold">{{ log.aroma }}</span></div>
            <div class="flex justify-between items-center"><span>After / 余韵</span><span class="font-mono text-coffee-espresso font-semibold">{{ log.aftertaste }}</span></div>
          </div>
        </div>

        <!-- ---- Flavor Tags (Editorial Style) ---- -->
        <div class="space-y-3" :class="justCreated ? 'animate-slide-up-delay-3' : ''">
          <div class="flex items-center gap-2.5 select-none">
            <div class="h-px flex-1 max-w-[24px]" style="background: linear-gradient(to right, rgba(92,61,46,0.35), transparent);"></div>
            <h3 class="text-[10px] uppercase tracking-[0.2em] font-bold text-coffee-espresso">感官风味标签 / Flavor Tags</h3>
            <div class="h-px flex-1" style="background: linear-gradient(to left, rgba(92,61,46,0.35), transparent);"></div>
          </div>
          <div class="flex flex-wrap gap-2 select-none">
            <span 
              v-for="tag in log.flavor_tags" 
              :key="tag" 
              class="px-3.5 py-1.5 text-[11px] bg-coffee-warmWhite text-coffee-espresso border border-coffee-latte/40 rounded-full font-mono uppercase tracking-wider shadow-sm"
            >
              ★ {{ tag }}
            </span>
          </div>
        </div>

        <!-- ---- Lifestyle Tags (Mood / Scene / Pairing) ---- -->
        <div v-if="hasLifestyleTags" class="space-y-3">
          <div class="flex items-center gap-2.5 select-none">
            <div class="h-px flex-1 max-w-[24px]" style="background: linear-gradient(to right, rgba(92,61,46,0.35), transparent);"></div>
            <h3 class="text-[10px] uppercase tracking-[0.2em] font-bold text-coffee-espresso">生活标签 / Lifestyle</h3>
            <div class="h-px flex-1" style="background: linear-gradient(to left, rgba(92,61,46,0.35), transparent);"></div>
          </div>
          <div class="space-y-2.5">
            <!-- Mood Tags -->
            <div v-if="log.mood_tags?.length" class="space-y-1.5">
              <span class="text-[9px] uppercase tracking-wider text-amber-700 font-semibold select-none">心情 Mood</span>
              <div class="flex flex-wrap gap-1.5">
                <span v-for="t in log.mood_tags" :key="t" class="px-2.5 py-1 text-[10px] bg-amber-100 text-amber-800 border border-amber-300/60 rounded-full font-medium">{{ lifestyleTagLabel('mood', t) }}</span>
              </div>
            </div>
            <!-- Scene Tags -->
            <div v-if="log.scene_tags?.length" class="space-y-1.5">
              <span class="text-[9px] uppercase tracking-wider text-sky-700 font-semibold select-none">场景 Scene</span>
              <div class="flex flex-wrap gap-1.5">
                <span v-for="t in log.scene_tags" :key="t" class="px-2.5 py-1 text-[10px] bg-sky-100 text-sky-800 border border-sky-300/60 rounded-full font-medium">{{ lifestyleTagLabel('scene', t) }}</span>
              </div>
            </div>
            <!-- Pairing Tags -->
            <div v-if="log.pairing_tags?.length" class="space-y-1.5">
              <span class="text-[9px] uppercase tracking-wider text-rose-700 font-semibold select-none">搭配 Pairing</span>
              <div class="flex flex-wrap gap-1.5">
                <span v-for="t in log.pairing_tags" :key="t" class="px-2.5 py-1 text-[10px] bg-rose-100 text-rose-800 border border-rose-300/60 rounded-full font-medium">{{ lifestyleTagLabel('pairing', t) }}</span>
              </div>
            </div>
          </div>
        </div>

        <!-- ---- Coffee Bean Archive ---- -->
        <div v-if="hasBeanInfo" class="space-y-3">
          <div class="flex items-center gap-2.5 select-none">
            <div class="h-px flex-1 max-w-[24px]" style="background: linear-gradient(to right, rgba(92,61,46,0.35), transparent);"></div>
            <h3 class="text-[10px] uppercase tracking-[0.2em] font-bold text-coffee-espresso">咖啡豆档案 / Bean Archive</h3>
            <div class="h-px flex-1" style="background: linear-gradient(to left, rgba(92,61,46,0.35), transparent);"></div>
          </div>
          <div class="p-4 bg-coffee-cream/30 border border-coffee-latte/30 rounded-sm space-y-2">
            <div v-if="log.bean" class="space-y-1.5">
              <div class="font-serif text-sm font-semibold text-coffee-espresso">{{ log.bean.name }}</div>
              <div class="grid grid-cols-2 gap-x-4 gap-y-1 text-[10px]">
                <div v-if="log.bean.origin" class="flex justify-between"><span class="text-coffee-softGray">产地</span><span class="text-coffee-espresso font-medium">{{ log.bean.origin }}</span></div>
                <div v-if="log.bean.processing_method" class="flex justify-between"><span class="text-coffee-softGray">处理法</span><span class="text-coffee-espresso font-medium">{{ log.bean.processing_method }}</span></div>
                <div v-if="log.bean.roast_level" class="flex justify-between"><span class="text-coffee-softGray">烘焙度</span><span class="text-coffee-espresso font-medium">{{ log.bean.roast_level }}</span></div>
                <div v-if="log.bean.roaster" class="flex justify-between"><span class="text-coffee-softGray">烘焙商</span><span class="text-coffee-espresso font-medium">{{ log.bean.roaster }}</span></div>
              </div>
            </div>
            <div v-if="hasBrewParams" class="pt-2 mt-2 border-t border-coffee-cream grid grid-cols-3 gap-2 text-[10px]">
              <div v-if="log.brew_ratio" class="text-center"><div class="text-coffee-espresso font-mono font-semibold">{{ log.brew_ratio }}</div><div class="text-coffee-softGray">粉水比</div></div>
              <div v-if="log.water_temp" class="text-center"><div class="text-coffee-espresso font-mono font-semibold">{{ log.water_temp }}</div><div class="text-coffee-softGray">水温</div></div>
              <div v-if="log.grind_size" class="text-center"><div class="text-coffee-espresso font-mono font-semibold">{{ log.grind_size }}</div><div class="text-coffee-softGray">研磨度</div></div>
            </div>
          </div>
        </div>

        <!-- ---- Taste Notes Diary ---- -->
        <div class="space-y-2.5" :class="justCreated ? 'animate-slide-up-delay-3' : ''">
          <div class="flex items-center gap-2.5 select-none">
            <div class="h-px flex-1 max-w-[24px]" style="background: linear-gradient(to right, rgba(92,61,46,0.35), transparent);"></div>
            <h3 class="text-[10px] uppercase tracking-[0.2em] font-bold text-coffee-espresso">感官味觉日记 / Diary</h3>
            <div class="h-px flex-1" style="background: linear-gradient(to left, rgba(92,61,46,0.35), transparent);"></div>
          </div>
          <p class="text-[13px] text-coffee-brown leading-[1.75] font-light font-serif italic pl-2 border-l-2 border-coffee-latte/30">
            {{ log.notes }}
          </p>
        </div>

        <!-- ---- Share CTA (prominent, shown after creation) ---- -->
        <Transition name="slide-fade">
          <div v-if="justCreated" class="p-5 bg-coffee-espresso rounded-sm space-y-3">
            <div class="flex items-center gap-2.5">
              <Share2 class="w-4 h-4 text-coffee-latte" />
              <span class="text-[10px] uppercase tracking-[0.2em] font-bold text-coffee-warmWhite">分享这杯咖啡</span>
            </div>
            <p class="text-[11px] text-coffee-latte/80 font-serif italic leading-relaxed">把这段美好风味变成一张精致卡片，分享给朋友吧。</p>
            <router-link 
              :to="'/share/' + log.id" 
              class="block w-full py-3 text-[10px] uppercase tracking-widest font-semibold text-coffee-espresso hover:text-coffee-brown transition-all rounded-sm flex items-center justify-center gap-1.5 bg-coffee-warmWhite"
            >
              <span>生成分享海报</span>
              <span>→</span>
            </router-link>
          </div>
        </Transition>

        <!-- ---- Bottom Actions ---- -->
        <div class="space-y-3 pt-2 pb-4 select-none">
          <!-- Share button (always visible) -->
          <router-link 
            :to="'/share/' + log.id" 
            class="w-full py-3.5 text-[10px] uppercase tracking-widest font-semibold bg-coffee-espresso text-coffee-warmWhite hover:bg-coffee-brown transition-all rounded-sm flex items-center justify-center gap-2"
          >
            <Share2 class="w-4 h-4" />
            <span>生成分享海报</span>
          </router-link>
          <div class="grid grid-cols-2 gap-3">
            <router-link 
              :to="'/create?from_log_id=' + log.id" 
              class="py-3 text-[10px] uppercase tracking-widest font-semibold bg-coffee-cream text-coffee-espresso hover:bg-coffee-latte transition-all rounded-sm flex items-center justify-center gap-1.5"
            >
              <Plus class="w-3.5 h-3.5" />
              <span>复刻这杯</span>
            </router-link>
            <button 
              @click="handleDelete" 
              :disabled="isDeleting"
              class="py-3 text-[10px] uppercase tracking-widest font-semibold border border-red-200 text-red-700 hover:bg-red-50 transition-all rounded-sm"
            >
              {{ isDeleting ? '删除中...' : '删除记录' }}
            </button>
          </div>
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
import { ref, computed, onMounted, onUnmounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useCoffeeLogStore } from '@/stores/coffeeLog'
import { coffeeTypeShortLabel, moodLabel, LIFESTYLE_MOOD_TAGS, LIFESTYLE_SCENE_TAGS, LIFESTYLE_PAIRING_TAGS } from '@/constants/coffee'
import FlavorRadarChart from '@/components/charts/FlavorRadarChart.vue'
import { ArrowLeft, Share2, MapPin, Coffee, Check, Plus } from 'lucide-vue-next'

const props = defineProps<{
  id: string
}>()

const router = useRouter()
const route = useRoute()
const store = useCoffeeLogStore()

const log = computed(() => store.getLogById(parseInt(props.id)))

// Lifestyle tags helpers
const hasLifestyleTags = computed(() => {
  const l = log.value
  if (!l) return false
  return (l.mood_tags?.length ?? 0) > 0 || (l.scene_tags?.length ?? 0) > 0 || (l.pairing_tags?.length ?? 0) > 0
})

const hasBeanInfo = computed(() => {
  const l = log.value
  if (!l) return false
  return !!l.bean || !!l.bean_id
})

const hasBrewParams = computed(() => {
  const l = log.value
  if (!l) return false
  return !!l.brew_ratio || !!l.water_temp || !!l.grind_size
})

const lifestyleTagLabel = (type: 'mood' | 'scene' | 'pairing', val: string) => {
  const list = type === 'mood' ? LIFESTYLE_MOOD_TAGS : type === 'scene' ? LIFESTYLE_SCENE_TAGS : LIFESTYLE_PAIRING_TAGS
  return list.find(t => t.val === val)?.label ?? val
}

// Celebration state
const justCreated = computed(() => route.query.just_created === 'true')
const showCelebration = ref(false)
const isDeleting = ref(false)
let celebrationTimer: ReturnType<typeof setTimeout> | null = null

// Milestone config
const milestones = [3, 5, 10, 15, 20]
const nextMilestone = computed(() => {
  const count = store.monthBrews
  for (const m of milestones) {
    if (count < m) return m
  }
  return milestones[milestones.length - 1]
})
const progressPercent = computed(() => {
  const count = store.monthBrews
  const target = nextMilestone.value
  return Math.min(Math.round((count / target) * 100), 100)
})
const milestoneMessage = computed(() => {
  const count = store.monthBrews
  if (count >= 20) return '惊人的咖啡生活节奏，你是本月的咖啡大师。'
  if (count >= 15) return '离月度大师只差一步，继续记录你的风味旅程。'
  if (count >= 10) return '双位数达成！你的咖啡手账已经是一本精彩的生活杂志。'
  if (count >= 5) return '五杯里程碑！坚持记录，风味故事正在成型。'
  if (count >= 3) return '三杯起步，你的咖啡月度故事已经开篇。'
  const remaining = nextMilestone.value - count
  return `再记录 ${remaining} 杯，就能解锁下一个里程碑。`
})

// Celebration particle styles
const particleStyle = (i: number) => {
  const colors = ['#D7C4A8', '#E76F51', '#7A5638', '#F7F3EC', '#C4A882']
  const left = 10 + ((i * 37) % 80)
  const delay = i * 0.15
  const duration = 1.8 + (i % 3) * 0.4
  const size = 4 + (i % 4) * 2
  return {
    left: left + '%',
    animationDelay: delay + 's',
    animationDuration: duration + 's',
    width: size + 'px',
    height: size + 'px',
    backgroundColor: colors[i % colors.length]
  }
}

// Fetch from API if not in local cache
onMounted(async () => {
  if (!log.value) {
    try {
      await store.fetchLogById(parseInt(props.id))
    } catch {
      // log not found, will show empty state
    }
  }

  // Refresh stats when arriving from creation
  if (justCreated.value) {
    await store.fetchStats()
    showCelebration.value = true
    celebrationTimer = setTimeout(() => {
      showCelebration.value = false
    }, 3000)
  }
})

onUnmounted(() => {
  if (celebrationTimer) {
    clearTimeout(celebrationTimer)
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
    const id = parseInt(props.id)
    if (Number.isNaN(id) || isDeleting.value) return
    isDeleting.value = true
    try {
      await store.deleteLog(id, { removeFromCache: false })
      await router.replace('/home')
      store.removeLogFromCache(id)
      void Promise.all([
        store.fetchLogs({ page: 1, page_size: 10 }),
        store.fetchStats()
      ]).catch((error) => {
        console.error('Failed to refresh data after delete:', error)
      })
    } catch (e: any) {
      alert(e.message || '删除失败')
    } finally {
      isDeleting.value = false
    }
  }
}
</script>

<style scoped>
/* Celebration banner transition */
.celebration-enter-active {
  transition: all 0.4s cubic-bezier(0.34, 1.56, 0.64, 1);
}
.celebration-leave-active {
  transition: all 0.5s ease-in;
}
.celebration-enter-from {
  opacity: 0;
  transform: translateY(-20px);
}
.celebration-leave-to {
  opacity: 0;
  transform: translateY(-10px);
}

/* Slide-fade transition for progress card & share CTA */
.slide-fade-enter-active {
  transition: all 0.5s cubic-bezier(0.22, 1, 0.36, 1);
}
.slide-fade-leave-active {
  transition: all 0.3s ease-in;
}
.slide-fade-enter-from {
  opacity: 0;
  transform: translateY(12px);
}
.slide-fade-leave-to {
  opacity: 0;
  transform: translateY(-8px);
}

/* Fade transition */
.fade-enter-active { transition: opacity 0.3s ease; }
.fade-leave-active { transition: opacity 0.5s ease; }
.fade-enter-from, .fade-leave-to { opacity: 0; }

/* Celebration particles */
.celebration-particle {
  position: absolute;
  bottom: 0;
  border-radius: 50%;
  animation: particle-rise ease-out forwards;
  pointer-events: none;
}

@keyframes particle-rise {
  0% {
    opacity: 0.9;
    transform: translateY(0) scale(1);
  }
  60% {
    opacity: 0.6;
  }
  100% {
    opacity: 0;
    transform: translateY(-60vh) scale(0.3);
  }
}

/* Entrance animations for just-created state */
.animate-fade-in {
  animation: detail-fade-in 0.6s ease-out;
}
.animate-slide-up {
  animation: detail-slide-up 0.5s ease-out;
}
.animate-slide-up-delay-1 {
  animation: detail-slide-up 0.5s ease-out 0.15s both;
}
.animate-slide-up-delay-2 {
  animation: detail-slide-up 0.5s ease-out 0.3s both;
}
.animate-slide-up-delay-3 {
  animation: detail-slide-up 0.5s ease-out 0.4s both;
}

@keyframes detail-fade-in {
  from { opacity: 0; }
  to { opacity: 1; }
}
@keyframes detail-slide-up {
  from { opacity: 0; transform: translateY(16px); }
  to { opacity: 1; transform: translateY(0); }
}
</style>
