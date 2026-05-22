<template>
  <div class="flex-1 w-full flex flex-col justify-between bg-coffee-warmWhite text-coffee-charcoal">
    
    <!-- Header -->
    <div class="px-6 py-5 flex justify-between items-end border-b border-coffee-cream bg-coffee-warmWhite sticky top-0 z-10 select-none">
      <div>
        <span class="text-[9px] uppercase tracking-[0.25em] font-semibold text-coffee-softGray">Account & Journal</span>
        <h2 class="font-serif text-2xl font-light text-coffee-espresso leading-none mt-1">手账主页</h2>
      </div>
      <button @click="handleLogout" class="text-xs uppercase tracking-widest text-red-600 font-semibold hover:text-red-800 transition-colors">
        退出登录
      </button>
    </div>

    <!-- Scrollable Body Contents -->
    <div class="flex-1 overflow-y-auto px-6 py-6 space-y-6 pb-24 scrollbar-none">
      
      <!-- Tactile Card: User Info -->
      <div class="p-6 bg-coffee-cream rounded-2xl double-border space-y-5 text-center relative overflow-hidden">
        <div 
          @click="showAvatarModal = true"
          class="w-20 h-20 mx-auto rounded-full overflow-hidden border-2 border-coffee-espresso relative group cursor-pointer select-none"
        >
          <img :src="userAvatar" alt="Avatar" class="w-full h-full object-cover">
          <div class="absolute inset-0 bg-black/40 opacity-0 group-hover:opacity-100 flex items-center justify-center text-white text-[10px] transition-opacity">
            修改头像
          </div>
        </div>

        <div class="space-y-1">
          <!-- Inline Nickname Edit -->
          <div class="flex items-center justify-center gap-2">
            <template v-if="isEditing">
              <input 
                type="text" 
                v-model="editNickname" 
                @keyup.enter="saveNickname"
                class="px-2 py-1 text-center font-serif text-xl border border-coffee-brown bg-white/80 focus:outline-none rounded-lg max-w-[150px]"
                autofocus
              >
              <button @click="saveNickname" class="text-coffee-espresso hover:text-coffee-brown">
                <Check class="w-4 h-4" />
              </button>
            </template>
            <template v-else>
              <h3 class="font-serif text-2xl font-light text-coffee-espresso italic truncate max-w-[200px]">{{ authStore.user?.nickname }}</h3>
              <button @click="startEdit" class="text-coffee-softGray hover:text-coffee-espresso transition-colors">
                <Edit2 class="w-3.5 h-3.5" />
              </button>
            </template>
          </div>
          <p class="text-[10px] text-coffee-softGray font-mono tracking-wider">{{ authStore.user?.email }}</p>
        </div>

        <div class="text-[9px] text-coffee-softGray font-bold uppercase tracking-widest pt-2 border-t border-coffee-cream/60">
          味觉生活手账 • 旅程第 {{ daysSinceJoined }} 天
        </div>
      </div>

      <!-- Quick Statistics Insight -->
      <div class="space-y-3">
        <h4 class="text-[10px] uppercase tracking-widest font-semibold text-coffee-softGray mb-1 select-none">味觉手账数据</h4>
        <div class="grid grid-cols-2 gap-4">
          <div class="p-4 bg-coffee-cream/30 border border-coffee-cream/80 rounded-xl">
            <span class="text-[9px] uppercase tracking-wider text-coffee-softGray block">冲煎日志</span>
            <div class="text-3xl font-serif text-coffee-espresso font-light mt-1">{{ logsStore.logs.length }} 篇</div>
          </div>
          <div class="p-4 bg-coffee-cream/30 border border-coffee-cream/80 rounded-xl">
            <span class="text-[9px] uppercase tracking-wider text-coffee-softGray block">首记日期</span>
            <div class="text-[13px] font-serif text-coffee-espresso font-light mt-2 truncate">{{ firstBrewDate }}</div>
          </div>
        </div>
      </div>

      <!-- Editorial Settings Menu -->
      <div class="space-y-2">
        <h4 class="text-[10px] uppercase tracking-widest font-semibold text-coffee-softGray mb-1 select-none">系统偏好</h4>
        
        <div class="editorial-border bg-coffee-cream/10 divide-y divide-coffee-cream">
          <div class="p-4 flex justify-between items-center text-xs">
            <div class="space-y-0.5">
              <span class="font-medium text-coffee-espresso">每日推送提醒</span>
              <p class="text-[10px] text-coffee-softGray">于每日 10:00 提醒你记录今日咖啡风味</p>
            </div>
            <div class="w-8 h-4 bg-coffee-cream rounded-full relative cursor-pointer p-0.5" @click="togglePush">
              <div class="w-3 h-3 bg-coffee-espresso rounded-full transition-all" :class="pushEnabled ? 'translate-x-4' : ''"></div>
            </div>
          </div>

          <router-link to="/coffee-shops" class="p-4 flex justify-between items-center text-xs cursor-pointer hover:bg-coffee-cream/15 transition-colors">
            <div class="space-y-0.5">
              <span class="font-medium text-coffee-espresso">咖啡店收藏</span>
              <p class="text-[10px] text-coffee-softGray">浏览和管理你记录过的咖啡店</p>
            </div>
            <ChevronRight class="w-4 h-4 text-coffee-softGray" />
          </router-link>

          <div class="p-4 flex justify-between items-center text-xs cursor-pointer hover:bg-coffee-cream/15 transition-colors" @click="exportBackup">
            <div class="space-y-0.5">
              <span class="font-medium text-coffee-espresso">导出数据备份</span>
              <p class="text-[10px] text-coffee-softGray">一键导出 JSON 格式的手账完整历史</p>
            </div>
            <ChevronRight class="w-4 h-4 text-coffee-softGray" />
          </div>

          <div class="p-4 flex justify-between items-center text-xs cursor-pointer hover:bg-coffee-cream/15 transition-colors" @click="showAbout">
            <div class="space-y-0.5">
              <span class="font-medium text-coffee-espresso">关于 My Coffee Log</span>
              <p class="text-[10px] text-coffee-softGray">Version 1.0.0 (Nordic Minimal)</p>
            </div>
            <ChevronRight class="w-4 h-4 text-coffee-softGray" />
          </div>
        </div>
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
      <router-link to="/stats" class="flex-1 flex flex-col items-center gap-0.5 text-coffee-softGray hover:text-coffee-brown transition-colors">
        <BarChart3 class="w-5 h-5" />
        <span class="text-[9.5px] tracking-widest font-medium">咖迹</span>
      </router-link>
      <router-link to="/profile" class="flex-1 flex flex-col items-center gap-0.5 text-coffee-brown">
        <User class="w-5 h-5" />
        <span class="text-[9.5px] tracking-widest font-medium">我的</span>
      </router-link>
    </div>

    <!-- Avatar Selector Modal — Redesigned -->
    <div v-if="showAvatarModal" class="fixed inset-0 bg-black/65 backdrop-blur-sm z-50 flex items-center justify-center p-6">
      <div class="bg-coffee-warmWhite w-full max-w-[300px] rounded-2xl shadow-2xl overflow-hidden">

        <!-- Header -->
        <div class="px-5 pt-5 pb-4 flex justify-between items-center border-b border-coffee-cream/70">
          <span class="text-[10px] uppercase tracking-widest font-bold text-coffee-softGray select-none">修改手账头像</span>
          <button @click="closeAvatarModal" class="text-coffee-softGray hover:text-coffee-espresso transition-colors">
            <X class="w-4 h-4" />
          </button>
        </div>

        <!-- Preview bar -->
        <div class="px-5 py-4 flex items-center gap-3 bg-coffee-cream/20 border-b border-coffee-cream/50">
          <div class="w-11 h-11 rounded-full overflow-hidden border-2 border-coffee-espresso flex-shrink-0 shadow-sm">
            <img :src="selectedAvatarUrl || userAvatar" class="w-full h-full object-cover">
          </div>
          <div class="select-none">
            <p class="text-[11px] font-semibold text-coffee-espresso leading-none">头像预览</p>
            <p class="text-[9px] text-coffee-softGray mt-1">从下方选择预设或上传本地相片</p>
          </div>
        </div>

        <!-- Unified 3×3 Grid: first cell = upload, rest = presets -->
        <div class="p-5">
          <div class="grid grid-cols-3 gap-3">

            <!-- Upload Cell -->
            <div
              @click="triggerFileSelect"
              class="aspect-square rounded-full border-2 border-dashed cursor-pointer transition-all flex flex-col items-center justify-center select-none relative overflow-hidden"
              :class="localUploadedUrl ? 'border-coffee-espresso' : 'border-coffee-latte/50 hover:border-coffee-brown bg-coffee-cream/30 hover:bg-coffee-cream/60'"
            >
              <template v-if="isUploadingLocal">
                <div class="w-4 h-4 border-2 border-coffee-espresso border-t-transparent rounded-full animate-spin"></div>
              </template>
              <template v-else-if="localUploadedUrl">
                <img :src="localUploadedUrl" class="w-full h-full object-cover absolute inset-0 rounded-full">
                <div class="absolute inset-0 bg-coffee-espresso/20 flex items-center justify-center rounded-full">
                  <Check class="w-4 h-4 text-white" />
                </div>
              </template>
              <template v-else>
                <Plus class="w-5 h-5 text-coffee-softGray" />
                <span class="text-[8px] text-coffee-softGray mt-1 text-center leading-tight px-1">本地上传</span>
              </template>
            </div>
            <!-- Hidden File Input -->
            <input type="file" ref="fileInput" accept="image/*" @change="handleFileChange" class="hidden">

            <!-- Preset Avatar Cells -->
            <div
              v-for="(url, idx) in presetAvatars"
              :key="idx"
              @click="selectPresetAvatar(url)"
              class="aspect-square rounded-full overflow-hidden border-2 cursor-pointer transition-all relative"
              :class="selectedAvatarUrl === url && !localUploadedUrl
                ? 'border-coffee-espresso scale-105 shadow-sm'
                : 'border-transparent opacity-70 hover:opacity-100 hover:scale-[1.04]'"
            >
              <img :src="url" class="w-full h-full object-cover">
              <div v-if="selectedAvatarUrl === url && !localUploadedUrl" class="absolute inset-0 bg-coffee-espresso/25 flex items-center justify-center">
                <Check class="w-3.5 h-3.5 text-white" />
              </div>
            </div>

          </div>
        </div>

        <!-- Save Button -->
        <div class="px-5 pb-5">
          <button
            @click="saveAvatar"
            :disabled="isSavingAvatar"
            class="w-full py-3 bg-coffee-espresso text-coffee-warmWhite hover:bg-coffee-brown disabled:opacity-60 transition-colors rounded-xl text-xs font-semibold tracking-widest uppercase flex items-center justify-center gap-1.5"
          >
            <template v-if="isSavingAvatar">
              <div class="w-3.5 h-3.5 border-2 border-white border-t-transparent rounded-full animate-spin"></div>
              <span>正在保存...</span>
            </template>
            <template v-else>
              <span>保存头像配置</span>
            </template>
          </button>
        </div>

      </div>
    </div>

  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { useCoffeeLogStore } from '@/stores/coffeeLog'
