<template>
  <section class="wishlist-block" aria-label="Wishlist">
    <div class="block-header">
      <div>
        <h2>Желания</h2>
        <p v-if="ownerUsername" class="block-meta">Для @{{ ownerUsername }}</p>
      </div>
      <span class="counter">{{ items.length }}</span>
    </div>

    <div v-if="error && items.length === 0" class="error-state">{{ error }}</div>

    <template v-else>
      <div v-if="error" class="error-banner">{{ error }}</div>

      <div v-if="items.length === 0" class="empty-state">
        В этом wishlist пока нет добавленных желаний.
      </div>

      <ul v-else class="wishlist-list">
        <li v-for="item in items" :key="item.id || item.title" class="wishlist-item">
          <WishlistCard :item="item" />
        </li>
      </ul>
    </template>
  </section>
</template>

<script>
import WishlistCard from './WishlistCard.vue'

export default {
  name: 'WishlistItemsBlock',
  components: {
    WishlistCard
  },
  props: {
    items: {
      type: Array,
      default: () => []
    },
    error: {
      type: String,
      default: ''
    },
    ownerUsername: {
      type: String,
      default: ''
    }
  }
}
</script>

<style scoped>
.wishlist-block {
  border-radius: 20px;
  background:
    radial-gradient(circle at top right, rgba(56, 189, 248, 0.14), transparent 30%),
    linear-gradient(180deg, rgba(17, 26, 44, 0.96), rgba(12, 19, 33, 0.96));
  border: 1px solid var(--border-primary);
  min-height: 260px;
  padding: 16px;
  display: flex;
  flex-direction: column;
  gap: 14px;
}

.block-header {
  display: flex;
  align-items: flex-start;
  justify-content: space-between;
  gap: 12px;
}

.block-header h2 {
  font-size: 18px;
  font-weight: 700;
  color: var(--text-primary);
}

.block-meta {
  margin-top: 4px;
  color: var(--text-secondary);
  font-size: 13px;
}

.counter {
  min-width: 32px;
  padding: 4px 10px;
  border-radius: 999px;
  background: rgba(56, 189, 248, 0.14);
  color: var(--accent-primary);
  font-size: 13px;
  font-weight: 700;
  text-align: center;
}

.wishlist-list {
  list-style: none;
  margin: 0;
  padding: 0;
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.wishlist-item {
  min-width: 0;
}

.empty-state,
.error-state {
  margin-top: 8px;
  border: 1px dashed var(--border-secondary);
  border-radius: 12px;
  min-height: 170px;
  display: grid;
  place-items: center;
  color: var(--text-secondary);
  font-size: 14px;
  text-align: center;
  padding: 16px;
  line-height: 1.5;
}

.error-banner {
  border-radius: 12px;
  border: 1px solid rgba(252, 165, 165, 0.45);
  background: rgba(127, 29, 29, 0.18);
  color: #fca5a5;
  font-size: 13px;
  line-height: 1.45;
  padding: 10px 12px;
}

.error-state {
  color: #fca5a5;
  border-color: rgba(252, 165, 165, 0.45);
  background: rgba(127, 29, 29, 0.18);
}
</style>
