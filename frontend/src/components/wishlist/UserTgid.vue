<template>
  <p class="tgid-line">TG ID: <strong>{{ tgid }}</strong></p>
</template>

<script>
export default {
  name: 'UserTgid',
  props: {
    tgid: {
      type: [String, Number],
      default: '—'
    }
  }
}
</script>

<style scoped>
.tgid-line {
  color: var(--text-secondary);
  font-size: 14px;
}

.tgid-line strong {
  color: var(--text-primary);
}
</style>
