<template>
  <header class="hdr">
    <div class="hdr-left">
      <div class="logo"><span class="spark">*</span></div>
      <nav class="nav">
        <RouterLink to="/" class="nav-link">Панель управления</RouterLink>
        <a href="#" class="nav-link">Проекты</a>
        <a href="#" class="nav-link">Дефекты</a>
        <a href="#" class="nav-link">Пользователи</a>
        <a href="#" class="nav-link">Отчеты</a>
      </nav>
    </div>

    <div class="hdr-right">
      <div class="user" ref="menuRoot">
        <button class="avatar" @click="toggle" @keydown.down.prevent="open" aria-haspopup="menu" :aria-expanded="openMenu">
          👤
        </button>

        <transition name="fade">
          <ul v-if="openMenu" class="menu" role="menu">
            <li role="menuitem">
              <RouterLink class="mi" to="/settings" @click.native="close">Настройки</RouterLink>
            </li>
            <li role="menuitem">
              <button class="mi danger" @click="logout">Выйти</button>
            </li>
          </ul>
        </transition>
      </div>
    </div>
  </header>
</template>

<script setup>
import { ref, onMounted, onBeforeUnmount } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()
const openMenu = ref(false)
const menuRoot = ref(null)

function toggle(){ openMenu.value = !openMenu.value }
function open(){ openMenu.value = true }
function close(){ openMenu.value = false }

function onClickOutside(e){
  if (!menuRoot.value) return
  if (!menuRoot.value.contains(e.target)) close()
}
function onEsc(e){
  if (e.key === 'Escape') close()
}

function logout(){
  localStorage.removeItem('token')
  localStorage.removeItem('user_id')
  close()
  router.replace({ name: 'auth' })
}

onMounted(()=>{
  document.addEventListener('click', onClickOutside)
  document.addEventListener('keydown', onEsc)
})
onBeforeUnmount(()=>{
  document.removeEventListener('click', onClickOutside)
  document.removeEventListener('keydown', onEsc)
})
</script>

<style scoped>
.hdr{
  height:56px; display:flex; align-items:center; justify-content:space-between;
  padding:0 20px; border-bottom:1px solid rgba(255,255,255,.06); background:#0b0d12;
  position:sticky; top:0; z-index:20;
}
.hdr-left{display:flex; align-items:center; gap:14px}
.logo{
  width:28px; height:28px; border-radius:6px; display:grid; place-items:center; background:#1d2229;
  border:1px solid rgba(255,255,255,.06);
}
.spark{color:#e07a33; font-weight:900}
.nav{display:flex; gap:16px}
.nav-link{color:#c7cbd3; text-decoration:none; font-size:14px}
.nav-link.router-link-active, .nav-link:hover{color:#e7e9ee}

.hdr-right{display:flex; align-items:center; gap:10px}
.user{position:relative}
.avatar{
  width:30px; height:30px; display:grid; place-items:center;
  background:#1d2229; border:1px solid rgba(255,255,255,.06);
  border-radius:50%; cursor:pointer; color:#e7e9ee;
}

.menu{
  position:absolute; right:0; top:40px; min-width:180px;
  background:#1d2229; border:1px solid rgba(255,255,255,.08);
  border-radius:10px; padding:6px; box-shadow:0 10px 24px rgba(0,0,0,.45);
}
.mi{
  width:100%; text-align:left; background:transparent; border:none; color:#e7e9ee;
  padding:10px 10px; border-radius:8px; cursor:pointer; display:block; text-decoration:none;
}
.mi:hover{ background:#0f141a }
.mi.danger{ color:#ffb3b3 }

.fade-enter-active, .fade-leave-active{ transition:opacity .12s ease, transform .12s ease }
.fade-enter-from, .fade-leave-to{ opacity:0; transform:translateY(-4px) }
</style>