import axios from '../lib/axios'

const USER_API_PREFIX = '/api/v1/user'

function normalizeBaseUrl(baseUrl) {
  return baseUrl.endsWith('/') ? baseUrl.slice(0, -1) : baseUrl
}

function getApiBaseUrl() {
  const configuredBaseUrl = import.meta.env.VITE_API_BASE_URL?.trim()

  return configuredBaseUrl ? normalizeBaseUrl(configuredBaseUrl) : undefined
}

const httpClient = axios.create({
  baseURL: getApiBaseUrl(),
  headers: {
    Accept: 'application/json'
  }
})

function hasOwnProperty(target, key) {
  return Object.prototype.hasOwnProperty.call(target, key)
}

function normalizeBoolean(value) {
  if (typeof value === 'string') {
    const normalizedValue = value.trim().toLowerCase()

    if (normalizedValue === 'true') {
      return true
    }

    if (normalizedValue === 'false') {
      return false
    }
  }

  return Boolean(value)
}

function buildAuthorizationHeader(token, mode = 'bearer') {
  if (token === undefined || token === null) {
    return null
  }

  const tokenValue = String(token).trim()
  if (!tokenValue) {
    return null
  }

  return mode === 'raw' ? tokenValue : `Bearer ${tokenValue}`
}

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

async function request(
  path,
  { method = 'GET', token, authorizationMode = 'bearer', body, params } = {}
) {
  const headers = {}

  const authorizationHeader = buildAuthorizationHeader(token, authorizationMode)
  if (authorizationHeader) {
    headers.Authorization = authorizationHeader
  }

  try {
    const response = await httpClient.request({
      url: path,
      method,
      headers,
      data: body,
      params
    })

    if (response.status === 204) {
      return null
    }

    return response.data ?? null
  } catch (error) {
    if (!axios.isAxiosError(error) || !error.response) {
      throw new Error('Не удалось связаться с Go-сервером')
    }

    throw createApiError(error.response.status, error.response.data)
  }
}

function normalizeSubscription(subscription = {}) {
  const hasAcceptedValue =
    hasOwnProperty(subscription, 'isAccepted') || hasOwnProperty(subscription, 'is_accepted')

  return {
    tgId: String(subscription.tgId ?? subscription.tg_id ?? ''),
    username: subscription.username ?? '',
    isAccepted: hasAcceptedValue
      ? normalizeBoolean(subscription.isAccepted ?? subscription.is_accepted)
      : false
  }
}

function normalizeSubscriptions(subscriptions) {
  return Array.isArray(subscriptions) ? subscriptions.map(normalizeSubscription) : []
}

export async function getSub({ tgId, token } = {}) {
  const payload = await request(`${USER_API_PREFIX}/${encodeURIComponent(String(tgId))}`, {
    token
  })

  return {
    username: payload?.username ?? '',
    followerTo: normalizeSubscriptions(payload?.followerTo ?? payload?.follower_to),
    followerFrom: normalizeSubscriptions(payload?.followerFrom ?? payload?.follower_from)
  }
}

export async function createUser({ tgId, username }) {
  await request(USER_API_PREFIX, {
    method: 'POST',
    body: {
      tgId: String(tgId),
      username
    }
  })
}

export async function subscribeToUser({ toUsername, token }) {
  await request(`${USER_API_PREFIX}/subscribe`, {
    method: 'POST',
    token,
    authorizationMode: 'bearer',
    body: {
      toUsername
    }
  })
}

export async function acceptSubscription({ acceptedTgId, token }) {
  await request(`${USER_API_PREFIX}/accepted-sub`, {
    method: 'PUT',
    token,
    authorizationMode: 'raw',
    body: {
      acceptedTgId: String(acceptedTgId)
    }
  })
}

export default {
  getSub,
  createUser,
  subscribeToUser,
  acceptSubscription
}
