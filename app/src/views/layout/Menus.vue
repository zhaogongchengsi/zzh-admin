<script setup lang="ts">
import { ref } from "vue"
import { useMenuStore } from '@/pinia'
import { useRouter } from "vue-router";
const menuStore = useMenuStore()
const collapsed = ref(false)
const router = useRouter()
const onRoute = (href: string) => {
  router.push(href)
}

</script>


<template>
  <div>
    <div class="logo-container">
      <div class="logo-img">
        <img src="/logo.svg" alt="logo" />
      </div>
      <div class="logo-title">
      <blog-typography-title :heading="6" :ellipsis="false">
        超级管理平台
      </blog-typography-title>
      </div>
    </div>
    <blog-menu
      :style="{ width: '200px', borderRadius: '4px' }"
      :theme="menuStore.theme" 
      :collapsed="menuStore.collapsed"
      :default-open-keys="[0]"
      :default-selected-keys="['1-0']"
    >
      <blog-menu-item key="0" @click="onRoute('/home')" >
          <template #icon>
            <icon-font type="icon-a-fangzishouye" />
          </template>
          控制台
      </blog-menu-item>
      <blog-sub-menu key="1">
        <template #icon>
          <icon-font type="icon-wenzhang" />
        </template>
        <template #title>文章管理</template>
        <blog-menu-item key="1-0" @click="onRoute('/article_list')" >
          <template #icon>
            <icon-font type="icon-yuwen" />
          </template>
          文章列表
        </blog-menu-item>
        <blog-menu-item key="1-1" @click="onRoute('/article_create')" >
          <template #icon>
            <icon-font type="icon-bianji" />
          </template>
          创建文章
        </blog-menu-item>
        <blog-menu-item key="1-2" @click="onRoute('/tags')" >
          <template #icon>
            <icon-font type="icon-yanshoubiaoqianguanli" />
          </template>
          标签管理
        </blog-menu-item>
      </blog-sub-menu>
      <blog-sub-menu key="2">
        <template #icon>
          <icon-font type="icon-kujialeqiyezhan_sucaishoucang" />
        </template>
        <template #title>素材管理</template>
        <blog-menu-item key="2-0" @click="onRoute('/media_admin')" >
          <template #icon>
            <icon-font type="icon-sucai" />
          </template>
          素材库
        </blog-menu-item>
        <blog-menu-item key="2-1" @click="onRoute('/media_upload')" >
          <template #icon>
            <icon-font type="icon-shangchuan" />
          </template>
          上传素材
        </blog-menu-item>
      </blog-sub-menu>
    </blog-menu>
  </div>
</template>

<style scoped lang="scss">
.menu-demo {
  box-sizing: border-box;
  width: 100%;
  padding: 40px;
  background-color: var(--color-neutral-2);
}

.logo-container {
  height: var(--header-tools-height);
  display: flex;
  align-items: center;
  .logo-img {
    width: 48px;
    height: var(--header-tools-height);
    flex-shrink: 0;
    display: flex;
    align-items: center;
    justify-content: center;
    img {
      width: 70%;
      height: 70%;
    }
  }

  .logo-title {
    width: calc(200px - 48px);
    height: var(--header-tools-height);
    flex-shrink: 0;
    display: flex;
    align-items: center;
    justify-content: center;
    font-size: 18px;
    color: #fff;
  }
}
</style>