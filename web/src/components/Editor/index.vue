<script setup>
import word from './word.vue'
import md from './md.vue'


const props = defineProps({
  EditorType: {
    type: String,
    default: "md",
  },
})

const emit = defineEmits(['onSave'])
const onSelectoEditor = (type) => {
  
}


const onSaveContent = (data, type) => {
  emit('onSave', data, type)
}

const onPreview = (data) => {

}

</script>


<template>
  <div class="editor-container">
    <word v-if="props.EditorType === 'word'" @onSave="onSaveContent" />
    <md v-else-if="props.EditorType === 'md'" @onSave="onSaveContent" />
  </div>
</template>

<style src="./style.scss"></style>