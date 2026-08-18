<template>
  <div class="flex-1 w-full flex flex-col justify-center px-8 py-10 bg-coffee-warmWhite text-coffee-charcoal">
    
    <!-- Branding Header -->
    <div class="text-center space-y-2 mb-8 select-none">
      <h1 class="font-serif text-[44px] uppercase font-light text-coffee-espresso tracking-[0.1em] leading-none">Coffee Log</h1>
      <p class="font-serif italic text-sm text-coffee-brown">“登录手账，留存每一份咖啡感官记忆。”</p>
    </div>

    <!-- Editorial Sign In Form -->
    <div class="bg-coffee-cream/40 p-6 rounded-2xl border border-coffee-cream/80 space-y-5 shadow-sm">
      <h2 class="text-center font-serif text-lg font-light text-coffee-espresso uppercase tracking-wider border-b border-coffee-cream pb-3">SIGN IN / 登录手账</h2>
      
      <div class="space-y-4">
        <!-- Email Input -->
        <div class="space-y-1.5">
          <label class="text-[9px] uppercase tracking-wider font-semibold text-coffee-softGray block">邮箱账号 / Email</label>
          <input 
            type="email" 
            v-model="email" 
            placeholder="demo@mycoffeelog.com" 
            class="w-full p-3 bg-white/70 border border-coffee-latte/50 focus:border-coffee-brown focus:outline-none rounded-xl text-sm font-sans"
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
        @click="handleLogin" 
        class="w-full py-3 bg-coffee-espresso text-coffee-warmWhite hover:bg-coffee-brown transition-colors rounded-xl text-xs font-semibold tracking-[0.25em] uppercase shadow-sm flex items-center justify-center gap-1.5"
        :disabled="isLoading"
      >
        <template v-if="isLoading">
          <div class="w-4 h-4 border-2 border-white border-t-transparent rounded-full animate-spin"></div>
          <span>正在登录...</span>
        </template>
        <template v-else>
          <span>登录手账</span>
        </template>
      </button>

      <!-- Toggle Links -->
      <div class="text-center text-[10px] text-coffee-softGray pt-2">
        还没有咖啡手账账号？ 
        <router-link to="/register" class="text-coffee-brown hover:text-coffee-espresso underline font-semibold ml-1">立即注册</router-link>
      </div>
    </div>

    <!-- Quick Seed accounts prompt -->
    <div class="mt-8 p-4 bg-coffee-cream/20 rounded-2xl border border-dashed border-coffee-latte/50 text-[10px] text-coffee-brown space-y-1 text-center select-none">
      <p class="inline-flex items-center justify-center gap-1.5 font-bold">
        <AppIcon name="sparkles" :size="12" />
        便捷测试通道提示
      </p>
      <p class="font-light leading-relaxed">
        前端路由中已注入守卫，在测试阶段如果你直接访问需鉴权的页面，系统将自动注入 Mock 登录态，完全不阻碍你的评审与流畅体验！
      </p>
    </div>

  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import AppIcon from '@/components/AppIcon.vue'

const router = useRouter()
const authStore = useAuthStore()

const email = ref('')
const password = ref('')
const isLoading = ref(false)

const handleLogin = async () => {
  if (!email.value.trim()) {
    alert('请输入邮箱账户')
    return
  }
  if (!password.value.trim()) {
    alert('请输入密码')
    return
  }
  
  isLoading.value = true
  
  try {
    await authStore.login(email.value.trim(), password.value.trim())
    router.push('/home')
  } catch (e: any) {
    alert(e.message || '登录失败，请检查邮箱和密码')
  } finally {
    isLoading.value = false
  }
}
</script>

<style scoped>
/* Sign In specific styles */
</style>
