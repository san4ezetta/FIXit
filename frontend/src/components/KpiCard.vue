<template>
  <div class="kpi">
    <div class="kpi-top">
      <div class="kpi-title">{{ title }}</div>
      <slot name="icon"></slot>
    </div>
    <div class="kpi-main">
      <div class="kpi-value">{{ value }}</div>
      <div v-if="caption" class="kpi-caption">{{ caption }}</div>
    </div>
  </div>
</template>

<script setup>
defineProps({
  title: String,
  value: [String, Number],
  caption: String
})
</script>

<style scoped>
.kpi{
  background:#1d2229; border:1px solid rgba(255,255,255,.06);
  border-radius:10px; padding:16px; min-height:110px;
}
.kpi-top{display:flex; justify-content:space-between; align-items:center; margin-bottom:8px}
.kpi-title{color:#c7cbd3; font-size:14px}
.kpi-main{display:flex; flex-direction:column; gap:4px}
.kpi-value{color:#e7e9ee; font-size:34px; font-weight:800; line-height:1}
.kpi-caption{color:#c7cbd3; font-size:13px}
</style>