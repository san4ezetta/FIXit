<template>
  <div class="page">
    <div class="card">
      <div class="logo">
        <img src="../assets/img/logo.png" alt="">
      </div>

      <h1 class="title">
        Добро пожаловать в сервис<br />
        <span class="brand">FIXit</span>
      </h1>

      <div class="tabs">
        <button
            class="tab"
            :class="{ active: activeTab === 'login' }"
            @click="activeTab='login'"
        >
          Вход
        </button>
        <button
            class="tab"
            :class="{ active: activeTab === 'register' }"
            @click="activeTab='register'"
        >
          Регистрация
        </button>
      </div>

      <form v-if="activeTab==='login'" class="form" @submit.prevent="noop">
        <label class="label">Электронная почта</label>
        <div class="input">
          <MailIcon />
          <input type="email" placeholder="example@sitedefect.com" />
        </div>

        <label class="label">Пароль</label>
        <div class="input">
          <LockIcon />
          <input type="password" placeholder="Введите ваш пароль" />
        </div>

        <button class="btn" type="submit" disabled>Войти</button>

        <a href="#" class="link">Забыли пароль?</a>
      </form>

      <form v-else class="form" @submit.prevent="submitRegister">
        <div class="grid">
          <div class="col">
            <label class="label">Имя</label>
            <div class="input">
              <UserIcon />
              <input v-model.trim="form.name" type="text" placeholder="Иван" required />
            </div>
          </div>
          <div class="col">
            <label class="label">Фамилия</label>
            <div class="input">
              <UserIcon />
              <input v-model.trim="form.surname" type="text" placeholder="Иванов" required />
            </div>
          </div>
        </div>

        <label class="label">Отчество</label>
        <div class="input">
          <UserIcon />
          <input v-model.trim="form.patronymic" type="text" placeholder="Иванович" />
        </div>

        <label class="label">Электронная почта</label>
        <div class="input">
          <MailIcon />
          <input v-model.trim="form.email" type="email" placeholder="example@sitedefect.com" required />
        </div>

        <label class="label">Пароль</label>
        <div class="input">
          <LockIcon />
          <input v-model="form.password" type="password" placeholder="Введите ваш пароль" required />
        </div>

        <label class="label">Повторите пароль</label>
        <div class="input">
          <LockIcon />
          <input v-model="form.password2" type="password" placeholder="Повторите пароль" required />
        </div>

        <div v-if="message" :class="['alert', message.type]">{{ message.text }}</div>

        <button class="btn" type="submit" :disabled="submitting">
          {{ submitting ? 'Отправка…' : 'Зарегистрироваться' }}
        </button>

        <a href="#" class="link">Забыли пароль?</a>
      </form>
    </div>
  </div>
</template>

<script setup>

import { ref } from 'vue'
import MailIcon from './components/MailIcon.vue'
import LockIcon from './components/LockIcon.vue'
import UserIcon from './components/UserIcon.vue'

const activeTab = ref('login')
const submitting = ref(false)
const message = ref(null)

const form = ref({
  name: '',
  surname: '',
  patronymic: '',
  email: '',
  password: '',
  password2: ''
})

function setMsg(type, text) {
  message.value = {type, text}
  setTimeout(() => {
    if (message.value?.text === text) message.value = null
  }, 5000)
}

async function submitRegister() {
  message.value = null

  if (form.value.password !== form.value.password2) {
    setMsg('error', 'Пароли не совпадают')
    return
  }

  submitting.value = true
  try {
    const res = await fetch('/api/register', {
      method: 'POST',
      headers: {'Content-Type': 'application/json', 'Accept': 'application/json'},
      body: JSON.stringify(form.value)
    })

    if (res.status === 400) {
      setMsg('error', 'Проверьте данные')
      return
    }
    if (res.status === 500) {
      setMsg('error', 'Что-то пошло не так')
      return
    }
    if (!res.ok) {
      let text = 'Ошибка'
      try {
        text = (await res.json())?.message || (await res.text()) || text
      } catch {
      }
      setMsg('error', text)
      return
    }

    setMsg('success', 'Регистрация успешна')
    // window.location.href = '/';
  } catch (e) {
    setMsg('error', 'Что-то пошло не так')
  } finally {
    submitting.value = false
  }
}

function noop() {
}
</script>

<style scoped>
.page {
  min-height: 100vh;
  background: #0b0d12;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 24px;
}

.card {
  width: 520px;
  max-width: 92vw;
  background: #1d2229;
  border-radius: 12px;
  box-shadow: 0 8px 28px rgba(0, 0, 0, .45);
  padding: 36px 40px 28px;
  border: 1px solid rgba(255, 255, 255, .06);
}

.logo {
  width: 100px;
  height: 56px;
  margin: 0 auto 16px;
}

.logo svg {
  fill: #e07a33;
  width: 100%;
  height: 100%
}

.title {
  text-align: center;
  color: #e6e7eb;
  font-weight: 800;
  line-height: 1.25;
  margin: 0 0 18px;
  font-size: 28px;
}

.brand {
  display: inline-block
}

.tabs {
  display: flex;
  gap: 8px;
  background: #20262e;
  border-radius: 8px;
  padding: 4px;
  margin: 8px 0 18px;
}

.tab {
  flex: 1;
  background: transparent;
  color: #c7cbd3;
  border: none;
  padding: 10px 0;
  border-radius: 6px;
  cursor: pointer;
  font-weight: 600;
}

.tab.active {
  background: #0f141a;
  color: #e7e9ee;
  box-shadow: inset 0 0 0 1px rgba(255, 255, 255, .06);
}

.form {
  display: flex;
  flex-direction: column;
  gap: 10px
}

.label {
  color: #c7cbd3;
  font-size: 13px;
  margin-top: 6px
}

.input {
  display: flex;
  align-items: center;
  background: #0f141a;
  border: 1px solid rgba(255, 255, 255, .08);
  border-radius: 6px;
  padding: 8px 10px;
  gap: 8px;
}

.input input {
  background: transparent;
  border: none;
  outline: none;
  color: #e7e9ee;
  width: 100%;
  height: 28px;
}

.btn {
  margin-top: 10px;
  background: #0f141a;
  color: #e7e9ee;
  border: 1px solid rgba(255, 255, 255, .1);
  border-radius: 6px;
  padding: 10px 12px;
  cursor: pointer;
  font-weight: 700;
}

.btn:disabled {
  opacity: .6;
  cursor: not-allowed
}

.link {
  margin: 10px auto 0;
  color: #e07a33;
  text-decoration: none;
  font-size: 14px;
}

.footer {
  margin-top: 20px;
  color: #a9afbb;
  font-size: 12px;
  opacity: .8;
}

.heart {
  color: #7f5af0
}

.grid {
  display: grid;
  grid-template-columns:1fr 1fr;
  gap: 10px
}

.alert {
  margin-top: 6px;
  padding: 10px 12px;
  border-radius: 6px;
  font-size: 14px;
}

.alert.error {
  background: rgba(224, 61, 61, .12);
  color: #ffb3b3;
  border: 1px solid rgba(224, 61, 61, .35)
}

.alert.success {
  background: rgba(56, 173, 80, .12);
  color: #b6f8c8;
  border: 1px solid rgba(56, 173, 80, .35)
}

/* иконки */
svg.icon {
  width: 18px;
  height: 18px;
  opacity: .85;
  flex-shrink: 0
}
</style>