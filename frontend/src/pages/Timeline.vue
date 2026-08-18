<template>
  <div class="flex-1 w-full flex flex-col justify-between bg-coffee-warmWhite text-coffee-charcoal">
    
    <!-- Timeline Header & Multi-Filters -->
    <div class="px-6 py-5 flex flex-col gap-4 border-b border-coffee-cream/80 bg-coffee-warmWhite/80 backdrop-blur-md sticky top-0 z-20 select-none">
      <div class="flex justify-between items-end">
        <div>
          <span class="text-[9px] uppercase tracking-[0.25em] font-semibold text-coffee-softGray">每日风味典藏</span>
          <h2 class="font-serif text-2xl font-light text-coffee-espresso leading-none mt-1">咖啡时间线</h2>
        </div>
        <img src="https://images.unsplash.com/photo-1534528741775-53994a69daeb?auto=format&fit=crop&q=80&w=120" alt="Avatar" class="w-8 h-8 rounded-full object-cover border border-coffee-latte">
      </div>

      <!-- Filter Pills Bar (Asymmetric, Editorial Scroll) -->
      <div class="flex gap-2 overflow-x-auto pb-1 -mx-2 px-2 scrollbar-none text-[10px] font-medium tracking-wide">
        <button 
          v-for="f in filterTypes" 
          :key="f.val"
          @click="activeFilter = f.val"
          class="px-3 py-1 rounded-full cursor-pointer transition-all flex-shrink-0"
          :class="activeFilter === f.val 
            ? 'bg-coffee-espresso text-coffee-warmWhite' 
            : 'bg-coffee-cream text-coffee-espresso hover:bg-coffee-latte/40'"
        >
          {{ f.label }}
        </button>
      </div>
    </div>

    <!-- Scrollable Timeline Items (Asymmetric Layout Grid) -->
    <div class="flex-1 overflow-y-auto px-6 py-5 space-y-8 pb-24 scrollbar-none">
      
      <!-- Grouped Months Loop -->
      <div v-for="group in groupedLogs" :key="group.month" class="space-y-6">
        
        <!-- Month Header -->
        <div class="flex items-center gap-3 select-none">
          <span class="font-serif text-sm uppercase tracking-widest text-coffee-espresso font-semibold">{{ group.month }}</span>
          <div class="h-[1px] bg-coffee-cream flex-1"></div>
          <span class="text-[10px] text-coffee-softGray font-mono font-semibold">{{ group.items.length }} 记</span>
        </div>

        <!-- Asymmetric Feed Loop -->
        <div class="space-y-6">
          <div v-for="(log, idx) in group.items" :key="log.id">
            
            <!-- Type A: Portrait Photo left side (Alternate 1) -->
            <div v-if="isModZero(idx, 3)" class="group flex gap-5 items-stretch relative">
            <router-link 
              :to="'/coffee/' + log.id" 
              class="flex gap-5 items-stretch flex-1"
            >
              <div class="w-2/5 aspect-[3/4] overflow-hidden rounded-xl bg-neutral-100 border border-coffee-cream flex-shrink-0 select-none">
                <img :src="log.image_url" class="w-full h-full object-cover filter saturate-50 group-hover:scale-105 transition-transform duration-700">
              </div>
              <div class="flex-1 flex flex-col justify-between py-1">
                <div class="space-y-1">
                  <div class="flex justify-between items-baseline text-[8px] font-mono text-coffee-softGray select-none">
                    <span>{{ getDayNum(log.drink_date) }} {{ getMonthAbbr(log.drink_date) }}</span>
                    <span class="uppercase">{{ coffeeTypeLabel(log.coffee_type) }}</span>
                  </div>
                  <h4 class="font-serif text-lg font-light text-coffee-espresso group-hover:text-coffee-brown italic leading-tight transition-colors break-words">{{ log.coffee_name }}</h4>
                  <p class="text-[10px] text-coffee-brown leading-relaxed font-light line-clamp-3 font-serif">
                    “{{ log.notes }}”
                  </p>
                </div>
                <div class="flex justify-between items-center text-[9px] text-coffee-softGray select-none">
                  <span class="truncate max-w-[120px]">在 {{ log.shop_name.split(',')[0] }}</span>
                  <span class="inline-flex items-center gap-1 font-mono">
                    <AppIcon :name="moodIconName(log.mood)" :size="11" />
                    {{ moodLabel(log.mood) }}
                  </span>
                </div>
              </div>
              </router-link>
              <router-link :to="'/create?from_log_id=' + log.id" class="absolute bottom-1 right-0 opacity-0 group-hover:opacity-100 transition-opacity p-1.5 rounded-full bg-coffee-cream/80 hover:bg-coffee-latte text-coffee-espresso" title="复刻这杯">
                <Copy class="w-3 h-3" />
              </router-link>
            </div>

            <!-- Type B: Text Only Card (Alternate 2) -->
            <div v-else-if="isModOne(idx, 3)" class="group relative">
            <router-link 
              :to="'/coffee/' + log.id" 
              class="block p-5 bg-coffee-cream/30 rounded-2xl border border-coffee-cream/70 hover:bg-coffee-cream/45 transition-colors"
            >
              <div class="space-y-2">
                <div class="flex justify-between items-center text-[8px] font-mono text-coffee-softGray select-none">
                  <span>{{ getDayNum(log.drink_date) }} {{ getMonthAbbr(log.drink_date) }}</span>
                  <span class="uppercase">{{ coffeeTypeLabel(log.coffee_type) }}</span>
                </div>
                <h4 class="font-serif text-base font-light text-coffee-espresso italic group-hover:text-coffee-brown transition-colors break-words">{{ log.coffee_name }}</h4>
                <p class="font-serif italic text-xs text-coffee-brown leading-relaxed line-clamp-3">
                  “{{ log.notes }}”
                </p>
                <div class="flex justify-between items-center text-[9px] text-coffee-softGray pt-1.5 border-t border-coffee-cream/40 select-none">
                  <span class="truncate max-w-[150px]">在 {{ log.shop_name.split(',')[0] }}</span>
                  <span class="inline-flex items-center gap-1 font-mono uppercase">
                    <AppIcon :name="moodIconName(log.mood)" :size="11" />
                    {{ moodLabel(log.mood) }}
                  </span>
                </div>
              </div>
            </router-link>
              <router-link :to="'/create?from_log_id=' + log.id" class="absolute top-3 right-3 opacity-0 group-hover:opacity-100 transition-opacity p-1.5 rounded-full bg-coffee-cream/80 hover:bg-coffee-latte text-coffee-espresso" title="复刻这杯">
                <Copy class="w-3 h-3" />
              </router-link>
            </div>

            <!-- Type C: Wide card with Landscape photo (Alternate 3) -->
            <div v-else class="group space-y-2 relative">
            <router-link 
              :to="'/coffee/' + log.id" 
              class="block space-y-2"
            >
              <div class="w-full h-40 overflow-hidden rounded-xl bg-neutral-100 border border-coffee-cream select-none">
                <img :src="log.image_url" class="w-full h-full object-cover filter saturate-50 group-hover:scale-[1.02] transition-transform duration-700">
              </div>
              <div class="space-y-1">
                <div class="flex justify-between items-center text-[8px] font-mono text-coffee-softGray select-none">
                  <span>{{ getDayNum(log.drink_date) }} {{ getMonthAbbr(log.drink_date) }}</span>
                  <span class="uppercase">{{ coffeeTypeLabel(log.coffee_type) }}</span>
                </div>
                <div class="flex justify-between items-baseline">
                  <h4 class="font-serif text-base font-light text-coffee-espresso group-hover:text-coffee-brown italic leading-none transition-colors break-words">{{ log.coffee_name }}</h4>
                  <span class="text-[9px] text-coffee-softGray font-serif select-none truncate max-w-[160px]">{{ log.shop_name.split(',')[0] }}</span>
                </div>
                <p class="text-[10px] text-coffee-brown font-light truncate leading-relaxed">
                  “{{ log.notes }}”
                </p>
              </div>
            </router-link>
              <router-link :to="'/create?from_log_id=' + log.id" class="absolute top-2 right-2 opacity-0 group-hover:opacity-100 transition-opacity p-1.5 rounded-full bg-coffee-cream/80 hover:bg-coffee-latte text-coffee-espresso" title="复刻这杯">
                <Copy class="w-3 h-3" />
              </router-link>
            </div>

          </div>
        </div>

      </div>

      <!-- Empty state inside Timeline -->
      <div v-if="groupedLogs.length === 0" class="text-center py-16 text-coffee-softGray space-y-3 select-none">
        <AppIcon name="coffee" :size="30" :stroke-width="1.35" class="mx-auto" />
        <p class="text-xs uppercase tracking-widest">No matching logs found</p>
      </div>

    </div>

    <!-- Sticky Bottom Navigation Bar -->
    <div class="relative h-16 border-t border-coffee-cream/60 bg-coffee-warmWhite flex items-center z-30 select-none">
      <router-link to="/home" class="flex-1 flex flex-col items-center gap-0.5 text-coffee-softGray hover:text-coffee-brown transition-colors">
        <BookOpen class="w-5 h-5" />
        <span class="text-[9.5px] tracking-widest font-medium">咖啡日志</span>
      </router-link>
      <router-link to="/timeline" class="flex-1 flex flex-col items-center gap-0.5 text-coffee-brown">
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
import { ref, computed, onMounted } from 'vue'
import { useCoffeeLogStore, CoffeeLog } from '@/stores/coffeeLog'
import { coffeeTypeLabel, moodLabel, moodIconName } from '@/constants/coffee'
import AppIcon from '@/components/AppIcon.vue'
import { BookOpen, Calendar, BarChart3, Plus, User, Copy } from 'lucide-vue-next'

