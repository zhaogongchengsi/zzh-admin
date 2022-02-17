<template>
  <div>
    <a-upload
      list-type="picture-card"
      action="/"
      :default-file-list="files"
      image-preview
    />
  </div>
</template>

<script lang="ts" setup>

const files:Array<File> = []
</script>