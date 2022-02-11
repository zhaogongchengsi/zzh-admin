<template>
  <div :class="['create-container']" v-dark:dark="menuStore.theme">
    <v-md-editor 
      v-model="mdString" 
      height="100%"
      mode="edit"
      @save="save"
    ></v-md-editor>
  </div>
</template>

<script setup lang="ts">
import { ref } from "vue";
const mdString = ref("");
import { useMenuStore } from "@/pinia";
const menuStore = useMenuStore();
const save = (text:string, html:string) => {
  console.log(text, html)
}

</script>

<style lang="scss" scoped>
.create-container {
  height: 100%;
}

.create-container-dark:deep(.v-md-editor),
.create-container-dark:deep(.v-md-textarea-editor textarea) {
  background-color: var(--color-menu-dark-bg);
  color: #fff;
}

</style>