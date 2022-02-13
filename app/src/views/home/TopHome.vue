<template>
  <a-row align="center" :gutter="30">
    <a-col :span="18">
      <div class="left-cards">
        <div>
          <a-typography-title :heading="5">哟！ 来啦</a-typography-title>
          <a-typography-paragraph :ellipsis="ellipsis">又来给博客添砖加瓦了？今天写点啥呢？</a-typography-paragraph>
        </div>
        <div class="card-list">
          <a-row :gutter="24">
            <a-col :span="8">
              <div class="home-card color-change-1x">
                <h4 class="card-title">
                  
                  <a-statistic title="留言数" :value="125670" show-group-separator>
                    <template #suffix>
                     <icon-font type="icon-liuyan" />
                    </template>
                  </a-statistic>
                </h4>
              </div>
            </a-col>
            <a-col :span="8">
              <div class="home-card color-change-2x">
                <h4 class="card-title">
                  
                  <a-statistic title="点赞数" :value="125670" show-group-separator>
                    <template #suffix>
                      <icon-font type="icon-mianxingdianzan" />
                    </template>
                  </a-statistic>
                </h4>
              </div>
            </a-col>
            <a-col :span="8">
              <div class="home-card color-change-3x">
                <h4 class="card-title">
              
                  <a-statistic title="访问量" :value="125670" show-group-separator>
                    <template #suffix>
                      <icon-font type="icon-jiankongchakanshenqing" />
                    </template>
                  </a-statistic>
                </h4>
              </div>
            </a-col>
          </a-row>
        </div>
      </div>
    </a-col>
    <a-col :span="6">
      <div class="home-img">
        <img src="/bg3.svg" alt />
      </div>
    </a-col>
  </a-row>
</template>
<script lang="ts" setup>
const ellipsis = {
  rows: 2,
  showTooltip: true,
};
</script>
<style src="./animista.css" ></style>
<style lang="scss" scoped>
.home-img {
  height: 100%;
  img {
    // width: 100%;
    height: 100%;
  }
}
.left-cards {
  display: flex;
  flex-direction: column;
  height: 100%;
  .card-list {
    flex: 1;
  }
}
.home-card {
  width: 100%;
//   height: 100px;
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
    --color-text-2: #fff;
    --color-text-1: #fff;
  }
}
.bg-color1 {
  background-image: radial-gradient(
    circle 248px at center,
    #16d9e3 0%,
    #30c7ec 47%,
    #46aef7 100%
  );
}
.bg-color2 {
  background-image: linear-gradient(
    to right,
    #eea2a2 0%,
    #bbc1bf 19%,
    #57c6e1 42%,
    #b49fda 79%,
    #7ac5d8 100%
  );
}
.bg-color3 {
  background-image: linear-gradient(to top, #37ecba 0%, #72afd3 100%);
}
</style>