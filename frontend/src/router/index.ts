import { createRouter, createWebHistory, RouteLocationNormalized, NavigationGuardNext } from 'vue-router'
import Splash from '@/pages/Splash.vue'
import Home from '@/pages/Home.vue'

// Lazy-loaded pages for better bundle splitting
const CoffeeDetail = () => import('@/pages/CoffeeDetail.vue')
const Timeline = () => import('@/pages/Timeline.vue')
const ShareCard = () => import('@/pages/ShareCard.vue')
const CreateCoffeeLog = () => import('@/pages/CreateCoffeeLog.vue')
const Stats = () => import('@/pages/Stats.vue')
const MonthlyReview = () => import('@/pages/MonthlyReview.vue')
const CoffeeShops = () => import('@/pages/CoffeeShops.vue')
const CoffeeShopDetail = () => import('@/pages/CoffeeShopDetail.vue')
const Login = () => import('@/pages/Login.vue')
const Register = () => import('@/pages/Register.vue')
const Profile = () => import('@/pages/Profile.vue')
const RecordSuccess = () => import('@/pages/RecordSuccess.vue')
const Onboarding = () => import('@/pages/Onboarding.vue')
const WeeklyReview = () => import('@/pages/WeeklyReview.vue')

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
    path: '/onboarding',
    name: 'onboarding',
    component: Onboarding
  },
  {
    path: '/weekly-review',
    name: 'weekly-review',
    component: WeeklyReview
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
    path: '/coffee/:id/success',
    name: 'record-success',
    component: RecordSuccess,
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
    path: '/monthly-review',
    name: 'monthly-review',
    component: MonthlyReview
  },
  {
    path: '/coffee-shops',
    name: 'coffee-shops',
    component: CoffeeShops
  },
  {
    path: '/coffee-shops/:id',
    name: 'coffee-shop-detail',
    component: CoffeeShopDetail,
    props: true
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
router.beforeEach(async (to: RouteLocationNormalized, _from: RouteLocationNormalized, next: NavigationGuardNext) => {
  const token = localStorage.getItem('token')
  const publicPages = ['splash', 'login', 'register']
  const authRequired = !publicPages.includes(to.name as string)

  if (authRequired && !token) {
    next({ name: 'login' })
  } else if (to.name === 'onboarding' && !token) {
    next({ name: 'login' })
  } else if (to.name === 'home' && token) {
    // Check onboarding status for new users
    try {
      const response = await fetch('/api/v1/users/me', {
        headers: { 'Authorization': `Bearer ${token}` }
      })
      if (response.ok) {
        const result = await response.json()
        const user = result.data
        if (!user.onboarding_completed && !user.first_record_at) {
          next({ name: 'onboarding' })
          return
        }
      }
    } catch {
      // Continue to home if check fails
    }
    next()
  } else {
    next()
  }
})

export default router
