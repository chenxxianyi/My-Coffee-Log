<template>
  <div class="flex items-center justify-center select-none" :style="{ width: size + 'px', height: size + 'px' }">
    <svg :viewBox="`0 0 ${size} ${size}`" class="overflow-visible w-full h-full">
      <!-- 1. Concentric background grid polygons -->
      <polygon
        v-for="i in maxVal"
        :key="'grid-' + i"
        :points="getGridPoints(i)"
        fill="none"
        stroke="rgba(122, 86, 56, 0.12)"
        :stroke-width="0.8"
      />

      <!-- 2. Axes lines and text labels -->
      <g v-for="(dim, index) in dimensions" :key="'axis-' + index">
        <!-- Axis line -->
        <line
          :x1="cx"
          :y1="cy"
          :x2="getAxisEndPoint(index).x"
          :y2="getAxisEndPoint(index).y"
          stroke="rgba(122, 86, 56, 0.15)"
          :stroke-width="0.8"
        />
        <!-- Axis Text label (Shown only if showLabels is true) -->
        <text
          v-if="showLabels"
          :x="getLabelPosition(index).x"
          :y="getLabelPosition(index).y"
          font-family="Plus Jakarta Sans"
          :font-size="labelFontSize + 'px'"
          font-weight="600"
          fill="#7A5638"
          :text-anchor="getLabelAnchor(index)"
          dominant-baseline="middle"
        >
          {{ dim }}
        </text>
      </g>

      <!-- 3. Flavor Data Polygon (Latte color fill & brown stroke) -->
      <polygon
        :points="getDataPoints"
        fill="rgba(215, 196, 168, 0.45)"
        stroke="#7A5638"
        :stroke-width="1.2"
      />

      <!-- 4. Tiny dots on vertices for high editorial visual details -->
      <circle
        v-for="(val, index) in values"
        :key="'dot-' + index"
        :cx="getDotPosition(val, index).x"
        :cy="getDotPosition(val, index).y"
        :r="dotRadius"
        fill="#7A5638"
        stroke="#F7F3EC"
        :stroke-width="0.5"
      />
    </svg>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'

interface Props {
  values: number[]
  dimensions?: string[]
  size?: number
  maxVal?: number
  showLabels?: boolean
  labelFontSize?: number
  dotRadius?: number
}

const props = withDefaults(defineProps<Props>(), {
  dimensions: () => ["Acid", "Bitter", "Sweet", "Body", "Aroma", "After"],
  size: 110,
  maxVal: 5,
  showLabels: true,
  labelFontSize: 7.5,
  dotRadius: 2.2
})

// Trigonometry Center Coordinates
const cx = computed(() => props.size / 2)
const cy = computed(() => props.size / 2)
// Outer radius padding
const r = computed(() => (props.size / 2) * 0.72)

const axesCount = computed(() => props.dimensions.length)
const angleStep = computed(() => (2 * Math.PI) / axesCount.value)

// Grid Polygons coordinates
const getGridPoints = (step: number) => {
  const stepRadius = r.value * (step / props.maxVal)
  const points: string[] = []
  for (let a = 0; a < axesCount.value; a++) {
    const angle = a * angleStep.value - Math.PI / 2
    const px = cx.value + stepRadius * Math.cos(angle)
    const py = cy.value + stepRadius * Math.sin(angle)
    points.push(`${px},${py}`)
  }
  return points.join(' ')
}

// Axis end coordinates
const getAxisEndPoint = (index: number) => {
  const angle = index * angleStep.value - Math.PI / 2
  return {
    x: cx.value + r.value * Math.cos(angle),
    y: cy.value + r.value * Math.sin(angle)
  }
}

// Axis label text coordinates
const getLabelPosition = (index: number) => {
  const angle = index * angleStep.value - Math.PI / 2
  const labelPadding = props.labelFontSize * 1.4
  return {
    x: cx.value + (r.value + labelPadding) * Math.cos(angle),
    y: cy.value + (r.value + labelPadding) * Math.sin(angle)
  }
}

// Label text-anchor positioning helper
const getLabelAnchor = (index: number) => {
  const angle = index * angleStep.value - Math.PI / 2
  const cos = Math.cos(angle)
  if (cos > 0.05) return 'start'
  if (cos < -0.05) return 'end'
  return 'middle'
}

// Vertices scoring coordinates helper
const getDotPosition = (val: number, index: number) => {
  const scoreRadius = r.value * (val / props.maxVal)
  const angle = index * angleStep.value - Math.PI / 2
  return {
    x: cx.value + scoreRadius * Math.cos(angle),
    y: cy.value + scoreRadius * Math.sin(angle)
  }
}

// Data Polygon coordinates string
const getDataPoints = computed(() => {
  const points: string[] = []
  for (let a = 0; a < axesCount.value; a++) {
    const val = props.values[a] ?? 0
    const scoreRadius = r.value * (val / props.maxVal)
    const angle = a * angleStep.value - Math.PI / 2
    const dx = cx.value + scoreRadius * Math.cos(angle)
    const dy = cy.value + scoreRadius * Math.sin(angle)
    points.push(`${dx},${dy}`)
  }
  return points.join(' ')
})
</script>

<style scoped>
/* Scoped chart style overrides */
</style>
