<template>
  <div class="imgs-container">
    <ul class="img-list">
      <li v-for="img of fileList" :key="img.ID" class="img-item">
        <a-image :src="`http://${host}/${img.sava_path}${img.file_name}${img.file_ext}`" :alt="img.file_name" height="200" width="300"  />
      </li>
    </ul>
  </div>
</template>

<script lang="ts" setup>
import { ref, onMounted } from "vue";
import { uploadFile, getfiles } from '@/api/upload'
import { file } from "@/types/response";

const fileList = ref<file[]>([])
const host = ref('')

onMounted(async () => {
 const res = await getfiles('image', 10, 0)
 fileList.value = res.file_list
  host.value = res.host
})




</script>

<style lang="scss" scoped>
.img-list {
  display: grid;
  grid-template-columns: repeat(auto-fill, 300px);
  grid-template-rows: repeat(auto-fill, 200px);
  grid-row-gap: 15px;
  grid-column-gap: 15px;
  justify-content: space-around;
  .img-item {
    width: 100%;
    height: 100%;
    display: flex;
      align-items: center;
      justify-content: center;
      overflow: hidden;
      cursor: pointer;
  }
}

.img-item:deep(.arco-image) {
  flex: 1;
  img {
    width: 100% !important;
    object-fit: cover;
  }
}
</style>