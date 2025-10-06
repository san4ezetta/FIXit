export async function authFetch(url, options = {}) {
    const token = localStorage.getItem('token')
    const headers = new Headers(options.headers || {})
    if (token) headers.set('Authorization', `Bearer ${token}`)
    // JSON по умолчанию
    if (!headers.has('Accept')) headers.set('Accept', 'application/json')
    const res = await fetch(url, { ...options, headers })
    return res
}