import request from '@/api/request'
import { getLocalDateString } from '@/utils/date'
import { BookOpen, Calendar, BarChart3, Plus, ChevronRight, Edit2, Check, X, User } from 'lucide-vue-next'

const router = useRouter()
const authStore = useAuthStore()
const logsStore = useCoffeeLogStore()

// State
const isEditing = ref(false)
const editNickname = ref('')
const pushEnabled = ref(true)

const showAvatarModal = ref(false)
const isSavingAvatar = ref(false)
const selectedAvatarUrl = ref('')

// File upload states
const fileInput = ref<HTMLInputElement | null>(null)
const isUploadingLocal = ref(false)
const localUploadedUrl = ref('')

// Pre-curated premium high-res coffee life photo presets
const presetAvatars = [
  'https://images.unsplash.com/photo-1534528741775-53994a69daeb?auto=format&fit=crop&q=80&w=150', // Original profile girl
  'https://images.unsplash.com/photo-1507003211169-0a1dd7228f2d?auto=format&fit=crop&q=80&w=150', // Boy profile
  'https://images.unsplash.com/photo-1492562080023-ab3db95bfbce?auto=format&fit=crop&q=80&w=150', // Casual smart boy
  'https://images.unsplash.com/photo-1494790108377-be9c29b29330?auto=format&fit=crop&q=80&w=150', // Smiling girl
  'https://images.unsplash.com/photo-1442512595331-e89e73853f31?auto=format&fit=crop&q=80&w=150', // Pure espresso dripper
  'https://images.unsplash.com/photo-1514432324607-a09d9b4aefdd?auto=format&fit=crop&q=80&w=150', // Cozy cup book
  'https://images.unsplash.com/photo-1495474472287-4d71bcdd2085?auto=format&fit=crop&q=80&w=150', // Warm latte art
  'https://images.unsplash.com/photo-1509042239860-f550ce710b93?auto=format&fit=crop&q=80&w=150'  // Coffee roasting beans
]

