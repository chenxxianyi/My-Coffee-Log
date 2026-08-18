<template>
  <div class="register-page flex-1 w-full flex flex-col justify-center px-6 sm:px-8 py-8 bg-coffee-warmWhite text-coffee-charcoal overflow-y-auto">
    <!-- Branding Header -->
    <header class="text-center space-y-2 mb-6 select-none">
      <h1 class="font-serif text-[42px] uppercase font-light text-coffee-espresso tracking-[0.1em] leading-none">Coffee Log</h1>
      <p class="font-serif italic text-sm text-coffee-brown">“从今天起，开启属于你的咖啡感官风味志。”</p>
    </header>

    <!-- Editorial Sign Up Form -->
    <form
      class="bg-coffee-cream/40 p-5 sm:p-6 rounded-2xl border border-coffee-cream/80 space-y-4 shadow-sm"
      novalidate
      @submit.prevent="handleRegister"
    >
      <h2 class="text-center font-serif text-lg font-light text-coffee-espresso uppercase tracking-wider border-b border-coffee-cream pb-3">
        CREATE ACCOUNT / 注册账户
      </h2>

      <div class="space-y-3.5">
        <!-- Nickname Input -->
        <div class="space-y-1.5">
          <label for="nickname" class="field-label">手账昵称 / Nickname</label>
          <input
            id="nickname"
            ref="nicknameInput"
            v-model="nickname"
            type="text"
            name="nickname"
            autocomplete="nickname"
            maxlength="100"
            placeholder="风味探索者"
            class="field-input"
            :class="{ 'field-input--error': nicknameError }"
            :aria-invalid="Boolean(nicknameError)"
            :aria-describedby="nicknameError ? 'nickname-error' : undefined"
            @blur="touched.nickname = true"
            @input="formError = ''"
          >
          <p v-if="nicknameError" id="nickname-error" class="field-message field-message--error" role="alert">
            {{ nicknameError }}
          </p>
        </div>

        <!-- Email Input -->
        <div class="space-y-1.5">
          <label for="email" class="field-label">邮箱账号 / Email</label>
          <input
            id="email"
            ref="emailInput"
            v-model="email"
            type="email"
            name="email"
            inputmode="email"
            autocomplete="email"
            placeholder="demo@mycoffeelog.com"
            class="field-input"
            :class="{ 'field-input--error': emailError }"
            :aria-invalid="Boolean(emailError)"
            :aria-describedby="emailError ? 'email-error' : undefined"
            @blur="touched.email = true"
            @input="formError = ''"
          >
          <p v-if="emailError" id="email-error" class="field-message field-message--error" role="alert">
            {{ emailError }}
          </p>
        </div>

        <!-- Password Input -->
        <div class="space-y-1.5">
          <label for="password" class="field-label">登录密码 / Password</label>
          <div class="relative">
            <input
              id="password"
              ref="passwordInput"
              v-model="password"
              :type="showPassword ? 'text' : 'password'"
              name="password"
              autocomplete="new-password"
              placeholder="至少 6 位字符"
              class="field-input pr-11"
              :class="{ 'field-input--error': passwordError }"
              :aria-invalid="Boolean(passwordError)"
              :aria-describedby="passwordError ? 'password-error' : 'password-hint'"
              @blur="touched.password = true"
              @input="handlePasswordInput"
            >
            <button
              type="button"
              class="password-toggle"
              :aria-label="showPassword ? '隐藏密码' : '显示密码'"
              :title="showPassword ? '隐藏密码' : '显示密码'"
              @click="showPassword = !showPassword"
            >
              <EyeOff v-if="showPassword" :size="17" aria-hidden="true" />
              <Eye v-else :size="17" aria-hidden="true" />
            </button>
          </div>
          <p v-if="passwordError" id="password-error" class="field-message field-message--error" role="alert">
            {{ passwordError }}
          </p>
          <p v-else id="password-hint" class="field-message">建议使用字母、数字或符号组合</p>
        </div>

        <!-- Confirm Password Input -->
        <div class="space-y-1.5">
          <label for="confirm-password" class="field-label">确认密码 / Confirm Password</label>
          <div class="relative">
            <input
              id="confirm-password"
              ref="confirmPasswordInput"
              v-model="confirmPassword"
              :type="showConfirmPassword ? 'text' : 'password'"
              name="confirm-password"
              autocomplete="new-password"
              placeholder="再次输入登录密码"
              class="field-input pr-11"
              :class="{
                'field-input--error': confirmPasswordError,
                'field-input--success': passwordsMatch
              }"
              :aria-invalid="Boolean(confirmPasswordError)"
              :aria-describedby="confirmPasswordMessage ? 'confirm-password-message' : undefined"
              @blur="touched.confirmPassword = true"
              @input="formError = ''"
            >
            <button
              type="button"
              class="password-toggle"
              :aria-label="showConfirmPassword ? '隐藏确认密码' : '显示确认密码'"
              :title="showConfirmPassword ? '隐藏确认密码' : '显示确认密码'"
              @click="showConfirmPassword = !showConfirmPassword"
            >
              <EyeOff v-if="showConfirmPassword" :size="17" aria-hidden="true" />
              <Eye v-else :size="17" aria-hidden="true" />
            </button>
          </div>
          <p
            v-if="confirmPasswordMessage"
            id="confirm-password-message"
            class="field-message"
            :class="passwordsMatch ? 'field-message--success' : 'field-message--error'"
            :role="confirmPasswordError ? 'alert' : undefined"
            aria-live="polite"
          >
            <CheckCircle2 v-if="passwordsMatch" :size="12" aria-hidden="true" />
            {{ confirmPasswordMessage }}
          </p>
        </div>
      </div>

      <p
        v-if="formError"
        class="rounded-xl border border-coffee-latte/50 bg-white/60 px-3 py-2.5 text-[11px] leading-relaxed text-coffee-latte"
        role="alert"
      >
        {{ formError }}
      </p>

      <!-- Action Button -->
      <button
        type="submit"
        class="w-full min-h-12 py-3 bg-coffee-espresso text-coffee-warmWhite hover:bg-coffee-brown focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-coffee-latte focus-visible:ring-offset-2 transition-colors rounded-xl text-xs font-semibold tracking-[0.25em] uppercase shadow-sm flex items-center justify-center gap-1.5 disabled:opacity-60 disabled:cursor-not-allowed"
        :disabled="isLoading"
      >
        <template v-if="isLoading">
          <span class="w-4 h-4 border-2 border-white border-t-transparent rounded-full animate-spin" aria-hidden="true"></span>
          <span>正在创建账号...</span>
        </template>
        <span v-else>创建并开启手账</span>
      </button>

      <!-- Toggle Links -->
      <div class="text-center text-[10px] text-coffee-softGray pt-1">
        已有咖啡手账账户？
        <router-link to="/login" class="text-coffee-brown hover:text-coffee-espresso focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-coffee-latte rounded-sm underline font-semibold ml-1">
          立即登录
        </router-link>
      </div>
    </form>
  </div>
