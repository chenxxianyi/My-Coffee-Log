<template>
  <div class="flex-1 w-full flex flex-col justify-center px-8 py-10 bg-coffee-warmWhite text-coffee-charcoal">
    
    <!-- Branding Header -->
    <div class="text-center space-y-2 mb-8 select-none">
      <h1 class="font-serif text-[44px] uppercase font-light text-coffee-espresso tracking-[0.1em] leading-none">Coffee Log</h1>
      <p class="font-serif italic text-sm text-coffee-brown">“从今天起，开启属于你的咖啡感官风味志。”</p>
    </div>

    <!-- Editorial Sign Up Form -->
    <div class="bg-coffee-cream/40 p-6 rounded-2xl border border-coffee-cream/80 space-y-5 shadow-sm">
      <h2 class="text-center font-serif text-lg font-light text-coffee-espresso uppercase tracking-wider border-b border-coffee-cream pb-3">CREATE ACCOUNT / 注册账户</h2>
      
      <div class="space-y-3.5">
        <!-- Nickname Input -->
        <div class="space-y-1.5">
          <label class="text-[9px] uppercase tracking-wider font-semibold text-coffee-softGray block">手账昵称 / Nickname</label>
          <input 
            type="text" 
            v-model="nickname" 
            placeholder="风味探索者" 
            class="w-full p-3 bg-white/70 border border-coffee-latte/50 focus:border-coffee-brown focus:outline-none rounded-xl text-sm"
          >
        </div>

        <!-- Email Input -->
        <div class="space-y-1.5">
          <label class="text-[9px] uppercase tracking-wider font-semibold text-coffee-softGray block">邮箱账号 / Email</label>
          <input 
            type="email" 
            v-model="email" 
            placeholder="demo@mycoffeelog.com" 
            class="w-full p-3 bg-white/70 border border-coffee-latte/50 focus:border-coffee-brown focus:outline-none rounded-xl text-sm"
          >
        </div>

        <!-- Password Input -->
        <div class="space-y-1.5">
          <label class="text-[9px] uppercase tracking-wider font-semibold text-coffee-softGray block">登录密码 / Password</label>
          <input 
            type="password" 
            v-model="password" 
            placeholder="••••••" 
            class="w-full p-3 bg-white/70 border border-coffee-latte/50 focus:border-coffee-brown focus:outline-none rounded-xl text-sm"
          >
        </div>
      </div>

      <!-- Action Button -->
      <button 
        @click="handleRegister" 
        class="w-full py-3 bg-coffee-espresso text-coffee-warmWhite hover:bg-coffee-brown transition-colors rounded-xl text-xs font-semibold tracking-[0.25em] uppercase shadow-sm flex items-center justify-center gap-1.5"
        :disabled="isLoading"
      >
        <template v-if="isLoading">
          <div class="w-4 h-4 border-2 border-white border-t-transparent rounded-full animate-spin"></div>
          <span>正在创建账号...</span>
        </template>
        <template v-else>
          <span>创建并开启手账</span>
        </template>
      </button>

      <!-- Toggle Links -->
      <div class="text-center text-[10px] text-coffee-softGray pt-2">
        已有咖啡手账账户？ 
        <router-link to="/login" class="text-coffee-brown hover:text-coffee-espresso underline font-semibold ml-1">立即登录</router-link>
      </div>
    </div>

  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

const router = useRouter()
const authStore = useAuthStore()

const nickname = ref('')
const email = ref('')
const password = ref('')
const isLoading = ref(false)

const handleRegister = async () => {
  if (!nickname.value.trim()) {
    alert('请输入昵称')
    return
  }
  if (!email.value.trim()) {
    alert('请输入邮箱')
    return
  }
  if (!password.value.trim()) {
    alert('请输入密码')
    return
  }
  
  isLoading.value = true
  
  try {
    await authStore.register(email.value.trim(), password.value.trim(), nickname.value.trim())
    router.push('/home')
  } catch (e: any) {
    alert(e.message || '注册失败，请稍后重试')
  } finally {
    isLoading.value = false
  }
}
</script>

<style scoped>
/* Sign Up specific styles */
</style>
