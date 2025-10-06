import { createRouter, createWebHistory } from 'vue-router'
import AuthPage from '../pages/AuthPage.vue'
import UserPage from '../pages/UserPage.vue'

// простая проверка токена; дополнительно можно проверить exp
function isAuthed() {
    const t = localStorage.getItem('token')
    if (!t) return false

    // Необязательная проверка exp из JWT (если токен — JWT):
    try {
        const payload = JSON.parse(atob(t.split('.')[1] || ''))
        if (payload && typeof payload.exp === 'number') {
            const now = Math.floor(Date.now() / 1000)
            if (payload.exp < now) return false
        }
    } catch (_) {
        // если не JWT — просто считаем наличие валидным
    }
    return true
}

const routes = [
    { path: '/', name: 'auth', component: AuthPage, meta: { guestOnly: true } },
    { path: '/user/:id', name: 'user', component: UserPage, props: true, meta: { requiresAuth: true } },
    { path: '/settings', name: 'settings', component: () => import('../pages/SettingsStub.vue'), meta: { requiresAuth: true } },
    { path: '/:pathMatch(.*)*', redirect: '/' },
]

const router = createRouter({
    history: createWebHistory(),
    routes,
    scrollBehavior: () => ({ top: 0 }),
})

router.beforeEach((to) => {
    const authed = isAuthed()

    if (to.meta.requiresAuth && !authed) {
        return { name: 'auth', replace: true }
    }

    if (to.meta.guestOnly && authed) {
        const uid = localStorage.getItem('user_id') || 'me'
        return { name: 'user', params: { id: uid }, replace: true }
    }
})

export default router