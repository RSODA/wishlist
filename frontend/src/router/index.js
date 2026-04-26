import { reactive } from 'vue'

function normalizeUsername(value) {
  return String(value ?? '').trim().replace(/^@+/, '')
}

function parseQuery(search = '') {
  const searchParams = new URLSearchParams(search)

  return {
    username: normalizeUsername(searchParams.get('username') ?? ''),
    tgId: String(searchParams.get('tg_id') ?? '').trim()
  }
}

function buildSearch(query = {}) {
  const searchParams = new URLSearchParams()
  const preservedQuery = parseQuery(window.location.search)
  const tgId = String(query.tgId ?? preservedQuery.tgId ?? '').trim()
  const username = normalizeUsername(query.username ?? preservedQuery.username)

  if (tgId) {
    searchParams.set('tg_id', tgId)
  }

  if (username) {
    searchParams.set('username', username)
  }

  const queryString = searchParams.toString()
  return queryString ? `?${queryString}` : ''
}

function resolveRoute(location = window.location) {
  const pathname = location.pathname || '/'
  const wishlistMatch = pathname.match(/^\/wishlist\/([^/]+)$/)

  if (wishlistMatch) {
    return {
      name: 'wishlist-view',
      params: {
        targetTgId: decodeURIComponent(wishlistMatch[1])
      },
      query: parseQuery(location.search),
      meta: {}
    }
  }

  return {
    name: 'subscriptions',
    params: {},
    query: parseQuery(location.search),
    meta: {
      keepAlive: true
    }
  }
}

export const routeState = reactive(resolveRoute())

function syncRoute() {
  const nextRoute = resolveRoute()

  routeState.name = nextRoute.name
  routeState.params = nextRoute.params
  routeState.query = nextRoute.query
  routeState.meta = nextRoute.meta
}

function buildUrl(location) {
  if (typeof location === 'string') {
    return location
  }

  if (location?.name === 'wishlist-view') {
    const targetTgId = encodeURIComponent(String(location.params?.targetTgId ?? ''))
    const search = buildSearch(location.query)

    return `/wishlist/${targetTgId}${search}`
  }

  return `/${buildSearch(location.query)}`
}

const router = {
  push(location) {
    const nextUrl = buildUrl(location)
    window.history.pushState({}, '', nextUrl)
    syncRoute()
  },
  install(app) {
    window.addEventListener('popstate', syncRoute)

    app.config.globalProperties.$router = router

    Object.defineProperty(app.config.globalProperties, '$route', {
      get() {
        return routeState
      }
    })
  }
}

export default router
