<script setup>
import MdEditor from 'md-editor-v3';
import 'md-editor-v3/lib/style.css';
import { ref } from 'vue';
// import sanitizeHtml from 'sanitize-html'; 防止xxs 

const props = defineProps({
  filterXxs: {
    type: Boolean,
    default: false
  }
})

const emit = defineEmits(['onSave'])


const text = ref('')

const toolbars = [
  'bold',
  'underline',
  'italic',
  '-',
  'strikeThrough',
  'sub',
  'sup',
  'quote',
  'unorderedList',
  'orderedList',
  '-',
  'codeRow',
  'code',
  'link',
  'image',
  'table',
  'mermaid',
  'katex',
  '-',
  'revoke',
  'next',
  'save',
  '=',
  'pageFullscreen',
  'fullscreen',
  'preview',
  'htmlPreview',
  'catalog',
  '加粗',
  '下划线',
  '斜体',
  '删除线',
  '下标',
  '上标',
  '引用',
  '无序列表',
  '有序列表',
  '行内代码',
  '块级代码',
  '链接',
  '图片',
  '表格',
  '图表',
  '公式',
  '后退一步',
  '前进一步',
  '保存',
  '页面内全屏',
  '屏幕全屏',
  '内容预览',
  'html代码预览',
  '目录',
  '源码地址'
]

const save = () => {
  emit("onSave", text.value, 'md')
}

const upImg = (files, callback) => {

  console.log(files)

  //  callback(res.map((item) => {
     
  //  }));
}

const Change = (html) => {
  text.value = html
}

const sanIt = (html) => {
  let _html = sanitizeHtml(html)
  return _html
}

</script>

<template>
  <div>
    <md-editor
      v-model="text"
      previewTheme="vuepress"
      :toolbars="toolbars"
      editorClass="editor-blog"
      @onSave="save"
      @onUploadImg="upImg"
      @onChange="Change"
    />
  </div>
</template>