
<template>
  <div class="home-container" v-dark:dark="menuStore.theme">
    <top-home />
    <a-divider orientation="left">快捷入口</a-divider>
    <quick-entrance />
    <a-divider orientation="left">数据统计</a-divider>
    <a-row :gutter="24">
        <a-col :span="19">
            <statistic />
        </a-col>
        <a-col :span="5">
            <pian />
        </a-col>
    </a-row>
  </div>
</template>

<script lang="ts" setup>
import TopHome from './TopHome.vue'
import QuickEntrance from './QuickEntrance.vue'
import pian from './plan.vue'
import statistic from './Statistics.vue'
import { useMenuStore } from "@/pinia";
const menuStore = useMenuStore();
const ellipsis = {
  rows: 2,
  showTooltip: true,
};
</script>

<style lang="scss" scoped>
.home-container {
  box-sizing: border-box;
  padding: 15px;
}
.home-img {
  img {
    width: 100%;
  }
}

.home-card {
    width: 100%;
    height: 100px;
    border-radius: 5px;
    box-sizing: border-box;
    padding: 10px;
    &:hover {
        box-shadow: 0 2px 5px 0 rgba(141, 137, 137, 0.678);
    }

    .card-title {
        color: #fff;
        font-weight: bold;
        font-size: 20px;
    }
}
.bg-color1 {
    background-image: radial-gradient(circle 248px at center, #16d9e3 0%, #30c7ec 47%, #46aef7 100%);
}
.bg-color2 {
    background-image: linear-gradient(to right, #eea2a2 0%, #bbc1bf 19%, #57c6e1 42%, #b49fda 79%, #7ac5d8 100%);
}
.bg-color3 {
   background-image: linear-gradient(to top, #37ecba 0%, #72afd3 100%);
}
</style>