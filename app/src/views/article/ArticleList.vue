<template>
  <div class="article-list" v-dark:dark="menuStore.theme">
  <ul class="article-ul">
    <li 
      class="article-item"
      v-for="item in articles"
      :key="item.ID"
    >
      <div class="article-card">
        <span class="art-title">{{item.article_title}}</span>
        <p class="art-con">{{item.article_context}}</p>
        <div class="article-foot">
          <span>{{item.article_author}}</span>
        </div>
      </div>
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
import { getArticleList } from '@/api/article'
import { onMounted, ref } from "vue";
import type { article } from '@/types/response'

const articles = ref<article[]>([])

onMounted(async () => {
  try {
      const res = await getArticleList()
      articles.value = res.article_list
      console.log(res)
  } catch (e) {

  }

})

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

.article-item {
  border: 1px solid #ccc;
  padding: 10px;
  border: 1px solid var(--color-border-3);
}

.article-card {
  display: flex;
  flex-direction: column;

  .art-title {
    font-size: 18px;
  }

  .art-con {
    line-height: 30px ;
    height: 60px;
  }
}
</style>