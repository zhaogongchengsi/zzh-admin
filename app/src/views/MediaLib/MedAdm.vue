<template>
  <div>
    <a-tabs :animation="true" @change="onTabChange" :default-active-key="upTable">
      <template #extra>
        <a-button @click="onUpdate" type="primary">
          <template #icon>
            <icon-font type="icon-shangchuan" />
          </template>
          上传
        </a-button>
      </template>
      <a-tab-pane key="image">
        <template #title>
          <a-space>
            <icon-font type="icon-tupian1" />
            <span>图片素材</span>
          </a-space>
        </template>
        <div class="tab-box">
          <photo-wall />
        </div>
      </a-tab-pane>
      <a-tab-pane key="music">
        <template #title>
          <a-space>
            <icon-font type="icon-yinle1" />
            <span>音乐素材</span>
          </a-space>
        </template>
        <div class="tab-box">
          <misoc-wall />
        </div>
      </a-tab-pane>
      <a-tab-pane key="qita">
        <template #title>
          <a-space>
            <icon-font type="icon-qita" />
            <span>其他素材</span>
          </a-space>
        </template>
        <div class="tab-box">其他素材 功能待做</div>
      </a-tab-pane>
    </a-tabs>
    <a-modal v-model:visible="up_visible" @cancel="up_visible = false" :footer="false">
      <template #title>素材上传</template>
      <div>
        <a-upload
          :list-type="'picture'"
          :default-file-list="files"
          :custom-request="upload"
          :multiple="false"
        />
      </div>
    </a-modal>
  </div>
</template>


<script lang="ts" setup>
import PhotoWall from "./PhotoWall.vue";
import MisocWall from "./MusicWall.vue";
import { ref, onMounted } from "vue";
import { uploadFile, getfiles } from '@/api/upload'


let upTable = ref<string>("image");
const host:string = "" 
const files = [];
const onTabChange = (key: string) => {
  upTable.value = key;
};
const up_visible = ref<boolean>(false)
onMounted(async () => {
   
})
const onUpdate = () => (up_visible.value = true);
const upload = async (options) => {
  const isOk = await uploadFile(options)
  isOk === true ? up_visible.value = false : "";
};
</script>


<style lang="scss" scoped>
.tab-box {
  box-sizing: border-box;
  padding: 0 15px;
}
</style>