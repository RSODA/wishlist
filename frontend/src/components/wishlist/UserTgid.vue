<template>
  <div class="user-meta">
    <p class="tgid-line">TG ID: <strong>{{ tgid }}</strong></p>
    <p v-if="username" class="username-line">@{{ username }}</p>
  </div>
</template>

<script>
export default {
  name: 'UserTgid',
  props: {
    tgid: {
      type: [String, Number],
      default: '—'
    },
    username: {
      type: String,
      default: ''
    }
  }
}
</script>

<style scoped>
.user-meta {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.tgid-line {
  color: var(--text-secondary);
  font-size: 14px;
}

.tgid-line strong {
  color: var(--text-primary);
}

.username-line {
  color: var(--text-secondary);
  font-size: 14px;
}
</style>
