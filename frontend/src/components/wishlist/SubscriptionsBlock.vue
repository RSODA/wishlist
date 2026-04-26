<template>
  <section class="subscriptions-block" aria-label="Блок подписок">
    <div class="block-header">
      <div>
        <h2>Подписки</h2>
        <p v-if="username" class="block-meta">@{{ username }}</p>
      </div>
      <span class="counter">{{ displayedSubscriptions.length }}</span>
    </div>

    <div v-if="error && totalSubscriptionsCount === 0" class="error-state">{{ error }}</div>

    <template v-else>
      <div v-if="error" class="error-banner">{{ error }}</div>

      <div class="tabs-row" role="tablist" aria-label="Фильтр подписок">
        <button
          type="button"
          class="tab-btn"
          :class="{ active: activeAcceptedFilter === true }"
          @click="openAcceptedTab"
        >
          <span>Принятые</span>
          <span class="tab-count">{{ acceptedSubscriptions.length }}</span>
        </button>

        <button
          type="button"
          class="tab-btn"
          :class="{ active: activeAcceptedFilter === false }"
          @click="openWaitingTab"
        >
          <span>Ожидают</span>
          <span class="tab-count">{{ waitingSubscriptionsCount }}</span>
        </button>
      </div>

      <div v-if="!activeAcceptedFilter" class="pending-switcher" role="tablist" aria-label="Тип ожидающих заявок">
        <button
          type="button"
          class="pending-switch-btn"
          :class="{ active: waitingDisplayMode === 'incoming' }"
          @click="waitingDisplayMode = 'incoming'"
        >
          <span>Ожидают подтверждения</span>
          <span class="pending-switch-count">{{ pendingSubscriptions.length }}</span>
        </button>

        <button
          type="button"
          class="pending-switch-btn"
          :class="{ active: waitingDisplayMode === 'sent' }"
          @click="waitingDisplayMode = 'sent'"
        >
          <span>Отправленные</span>
          <span class="pending-switch-count">{{ sentSubscriptions.length }}</span>
        </button>
      </div>

      <div v-if="displayedSubscriptions.length === 0" class="empty-state">
        {{ emptyStateMessage }}
      </div>

      <ul v-else class="subscriptions-list">
        <li
          v-for="(subscription, index) in displayedSubscriptions"
          :key="subscription.tgId || subscription.username || index"
          class="subscription-item"
        >
          <div class="subscription-copy">
            <p class="subscription-name">
              {{ subscription.username ? `@${subscription.username}` : 'Без username' }}
            </p>
            <p class="subscription-id">TG ID: {{ subscription.tgId || '—' }}</p>
          </div>

          <div class="subscription-actions">
            <button
              v-if="showWishlistAction"
              type="button"
              class="wishlist-btn"
              @click="$emit('open-wishlist', subscription)"
            >
              Открыть wishlist
            </button>

            <button
              v-if="showAcceptAction"
              type="button"
              class="confirm-btn"
              :disabled="isAcceptingSubscription(subscription)"
              @click="$emit('accept-subscription', subscription.tgId)"
            >
              {{ isAcceptingSubscription(subscription) ? 'Подтверждение...' : 'Подтвердить' }}
            </button>

            <span class="status-pill" :class="statusPillClass">
              {{ statusPillLabel }}
            </span>
          </div>
        </li>
      </ul>
    </template>
  </section>
</template>

<script>
export default {
  name: 'SubscriptionsBlock',
  emits: ['accept-subscription', 'open-wishlist'],
  props: {
    acceptedSubscriptions: {
      type: Array,
      default: () => []
    },
    pendingSubscriptions: {
      type: Array,
      default: () => []
    },
    sentSubscriptions: {
      type: Array,
      default: () => []
    },
    error: {
      type: String,
      default: ''
    },
    username: {
      type: String,
      default: ''
    },
    acceptingSubscriptionTgId: {
      type: [String, Number],
      default: ''
    }
  },
  data() {
    return {
      activeAcceptedFilter: true,
      waitingDisplayMode: 'incoming'
    }
  },
  computed: {
    waitingSubscriptionsCount() {
      return this.pendingSubscriptions.length + this.sentSubscriptions.length
    },
    displayedSubscriptions() {
      if (this.activeAcceptedFilter) {
        return this.acceptedSubscriptions
      }

      return this.waitingDisplayMode === 'incoming'
        ? this.pendingSubscriptions
        : this.sentSubscriptions
    },
    totalSubscriptionsCount() {
      return this.acceptedSubscriptions.length + this.waitingSubscriptionsCount
    },
    showAcceptAction() {
      return !this.activeAcceptedFilter && this.waitingDisplayMode === 'incoming'
    },
    showWishlistAction() {
      return this.activeAcceptedFilter
    },
    statusPillClass() {
      if (this.activeAcceptedFilter) {
        return 'accepted'
      }

      return this.waitingDisplayMode === 'incoming' ? 'pending' : 'sent'
    },
    statusPillLabel() {
      if (this.activeAcceptedFilter) {
        return 'Принята'
      }

      return this.waitingDisplayMode === 'incoming' ? 'Ожидает' : 'Отправлена'
    },
    emptyStateMessage() {
      if (this.activeAcceptedFilter) {
        return 'Пока нет подтверждённых подписок.'
      }

      if (this.waitingDisplayMode === 'incoming') {
        return 'Сейчас нет входящих заявок, ожидающих подтверждения.'
      }

      return 'Сейчас нет отправленных заявок, ожидающих подтверждения.'
    }
  },
  methods: {
    openAcceptedTab() {
      this.activeAcceptedFilter = true
    },
    openWaitingTab() {
      this.activeAcceptedFilter = false
      this.waitingDisplayMode = this.pendingSubscriptions.length > 0 || this.sentSubscriptions.length === 0
        ? 'incoming'
        : 'sent'
    },
    isAcceptingSubscription(subscription) {
      return String(subscription.tgId ?? '') === String(this.acceptingSubscriptionTgId ?? '')
    }
  }
}
</script>

