<template>
  <div class="loading-overlay" role="status" aria-live="polite">
    <div class="loading-card">Загрузка...</div>
  </div>
</template>

<script>
export default {
  name: 'LoadingOverlay'
}
</script>

<style scoped>
.loading-overlay {
  position: fixed;
  inset: 0;
  z-index: 1000;
  background: rgba(2, 6, 23, 0.78);
  display: grid;
  place-items: center;
  backdrop-filter: blur(4px);
}

.loading-card {
  background: var(--surface-primary);
  color: var(--text-primary);
  border: 1px solid var(--border-primary);
  border-radius: 14px;
  padding: 14px 20px;
  font-size: 16px;
  font-weight: 600;
  box-shadow: 0 12px 32px rgba(0, 0, 0, 0.35);
}
</style>
