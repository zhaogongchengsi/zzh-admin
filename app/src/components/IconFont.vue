<template>
    <icon-font :size="props.size" :type="props.type" :spin="props.spin" />
</template>
<script setup lang="ts">
import { Icon } from '@arco-design/web-vue';
import { defineProps } from 'vue';
const props = defineProps({
    type: String,
    size: Number,
    spin:{
        type: Boolean,
        default: false
    }
})
const IconFont = Icon.addFromIconFontCn({ src: 'https://at.alicdn.com/t/font_3178565_a9xd4lsdgje.js' });
</script>