// Computed
const userAvatar = computed(() => authStore.user?.avatar_url || 'https://images.unsplash.com/photo-1534528741775-53994a69daeb?auto=format&fit=crop&q=80&w=120')

const daysSinceJoined = computed(() => {
  if (!authStore.user) return 1
  const joined = new Date(1779344666 * 1000) // Default startup timestamp
  const now = new Date()
  const diffTime = Math.abs(now.getTime() - joined.getTime())
  const diffDays = Math.ceil(diffTime / (1000 * 60 * 60 * 24))
  return diffDays || 1
})

const firstBrewDate = computed(() => {
  if (logsStore.logs.length === 0) return '尚未冲煮第一杯'
  const dates = logsStore.logs.map(l => new Date(l.drink_date).getTime())
  const earliest = new Date(Math.min(...dates))
  return earliest.toLocaleDateString('zh-CN', { year: 'numeric', month: 'long', day: 'numeric' })
})

onMounted(async () => {
  if (logsStore.logs.length === 0) {
    try {
      await logsStore.fetchLogs({ page: 1, page_size: 50 })
    } catch (e) {
      console.error('Failed to fetch logs:', e)
    }
  }
})

// Methods
const handleLogout = () => {
  if (confirm('确认退出登录吗？退出后本地手账历史将会清除。')) {
    authStore.logout()
    router.push('/login')
  }
}