<style scoped>
.subscriptions-block {
  border-radius: 16px;
  background: var(--surface-primary);
  border: 1px solid var(--border-primary);
  min-height: 260px;
  padding: 16px;
  display: flex;
  flex-direction: column;
  gap: 12px;
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

.tabs-row,
.pending-switcher {
  display: grid;
  grid-template-columns: repeat(2, minmax(0, 1fr));
  gap: 10px;
}

.tab-btn,
.pending-switch-btn {
  min-height: 44px;
  border: 1px solid var(--border-primary);
  border-radius: 14px;
  background: rgba(15, 23, 42, 0.45);
  color: var(--text-secondary);
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 10px;
  padding: 0 14px;
  font-size: 14px;
  font-weight: 700;
  cursor: pointer;
  transition: border-color 0.2s ease, background 0.2s ease, color 0.2s ease;
}

.tab-btn.active,
.pending-switch-btn.active {
  border-color: rgba(56, 189, 248, 0.5);
  background: rgba(56, 189, 248, 0.14);
  color: var(--text-primary);
}

.tab-count,
.pending-switch-count {
  min-width: 24px;
  height: 24px;
  border-radius: 999px;
  background: rgba(148, 163, 184, 0.12);
  display: inline-grid;
  place-items: center;
  font-size: 12px;
}

.tab-btn.active .tab-count,
.pending-switch-btn.active .pending-switch-count {
  background: rgba(14, 165, 233, 0.22);
  color: #bfe9ff;
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

.subscriptions-list {
  list-style: none;
  margin: 0;
  padding: 0;
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.subscription-item {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 12px;
  padding: 14px;
  border-radius: 14px;
  border: 1px solid var(--border-secondary);
  background: rgba(15, 23, 42, 0.38);
}

.subscription-copy {
  min-width: 0;
}

.subscription-name {
  color: var(--text-primary);
  font-size: 15px;
  font-weight: 700;
  line-height: 1.4;
  word-break: break-word;
}

.subscription-id {
  margin-top: 4px;
  color: var(--text-secondary);
  font-size: 12px;
}

.subscription-actions {
  display: flex;
  align-items: center;
  gap: 10px;
  flex-shrink: 0;
}

.confirm-btn {
  min-height: 36px;
  padding: 0 14px;
  border: none;
  border-radius: 10px;
  background: linear-gradient(135deg, #38bdf8, #0ea5e9);
  color: #04131d;
  font-size: 13px;
  font-weight: 700;
  cursor: pointer;
  transition: transform 0.2s ease, opacity 0.2s ease;
}

.wishlist-btn {
  min-height: 36px;
  padding: 0 14px;
  border: 1px solid var(--border-primary);
  border-radius: 10px;
  background: rgba(56, 189, 248, 0.1);
  color: var(--text-primary);
  font-size: 13px;
  font-weight: 700;
  cursor: pointer;
}

.confirm-btn:disabled {
  cursor: wait;
  opacity: 0.72;
  transform: none;
}

.status-pill {
  min-height: 30px;
  padding: 0 12px;
  border-radius: 999px;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  font-size: 12px;
  font-weight: 700;
}

.status-pill.accepted {
  background: rgba(74, 222, 128, 0.14);
  color: #86efac;
}

.status-pill.pending {
  background: rgba(251, 191, 36, 0.14);
  color: #fcd34d;
}

.status-pill.sent {
  background: rgba(191, 219, 254, 0.14);
  color: #bfdbfe;
}

@media (max-width: 640px) {
  .subscription-item {
    flex-direction: column;
    align-items: flex-start;
  }

  .subscription-actions {
    width: 100%;
    justify-content: space-between;
  }
}
</style>
