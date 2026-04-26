<template>
  <div class="modal-backdrop" @click.self="handleBackdropClick">
    <section class="modal-card" role="dialog" aria-modal="true" aria-labelledby="add-user-title">
      <button
        type="button"
        class="close-btn"
        aria-label="Закрыть"
        :disabled="isSubmitting"
        @click="$emit('close')"
      >
        ×
      </button>

      <p class="modal-kicker">Новая подписка</p>
      <h2 id="add-user-title" class="modal-title">Добавить пользователя</h2>
      <p class="modal-copy">
        Введите Telegram username пользователя, которого хотите добавить к себе. Можно вводить
        как <strong>@username</strong>, так и без символа <strong>@</strong>.
      </p>

      <form class="modal-form" @submit.prevent="submitForm">
        <label class="field-label" for="subscription-username">Username</label>

        <div class="input-shell">
          <span class="input-prefix">@</span>
          <input
            id="subscription-username"
            ref="usernameInput"
            v-model="localUsername"
            class="username-input"
            type="text"
            inputmode="text"
            autocomplete="off"
            placeholder="nickname"
            :disabled="isSubmitting"
          />
        </div>

        <p v-if="error" class="form-error">{{ error }}</p>

        <div class="actions-row">
          <button type="button" class="secondary-btn" :disabled="isSubmitting" @click="$emit('close')">
            Отмена
          </button>
          <button type="submit" class="primary-btn" :disabled="isSubmitting">
            {{ isSubmitting ? 'Добавление...' : 'Добавить' }}
          </button>
        </div>
      </form>
    </section>
  </div>
</template>

<script>
export default {
  name: 'AddSubscriptionModal',
  emits: ['close', 'submit'],
  props: {
    error: {
      type: String,
      default: ''
    },
    isSubmitting: {
      type: Boolean,
      default: false
    }
  },
  data() {
    return {
      localUsername: ''
    }
  },
  mounted() {
    document.addEventListener('keydown', this.handleKeydown)
    this.focusInput()
  },
  beforeUnmount() {
    document.removeEventListener('keydown', this.handleKeydown)
  },
  methods: {
    focusInput() {
      this.$nextTick(() => {
        this.$refs.usernameInput?.focus()
      })
    },
    handleBackdropClick() {
      if (!this.isSubmitting) {
        this.$emit('close')
      }
    },
    handleKeydown(event) {
      if (event.key === 'Escape' && !this.isSubmitting) {
        this.$emit('close')
      }
    },
    submitForm() {
      this.$emit('submit', this.localUsername)
    }
  }
}
</script>

<style scoped>
.modal-backdrop {
  position: fixed;
  inset: 0;
  z-index: 1100;
  background: rgba(2, 6, 23, 0.82);
  backdrop-filter: blur(10px);
  display: grid;
  place-items: center;
  padding: 20px 16px;
}

.modal-card {
  position: relative;
  width: min(100%, 420px);
  border-radius: 22px;
  padding: 22px;
  background:
    radial-gradient(circle at top right, rgba(56, 189, 248, 0.18), transparent 34%),
    linear-gradient(180deg, rgba(17, 26, 44, 0.98), rgba(12, 19, 33, 0.98));
  border: 1px solid rgba(97, 149, 219, 0.32);
  box-shadow: 0 24px 60px rgba(0, 0, 0, 0.4);
}

.close-btn {
  position: absolute;
  top: 14px;
  right: 14px;
  width: 36px;
  height: 36px;
  border: 0;
  border-radius: 999px;
  background: rgba(148, 163, 184, 0.12);
  color: var(--text-secondary);
  font-size: 24px;
  line-height: 1;
  cursor: pointer;
}

.close-btn:disabled {
  cursor: not-allowed;
  opacity: 0.55;
}

.modal-kicker {
  color: var(--accent-primary);
  font-size: 12px;
  font-weight: 700;
  text-transform: uppercase;
  letter-spacing: 0.08em;
}

.modal-title {
  margin-top: 8px;
  font-size: 26px;
  line-height: 1.1;
  color: var(--text-primary);
}

.modal-copy {
  margin-top: 10px;
  color: var(--text-secondary);
  font-size: 14px;
  line-height: 1.5;
}

.modal-form {
  margin-top: 22px;
}

.field-label {
  display: block;
  margin-bottom: 8px;
  color: var(--text-primary);
  font-size: 14px;
  font-weight: 600;
}

.input-shell {
  display: flex;
  align-items: center;
  gap: 10px;
  border-radius: 14px;
  border: 1px solid var(--border-primary);
  background: rgba(15, 23, 42, 0.78);
  padding: 0 14px;
}

.input-prefix {
  color: var(--accent-primary);
  font-size: 18px;
  font-weight: 700;
}

.username-input {
  width: 100%;
  min-height: 52px;
  border: 0;
  outline: 0;
  background: transparent;
  color: var(--text-primary);
  font-size: 16px;
}

.username-input::placeholder {
  color: rgba(148, 163, 184, 0.75);
}

.form-error {
  margin-top: 10px;
  border-radius: 12px;
  border: 1px solid rgba(252, 165, 165, 0.3);
  background: rgba(127, 29, 29, 0.22);
  color: #fecaca;
  padding: 10px 12px;
  font-size: 13px;
  line-height: 1.45;
}

.actions-row {
  margin-top: 18px;
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 10px;
}

.secondary-btn,
.primary-btn {
  min-height: 46px;
  border-radius: 14px;
  border: 1px solid transparent;
  font-size: 14px;
  font-weight: 700;
  cursor: pointer;
}

.secondary-btn {
  background: rgba(148, 163, 184, 0.08);
  border-color: var(--border-primary);
  color: var(--text-primary);
}

.primary-btn {
  background: linear-gradient(135deg, #38bdf8, #0ea5e9);
  color: #ffffff;
  box-shadow: 0 14px 28px rgba(14, 165, 233, 0.24);
}

.secondary-btn:disabled,
.primary-btn:disabled {
  cursor: not-allowed;
  opacity: 0.65;
}
</style>