const startEdit = () => {
  editNickname.value = authStore.user?.nickname || ''
  isEditing.value = true
}

const saveNickname = async () => {
  const name = editNickname.value.trim()
  if (!name) return
  try {
    await authStore.updateNickname(name)
    isEditing.value = false
  } catch (e: any) {
    alert(e.message || '更新失败')
  }
}

const togglePush = () => {
  pushEnabled.value = !pushEnabled.value
}

const exportBackup = () => {
  const dataStr = "data:text/json;charset=utf-8," + encodeURIComponent(JSON.stringify(logsStore.logs, null, 2))
  const downloadAnchor = document.createElement('a')
  downloadAnchor.setAttribute("href",     dataStr)
  downloadAnchor.setAttribute("download", `mcl_backup_${getLocalDateString()}.json`)
  document.body.appendChild(downloadAnchor)
  downloadAnchor.click()
  downloadAnchor.remove()
}

const showAbout = () => {
  alert('My Coffee Log - 专为咖啡爱好者设计的极简感官手账。\n致力于通过北欧风极简视觉和智能感官算法，为您留存每一次与美味的温热邂逅。')
}

// Avatar Modal Actions
const closeAvatarModal = () => {
  showAvatarModal.value = false
  selectedAvatarUrl.value = ''
  localUploadedUrl.value = ''
}

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
  // 5MB limit validation on frontend
  if (file.size > 5 * 1024 * 1024) {
    alert('照片大小不能超过 5MB 喔！')
    return
  }

  isUploadingLocal.value = true
  const formData = new FormData()
  formData.append('file', file)

  try {
    const res = await request.post('/upload', formData, {
      headers: {
        'Content-Type': 'multipart/form-data'
      }
    }) as any

    // The backend returns { code: 0, message: 'success', data: { url: '/uploads/xxx.png' } }
    // Our request.ts response interceptor auto-unwraps this, returning res = { url: '/uploads/xxx.png' }
    if (res && res.url) {
      localUploadedUrl.value = res.url
      selectedAvatarUrl.value = res.url
    } else {
      throw new Error('未获取到返回的图片地址')
    }
  } catch (e: any) {
    alert(e.message || '相片上传失败，请重试')
  } finally {
    isUploadingLocal.value = false
    // Clear value to allow same file selection
    target.value = ''
  }
}

const selectPresetAvatar = (url: string) => {
  selectedAvatarUrl.value = url
  localUploadedUrl.value = '' // Clear local upload preview
}

const saveAvatar = async () => {
  const url = selectedAvatarUrl.value.trim()
  if (!url) {
    alert('请选择一款预设头像或上传本地相片')
    return
  }
  isSavingAvatar.value = true
  try {
    await authStore.updateAvatar(url)
    showAvatarModal.value = false
    closeAvatarModal()
  } catch (e: any) {
    alert(e.message || '头像更新失败')
  } finally {
    isSavingAvatar.value = false
  }
}
</script>

<style scoped>
.editorial-border {
  border: 1px solid rgba(92, 61, 46, 0.2);
  border-radius: 16px;
}
</style>
