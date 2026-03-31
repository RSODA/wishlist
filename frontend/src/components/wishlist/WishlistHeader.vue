<template>
  <header class="top-area">
    <div>
      <h1 class="title">Wishlist</h1>
      <p class="subtitle">Управление подписками</p>
    </div>

    <div class="actions-row">
      <button type="button" class="action-btn search-btn" aria-label="Поиск">
        <span class="action-icon">🔍</span>
        <span>Поиск</span>
      </button>

      <button type="button" class="action-btn plus-btn" aria-label="Добавить">
        +
      </button>
    </div>
  </header>
</template>

<script>
export default {
  name: 'WishlistHeader'
}
</script>

<style scoped>
.top-area {
  display: flex;
  flex-direction: column;
  gap: 14px;
}

.title {
  font-size: 28px;
  line-height: 1.2;
  font-weight: 700;
  color: var(--text-primary);
}

.subtitle {
  margin-top: 2px;
  color: var(--text-secondary);
  font-size: 14px;
}

.actions-row {
  display: flex;
  gap: 10px;
}

.action-btn {
  border: 1px solid transparent;
  border-radius: 12px;
  font-weight: 600;
  cursor: pointer;
  -webkit-tap-highlight-color: transparent;
}

.search-btn {
  flex: 1;
  min-height: 44px;
  background: var(--surface-secondary);
  color: var(--text-primary);
  border-color: var(--border-primary);
  padding: 0 14px;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
}

.action-icon {
  font-size: 14px;
}

.plus-btn {
  width: 44px;
  min-height: 44px;
  background: var(--accent-primary);
  color: #ffffff;
  font-size: 24px;
  line-height: 1;
}
</style>
