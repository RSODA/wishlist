<script>
import LoadingOverlay from '../components/wishlist/LoadingOverlay.vue'
import AddSubscriptionModal from '../components/wishlist/AddSubscriptionModal.vue'
import WishlistHeader from '../components/wishlist/WishlistHeader.vue'
import SubscriptionsBlock from '../components/wishlist/SubscriptionsBlock.vue'
import UserTgid from '../components/wishlist/UserTgid.vue'
import { acceptSubscription, createUser, getSub, subscribeToUser } from '../api/userApi'
import { getTelegramContext } from '../utils/telegram'

export default {
  name: 'SubscriptionsPage',
  components: {
    LoadingOverlay,
    AddSubscriptionModal,
    WishlistHeader,
    SubscriptionsBlock,
    UserTgid
  },
  data() {
    const telegramContext = getTelegramContext()

    return {
      isLoading: false,
      tgid: telegramContext.tgid,
      username: telegramContext.username,
      acceptedSubscriptions: [],
      pendingSubscriptions: [],
      sentSubscriptions: [],
      loadError: '',
      isAddUserModalOpen: false,
      isSubmittingSubscription: false,
      acceptingSubscriptionTgId: '',
      addSubscriptionError: '',
      isInitialized: false
    }
  },
  computed: {
    displayTgid() {
      return this.tgid || '—'
    },
    subscriptions() {
      return [
        ...this.acceptedSubscriptions,
        ...this.pendingSubscriptions,
        ...this.sentSubscriptions
      ]
    }
  },
  async activated() {
    if (!this.isInitialized) {
      await this.initializeApp()
    }
  },
  async mounted() {
    if (!this.isInitialized) {
      await this.initializeApp()
    }
  },
  methods: {
    normalizeUsername(value) {
      return String(value ?? '').trim().replace(/^@+/, '')
    },
    getRegistrationUsername() {
      const normalizedUsername = this.normalizeUsername(this.username)

      return normalizedUsername || `user_${this.tgid}`
    },
    isUserAlreadyExistsError(error) {
      return error?.status === 409 || error?.message === 'user is exist'
    },
    buildSubscriptionKey(subscription) {
      const subscriptionTgId = String(subscription?.tgId ?? '').trim()
      if (subscriptionTgId) {
        return `tgid:${subscriptionTgId}`
      }

      const normalizedUsername = this.normalizeUsername(subscription?.username).toLowerCase()
      return normalizedUsername ? `username:${normalizedUsername}` : ''
    },
    sanitizeSubscriptions(subscriptions = [], { excludeCurrentUser = false } = {}) {
      const currentTgId = String(this.tgid ?? '').trim()
      const currentUsername = this.normalizeUsername(this.username).toLowerCase()

      return subscriptions.filter((subscription) => {
        if (!subscription) {
          return false
        }

        const subscriptionTgId = String(subscription.tgId ?? '').trim()
        const normalizedUsername = this.normalizeUsername(subscription.username)

        if (!subscriptionTgId && !normalizedUsername) {
          return false
        }

        if (!excludeCurrentUser) {
          return true
        }

        const matchesCurrentTgId = currentTgId && subscriptionTgId === currentTgId
        const matchesCurrentUsername =
          currentUsername && normalizedUsername.toLowerCase() === currentUsername

        return !matchesCurrentTgId && !matchesCurrentUsername
      })
    },
    dedupeSubscriptions(subscriptions = []) {
      const uniqueSubscriptions = new Map()

      subscriptions.forEach((subscription) => {
        const key = this.buildSubscriptionKey(subscription)
        if (!key) {
          return
        }

        const existingSubscription = uniqueSubscriptions.get(key)
        if (!existingSubscription) {
          uniqueSubscriptions.set(key, { ...subscription })
          return
        }

        uniqueSubscriptions.set(key, {
          ...existingSubscription,
          ...subscription,
          tgId: subscription.tgId || existingSubscription.tgId,
          username: subscription.username || existingSubscription.username,
          isAccepted: subscription.isAccepted ?? existingSubscription.isAccepted
        })
      })

      return Array.from(uniqueSubscriptions.values())
    },
    getSubscriptionByTgId(tgId) {
      return this.subscriptions.find((subscription) => String(subscription.tgId) === String(tgId))
    },
    openAddUserModal() {
      if (!this.tgid) {
        return
      }

      this.addSubscriptionError = ''
      this.isAddUserModalOpen = true
    },
    closeAddUserModal() {
      if (this.isSubmittingSubscription) {
        return
      }

      this.addSubscriptionError = ''
      this.isAddUserModalOpen = false
    },
    hasSubscription(username) {
      const normalizedUsername = this.normalizeUsername(username).toLowerCase()

      return this.subscriptions.some((subscription) => {
        return this.normalizeUsername(subscription.username).toLowerCase() === normalizedUsername
      })
    },
    getAddSubscriptionErrorMessage(error, username) {
      if (error?.message === 'user to subscription not found') {
        return `Пользователь @${username} не найден`
      }

      if (error?.message === 'invalid argument') {
        return 'Укажите корректный username пользователя'
      }

      if (error?.message === 'subscribe error') {
        return `Не удалось добавить @${username}. Попробуйте ещё раз.`
      }

      return error?.message || 'Не удалось добавить пользователя'
    },
    getAcceptSubscriptionErrorMessage(error, username, tgId) {
      if (error?.message === 'user not found') {
        return username
          ? `Не удалось найти заявку от @${username}`
          : `Не удалось найти заявку с TG ID ${tgId}`
      }

      if (error?.message === 'invalid argument(-s)') {
        return 'Не удалось подтвердить подписку из-за некорректных данных'
      }

      if (error?.message === 'err accepted sub') {
        return username
          ? `Не удалось подтвердить @${username}. Попробуйте ещё раз.`
          : 'Не удалось подтвердить подписку. Попробуйте ещё раз.'
      }

      return error?.message || 'Не удалось подтвердить подписку'
    },
    goToWishlist(subscription) {
      const targetTgId = String(subscription?.tgId ?? '').trim()
      if (!targetTgId) {
        return
      }

      const normalizedUsername = this.normalizeUsername(subscription?.username)

      this.$router.push({
        name: 'wishlist-view',
        params: {
          targetTgId
        },
        query: normalizedUsername ? { username: normalizedUsername } : {}
      })
    },
    goToOwnWishlist() {
      if (!this.tgid) {
        return
      }

      const normalizedUsername = this.normalizeUsername(this.username)

      this.$router.push({
        name: 'wishlist-view',
        params: {
          targetTgId: this.tgid
        },
        query: normalizedUsername ? { username: normalizedUsername } : {}
      })
    },
    async initializeApp() {
      this.isInitialized = true

      if (!this.tgid) {
        this.loadError =
          'Не удалось определить Telegram ID. Для локальной проверки можно открыть страницу с параметром ?tg_id=123456.'
        return
      }

      this.isLoading = true
      this.loadError = ''

      let registrationError = null

      try {
        await this.registerUser()
      } catch (error) {
        registrationError = error
        console.error('Failed to register user', error)
      }

      try {
        await this.fetchSubscriptions()
      } catch (error) {
        console.error('Failed to load subscriptions', error)
        this.acceptedSubscriptions = []
        this.pendingSubscriptions = []
        this.sentSubscriptions = []
        this.loadError = error.message || registrationError?.message || 'Не удалось загрузить подписки'
      } finally {
        this.isLoading = false
      }
    },
    async registerUser() {
      try {
        await createUser({
          tgId: this.tgid,
          username: this.getRegistrationUsername()
        })
      } catch (error) {
        if (this.isUserAlreadyExistsError(error)) {
          return
        }

        throw error
      }
    },
    async fetchSubscriptions() {
      this.loadError = ''

      const response = await getSub({
        tgId: this.tgid,
        token: this.tgid
      })

      const incomingSubscriptions = this.sanitizeSubscriptions(response?.followerTo ?? [], {
        excludeCurrentUser: true
      })
      const outgoingSubscriptions = this.sanitizeSubscriptions(response?.followerFrom ?? [], {
        excludeCurrentUser: true
      })

      this.acceptedSubscriptions = this.dedupeSubscriptions(
        outgoingSubscriptions.filter((subscription) => subscription.isAccepted)
      )

      this.pendingSubscriptions = this.dedupeSubscriptions(
        incomingSubscriptions.filter((subscription) => !subscription.isAccepted)
      )

      this.sentSubscriptions = this.dedupeSubscriptions(
        outgoingSubscriptions.filter((subscription) => !subscription.isAccepted)
      )

      const nextUsername = this.normalizeUsername(response?.username)
      if (nextUsername) {
        this.username = nextUsername
      }
    },
    async handleAddSubscription(username) {
      const normalizedUsername = this.normalizeUsername(username)
      const currentUsername = this.normalizeUsername(this.username).toLowerCase()

      if (!normalizedUsername) {
        this.addSubscriptionError = 'Введите username пользователя'
        return
      }

      if (currentUsername && normalizedUsername.toLowerCase() === currentUsername) {
        this.addSubscriptionError = 'Нельзя подписаться на самого себя'
        return
      }

      if (this.hasSubscription(normalizedUsername)) {
        this.addSubscriptionError = `@${normalizedUsername} уже есть в вашем списке`
        return
      }

      this.isSubmittingSubscription = true
      this.addSubscriptionError = ''

      try {
        await subscribeToUser({
          toUsername: normalizedUsername,
          token: this.tgid
        })

        await this.fetchSubscriptions()
        this.isAddUserModalOpen = false
      } catch (error) {
        console.error('Failed to add subscription', error)
        this.addSubscriptionError = this.getAddSubscriptionErrorMessage(error, normalizedUsername)
      } finally {
        this.isSubmittingSubscription = false
      }
    },
    async handleAcceptSubscription(subscriptionTgId) {
      const subscription = this.getSubscriptionByTgId(subscriptionTgId)
      const normalizedUsername = this.normalizeUsername(subscription?.username)

      this.acceptingSubscriptionTgId = String(subscriptionTgId)
      this.loadError = ''

      try {
        await acceptSubscription({
          acceptedTgId: subscriptionTgId,
          token: this.tgid
        })

        await this.fetchSubscriptions()
      } catch (error) {
        console.error('Failed to accept subscription', error)
        this.loadError = this.getAcceptSubscriptionErrorMessage(error, normalizedUsername, subscriptionTgId)
      } finally {
        this.acceptingSubscriptionTgId = ''
      }
    }
  }
}
</script>

