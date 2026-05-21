import { createRouter, createWebHistory, RouteLocationNormalized, NavigationGuardNext } from 'vue-router'
import Splash from '@/pages/Splash.vue'
import Home from '@/pages/Home.vue'
import CoffeeDetail from '@/pages/CoffeeDetail.vue'
import Timeline from '@/pages/Timeline.vue'
import ShareCard from '@/pages/ShareCard.vue'
import CreateCoffeeLog from '@/pages/CreateCoffeeLog.vue'
import Stats from '@/pages/Stats.vue'
import Login from '@/pages/Login.vue'
import Register from '@/pages/Register.vue'
import Profile from '@/pages/Profile.vue'

const routes = [
  {
    path: '/',
    name: 'splash',
    component: Splash
  },
  {
    path: '/home',
    name: 'home',
    component: Home
  },
  {
    path: '/login',
    name: 'login',
    component: Login
  },
  {
    path: '/register',
    name: 'register',
    component: Register
  },
  {
    path: '/create',
    name: 'create',
    component: CreateCoffeeLog
  },
  {
    path: '/coffee/:id',
    name: 'detail',
    component: CoffeeDetail,
    props: true
  },
  {
    path: '/timeline',
    name: 'timeline',
    component: Timeline
  },
  {
    path: '/stats',
    name: 'stats',
    component: Stats
  },
  {
    path: '/profile',
    name: 'profile',
    component: Profile
  },
  {
    path: '/share/:id',
    name: 'share',
    component: ShareCard,
    props: true
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes,
  scrollBehavior() {
    return { top: 0 }
  }
})

// Navigation guard for auth
router.beforeEach((to: RouteLocationNormalized, _from: RouteLocationNormalized, next: NavigationGuardNext) => {
  const token = localStorage.getItem('token')
  const publicPages = ['splash', 'login', 'register']
  const authRequired = !publicPages.includes(to.name as string)

  if (authRequired && !token) {
    next({ name: 'login' })
  } else {
    next()
  }
})

export default router