const store = useCoffeeLogStore()

onMounted(async () => {
  if (store.logs.length === 0) {
    await store.fetchLogs({ page: 1, page_size: 50 })
  }
})
const activeFilter = ref('all')

const filterTypes = [
  { val: 'all', label: 'All Brews / 全部' },
  { val: 'Pour Over', label: 'Pour Over / 手冲' },
  { val: 'Latte', label: 'Latte / 拿铁' },
  { val: 'Americano', label: 'Americano / 美式' },
  { val: 'Cold Brew', label: 'Cold Brew / 冷萃' },
  { val: 'Espresso', label: 'Espresso / 浓缩' },
  { val: 'Dirty', label: 'Dirty / 脏咖啡' },
  { val: 'Cappuccino', label: 'Cappuccino / 卡布奇诺' },
  { val: 'Flat White', label: 'Flat White / 馥芮白' }
]

// Grouping logic inside computed (Groups logs by Year-Month)
interface GroupedMonth {
  month: string
  items: CoffeeLog[]
}

const groupedLogs = computed(() => {
  const filtered = activeFilter.value === 'all' 
    ? store.logs 
    : store.logs.filter((log: CoffeeLog) => log.coffee_type === activeFilter.value)
    
  const groups: Record<string, CoffeeLog[]> = {}
  
  filtered.forEach((log: CoffeeLog) => {
    const date = new Date(log.drink_date)
    const months = ["January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"]
    const monthStr = `${months[date.getMonth()]} ${date.getFullYear()}`
    
    if (!groups[monthStr]) {
      groups[monthStr] = []
    }
    groups[monthStr].push(log)
  })
  
  return Object.entries(groups).map(([month, items]) => ({
    month,
    items
  })) as GroupedMonth[]
})

// Modulo helpers to bypass template type-inference limitations
const isModZero = (idx: any, mod: number) => Number(idx) % mod === 0
const isModOne = (idx: any, mod: number) => Number(idx) % mod === 1

// Date helpers
const getDayNum = (dateStr: string) => {
  return new Date(dateStr).getDate()
}

const getMonthAbbr = (dateStr: string) => {
  const date = new Date(dateStr)
  const months = ["Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"]
  return months[date.getMonth()].toUpperCase()
}
</script>

<style scoped>
/* Timeline specific scoped styles */
</style>
