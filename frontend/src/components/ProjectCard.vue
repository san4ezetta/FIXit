<template>
  <div class="pc">
    <div class="pc-top">
      <div class="pc-title">{{ name }}</div>
      <span class="badge" :class="statusClass">{{ status }}</span>
    </div>
    <div class="pc-meta">Дефектов: {{ defects }}</div>
    <a href="#" class="more">Подробнее ↗</a>
  </div>
</template>

<script setup>
const props = defineProps({
  name: String,
  status: String,
  defects: String
})
const statusClass = computed(() => {
  if (props.status === 'Активен') return 'ok'
  if (props.status === 'Приостановлен') return 'warn'
  return 'muted'
})
import { computed } from 'vue'
</script>

<style scoped>
.pc{background:#1d2229;border:1px solid rgba(255,255,255,.06);border-radius:10px;padding:14px}
.pc-top{display:flex;justify-content:space-between;align-items:center;margin-bottom:8px}
.pc-title{color:#e7e9ee;font-weight:700}
.badge{padding:2px 8px;border-radius:999px;font-size:12px;border:1px solid rgba(255,255,255,.1)}
.badge.ok{background:#16351f;color:#8be19f;border-color:#1f6a34}
.badge.warn{background:#2e2320;color:#f7b199;border-color:#774d3e}
.badge.muted{background:#22262d;color:#c7cbd3;border-color:#3a3f47}
.pc-meta{color:#c7cbd3;margin-bottom:8px}
.more{color:#e07a33;text-decoration:none;font-size:14px}
</style>