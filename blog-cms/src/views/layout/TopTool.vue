<script setup lang="ts">
import { IconMinus, IconClose } from '@arco-design/web-vue/es/icon';

const onSysetm = (type: String) => {
  window.electron.doThing()
}

</script>

<template>
  <div class="toptool-container">
    <div class="toptool-right" id="toptool">
      <icon-minus @click="onSysetm('minus')" />
      <icon-close @click="onSysetm('close')" />
    </div>
  </div>
</template>

<style scoped lang="scss">
.toptool-container {
  height: var(--header-tools-height);
  background-color: var(--color-menu-dark-bg);
  -webkit-app-region: drag;
  display: flex;
  justify-content: end;
  .toptool-right {
    width: 80px;
    display: flex;
    align-items: center;
    justify-content:space-evenly;
  }
}

.toptool-container:deep(.arco-icon) {
  -webkit-app-region: no-drag;
  cursor: pointer;
  width: 1.2rem;
  height: 1.2rem;
  border-radius: 3px;
  &:hover {
    background-color: var(--color-secondary-hover);
    color: var(--color-text-2);
  }
}

</style>
