<template>
  <div class="article-list" v-dark:dark="menuStore.theme">
  <ul class="article-ul">
    <li class="article-item">
      <a-card title="Arco Card" hoverable>
        <a-typography-text>
            ByteDance's core product, Toutiao ("Headlines"), is a content platform in
            China and around the world. Toutiao started out as a news recommendation
            engine and gradually evolved into a platform delivering content in various
            formats.
        </a-typography-text>
        <template #actions>
          <span class="icon-hover"> <IconThumbUp /> </span>
          <span class="icon-hover"> <IconShareInternal /> </span>
          <span class="icon-hover"> <IconMore /> </span>
        </template>
        <a-card-meta>
          <template #avatar>
            <div
              :style="{ display: 'flex', alignItems: 'center', color: '#1D2129' }"
            >
              <a-avatar :size="24" :style="{ marginRight: '8px' }">
                A
              </a-avatar>
              <a-typography-text>Username</a-typography-text>
            </div>
          </template>
        </a-card-meta>
      </a-card>
    </li>
  </ul>
  </div>
</template>
<script setup lang="ts">
import { useMenuStore } from "@/pinia";
import {
  IconThumbUp,
  IconShareInternal,
  IconMore,
} from '@arco-design/web-vue/es/icon';
const menuStore = useMenuStore();
</script>
<style lang="scss">
.article-list {
  height: calc(100vh - 24px - 52px);
  padding: 15px;
  box-sizing: border-box;
}
.article-list-dark {
  background-color: var(--color-menu-dark-bg);
  color: #fff;
}

.article-ul {
  --space: 30px;
  width: 100%;
  display: grid;
  grid-template-columns: repeat(4,1fr);
  grid-row-gap: var(--space);
  grid-column-gap: var(--space);
}
</style>