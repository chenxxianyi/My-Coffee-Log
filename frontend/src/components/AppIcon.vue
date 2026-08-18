<template>
  <component
    :is="iconComponent"
    :size="size"
    :style="{ strokeWidth }"
    aria-hidden="true"
  />
</template>

<script setup lang="ts">
import { computed, type Component } from 'vue'
import {
  Armchair,
  BatteryLow,
  BookOpen,
  Building2,
  CalendarRange,
  CakeSlice,
  Citrus,
  CloudRain,
  Coffee,
  Flame,
  GraduationCap,
  Headphones,
  Heart,
  Home,
  Laptop,
  Leaf,
  Milk,
  Moon,
  Palmtree,
  PartyPopper,
  Plane,
  Rocket,
  Smile,
  Snail,
  Sparkles,
  Sun,
  Sunrise,
  Target,
  UsersRound,
  Waves,
  Zap
} from 'lucide-vue-next'

const props = withDefaults(defineProps<{
  name?: string
  size?: number
  strokeWidth?: number
}>(), {
  name: 'sparkles',
  size: 16,
  strokeWidth: 1.7
})

const iconMap: Record<string, Component> = {
  calm: Waves,
  focused: Target,
  tired: BatteryLow,
  happy: Smile,
  energetic: Zap,
  reflective: Moon,
  rainy: CloudRain,
  slow: Snail,
  productive: Rocket,
  morning: Sunrise,
  office: Building2,
  weekend: Palmtree,
  cafe: Coffee,
  travel: Plane,
  home: Home,
  study: GraduationCap,
  book: BookOpen,
  music: Headphones,
  work: Laptop,
  dessert: CakeSlice,
  alone: Armchair,
  friends: UsersRound,
  bright: Sun,
  floral: Leaf,
  smooth: Heart,
  bold: Flame,
  citrus: Citrus,
  creamy: Milk,
  social: PartyPopper,
  'monthly-review': CalendarRange,
  coffee: Coffee,
  sparkles: Sparkles
}

const iconComponent = computed(() => iconMap[props.name] ?? Sparkles)
</script>
