<script>
import { routeState } from './router'
import SubscriptionsPage from './pages/SubscriptionsPage.vue'
import WishlistPage from './pages/WishlistPage.vue'

export default {
  name: 'App',
  components: {
    SubscriptionsPage,
    WishlistPage
  },
  data() {
    return {
      routeState
    }
  },
  computed: {
    isWishlistPage() {
      return this.routeState.name === 'wishlist-view'
    }
  }
}
</script>

<template>
  <SubscriptionsPage v-show="!isWishlistPage" />
  <WishlistPage v-if="isWishlistPage" />
</template>
