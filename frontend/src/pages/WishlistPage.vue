<script>
import LoadingOverlay from '../components/wishlist/LoadingOverlay.vue'
import AddWishModal from '../components/wishlist/AddWishModal.vue'
import WishlistViewHeader from '../components/wishlist/WishlistViewHeader.vue'
import WishlistItemsBlock from '../components/wishlist/WishlistItemsBlock.vue'
import UserTgid from '../components/wishlist/UserTgid.vue'
import { createWish, getWishlist } from '../api/wishlistApi'
import { getTelegramContext } from '../utils/telegram'

export default {
  name: 'WishlistPage',
  components: {
    LoadingOverlay,
    AddWishModal,
    WishlistViewHeader,
    WishlistItemsBlock,
    UserTgid
  },
  data() {
    const telegramContext = getTelegramContext()

    return {
      isLoading: false,
      tgId: telegramContext.tgid,
      username: telegramContext.username,
      items: [],
      loadError: '',
      lastLoadedRouteKey: '',
      isAddWishModalOpen: false,
      isSubmittingWish: false,
      addWishError: ''
    }
  },
  computed: {
    targetTgId() {
      return String(this.$route.params.targetTgId ?? '').trim()
    },
    targetUsername() {
      return String(this.$route.query.username ?? '').trim().replace(/^@+/, '')
    },
    displayTgId() {
      return this.tgId || '—'
    },
    isOwnWishlist() {
      return Boolean(this.tgId) && this.targetTgId === this.tgId
    },
    routeKey() {
      return `${this.$route.name}:${this.targetTgId}:${this.targetUsername}`
    }
  },
  async mounted() {
    await this.fetchWishlist()
  },
  async updated() {
    if (this.routeKey && this.routeKey !== this.lastLoadedRouteKey) {
      await this.fetchWishlist()
    }
  },
  methods: {
    async fetchWishlist() {
      if (!this.tgId) {
        this.items = []
        this.loadError =
          'Не удалось определить Telegram ID. Для локальной проверки можно открыть страницу с параметром ?tg_id=123456.'
        this.lastLoadedRouteKey = this.routeKey
        return
      }

      if (!this.targetTgId) {
        this.items = []
        this.loadError = 'Не удалось определить владельца wishlist.'
        this.lastLoadedRouteKey = this.routeKey
        return
      }

      this.isLoading = true
      this.loadError = ''

      try {
        this.items = await getWishlist({
          fromTgId: this.tgId,
          toTgId: this.targetTgId,
          offset: 0
        })
        this.lastLoadedRouteKey = this.routeKey
      } catch (error) {
        console.error('Failed to load wishlist', error)
        this.items = []
        this.loadError = this.getWishlistErrorMessage(error)
        this.lastLoadedRouteKey = this.routeKey
      } finally {
        this.isLoading = false
      }
    },
    getWishlistErrorMessage(error) {
      if (error?.message === 'sub not confirmed' || error?.status === 403 || error?.grpcCode === 7) {
        return 'Этот wishlist пока недоступен: подписка ещё не подтверждена.'
      }

      if (error?.message === 'invalid id' || error?.grpcCode === 3) {
        return 'Не удалось открыть wishlist из-за некорректного идентификатора пользователя.'
      }

      if (error?.status === 404) {
        return 'Wishlist не найден.'
      }

      return error?.message || 'Не удалось загрузить wishlist.'
    },
    getAddWishErrorMessage(error) {
      if (error?.status === 400 || error?.grpcCode === 3 || error?.message === 'err invalid argument') {
        return 'Проверь поля: название, цена, ссылка и изображение должны быть заполнены корректно.'
      }

      if (error?.status === 404 || error?.message === 'user not found') {
        return 'Не удалось найти пользователя для этого wishlist.'
      }

      return error?.message || 'Не удалось добавить желание.'
    },
    openAddWishModal() {
      if (!this.isOwnWishlist || this.isSubmittingWish) {
        return
      }

      this.addWishError = ''
      this.isAddWishModalOpen = true
    },
    closeAddWishModal() {
      if (this.isSubmittingWish) {
        return
      }

      this.addWishError = ''
      this.isAddWishModalOpen = false
    },
    async handleAddWish(form) {
      const title = String(form?.title ?? '').trim()
      const price = Number(form?.price)
      const url = String(form?.url ?? '').trim()
      const picture = String(form?.picture ?? '').trim()

      if (!title) {
        this.addWishError = 'Укажи название желания.'
        return
      }

      if (!Number.isFinite(price) || price <= 0) {
        this.addWishError = 'Укажи корректную цену.'
        return
      }

      if (!url) {
        this.addWishError = 'Добавь ссылку на товар.'
        return
      }

      if (!picture) {
        this.addWishError = 'Прикрепи изображение.'
        return
      }

      this.isSubmittingWish = true
      this.addWishError = ''

      try {
        await createWish({
          tgId: this.tgId,
          title,
          price,
          url,
          picture
        })

        await this.fetchWishlist()
        this.isAddWishModalOpen = false
      } catch (error) {
        console.error('Failed to create wish', error)
        this.addWishError = this.getAddWishErrorMessage(error)
      } finally {
        this.isSubmittingWish = false
      }
    },
    goBack() {
      this.$router.push({ name: 'subscriptions' })
    }
  }
}
</script>

<template>
  <div class="page-shell">
    <LoadingOverlay v-if="isLoading" />

    <WishlistViewHeader
      :owner-tg-id="targetTgId"
      :owner-username="targetUsername"
      :can-add-wish="isOwnWishlist"
      :is-busy="isSubmittingWish"
      @back="goBack"
      @open-add-wish="openAddWishModal"
    />

    <main class="content-area">
      <WishlistItemsBlock
        :items="items"
        :error="loadError"
        :owner-username="targetUsername"
      />
      <UserTgid :tgid="displayTgId" :username="username" />
    </main>

    <AddWishModal
      v-if="isAddWishModalOpen"
      :error="addWishError"
      :is-submitting="isSubmittingWish"
      @close="closeAddWishModal"
      @submit="handleAddWish"
    />
  </div>
</template>

<style scoped>
.page-shell {
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
