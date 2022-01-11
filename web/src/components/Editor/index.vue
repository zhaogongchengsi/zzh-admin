<script setup>
import word from './word.vue'
import md from './md.vue'


const props = defineProps({
  EditorType: {
    type: String,
    default: "md",
  },
})

const emit = defineEmits(['onSelectoEditor'])
const onSelectoEditor = (type) => {
  emit('onSelectoEditor', type)
}


const onSaveContent = (data) => {

}

const onPreview = (data) => {

}

</script>


<template>
  <div class="editor-container">
    <el-affix :offset="0" target=".editor-container">
      <div class="editor-header">
        <el-button size="small" @click="onSaveContent">保存</el-button>
        <el-button size="small" @click="onPreview">预览</el-button>
      </div>
    </el-affix>
    <word v-if="props.EditorType === 'word'" @click="onSelectoEditor('word')" />
    <md v-else-if="props.EditorType === 'md'" @click="onSelectoEditor('md')" />
  </div>
</template>

<style src="./style.scss"></style>