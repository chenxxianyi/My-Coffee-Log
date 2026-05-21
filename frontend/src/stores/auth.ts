import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import * as authApi from '@/api/auth'

export interface User {
  id: number
  email: string
  nickname: string
  avatar_url: string
}

export const useAuthStore = defineStore('auth', () => {
  const token = ref<string | null>(localStorage.getItem('token'))
  const user = ref<User | null>(null)

  const isAuthenticated = computed(() => !!token.value)

  // Restore user info from API if token exists
  async function fetchUser() {
    if (!token.value) return
    try {
      const data = await authApi.getCurrentUser() as any
      user.value = {
        id: data.id,
        email: data.email,
        nickname: data.nickname,
        avatar_url: data.avatar_url || ''
      }
    } catch {
      logout()
    }
  }

  // Auto-fetch user on init if token exists
  if (token.value) {
    fetchUser()
  }

  async function login(email: string, password: string) {
    const res = await authApi.login(email, password) as any
    token.value = res.token
    localStorage.setItem('token', res.token)
    user.value = {
      id: res.user.id,
      email: res.user.email,
      nickname: res.user.nickname,
      avatar_url: res.user.avatar_url || ''
    }
  }

  async function register(email: string, password: string, nickname: string) {
    const res = await authApi.register(email, password, nickname) as any
    token.value = res.token
    localStorage.setItem('token', res.token)
    user.value = {
      id: res.user.id,
      email: res.user.email,
      nickname: res.user.nickname,
      avatar_url: res.user.avatar_url || ''
    }
  }

  function logout() {
    token.value = null
    user.value = null
    localStorage.removeItem('token')
  }

  async function updateNickname(name: string) {
    const data = await authApi.updateUser({ nickname: name }) as any
    if (user.value) {
      user.value.nickname = data.nickname
    }
  }

  async function updateAvatar(url: string) {
    const data = await authApi.updateUser({ avatar_url: url }) as any
    if (user.value) {
      user.value.avatar_url = data.avatar_url
    }
  }

  return {
    token,
    user,
    isAuthenticated,
    login,
    register,
    logout,
    updateNickname,
    updateAvatar,
    fetchUser
  }
})
