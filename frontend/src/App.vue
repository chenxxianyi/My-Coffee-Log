<template>
  <!-- Main Viewport centered as phone size on desktop, full-bleed on mobile -->
  <div class="paper-grain text-coffee-charcoal font-sans antialiased min-h-screen flex flex-col justify-between max-w-[430px] mx-auto border-x border-coffee-cream/80 shadow-sm relative overflow-hidden">
    
    <!-- Core Page View Router -->
    <router-view v-slot="{ Component }">
      <transition name="fade" mode="out-in">
        <component :is="Component" />
      </transition>
    </router-view>
    
  </div>
</template>

<script setup lang="ts">
// Root controller
</script>

<style>
/* App wide custom style overrides */
</style>
