<template>
  <header class="wishlist-header">
    <div class="top-row">
      <button type="button" class="back-btn" @click="$emit('back')">Назад</button>
      <button
        v-if="canAddWish"
        type="button"
        class="add-btn"
        :disabled="isBusy"
        @click="$emit('open-add-wish')"
      >
        {{ isBusy ? 'Сохраняем...' : 'Добавить желание' }}
      </button>
    </div>

    <div class="copy">
      <p class="eyebrow">{{ canAddWish ? 'Мой wishlist' : 'Просмотр wishlist' }}</p>
      <h1 class="title">
        {{ canAddWish ? 'Мои желания' : ownerUsername ? `@${ownerUsername}` : 'Пользователь' }}
      </h1>
      <p class="subtitle">TG ID: {{ ownerTgId || '—' }}</p>
    </div>
  </header>
</template>

<script>
export default {
  name: 'WishlistViewHeader',
  emits: ['back', 'open-add-wish'],
  props: {
    ownerTgId: {
      type: [String, Number],
      default: ''
    },
    ownerUsername: {
      type: String,
      default: ''
    },
    canAddWish: {
      type: Boolean,
      default: false
    },
    isBusy: {
      type: Boolean,
      default: false
    }
  }
}
</script>

<style scoped>
.wishlist-header {
  display: flex;
  flex-direction: column;
  gap: 14px;
}

.top-row {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 12px;
}

.back-btn {
  width: fit-content;
  min-height: 40px;
  padding: 0 14px;
  border: 1px solid var(--border-primary);
  border-radius: 999px;
  background: rgba(15, 23, 42, 0.45);
  color: var(--text-primary);
  font-size: 14px;
  font-weight: 700;
  cursor: pointer;
}

.add-btn {
  min-height: 40px;
  padding: 0 14px;
  border: 1px solid transparent;
  border-radius: 999px;
  background: linear-gradient(135deg, #38bdf8, #0ea5e9);
  color: #ffffff;
  font-size: 14px;
  font-weight: 700;
  cursor: pointer;
}

.add-btn:disabled {
  opacity: 0.65;
  cursor: not-allowed;
}

.eyebrow {
  color: var(--accent-primary);
  font-size: 12px;
  font-weight: 700;
  letter-spacing: 0.08em;
  text-transform: uppercase;
}

.title {
  margin-top: 6px;
  font-size: 30px;
  line-height: 1.1;
  color: var(--text-primary);
}

.subtitle {
  margin-top: 6px;
  color: var(--text-secondary);
  font-size: 14px;
}

@media (max-width: 640px) {
  .top-row {
    flex-direction: column;
    align-items: stretch;
  }

  .add-btn {
    width: 100%;
  }
}
</style>
