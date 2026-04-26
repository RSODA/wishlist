import axios from '../lib/axios'

function normalizeBaseUrl(baseUrl) {
  return baseUrl.endsWith('/') ? baseUrl.slice(0, -1) : baseUrl
}

function getApiBaseUrl() {
  const configuredWishlistBaseUrl = import.meta.env.VITE_WISHLIST_API_BASE_URL?.trim()
  if (configuredWishlistBaseUrl) {
    return normalizeBaseUrl(configuredWishlistBaseUrl)
  }

  if (import.meta.env.DEV) {
    const devWishlistProxyTarget = import.meta.env.VITE_DEV_WISH_PROXY_TARGET?.trim()
    return normalizeBaseUrl(devWishlistProxyTarget || 'http://localhost:8081')
  }

  const configuredBaseUrl = import.meta.env.VITE_API_BASE_URL?.trim()

  return configuredBaseUrl ? normalizeBaseUrl(configuredBaseUrl) : undefined
}

const httpClient = axios.create({
  baseURL: getApiBaseUrl(),
  headers: {
    Accept: 'application/json'
  }
})

function createApiError(status, payload) {
  const message =
    typeof payload === 'object' && payload?.message
      ? payload.message
      : `Request failed with status ${status}`

  const error = new Error(message)
  error.status = status
  error.payload = payload

  return error
}

async function request(path, { method = 'GET', body, params } = {}) {
  try {
    const response = await httpClient.request({
      url: path,
      method,
      data: body,
      params
    })

    if (response.status === 204) {
      return null
    }

    return response.data ?? null
  } catch (error) {
    if (!axios.isAxiosError(error) || !error.response) {
      throw new Error('Не удалось связаться с сервисом wishlist')
    }

    throw createApiError(error.response.status, error.response.data)
  }
}

function normalizePrice(value) {
  if (value === undefined || value === null || value === '') {
    return null
  }

  const parsedValue = Number(value)
  return Number.isFinite(parsedValue) ? parsedValue : null
}

function normalizeWishlistItem(wish = {}) {
  const status = wish.status ?? {}

  return {
    id: String(wish.id ?? ''),
    wishlistTgId: String(wish.tgId ?? wish.tg_id ?? ''),
    title: wish.title ?? '',
    price: normalizePrice(wish.price),
    url: wish.url ?? '',
    picture: wish.picture ?? '',
    statusLabel: status.status ?? status.title ?? '',
    statusOwnerTgId: String(status.tgId ?? status.tg_id ?? '')
  }
}

export async function getWishlist({ fromTgId, toTgId, offset = 0 } = {}) {
  const payload = await request(`/api/v1/users/${encodeURIComponent(String(fromTgId))}/wishes`, {
    params: {
      to_tg_id: String(toTgId),
      offset,
    }
  })

  const items = payload?.wishs ?? payload?.wishes ?? []

  return Array.isArray(items) ? items.map(normalizeWishlistItem) : []
}

export async function createWish({ tgId, title, price, url, picture } = {}) {
  const payload = await request(`/api/v1/users/${encodeURIComponent(String(tgId))}/wishes`, {
    method: 'POST',
    body: {
      tgId: String(tgId),
      title,
      price: String(price),
      url,
      picture
    }
  })

  return {
    id: String(payload?.id ?? '')
  }
}

export default {
  createWish,
  getWishlist
}
