<template>
  <div class="flex-1 w-full flex flex-col justify-between bg-coffee-warmWhite text-coffee-charcoal">
    
    <!-- Header -->
    <div class="px-6 py-5 flex justify-between items-end border-b border-coffee-cream bg-coffee-warmWhite sticky top-0 z-10 select-none">
      <div>
        <span class="text-[9px] uppercase tracking-[0.25em] font-semibold text-coffee-softGray">My Collection</span>
        <h2 class="font-serif text-2xl font-light text-coffee-espresso leading-none mt-1">咖啡店收藏</h2>
      </div>
      <button @click="showAddModal = true" class="w-8 h-8 rounded-full bg-coffee-espresso text-coffee-warmWhite flex items-center justify-center hover:bg-coffee-brown transition-colors">
        <Plus class="w-4 h-4" />
      </button>
    </div>

    <!-- Search Bar -->
    <div class="px-6 py-3 border-b border-coffee-cream/40 bg-coffee-warmWhite sticky top-[68px] z-10">
      <div class="relative">
        <Search class="w-3.5 h-3.5 absolute left-3 top-1/2 -translate-y-1/2 text-coffee-softGray" />
        <input 
          type="text" 
          v-model="searchQuery" 
          placeholder="搜索咖啡店..." 
          class="w-full pl-9 pr-3 py-2 bg-coffee-cream/40 border border-coffee-latte/50 focus:border-coffee-brown focus:outline-none rounded-sm text-xs transition-colors"
        >
      </div>
    </div>

    <!-- Scrollable Body -->
    <div class="flex-1 overflow-y-auto px-6 py-4 space-y-3 pb-24 scrollbar-none">

      <!-- Empty State -->
      <div v-if="shops.length === 0 && !isLoading" class="flex flex-col items-center justify-center py-16 space-y-4 select-none">
        <div class="text-5xl opacity-30">🏪</div>
        <p class="text-sm text-coffee-softGray font-light italic text-center leading-relaxed">
          还没有收藏的咖啡店。<br>记录咖啡时填写店铺名称，会自动收藏。
        </p>
        <button @click="showAddModal = true" class="px-5 py-2 text-[10px] uppercase tracking-widest font-semibold bg-coffee-espresso text-coffee-warmWhite rounded-sm hover:bg-coffee-brown transition-colors">
          手动添加
        </button>
      </div>

      <!-- Shop Cards -->
      <router-link 
        v-for="shop in shops" 
        :key="shop.id" 
        :to="`/coffee-shops/${shop.id}`"
        class="block p-4 rounded-sm border border-coffee-latte/30 hover:border-coffee-brown/40 transition-colors select-none"
        style="background: linear-gradient(135deg, rgba(215,196,168,0.08) 0%, rgba(231,111,81,0.02) 100%);"
      >
        <div class="flex items-start gap-3">
          <!-- Shop Image or Placeholder -->
          <div class="w-12 h-12 rounded-sm overflow-hidden flex-shrink-0 bg-coffee-cream/50 border border-coffee-latte/30">
            <img v-if="shop.image_url" :src="shop.image_url" class="w-full h-full object-cover">
            <div v-else class="w-full h-full flex items-center justify-center text-lg">☕</div>
          </div>
          <div class="flex-1 min-w-0">
            <div class="flex items-center gap-2">
              <h4 class="font-serif text-sm font-semibold text-coffee-espresso leading-tight truncate">{{ shop.name }}</h4>
              <div v-if="shop.rating > 0" class="flex items-center gap-0.5 flex-shrink-0">
                <Star v-for="i in shop.rating" :key="i" class="w-2.5 h-2.5 text-amber-500 fill-amber-500" />
              </div>
            </div>
            <div class="flex items-center gap-3 mt-1">
              <span v-if="shop.address" class="text-[10px] text-coffee-softGray truncate">{{ shop.address }}</span>
              <span class="text-[10px] text-coffee-brown font-semibold flex-shrink-0">{{ shop.visit_count }} 次到访</span>
            </div>
            <div v-if="shop.last_visit_at" class="text-[9px] text-coffee-softGray mt-0.5">
              最近到访: {{ shop.last_visit_at?.split('T')[0] }}
            </div>
          </div>
          <ChevronRight class="w-4 h-4 text-coffee-softGray flex-shrink-0 mt-1" />
        </div>
      </router-link>

      <!-- Load More -->
      <div v-if="hasMore" class="text-center py-3">
        <button @click="loadMore" class="text-[10px] uppercase tracking-widest font-semibold text-coffee-brown hover:text-coffee-espresso transition-colors">
          加载更多
        </button>
      </div>
    </div>

    <!-- Add Shop Modal -->
    <div v-if="showAddModal" class="fixed inset-0 bg-black/40 z-50 flex items-end justify-center" @click.self="showAddModal = false">
      <div class="w-full max-w-md bg-coffee-warmWhite rounded-t-2xl p-6 space-y-4 animate-slide-up">
        <div class="flex justify-between items-center">
          <h3 class="font-serif text-lg font-semibold text-coffee-espresso">添加咖啡店</h3>
          <button @click="showAddModal = false" class="w-7 h-7 rounded-full bg-coffee-cream flex items-center justify-center">
            <X class="w-3.5 h-3.5 text-coffee-brown" />
          </button>
        </div>
        <div class="space-y-3">
          <input v-model="newShop.name" type="text" placeholder="店铺名称 *" class="w-full p-3 bg-coffee-cream/40 border border-coffee-latte/60 focus:border-coffee-brown focus:outline-none rounded-sm text-sm">
          <input v-model="newShop.address" type="text" placeholder="地址 (选填)" class="w-full p-3 bg-coffee-cream/40 border border-coffee-latte/60 focus:border-coffee-brown focus:outline-none rounded-sm text-sm">
          <div class="flex items-center gap-2">
            <span class="text-[10px] text-coffee-softGray">评分</span>
            <button v-for="i in 5" :key="i" @click="newShop.rating = i" class="transition-transform" :class="newShop.rating >= i ? 'scale-110' : ''">
              <Star class="w-5 h-5 transition-colors" :class="newShop.rating >= i ? 'text-amber-500 fill-amber-500' : 'text-coffee-cream'" />
            </button>
          </div>
        </div>
        <button @click="addShop" :disabled="!newShop.name.trim() || isAdding" class="w-full py-3 text-[10px] uppercase tracking-widest font-semibold bg-coffee-espresso text-coffee-warmWhite hover:bg-coffee-brown transition-all rounded-sm disabled:opacity-50">
          {{ isAdding ? '添加中...' : '添加收藏' }}
        </button>
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
import { ref, reactive, computed, onMounted, watch } from 'vue'
import * as shopApi from '@/api/coffeeShop'
import type { CoffeeShop } from '@/api/coffeeShop'
import { BookOpen, Calendar, BarChart3, Plus, User, ChevronRight, X, Search, Star } from 'lucide-vue-next'

