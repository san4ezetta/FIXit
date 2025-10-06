<template>
  <div class="page">
    <div class="card">
      <div class="logo">
        <img src="../assets/img/logo.png" alt="">
      </div>

      <h1 class="title">
        Добро пожаловать в сервис<br/>
        <span class="brand">FIXit</span>
      </h1>

      <div class="tabs">
        <button class="tab" :class="{ active: activeTab==='login' }" @click="activeTab='login'">Вход</button>
        <button class="tab" :class="{ active: activeTab==='register' }" @click="activeTab='register'">Регистрация</button>
      </div>

      <LoginForm v-if="activeTab==='login'" @switch-to-register="activeTab='register'"/>
      <RegisterForm v-else @switch-to-login="activeTab='login'"/>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import LoginForm from '../components/LoginForm.vue'
import RegisterForm from '../components/RegisterForm.vue'

const activeTab = ref('login')
</script>

<style scoped>
.page{min-height:100vh;background:#0b0d12;display:flex;align-items:center;justify-content:center;padding:24px}
.card{width:520px;max-width:92vw;background:#1d2229;border-radius:12px;box-shadow:0 8px 28px rgba(0,0,0,.45);padding:36px 40px 28px;border:1px solid rgba(255,255,255,.06)}
.logo{width:56px;height:56px;margin:0 auto 16px}
.logo img{width:100%;height:100%;object-fit:contain}
.title{text-align:center;color:#e6e7eb;font-weight:800;line-height:1.25;margin:0 0 18px;font-size:28px}
.brand{display:inline-block}
.tabs{display:flex;gap:8px;background:#20262e;border-radius:8px;padding:4px;margin:8px 0 18px}
.tab{flex:1;background:transparent;color:#c7cbd3;border:none;padding:10px 0;border-radius:6px;cursor:pointer;font-weight:600}
.tab.active{background:#0f141a;color:#e7e9ee;box-shadow:inset 0 0 0 1px rgba(255,255,255,.06)}
</style>