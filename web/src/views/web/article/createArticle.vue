<script setup>
import { ref } from 'vue'
import Editor from '@/components/Editor/index.vue'
const activeName = ref('Word')


const handleClick = (tab, event) => {
  // console.log(tab, event)
}


</script>


<template>
  <div class="article-con">
      <el-tabs v-model="activeName" @tab-click="handleClick">
        <el-tab-pane label="Word" name="Word">
          <Editor EditorType="word"  />
        </el-tab-pane>
        <el-tab-pane label="Md" name="Md">
          <Editor EditorType="md"  />
        </el-tab-pane>
      </el-tabs>
  </div>
</template>

<style lang="scss" scoped>

.article-con {
  padding: 10px;
  box-sizing: border-box;
}

</style>