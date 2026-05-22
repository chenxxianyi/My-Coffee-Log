<template>
  <div class="flex-1 w-full flex flex-col justify-between bg-coffee-warmWhite text-coffee-charcoal">
    
    <!-- Header -->
    <div class="px-6 py-4 flex justify-between items-center border-b border-coffee-cream bg-coffee-warmWhite sticky top-0 z-10 select-none">
      <button @click="router.back()" class="text-coffee-brown hover:text-coffee-espresso">
        <ChevronLeft class="w-5 h-5" />
      </button>
      <span class="text-[10px] uppercase tracking-[0.2em] font-semibold text-coffee-espresso">Coffee Shop</span>
      <button @click="showEditModal = true" class="text-coffee-brown hover:text-coffee-espresso">
        <Pencil class="w-4 h-4" />
      </button>
    </div>

    <!-- Scrollable Body -->
    <div class="flex-1 overflow-y-auto pb-24 scrollbar-none">

      <!-- Shop Hero -->
      <div v-if="shop" class="relative">
        <!-- Cover Image -->
        <div class="h-40 bg-coffee-cream/40 overflow-hidden">
          <img v-if="shop.image_url" :src="shop.image_url" class="w-full h-full object-cover">
          <div v-else class="w-full h-full flex items-center justify-center text-6xl opacity-20">🏪</div>
        </div>
        <!-- Shop Info Overlay -->
        <div class="px-6 -mt-8 relative z-10">
          <div class="p-5 rounded-sm border border-coffee-latte/40 bg-coffee-warmWhite shadow-sm">
            <div class="flex items-start justify-between">
              <div>
                <h2 class="font-serif text-xl font-semibold text-coffee-espresso leading-tight">{{ shop.name }}</h2>
                <div v-if="shop.address" class="flex items-center gap-1.5 mt-1.5">
                  <MapPin class="w-3 h-3 text-coffee-softGray" />
                  <span class="text-[11px] text-coffee-softGray">{{ shop.address }}</span>
                </div>
              </div>
              <div v-if="shop.rating > 0" class="flex items-center gap-0.5">
                <Star v-for="i in shop.rating" :key="i" class="w-3.5 h-3.5 text-amber-500 fill-amber-500" />
              </div>
            </div>
            <!-- Stats Row -->
            <div class="flex items-center gap-5 mt-4 pt-3 border-t border-coffee-cream">
              <div class="text-center">
                <div class="text-xl font-serif text-coffee-espresso font-light">{{ shop.visit_count }}</div>
                <div class="text-[9px] uppercase tracking-wider text-coffee-softGray">到访次数</div>
              </div>
              <div class="text-center" v-if="shop.last_visit_at">
                <div class="text-sm font-serif text-coffee-espresso font-light">{{ shop.last_visit_at?.split('T')[0] }}</div>
                <div class="text-[9px] uppercase tracking-wider text-coffee-softGray">最近到访</div>
              </div>
              <div class="text-center">
                <div class="text-sm font-serif text-coffee-espresso font-light">{{ relatedLogs.length }}</div>
                <div class="text-[9px] uppercase tracking-wider text-coffee-softGray">咖啡记录</div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Loading State -->
      <div v-if="isLoading" class="flex items-center justify-center py-16">
        <div class="w-6 h-6 border-2 border-coffee-espresso border-t-transparent rounded-full animate-spin"></div>
      </div>

      <!-- Related Coffee Logs -->
      <div v-if="!isLoading && shop" class="px-6 mt-6 space-y-3">
        <div class="flex items-center gap-2.5 select-none">
          <div class="h-px flex-1 max-w-[24px]" style="background: linear-gradient(to right, rgba(92,61,46,0.35), transparent);"></div>
          <h3 class="text-[10px] uppercase tracking-[0.2em] font-bold text-coffee-espresso">这里的咖啡记录 / Coffee Logs</h3>
          <div class="h-px flex-1" style="background: linear-gradient(to left, rgba(92,61,46,0.35), transparent);"></div>
        </div>

        <div v-if="relatedLogs.length === 0" class="text-center py-8">
          <p class="text-[11px] text-coffee-softGray italic">暂无关联咖啡记录</p>
        </div>

        <router-link 
          v-for="log in relatedLogs" 
          :key="log.id"
          :to="`/coffee/${log.id}`"
          class="flex items-center gap-3 p-3 rounded-sm border border-coffee-latte/25 hover:border-coffee-brown/30 transition-colors"
        >
          <div class="w-10 h-10 rounded-sm overflow-hidden flex-shrink-0 bg-coffee-cream/40">
            <img v-if="log.image_url" :src="log.image_url" class="w-full h-full object-cover">
            <div v-else class="w-full h-full flex items-center justify-center text-sm">☕</div>
          </div>
          <div class="flex-1 min-w-0">
            <div class="text-xs font-serif font-medium text-coffee-espresso truncate">{{ log.coffee_name }}</div>
            <div class="text-[10px] text-coffee-softGray">{{ log.coffee_type }} · {{ log.drink_date }}</div>
          </div>
          <ChevronRight class="w-3.5 h-3.5 text-coffee-softGray flex-shrink-0" />
        </router-link>
      </div>
    </div>

    <!-- Edit Shop Modal -->
    <div v-if="showEditModal && shop" class="fixed inset-0 bg-black/40 z-50 flex items-end justify-center" @click.self="showEditModal = false">
      <div class="w-full max-w-md bg-coffee-warmWhite rounded-t-2xl p-6 space-y-4 animate-slide-up">
        <div class="flex justify-between items-center">
          <h3 class="font-serif text-lg font-semibold text-coffee-espresso">编辑咖啡店</h3>
          <button @click="showEditModal = false" class="w-7 h-7 rounded-full bg-coffee-cream flex items-center justify-center">
            <X class="w-3.5 h-3.5 text-coffee-brown" />
          </button>
        </div>
        <div class="space-y-3">
          <input v-model="editForm.name" type="text" placeholder="店铺名称" class="w-full p-3 bg-coffee-cream/40 border border-coffee-latte/60 focus:border-coffee-brown focus:outline-none rounded-sm text-sm">
          <input v-model="editForm.address" type="text" placeholder="地址" class="w-full p-3 bg-coffee-cream/40 border border-coffee-latte/60 focus:border-coffee-brown focus:outline-none rounded-sm text-sm">
          <div class="flex items-center gap-2">
            <span class="text-[10px] text-coffee-softGray">评分</span>
            <button v-for="i in 5" :key="i" @click="editForm.rating = i" class="transition-transform" :class="editForm.rating >= i ? 'scale-110' : ''">
              <Star class="w-5 h-5 transition-colors" :class="editForm.rating >= i ? 'text-amber-500 fill-amber-500' : 'text-coffee-cream'" />
            </button>
          </div>
        </div>
        <div class="flex gap-3">
          <button @click="deleteShop" class="py-3 px-4 text-[10px] uppercase tracking-widest font-semibold text-red-600 border border-red-200 hover:bg-red-50 transition-all rounded-sm">
            删除
          </button>
          <button @click="saveEdit" :disabled="isSaving" class="flex-1 py-3 text-[10px] uppercase tracking-widest font-semibold bg-coffee-espresso text-coffee-warmWhite hover:bg-coffee-brown transition-all rounded-sm disabled:opacity-50">
            {{ isSaving ? '保存中...' : '保存' }}
          </button>
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
      <router-link to="/profile" class="flex-1 flex flex-col items-center gap-0.5 text-coffee-softGray hover:text-coffee-brown transition-colors">
        <User class="w-5 h-5" />
        <span class="text-[9.5px] tracking-widest font-medium">我的</span>
      </router-link>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import * as shopApi from '@/api/coffeeShop'
