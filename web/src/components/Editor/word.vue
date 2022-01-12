
<script setup>
import { computed, onBeforeUnmount, ref } from 'vue'
import { Editor, Toolbar, getEditor, removeEditor } from '@wangeditor/editor-for-vue'
import cloneDeep from 'lodash.clonedeep'



const editorId = `w-e-${Math.random().toString().slice(-5)}`
const emit = defineEmits(['onSave'])

const defaultContent = []
const getDefaultContent = computed(() => cloneDeep(defaultContent)) // 注意，要深拷贝 defaultContent ，否则报错

const toolbarConfig = {}
const editorConfig = { placeholder: '请输入内容...' }

// 组件销毁时，也及时销毁编辑器
onBeforeUnmount(() => {
    const editor = getEditor(editorId)
    if (editor == null) return

    editor.destroy()
    removeEditor(editorId)
})

const getSaveText = () => {
  emit('onSave',getDefaultContent.value, "word")
}



</script>

<template>
  <div>
    <div class="editor-header">
        <el-button size="small" @click="getSaveText">保存</el-button>
        <el-button size="small" >预览</el-button>
    </div>
    <div style="border: 1px solid var(--el-border-color-base)">
      <Toolbar
        :editorId="editorId"
        :defaultConfig="toolbarConfig"
        mode="default"
        style="border-bottom: 1px solid var(--el-border-color-base)"
      />
      <Editor
        :editorId="editorId"
        :defaultConfig="editorConfig"
        :defaultContent="getDefaultContent"
        mode="default"
        style="min-height: var(--editor-min-height)"
      />
    </div>
  </div>
</template>

<style src="@wangeditor/editor/dist/css/style.css"></style>