function isAbsoluteUrl(url = '') {
  return /^https?:\/\//i.test(url)
}

function normalizeBaseUrl(baseURL = '') {
  return baseURL.endsWith('/') ? baseURL.slice(0, -1) : baseURL
}

function resolveUrl(baseURL, url = '') {
  if (isAbsoluteUrl(url) || !baseURL) {
    return url
  }

  const normalizedBaseUrl = normalizeBaseUrl(baseURL)
  return url.startsWith('/') ? `${normalizedBaseUrl}${url}` : `${normalizedBaseUrl}/${url}`
}

function isPlainObject(value) {
  return Object.prototype.toString.call(value) === '[object Object]'
}

function appendSearchParams(url = '', params) {
  if (!isPlainObject(params)) {
    return url
  }

  const searchParams = new URLSearchParams()

  Object.entries(params).forEach(([key, value]) => {
    if (value === undefined || value === null || value === '') {
      return
    }

    if (Array.isArray(value)) {
      value.forEach((item) => {
        if (item === undefined || item === null || item === '') {
          return
        }

        searchParams.append(key, String(item))
      })

      return
    }

    searchParams.append(key, String(value))
  })

  const queryString = searchParams.toString()
  if (!queryString) {
    return url
  }

  return `${url}${url.includes('?') ? '&' : '?'}${queryString}`
}

async function parseResponseBody(response) {
  if (response.status === 204) {
    return null
  }

  const responseText = await response.text()
  if (!responseText) {
    return null
  }

  try {
    return JSON.parse(responseText)
  } catch {
    return responseText
  }
}

function createAxiosError(message, config, response) {
  const error = new Error(message)
  error.name = 'AxiosError'
  error.config = config
  error.response = response
  error.isAxiosError = true

  return error
}

function createInstance(defaultConfig = {}) {
  const instance = {
    defaults: defaultConfig,
    async request(config = {}) {
      const finalConfig = {
        ...defaultConfig,
        ...config,
        headers: {
          ...(defaultConfig.headers ?? {}),
          ...(config.headers ?? {})
        }
      }

      const resolvedUrl = resolveUrl(finalConfig.baseURL, finalConfig.url)
      const url = appendSearchParams(resolvedUrl, finalConfig.params)
      const method = (finalConfig.method ?? 'GET').toUpperCase()
      const headers = { ...finalConfig.headers }
      let body = finalConfig.data

      if (isPlainObject(body)) {
        headers['Content-Type'] = headers['Content-Type'] ?? 'application/json'
        body = JSON.stringify(body)
      }

      let rawResponse

      try {
        rawResponse = await fetch(url, {
          method,
          headers,
          body
        })
      } catch (error) {
        throw createAxiosError(error.message || 'Network Error', finalConfig)
      }

      const data = await parseResponseBody(rawResponse)
      const response = {
        data,
        status: rawResponse.status,
        headers: Object.fromEntries(rawResponse.headers.entries()),
        config: finalConfig
      }

      if (!rawResponse.ok) {
        throw createAxiosError(`Request failed with status ${rawResponse.status}`, finalConfig, response)
      }

      return response
    },
    get(url, config = {}) {
      return instance.request({
        ...config,
        method: 'GET',
        url
      })
    }
  }

  return instance
}

const defaultInstance = createInstance()

const axios = {
  create(config = {}) {
    return createInstance(config)
  },
  request(config = {}) {
    return defaultInstance.request(config)
  },
  get(url, config = {}) {
    return defaultInstance.get(url, config)
  },
  isAxiosError(error) {
    return Boolean(error?.isAxiosError)
  }
}

export default axios
