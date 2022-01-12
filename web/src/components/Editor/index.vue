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


const onSaveContent = (data, type) => {

  console.log(data, type)
}

const onPreview = (data) => {

}

</script>


<template>
  <div class="editor-container">
    <word v-if="props.EditorType === 'word'" @click="onSelectoEditor('word')" />
    <md v-else-if="props.EditorType === 'md'" @click="onSelectoEditor('md')" @onSave="onSaveContent" />
  </div>
</template>

<style src="./style.scss"></style>