<script>
import LoadingOverlay from './components/wishlist/LoadingOverlay.vue'
import WishlistHeader from './components/wishlist/WishlistHeader.vue'
import SubscriptionsBlock from './components/wishlist/SubscriptionsBlock.vue'
import UserTgid from './components/wishlist/UserTgid.vue'

export default {
  name: 'App',
  components: {
    LoadingOverlay,
    WishlistHeader,
    SubscriptionsBlock,
    UserTgid
  },
  data() {
    return {
      isLoading: false,
      tgid: this.getTelegramUserId()
    }
  },
  methods: {
    getTelegramUserId() {
      const userId = window.Telegram?.WebApp?.initDataUnsafe?.user?.id
      return userId ?? '—'
    }
  }
}
</script>

<template>
  <div class="app-shell">
    <LoadingOverlay v-if="isLoading" />

    <WishlistHeader />

    <main class="content-area">
      <SubscriptionsBlock />
      <UserTgid :tgid="tgid" />
    </main>
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
