

<template>
  <div class="article-type">
    <div class="add-type"><a-button type="primary" @click="visible = true" >添加新文章类型</a-button></div>
    <a-table :data="data" row-key="ID">
      <template #columns>
        <a-table-column title="类型ID" data-index="ID" align="center"></a-table-column>
        <a-table-column title="类型名称" data-index="article_type" align="center"></a-table-column>
        <a-table-column title="类型LOGO" data-index="type_logo" align="center"></a-table-column>
        <a-table-column title="类型描述" data-index="article_type_desc" align="center"></a-table-column>
        <a-table-column title="属于该类型的文章" data-index="type_articles" align="center"></a-table-column>
      </template>
    </a-table>
  </div>
    <a-modal v-model:visible="visible" @ok="handleOk" @cancel="handleCancel">
    <template #title>
      新建文章类型
    </template>
      <a-form :model="form">
    <a-form-item field="name" label="类型名称">
      <a-input v-model="form.article_type" placeholder="请输入类型名称" />
    </a-form-item>
    <a-form-item field="post" label="类型描述">
      <a-input v-model="form.article_type_desc" placeholder="请输入文章描述" />
    </a-form-item>
    <a-form-item field="post" label="类型logo">
      <a-upload
        :list-type="'picture'"
        :custom-request="upload"
        :multiple="false"
         :show-file-list="false"
      >
        <template #upload-button>
          <div class="type-logo">
            <icon-font type="icon-shangchuan" size="30" v-if="src === ''" />
             <a-avatar :size="64" v-else>
               <img :src="src" alt="logo">
             </a-avatar>
          </div>
        </template>
      </a-upload>
    </a-form-item>
  </a-form>
  </a-modal>
</template>


<script setup lang="ts">
import { getArticleType, createArticleType } from "@/api/article"
import { onMounted, ref, reactive} from 'vue'
import type { ArticleType } from "@/types/response";
import { uploadFile, getfiles } from '@/api/upload'

const data = ref<ArticleType[]>([])
const visible = ref<boolean>(false)
const form = reactive<ArticleType>({
  article_type: "",
  article_type_desc: "",
  type_logo: ""
})
let src = ref<string>("")
onMounted(async () => {
  try {
    const t = await getArticleType()
    data.value = t
    console.log(t)
  } catch (e) {
    console.error(e)
  }
})

const upload = async (options) => {
  try {
    const isOk = await uploadFile(options)
    if (isOk) {
      src.value = `http://${isOk.host}/${isOk.file.sava_path + isOk.file.file_name+isOk.file.file_ext}`
      form.type_logo = isOk.file.ID
    }
  } catch (e) {
    console.error(e)
  }

};




const handleOk = async () => {
  try {
    const res = await createArticleType(form)
    if (res) {
      // 上传成功
      console.log(res)
    }
  } catch (e) {
    console.error(e)
  }

}

const handleCancel = () => {}



</script>


<style lang="scss" scoped>
.article-type {
  padding: 15px;
  box-sizing: border-box;
}
.add-type {
  margin-bottom: 15px;
}

.type-logo {
  width: 100px;
  height: 100px;
  background-color: rgba(241, 241, 241, 0.74);
  border: 1px dashed #ccc;
  display: flex;
  align-items: center;
  justify-content: center;
}
</style>