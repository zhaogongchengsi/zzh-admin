<template>
  <div class="todolist">
    <a-divider orientation="left">{{props.title}}</a-divider>
    <a-timeline labelPosition="relative">
      <a-timeline-item
        v-for="item in props.data"
        :key="item.id"
        :label="formatTime(item.planDate)"
        lineType="dashed"
      >
        <template #dot>
          <icon-font type="icon-yiwancheng-" v-if="item.isDone" />
          <icon-font type="icon-weiwancheng-" v-else />
        </template>
        <div class="todolist-item">
          <div class="todo-left">
            <span class="todolist-name">{{item.planName}}</span>

            <a-typography-text
              :ellipsis="ellipsis"
              :type="item.isDone === true ? 'success' : 'danger' "
            >{{item.planDescription}}</a-typography-text>
          </div>
          <div class="todo-right">
            <a-progress type="circle" status='success' :percent="1" size="small"/>
          </div>
        </div>
      </a-timeline-item>
    </a-timeline>
  </div>
</template>

<script lang="ts" setup>
import { defineProps } from "vue";
import { plan } from "./index";
import { formatTime } from "@/utils/index";
const props = defineProps({
  title: String,
  data: Array as () => Array<plan>,
});
const ellipsis = {
  rows: 2,
}
</script>

<style lang="scss" scoped>
.todolist {
  // height: 100vh;

  .todolist-title {
    color: var(--color-text-1);
  }

  .todolist-name {
    color: var(--color-text-1);
    font-size: 18px;
    display: block;
  }
  .todolist-item {
    width: 80%;
    border: 1px solid var(--color-border-3);
    box-sizing: border-box;
    border-radius: var(--border-radius-medium);
    padding: 10px;
    display: flex;
    justify-content: space-between;
    &:hover {
      box-shadow: 0 2px 5px 0 rgba(141, 137, 137, 0.678);
      cursor: pointer;
    }
  }

  .todo-left {
    padding-right: 30px;
    box-sizing: border-box;
  }

  .todo-right {
    align-self: center;
  }
}
</style>