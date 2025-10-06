<template>
  <form class="form" @submit.prevent="submitRegister">
    <label class="label">Имя</label>
    <div class="input">
      <UserIcon/>
      <input v-model.trim="form.name" type="text" placeholder="Иван" required/>
    </div>

    <label class="label">Фамилия</label>
    <div class="input">
      <UserIcon/>
      <input v-model.trim="form.surname" type="text" placeholder="Иванов" required/>
    </div>

    <label class="label">Отчество</label>
    <div class="input">
      <UserIcon/>
      <input v-model.trim="form.patronymic" type="text" placeholder="Иванович"/>
    </div>

    <label class="label">Электронная почта</label>
    <div class="input">
      <MailIcon/>
      <input v-model.trim="form.email" type="email" placeholder="example@sitedefect.com" required/>
    </div>

    <label class="label">Пароль</label>
    <div class="input">
      <LockIcon/>
      <input v-model="form.password" type="password" placeholder="Введите ваш пароль" required/>
    </div>

    <label class="label">Повторите пароль</label>
    <div class="input">
      <LockIcon/>
      <input v-model="form.password2" type="password" placeholder="Повторите пароль" required/>
    </div>

    <div v-if="message" :class="['alert', message.type]">{{ message.text }}</div>

    <button class="btn" type="submit" :disabled="submitting">
      {{ submitting ? 'Отправка…' : 'Зарегистрироваться' }}
    </button>

    <a href="#" class="link" @click.prevent="$emit('switch-to-login')">Уже есть аккаунт? Войти</a>
  </form>
</template>

<script setup>
import { ref } from 'vue'
import MailIcon from './MailIcon.vue'
import LockIcon from './LockIcon.vue'
import UserIcon from './UserIcon.vue'

const submitting = ref(false)
const message = ref(null)
const form = ref({ name:'', surname:'', patronymic:'', email:'', password:'', password2:'' })

function setMsg(type, text) {
  message.value = { type, text }
  setTimeout(() => { if (message.value?.text === text) message.value = null }, 5000)
}

async function submitRegister() {
  message.value = null
  if (form.value.password !== form.value.password2) {
    setMsg('error', 'Пароли не совпадают'); return
  }

  submitting.value = true
  try {
    const res = await fetch('/api/register', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json', 'Accept': 'application/json' },
      body: JSON.stringify(form.value)
    })

    if (res.status === 400) { setMsg('error', 'Проверьте данные'); return }
    if (res.status === 500) { setMsg('error', 'Что-то пошло не так'); return }
    if (!res.ok) {
      let text = 'Ошибка'
      try { text = (await res.json())?.message || (await res.text()) || text } catch {}
      setMsg('error', text); return
    }

    setMsg('success', 'Регистрация успешна')
  } catch (e) {
    setMsg('error', 'Что-то пошло не так')
  } finally {
    submitting.value = false
  }
}

defineEmits(['switch-to-login'])
</script>

<style scoped>
.form{display:flex;flex-direction:column;gap:10px}
.label{color:#c7cbd3;font-size:13px;margin-top:6px}
.input{display:flex;align-items:center;background:#0f141a;border:1px solid rgba(255,255,255,.08);border-radius:6px;padding:8px 10px;gap:8px}
.input input{background:transparent;border:none;outline:none;color:#e7e9ee;width:100%;height:28px}
.btn{margin-top:10px;background:#0f141a;color:#e7e9ee;border:1px solid rgba(255,255,255,.1);border-radius:6px;padding:10px 12px;cursor:pointer;font-weight:700}
.btn:disabled{opacity:.6;cursor:not-allowed}
.link{margin:10px auto 0;color:#e07a33;text-decoration:none;font-size:14px}
.alert{margin-top:6px;padding:10px 12px;border-radius:6px;font-size:14px}
.alert.error{background:rgba(224,61,61,.12);color:#ffb3b3;border:1px solid rgba(224,61,61,.35)}
.alert.success{background:rgba(56,173,80,.12);color:#b6f8c8;border:1px solid rgba(56,173,80,.35)}
</style>