<template>
  <div class="app-shell">
    <LoadingOverlay v-if="isLoading" />

    <WishlistHeader
      :can-add-user="Boolean(tgid)"
      :can-open-own-wishlist="Boolean(tgid)"
      @open-add-user="openAddUserModal"
      @open-own-wishlist="goToOwnWishlist"
    />

    <main class="content-area">
      <SubscriptionsBlock
        :accepted-subscriptions="acceptedSubscriptions"
        :pending-subscriptions="pendingSubscriptions"
        :sent-subscriptions="sentSubscriptions"
        :error="loadError"
        :username="username"
        :accepting-subscription-tg-id="acceptingSubscriptionTgId"
        @accept-subscription="handleAcceptSubscription"
        @open-wishlist="goToWishlist"
      />
      <UserTgid :tgid="displayTgid" :username="username" />
    </main>

    <AddSubscriptionModal
      v-if="isAddUserModalOpen"
      :error="addSubscriptionError"
      :is-submitting="isSubmittingSubscription"
      @close="closeAddUserModal"
      @submit="handleAddSubscription"
    />
  </div>
</template>

<style scoped>
.app-shell {
  min-height: 100vh;
  padding: calc(16px + env(safe-area-inset-top)) 16px calc(20px + env(safe-area-inset-bottom));
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.content-area {
  display: flex;
  flex-direction: column;
  gap: 14px;
}
</style>
