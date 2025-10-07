<template>
  <div class="page">
    <AppHeader/>

    <main class="container">
      <div class="role-tabs">
        <button class="rtab active">Инженер</button>
        <button class="rtab">Менеджер</button>
        <button class="rtab">Руководитель</button>
      </div>

      <h1 class="h1">Личный кабинет</h1>

      <section class="kpis">
        <KpiCard title="Открытые Дефекты" :value="32" caption="всего по проектам">
          <template #icon>🟠</template>
        </KpiCard>
        <KpiCard title="Проектов в Работе" :value="3" caption="активные проекты">
          <template #icon>🟡</template>
        </KpiCard>
        <KpiCard title="Дефекты по Приоритету" :value="'10 ВЫСОКИХ'" caption="требуют немедленного внимания">
          <template #icon>❗</template>
        </KpiCard>
      </section>

      <section class="panels">
        <div class="panel">
          <div class="panel-title">Распределение Дефектов</div>
          <DonutChart :segments="segments" :total="totalDefects" :size="240"/>
        </div>

        <div class="panel">
          <div class="panel-title">Краткий обзор статусов</div>
          <ul class="status-list">
            <li v-for="s in segments" :key="s.label">
              <span class="dot" :style="{background:s.color}"></span>
              <span class="sname">{{ s.label }}</span>
              <span class="sperc">{{ Math.round((s.value/totalDefects)*100) }}%</span>
            </li>
          </ul>
          <a href="#" class="link">Посмотреть все Дефекты ↗</a>
        </div>
      </section>

      <h2 class="h2">Список Активных Проектов</h2>
      <section class="projects">
        <ProjectCard name="ЖК 'Солнечный'" status="Активен" defects="15 / 120"/>
        <ProjectCard name="Бизнес-центр 'Вершина'" status="Активен" defects="8 / 80"/>
        <ProjectCard name="Торговый центр 'Галерея'" status="Приостановлен" defects="3 / 50"/>
      </section>

    </main>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import AppHeader from '../components/AppHeader.vue'
import KpiCard from '../components/KpiCard.vue'
import DonutChart from '../components/DonutChart.vue'
import ProjectCard from '../components/ProjectCard.vue'

const seed = Math.random()
function rnd(n){ return Math.floor((Math.random()+0.2) * n) }
const base = [rnd(10)+10, rnd(10)+10, rnd(10)+10, rnd(10)+10]
const sum = base.reduce((a,b)=>a+b,0)
const scaled = base.map(v => Math.round(v * (100/sum)))
const totalDefects = 100
const adjust = totalDefects - scaled.reduce((a,b)=>a+b,0)
scaled[0] += adjust

const segments = ref([
  { label:'Открыт',    color:'#e07a33', value: scaled[0] },
  { label:'В работе',  color:'#f0c24b', value: scaled[1] },
  { label:'Решен',     color:'#8be19f', value: scaled[2] },
  { label:'Закрыт',    color:'#a9afbb', value: scaled[3] },
])
</script>

<style scoped>
.page{min-height:100vh; background:#0b0d12; color:#e7e9ee}
.container{max-width:1180px; margin:0 auto; padding:18px 20px 32px}
.role-tabs{display:flex; gap:10px; margin:12px 0 10px}
.rtab{background:#1d2229; color:#c7cbd3; border:1px solid rgba(255,255,255,.08); border-radius:8px; padding:8px 12px; cursor:pointer}
.rtab.active{background:#0f141a; color:#e7e9ee}
.h1{font-size:28px; margin:10px 0 14px}
.kpis{display:grid; grid-template-columns:repeat(3,1fr); gap:14px; margin-bottom:18px}
.panels{display:grid; grid-template-columns:2fr 1.4fr; gap:14px; margin-bottom:18px}
.panel{background:#1d2229;border:1px solid rgba(255,255,255,.06); border-radius:10px; padding:16px}
.panel-title{color:#e7e9ee; font-weight:700; margin-bottom:10px}
.status-list{list-style:none; padding:0; margin:0; display:grid; gap:8px}
.status-list li{display:flex; align-items:center; justify-content:space-between}
.dot{width:10px;height:10px;border-radius:50%; margin-right:8px; display:inline-block}
.sname{flex:1; margin-left:8px}
.sperc{color:#c7cbd3}
.h2{font-size:18px; margin:16px 0 10px}
.projects{display:grid; grid-template-columns:repeat(3,1fr); gap:14px; margin-bottom:18px}
@media (max-width: 1100px){
  .kpis{grid-template-columns:1fr 1fr}
  .panels{grid-template-columns:1fr}
  .projects,.quick{grid-template-columns:1fr 1fr}
}
@media (max-width: 680px){
  .kpis,.projects,.quick{grid-template-columns:1fr}
}
</style>