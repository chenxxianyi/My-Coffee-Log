<template>
  <div class="flex-1 w-full flex flex-col justify-between px-8 py-12 bg-coffee-warmWhite text-coffee-charcoal">
    
    <!-- Top Branding Info -->
    <div class="flex justify-between items-center text-[10px] tracking-[0.25em] text-coffee-softGray font-bold mt-4 select-none">
      <span>每日风味志</span>
      <span>EST. 2026</span>
    </div>

    <!-- Center Logo & Motto -->
    <div class="flex flex-col items-center text-center my-auto py-12">
      <div class="relative mb-8 select-none">
        <span class="font-serif text-[130px] leading-none font-light italic text-coffee-latte tracking-tighter block">mcl</span>
        <span class="absolute top-[75px] left-1/2 -translate-x-1/2 font-serif text-[44px] uppercase font-light text-coffee-espresso tracking-[0.12em] leading-none block whitespace-nowrap">Coffee Log</span>
      </div>
      
      <div class="w-12 h-[1px] bg-coffee-latte my-4"></div>
      
      <p class="font-serif italic text-base text-coffee-brown max-w-[260px] leading-relaxed">
        “关于香气、情绪与静谧时光的感官日志”
      </p>
    </div>

    <!-- Bottom Button -->
    <div class="space-y-4 mb-4">
      <router-link
        to="/home"
        class="w-full py-4 bg-coffee-espresso text-coffee-warmWhite hover:bg-coffee-brown transition-all duration-300 rounded-sm text-xs font-semibold tracking-[0.25em] uppercase shadow-sm flex items-center justify-center gap-2"
      >
        <span>开启风味志</span>
        <ArrowRight class="w-4 h-4" />
      </router-link>
      <p class="text-[9px] text-coffee-softGray text-center tracking-wider font-light uppercase select-none">
        点击开启你的咖啡感官生活提案
      </p>
    </div>

  </div>
</template>

<script setup lang="ts">
import { ArrowRight } from 'lucide-vue-next'
</script>

<style scoped>
/* Page specific styling if any */
</style>
