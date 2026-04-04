const TELEGRAM_ID_QUERY_KEYS = ['tg_id', 'tgid', 'token', 'user_id']

function getFirstQueryValue(keys) {
  if (typeof window === 'undefined') {
    return ''
  }

  const searchParams = new URLSearchParams(window.location.search)

  for (const key of keys) {
    const value = searchParams.get(key)

    if (value) {
      return value
    }
  }

  return ''
}

export function getTelegramContext() {
  const telegramUser = window.Telegram?.WebApp?.initDataUnsafe?.user
  const fallbackId = getFirstQueryValue(TELEGRAM_ID_QUERY_KEYS) || import.meta.env.VITE_TG_ID?.trim() || ''
  const fallbackUsername = getFirstQueryValue(['username']) || import.meta.env.VITE_TG_USERNAME?.trim() || ''
  const tgid = telegramUser?.id ?? fallbackId

  return {
    tgid: tgid === undefined || tgid === null || tgid === '' ? '' : String(tgid),
    username: telegramUser?.username ?? fallbackUsername
  }
}
