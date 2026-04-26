<template>
  <article class="wish-card">
    <div class="wish-preview" :class="{ empty: !item.picture }">
      <img v-if="item.picture" :src="item.picture" :alt="item.title || 'Изображение желания'" class="wish-image" />
      <span v-else class="wish-placeholder">No image</span>
    </div>

    <div class="wish-copy">
      <div class="wish-head">
        <h3 class="wish-title">{{ item.title || 'Без названия' }}</h3>
        <span class="status-pill">{{ item.statusLabel || 'Без статуса' }}</span>
      </div>

      <p class="wish-price">{{ displayPrice }}</p>

      <a
        v-if="item.url"
        :href="item.url"
        target="_blank"
        rel="noreferrer noopener"
        class="wish-link"
      >
        Открыть ссылку
      </a>
    </div>
  </article>
</template>

<script>
export default {
  name: 'WishlistCard',
  props: {
    item: {
      type: Object,
      required: true
    }
  },
  computed: {
    displayPrice() {
      if (this.item.price === null || this.item.price === undefined) {
        return 'Цена не указана'
      }

      return new Intl.NumberFormat('ru-RU').format(this.item.price) + ' ₽'
    }
  }
}
</script>

<style scoped>
.wish-card {
  display: grid;
  grid-template-columns: 108px minmax(0, 1fr);
  gap: 14px;
  padding: 14px;
  border-radius: 16px;
  border: 1px solid var(--border-secondary);
  background: rgba(15, 23, 42, 0.38);
}

.wish-preview {
  min-height: 108px;
  border-radius: 12px;
  overflow: hidden;
  background: rgba(26, 39, 64, 0.9);
  border: 1px solid rgba(148, 163, 184, 0.14);
}

.wish-preview.empty {
  display: grid;
  place-items: center;
}

.wish-image {
  display: block;
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.wish-placeholder {
  color: var(--text-secondary);
  font-size: 12px;
  font-weight: 700;
  text-transform: uppercase;
  letter-spacing: 0.08em;
}

.wish-copy {
  min-width: 0;
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.wish-head {
  display: flex;
  align-items: flex-start;
  justify-content: space-between;
  gap: 10px;
}

.wish-title {
  color: var(--text-primary);
  font-size: 16px;
  line-height: 1.4;
  font-weight: 700;
  word-break: break-word;
}

.status-pill {
  flex-shrink: 0;
  min-height: 30px;
  padding: 0 12px;
  border-radius: 999px;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  background: rgba(191, 219, 254, 0.14);
  color: #bfdbfe;
  font-size: 12px;
  font-weight: 700;
}

.wish-price {
  color: var(--accent-primary);
  font-size: 15px;
  font-weight: 700;
}

.wish-link {
  width: fit-content;
  color: var(--text-primary);
  font-size: 14px;
  font-weight: 700;
  text-decoration: none;
  border-bottom: 1px solid rgba(226, 232, 240, 0.25);
}

@media (max-width: 640px) {
  .wish-card {
    grid-template-columns: 1fr;
  }

  .wish-preview {
    min-height: 180px;
  }

  .wish-head {
    flex-direction: column;
  }
}
</style>
