<template>
  <div :class="['create-container']" v-dark:dark="menuStore.theme">
    <v-md-editor v-model="mdString" height="100%" mode="edit" @save="save"></v-md-editor>
    <a-modal v-model:visible="visible" @ok="handleOk" @cancel="handleCancel">
      <template #title>创建文章</template>
      <div>
        <a-form :model="articleFrom" @submit="handleSubmit">
          <a-form-item field="articleTitle" label="文章标题">
            <a-input v-model="articleFrom.articleTitle" placeholder="请输入文章标题" />
          </a-form-item>
          <a-form-item field="articleName" label="文章名字">
            <a-input v-model="articleFrom.articleName" placeholder="请输入文章名字" />
          </a-form-item>
          <a-form-item field="articleAuthor" label="文章作者">
            <a-input v-model="articleFrom.articleAuthor" placeholder="请输入文章作者名字" />
          </a-form-item>
          <a-form-item field="articleType" label="文章类型">
            <a-input v-model="articleFrom.articleType" placeholder="请输入文章类型" />
          </a-form-item>
          <a-form-item field="fileName" label="文件名">
            <a-input v-model="articleFrom.fileName" placeholder="请输入文件名" />
          </a-form-item>
          <a-form-item field="articleStorageType" label="存储类型">
            <a-select
              v-model="articleFrom.articleStorageType"
              placeholder="文章存储类型"
              @change="onRadChange"
            >
              <a-option :value="articleStorageType.oos">对象存储</a-option>
              <a-option :value="articleStorageType.databases">数据库</a-option>
            </a-select>
          </a-form-item>
          <a-form-item field="articleTags" label="文章标签">
            <a-select
              v-model="articleFrom.articleTags"
              placeholder="选择文章标签"
              :multiple="true"
              :allow-search="false"
            >
              <a-option
                v-for="(item, index) in taglist"
                :key="index"
                :value="item"
              >{{item}}</a-option>
            </a-select>
          </a-form-item>
          <a-form-item field="article_desc" label="文章描述">
            <a-input v-model="articleFrom.article_desc" placeholder="请选择文章描述" />
          </a-form-item>
        </a-form>
      </div>
    </a-modal>
    <a-modal :visible="subVisible" @ok="okSubVisible" @cancel="clearSubVisible">
      <template #title>选择对象存储路径</template>
      <div>
        <a-cascader
          :options="options"
          default-value="datunli"
          expand-trigger="hover"
          placeholder="选择存储路径"
          v-model="strType"
          :path-mode="true"
          :allow-clear="true"
        />
      </div>
    </a-modal>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from "vue";
import { useMenuStore } from "@/pinia";
import {
  article_req,
  articleStorageType,
  CosTempKeyRequest,
  tags,
} from "@/types/request";
import { createArticle } from "@/api/article";
const mdString = ref("");
const visible = ref(false);
const subVisible = ref(false);
const articleFrom: article_req = reactive<article_req>({
  fileName: "",
  articleName: "",
  articleTitle: "",
  articleUrl: "",
  articleStorageType: "",
  articleAuthor: "",
  articleType: "",
  articleContext: "",
  articleTags: [],
  article_desc: "",
});
const strType = ref([]);
const taglist = ref<string[]>(['javascript', 'golang']);

const cosOpt: CosTempKeyRequest = {
  Region: "ap-nanjing",
  Bucket: "bloghtml-1301735126",
  action: "",
};

const options = [
  {
    value: "bloghtml-1301735126",
    label: "bloghtml-1301735126",
    children: [
      {
        value: "shili",
        label: "shili",
      },
    ],
  },
];

const menuStore = useMenuStore();
const save = (text: string, html: string) => {
  articleFrom.articleContext = text;
  visible.value = true;
};

const handleSubmit = () => {
  // visible.value = true;
};

const onRadChange = (value: string) => {
  if (value === articleStorageType.oos) {
    subVisible.value = true;
  }
};
// 创建
const handleOk = async () => {
  // console.log(articleFrom);
  let key = await createArticle(articleFrom,articleStorageType.oos, strType.value)
  console.log(key);
  // visible.value = false;
};
const handleCancel = () => {
  visible.value = false;
};

const okSubVisible = () => (subVisible.value = false);
const clearSubVisible = () => {
  subVisible.value = false;
  strType.value = [];
};
</script>

<style lang="scss" scoped>
.create-container {
  height: 100%;
}

.create-container-dark:deep(.v-md-editor),
.create-container-dark:deep(.v-md-textarea-editor textarea) {
  background-color: var(--color-menu-dark-bg);
  color: #fff;
}
</style>