const shops = ref<CoffeeShop[]>([])
const isLoading = ref(false)
const searchQuery = ref('')
const currentPage = ref(1)
const totalShops = ref(0)
const pageSize = 20

const hasMore = computed(() => shops.value.length < totalShops.value)

// Add modal
const showAddModal = ref(false)
const isAdding = ref(false)
const newShop = reactive({ name: '', address: '', rating: 0 })

async function fetchShops() {
  isLoading.value = true
  try {
    const res = await shopApi.getCoffeeShops({
      page: currentPage.value,
      page_size: pageSize,
      search: searchQuery.value || undefined
    })
    if (currentPage.value === 1) {
      shops.value = res.list
    } else {
      shops.value.push(...res.list)
    }
    totalShops.value = res.pagination.total
  } catch (e) {
    console.error('Failed to fetch shops:', e)
  } finally {
    isLoading.value = false
  }
}

function loadMore() {
  currentPage.value++
  fetchShops()
}

async function addShop() {
  if (!newShop.name.trim()) return
  isAdding.value = true
  try {
    await shopApi.createCoffeeShop({
      name: newShop.name.trim(),
      address: newShop.address.trim() || undefined,
      rating: newShop.rating || undefined
    })
    showAddModal.value = false
    newShop.name = ''
    newShop.address = ''
    newShop.rating = 0
    currentPage.value = 1
    fetchShops()
  } catch (e: any) {
    alert(e.message || '添加失败')
  } finally {
    isAdding.value = false
  }
}

watch(searchQuery, () => {
  currentPage.value = 1
  fetchShops()
})

onMounted(fetchShops)
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
