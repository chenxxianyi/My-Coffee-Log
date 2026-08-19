<template>
  <div class="flex-1 w-full flex flex-col bg-coffee-warmWhite text-coffee-charcoal min-h-screen">

    <!-- Header -->
    <div class="px-5 py-3.5 border-b border-coffee-cream flex justify-between items-center bg-coffee-warmWhite/95 backdrop-blur-sm sticky top-0 z-10 select-none">
      <div class="w-9"></div>
      <div class="text-center leading-none">
        <h1 class="font-serif text-[17px] font-semibold tracking-wide text-coffee-espresso">欢迎加入</h1>
        <p class="mt-1.5 text-[9px] tracking-[0.16em] text-coffee-softGray">让我们更了解你的咖啡偏好</p>
      </div>
      <button
        @click="handleSkip"
        class="text-[10px] text-coffee-softGray hover:text-coffee-brown transition-colors tracking-wider"
      >跳过</button>
    </div>

    <!-- Progress Bar -->
    <div class="h-1 bg-coffee-cream w-full flex select-none">
      <div class="h-full bg-coffee-brown transition-all duration-300" :style="{ width: (currentStep * 33.3) + '%' }"></div>
    </div>

    <!-- Step Content -->
    <div class="flex-1 overflow-y-auto px-6 py-8">

      <!-- Step 1: Preferred Coffee Types -->
      <div v-if="currentStep === 1" class="space-y-6 animate-fade-in">
        <div class="text-center mb-8">
          <h2 class="font-serif text-xl font-semibold text-coffee-espresso mb-2">你常喝什么咖啡？</h2>
          <p class="text-sm text-coffee-softGray">选择你喜欢的类型，我们会优先推荐</p>
        </div>

        <div class="grid grid-cols-2 gap-3">
          <button
            v-for="t in coffeeTypes"
            :key="t.val"
            @click="toggleCoffeeType(t.val)"
            type="button"
            class="p-4 border rounded-lg text-center transition-all duration-200 flex flex-col items-center gap-2"
            :class="selectedTypes.includes(t.val)
              ? 'bg-coffee-cream border-coffee-brown ring-1 ring-coffee-brown text-coffee-espresso'
              : 'bg-white/60 border-coffee-latte/50 hover:border-coffee-brown text-coffee-softGray'"
          >
            <span class="text-2xl">{{ t.icon }}</span>
            <span class="text-xs font-serif font-medium">{{ t.label }}</span>
          </button>
        </div>
      </div>

      <!-- Step 2: Preferred Log Mode -->
      <div v-if="currentStep === 2" class="space-y-6 animate-fade-in">
        <div class="text-center mb-8">
          <h2 class="font-serif text-xl font-semibold text-coffee-espresso mb-2">你想怎么记录？</h2>
          <p class="text-sm text-coffee-softGray">随时可以在设置中切换</p>
        </div>

        <div class="space-y-4">
          <button
            @click="preferredMode = 'quick'"
            type="button"
            class="w-full p-5 border rounded-lg text-left transition-all duration-200"
            :class="preferredMode === 'quick'
              ? 'bg-coffee-cream border-coffee-brown ring-1 ring-coffee-brown'
              : 'bg-white/60 border-coffee-latte/50 hover:border-coffee-brown'"
          >
            <div class="flex items-center gap-3">
              <span class="text-2xl">⚡</span>
              <div>
                <div class="text-sm font-serif font-semibold text-coffee-espresso">轻松记录</div>
                <div class="text-xs text-coffee-softGray mt-1">选好类型和心情，10秒完成</div>
              </div>
            </div>
          </button>

          <button
            @click="preferredMode = 'detailed'"
            type="button"
            class="w-full p-5 border rounded-lg text-left transition-all duration-200"
            :class="preferredMode === 'detailed'
              ? 'bg-coffee-cream border-coffee-brown ring-1 ring-coffee-brown'
              : 'bg-white/60 border-coffee-latte/50 hover:border-coffee-brown'"
          >
            <div class="flex items-center gap-3">
              <span class="text-2xl">📝</span>
              <div>
                <div class="text-sm font-serif font-semibold text-coffee-espresso">精细品鉴</div>
                <div class="text-xs text-coffee-softGray mt-1">记录风味、参数和感官评分</div>
              </div>
            </div>
          </button>
        </div>
      </div>

      <!-- Step 3: First Coffee (引导进入记录) -->
      <div v-if="currentStep === 3" class="space-y-6 animate-fade-in">
        <div class="text-center mb-8">
          <div class="inline-grid w-16 h-16 place-items-center rounded-full bg-coffee-cream/60 mb-4">
            <Coffee class="w-8 h-8 text-coffee-brown" />
          </div>
          <h2 class="font-serif text-xl font-semibold text-coffee-espresso mb-2">记录你的第一杯</h2>
          <p class="text-sm text-coffee-softGray">开启你的咖啡手账之旅</p>
        </div>

        <div class="rounded-lg border border-coffee-cream bg-white/60 p-5 text-center">
          <p class="text-sm text-coffee-softGray leading-relaxed">
            记录不需要很完美。<br>
            一杯咖啡、一个心情，就足以留下一段记忆。
          </p>
        </div>
      </div>

    </div>

    <!-- Bottom Actions -->
    <div class="border-t border-coffee-cream bg-coffee-warmWhite/95 backdrop-blur-sm sticky bottom-0 z-10 select-none p-5 space-y-3">
      <button
        @click="handleNext"
        class="w-full py-3 text-xs uppercase tracking-wider font-semibold bg-coffee-espresso text-coffee-warmWhite hover:bg-coffee-brown transition-all rounded-sm"
      >
        {{ currentStep === 3 ? '开始记录第一杯' : '下一步' }}
      </button>
    </div>

  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import request from '@/api/request'
import { Coffee } from 'lucide-vue-next'

const router = useRouter()

const currentStep = ref(1)
const selectedTypes = ref<string[]>([])
const preferredMode = ref<'quick' | 'detailed'>('quick')

const coffeeTypes = [
  { val: 'Pour Over', label: '手冲', icon: '☕' },
  { val: 'Latte', label: '拿铁', icon: '🥛' },
  { val: 'Americano', label: '美式', icon: '🫗' },
  { val: 'Cold Brew', label: '冷萃', icon: '🧊' },
  { val: 'Espresso', label: '浓缩', icon: '☕' },
  { val: 'Dirty', label: '脏咖啡', icon: '🫙' },
]

function toggleCoffeeType(type: string) {
  const idx = selectedTypes.value.indexOf(type)
  if (idx >= 0) {
    selectedTypes.value.splice(idx, 1)
  } else {
    selectedTypes.value.push(type)
  }
}

async function handleSkip() {
  await completeOnboarding(true)
}

async function handleNext() {
  if (currentStep.value < 3) {
    currentStep.value++
    return
  }

  await completeOnboarding(false)
}

async function completeOnboarding(skip: boolean) {
  try {
    await request.put('/users/me/onboarding', {
      preferred_log_mode: preferredMode.value,
      preferred_coffee_types: JSON.stringify(selectedTypes.value),
      skip
    })
  } catch (e) {
    console.error('Failed to save onboarding:', e)
  }
  router.push('/home')
}
</script>

<style scoped>
@keyframes fadeIn {
  from { opacity: 0; transform: translateY(8px); }
  to { opacity: 1; transform: translateY(0); }
}
.animate-fade-in {
  animation: fadeIn 0.3s ease-out;
}
</style>