</template>

<script setup lang="ts">
import { computed, nextTick, reactive, ref } from 'vue'
import { useRouter } from 'vue-router'
import { CheckCircle2, Eye, EyeOff } from 'lucide-vue-next'
import { useAuthStore } from '@/stores/auth'

const router = useRouter()
const authStore = useAuthStore()

const nickname = ref('')
const email = ref('')
const password = ref('')
const confirmPassword = ref('')
const showPassword = ref(false)
const showConfirmPassword = ref(false)
const isLoading = ref(false)
const submitted = ref(false)
const formError = ref('')
const nicknameInput = ref<HTMLInputElement | null>(null)
const emailInput = ref<HTMLInputElement | null>(null)
const passwordInput = ref<HTMLInputElement | null>(null)
const confirmPasswordInput = ref<HTMLInputElement | null>(null)
const touched = reactive({
  nickname: false,
  email: false,
  password: false,
  confirmPassword: false
})

const emailPattern = /^[^\s@]+@[^\s@]+\.[^\s@]+$/

const nicknameError = computed(() => {
  if (!submitted.value && !touched.nickname) return ''
  return nickname.value.trim() ? '' : '请输入手账昵称'
})

const emailError = computed(() => {
  if (!submitted.value && !touched.email) return ''
  const value = email.value.trim()
  if (!value) return '请输入邮箱账号'
  return emailPattern.test(value) ? '' : '请输入有效的邮箱地址'
})

const passwordError = computed(() => {
  if (!submitted.value && !touched.password) return ''
  if (!password.value) return '请输入登录密码'
  return password.value.length >= 6 ? '' : '密码至少需要 6 位字符'
})

