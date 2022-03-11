<template>
  <div class="article-list" v-dark:dark="menuStore.theme">
    <a-table 
      row-key="ID" 
      :data="articles" 
      :row-selection="rowSelection" 
      size="medium" 
      stripe
      :pagination="pagConf" 
    >
      <template #columns>
        <a-table-column title="创建时间" data-index="CreatedAt" align="center">
          <template #cell="{ record  }">
            <span>{{formatTime(record.CreatedAt)}}</span>
          </template>
        </a-table-column>
        <a-table-column title="文章ID" data-index="ID" align="center" :sortable="sortable"></a-table-column>
        <a-table-column title="文章作者" data-index="article_author" align="center"></a-table-column>
        <a-table-column title="文章描述" data-index="article_desc" align="center"></a-table-column>
        <a-table-column title="文章文件名" data-index="article_file_name" align="center"></a-table-column>
        <a-table-column title="文章存储类型" data-index="article_storage_type" align="center"></a-table-column>
        <a-table-column title="文章标题" data-index="article_title" align="center"></a-table-column>
        <a-table-column title="文章类型" data-index="article_type" align="center"></a-table-column>
        <a-table-column title="文章存储路径" data-index="article_url" align="center" width="300" ellipsis></a-table-column>
        <a-table-column title="文章标签" data-index="article_tags" align="center">
          <template #cell="{ record  }" >
            <a-trigger position="bottom" auto-fit-position>
              <span class="hover-tag">查看标签</span>
              <template #content>
                <div class="tags">
                  <a-space direction="vertical" align="center">
                    <a-tag v-for="(item,index) in record.article_tags" :key="index" :color="item.tag_color" >{{item.tag}}</a-tag>
                  </a-space>
                </div>
              </template>
            </a-trigger>
          </template>
        </a-table-column>
        <a-table-column title="点赞数" data-index="likes" align="center"></a-table-column>
        <a-table-column title="观看数" data-index="number_views" align="center"></a-table-column>
      </template>
    </a-table>
  </div>
</template>
<script setup lang="ts">
import { useMenuStore } from "@/pinia";
import { getArticleList } from "@/api/article";
import { onMounted, ref, reactive } from "vue";
import type { article } from "@/types/response";
import { formatTime, randomHexColor } from "@/utils/index";
const articles = ref<article[]>([]);
const pagConf = reactive({
  total: 0,
  current: 1,
  pageSize: 20,
})
onMounted(async () => {
  try {
    const res = await getArticleList(20, 0);
    articles.value = res.article_list;
    pagConf.total = res.count
  } catch (e) {
    console.log(e);
  }
});


// 过滤
const sortable ={
  sortDirections:['ascend','descend'],
  sorter: (a,b) => {
    return a < b
  }
}

const onitem = (type) => {
  console.log(type);
};
const menuStore = useMenuStore();

const rowSelection = {
  type: "checkbox",
  showCheckedAll: true,
};
</script>
<style lang="scss">
.article-list {
  height: calc(100vh - 24px - 52px);
  padding: 15px;
  box-sizing: border-box;
}
.hover-tag {
  cursor: pointer;
}
.tags {
  background-color: var(--color-bg-1);
  padding: 5px;
}
.article-list-dark {
  background-color: var(--color-menu-dark-bg);
  color: #fff;
  box-shadow: 0 2px 5px 0 rgb(0 0 0 / 8%);
}

.article-ul {
  --space: 30px;
  width: 100%;
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  grid-row-gap: var(--space);
  grid-column-gap: var(--space);
}

.article-item {
  border: 1px solid #ccc;
  padding: 10px;
  border: 1px solid var(--color-border-3);
  border-radius: 5px;
}

.article-card {
  display: flex;
  flex-direction: column;

  .art-title {
    font-size: 18px;
    display: block;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
  }

  .art-con {
    line-height: 30px;
    height: 60px;
  }
}
</style>