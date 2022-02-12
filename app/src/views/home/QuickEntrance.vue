<template>
  <div class="quick">
    <ul class="quick-ul">
      <li class="quick-item">
        <icon-font type="icon-bianji1" :size="size" />
        <span>创建文章</span>
      </li>
      <li class="quick-item">
        <icon-font type="icon-yuwen" :size="size" />
        <span>管理文章</span>
      </li>
      <li class="quick-item">
        <icon-font type="icon-sucai" :size="size" />
        <span>管理素材</span>
      </li>
      <li class="quick-item">
        <icon-font type="icon-shangchuan" :size="size" />
        <span>上传素材</span>
      </li>
    </ul>
  </div>
</template>

<script lang="ts" setup>
const size = 30;
</script>

<style lang="scss" scoped>
.quick {
  margin: 20px 0;
}
.quick-ul {
  display: flex;
  justify-content: space-around;
  .quick-item {
    width: 100px;
    height: 100px;
    border-radius: 5px;
    cursor: pointer;
    padding: 10px;
    box-sizing: border-box;
    color: #fff;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: space-evenly;
    &:hover {
      box-shadow: 0 2px 5px 0 rgba(141, 137, 137, 0.678);
    }
    &:nth-child(1) {
      background-image: radial-gradient(
        circle 248px at center,
        #16d9e3 0%,
        #30c7ec 47%,
        #46aef7 100%
      );
    }
    &:nth-child(2) {
      background-image: linear-gradient(
        45deg,
        #ff9a9e 0%,
        #fad0c4 99%,
        #fad0c4 100%
      );
    }
    &:nth-child(3) {
      background-image: linear-gradient(to top, #37ecba 0%, #72afd3 100%);
    }
    &:nth-child(4) {
      background-image: radial-gradient(
        circle 248px at center,
        #16d9e3 0%,
        #30c7ec 47%,
        #46aef7 100%
      );
    }
  }
}
</style>