const passwordsEqual = computed(() => (
  Boolean(confirmPassword.value) && password.value === confirmPassword.value
))

const passwordsMatch = computed(() => passwordsEqual.value && password.value.length >= 6)

const confirmPasswordError = computed(() => {
  if (!submitted.value && !touched.confirmPassword) return ''
  if (!confirmPassword.value) return '请再次输入登录密码'
  return passwordsEqual.value ? '' : '两次输入的密码不一致'
})

const confirmPasswordMessage = computed(() => {
  if (passwordsMatch.value) return '两次密码输入一致'
  return confirmPasswordError.value
})

const isFormValid = computed(() => (
  Boolean(nickname.value.trim()) &&
  emailPattern.test(email.value.trim()) &&
  password.value.length >= 6 &&
  passwordsEqual.value
))

const handlePasswordInput = () => {
  formError.value = ''
  if (confirmPassword.value) touched.confirmPassword = true
}

const getRegisterError = (error: unknown) => {
  const message = error instanceof Error ? error.message : ''
  if (/email already exists/i.test(message)) return '该邮箱已经注册，请直接登录或更换邮箱'
  return message || '注册失败，请稍后重试'
}

const handleRegister = async () => {
  submitted.value = true
  formError.value = ''

  if (!isFormValid.value) {
    await nextTick()
    const firstInvalidInput = [
      !nickname.value.trim() ? nicknameInput.value : null,
      !emailPattern.test(email.value.trim()) ? emailInput.value : null,
      password.value.length < 6 ? passwordInput.value : null,
      !passwordsEqual.value ? confirmPasswordInput.value : null
    ].find((input): input is HTMLInputElement => Boolean(input))

    firstInvalidInput?.focus()
    return
  }

  isLoading.value = true

  try {
    await authStore.register(email.value.trim(), password.value, nickname.value.trim())
    await router.push('/home')
  } catch (error: unknown) {
    formError.value = getRegisterError(error)
  } finally {
    isLoading.value = false
  }
}
</script>

<style scoped>
.field-label {
  display: block;
  color: #9c7b59;
  font-size: 9px;
  font-weight: 600;
  letter-spacing: 0.08em;
  text-transform: uppercase;
}

.field-input {
  width: 100%;
  border: 1px solid rgba(231, 111, 81, 0.5);
  border-radius: 0.75rem;
  background: rgba(255, 255, 255, 0.7);
  padding: 0.75rem;
  color: #2a1a0e;
  font-size: 0.875rem;
  line-height: 1.25rem;
  outline: none;
  transition: border-color 160ms ease, box-shadow 160ms ease, background-color 160ms ease;
}

.field-input::placeholder {
  color: rgba(156, 123, 89, 0.68);
}

.field-input.pr-11 {
  padding-right: 2.75rem;
}

.field-input:focus {
  border-color: #5c3d2e;
  background: rgba(255, 255, 255, 0.86);
  box-shadow: 0 0 0 3px rgba(231, 111, 81, 0.12);
}

.field-input--error,
.field-input--error:focus {
  border-color: #c9533c;
  box-shadow: 0 0 0 3px rgba(201, 83, 60, 0.1);
}

.field-input--success,
.field-input--success:focus {
  border-color: #73815d;
  box-shadow: 0 0 0 3px rgba(115, 129, 93, 0.1);
}

.password-toggle {
  position: absolute;
  top: 50%;
  right: 0.7rem;
  display: grid;
  width: 2rem;
  height: 2rem;
  place-items: center;
  transform: translateY(-50%);
  border-radius: 9999px;
  color: #9c7b59;
  transition: color 160ms ease, background-color 160ms ease;
}

.password-toggle:hover {
  color: #5c3d2e;
  background: rgba(253, 232, 194, 0.65);
}

.password-toggle:focus-visible {
  outline: 2px solid #e76f51;
  outline-offset: 1px;
}

.field-message {
  display: flex;
  min-height: 0.75rem;
  align-items: center;
  gap: 0.25rem;
  padding-left: 0.15rem;
  color: #9c7b59;
  font-size: 9px;
  line-height: 0.75rem;
}

.field-message--error {
  color: #b94832;
}

.field-message--success {
  color: #66734f;
}

@media (max-height: 740px) {
  .register-page {
    justify-content: flex-start;
    padding-top: 1.5rem;
    padding-bottom: 1.5rem;
  }
}
</style>