import type { CoffeeShop } from '@/api/coffeeShop'
import type { CoffeeLog } from '@/stores/coffeeLog'
import { BookOpen, Calendar, BarChart3, Plus, User, ChevronLeft, ChevronRight, X, Pencil, MapPin, Star } from 'lucide-vue-next'

const router = useRouter()
const route = useRoute()

const shop = ref<CoffeeShop | null>(null)
const relatedLogs = ref<CoffeeLog[]>([])
const isLoading = ref(true)

// Edit modal
const showEditModal = ref(false)
const isSaving = ref(false)
const editForm = reactive({ name: '', address: '', rating: 0 })

async function loadData() {
  isLoading.value = true
  const id = Number(route.params.id)
  try {
    const [shopData, logsRes] = await Promise.all([
      shopApi.getCoffeeShopById(id),
      shopApi.getShopLogs(id, { page: 1, page_size: 50 })
    ])
    shop.value = shopData
    relatedLogs.value = logsRes.list as any
    editForm.name = shopData.name
    editForm.address = shopData.address
    editForm.rating = shopData.rating
  } catch (e) {
    console.error('Failed to load shop:', e)
  } finally {
    isLoading.value = false
  }
}

async function saveEdit() {
  if (!shop.value) return
  isSaving.value = true
  try {
    const updated = await shopApi.updateCoffeeShop(shop.value.id, {
      name: editForm.name.trim() || undefined,
      address: editForm.address.trim() || undefined,
      rating: editForm.rating || undefined
    })
    shop.value = updated
    showEditModal.value = false
  } catch (e: any) {
    alert(e.message || '保存失败')
  } finally {
    isSaving.value = false
  }
}

async function deleteShop() {
  if (!shop.value) return
  if (!confirm('确定要删除这个咖啡店吗？')) return
  try {
    await shopApi.deleteCoffeeShop(shop.value.id)
    router.replace('/coffee-shops')
  } catch (e: any) {
    alert(e.message || '删除失败')
  }
}

onMounted(loadData)
</script>

<style scoped>
.animate-slide-up {
  animation: slideUp 0.25s ease-out;
}
@keyframes slideUp {
  from { transform: translateY(100%); }
  to { transform: translateY(0); }
}
</style>
