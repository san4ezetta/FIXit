<template>
  <div class="donut-wrap">
    <svg :width="size" :height="size" viewBox="0 0 100 100">
      <g transform="rotate(-90 50 50)">
        <template v-for="(s, i) in normalized" :key="i">
          <circle
              cx="50" cy="50" r="38"
              fill="transparent"
              :stroke="s.color"
              stroke-width="12"
              :stroke-dasharray="`${s.perc*2.387} ${circ - s.perc*2.387}`"
              :stroke-dashoffset="s.offset"
              stroke-linecap="butt"
          />
        </template>
      </g>
      <circle cx="50" cy="50" r="30" fill="#1d2229"/>
      <text x="50" y="50" text-anchor="middle" dominant-baseline="middle" class="label">
        {{ total }}
      </text>
    </svg>

    <ul class="legend">
      <li v-for="(s, i) in normalized" :key="i">
        <span class="dot" :style="{ background: s.color }"></span>
        <span class="legend-label">{{ s.label }}</span>
        <span class="legend-val">{{ Math.round(s.perc) }}%</span>
      </li>
    </ul>
  </div>
</template>

<script setup>
import { computed } from 'vue'
const props = defineProps({
  segments: { type: Array, required: true },
  total: { type: Number, default: 0 },
  size: { type: Number, default: 240 }
})

const circ = 238.7
const normalized = computed(() => {
  const sum = props.segments.reduce((a,b)=>a + (b.value||0), 0) || 1
  let acc = 0
  return props.segments.map(s => {
    const perc = (s.value/sum)*100
    const len = (perc/100)*circ
    const offset = -(acc/sum)*circ
    acc += s.value
    return { ...s, perc, offset }
  })
})
</script>

<style scoped>
.donut-wrap{display:grid; grid-template-columns: auto 1fr; gap:18px; align-items:center}
.legend{list-style:none; padding:0; margin:0; display:grid; gap:8px}
.dot{width:10px;height:10px;border-radius:50%; display:inline-block; margin-right:8px}
.legend-label{color:#e7e9ee; margin-right:8px}
.legend-val{color:#c7cbd3}
.label{fill:#e7e9ee; font-size:16px; font-weight:700}
@media (max-width: 860px){
  .donut-wrap{grid-template-columns: 1fr}
}
</style>