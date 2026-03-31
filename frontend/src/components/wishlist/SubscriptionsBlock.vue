<template>
  <section class="subscriptions-block" aria-label="Блок подписок">
    <div class="block-header">
      <h2>Подписки</h2>
    </div>

    <div class="empty-state">Список пока пуст</div>
  </section>
</template>

<script>
export default {
  name: 'SubscriptionsBlock'
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

.block-header h2 {
  font-size: 18px;
  font-weight: 700;
  color: var(--text-primary);
}

.empty-state {
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
}
</style>
