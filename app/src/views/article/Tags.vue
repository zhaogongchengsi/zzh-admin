<template>
  <div class="tags-container">
    <div class="tag-header">
      <a-button type="primary" @click="createTga">
        <template #icon>
          <icon-plus />
        </template>
        <template #default>添加标签</template>
      </a-button>
    </div>
    <a-list>
      <template #header>标签管理</template>
      <a-list-item v-for="t in taglist" :key="t.ID">
        <a-list-item-meta :title="t.tag" :description="t.tag_desc">
          <template #avatar>
            <div class="tag_color" :style="{backgroundColor: t.tag_color}">{{t.tag_color}}</div>
          </template>
        </a-list-item-meta>
        <a-collapse :default-active-key="[1]" accordion>
          <a-collapse-item header="标签下的文章" key="1">
            <div>标签文章</div>
          </a-collapse-item>
        </a-collapse>
      </a-list-item>
    </a-list>
    <a-modal v-model:visible="visible" @ok="handleOk" @cancel="visible = false">
      <template #title>添加标签</template>
      <div>
        <a-form :model="tagFrom" ref="fromIn">
          <a-form-item field="tag" label="标签名" :rules="[{required:true,message:'请输入标签名（必填）'}]">
            <a-input v-model="tagFrom.tag" placeholder="请输入标签名" />
          </a-form-item>
          <a-form-item field="tag_desc" label="标签描述">
            <a-input v-model="tagFrom.tag_desc" placeholder="请输入标签描述" />
          </a-form-item>
          <a-form-item field="tag_color" label="标签颜色">
            <div>
              <input v-model="tagFrom.tag_color" placeholder="请输入标签颜色" type="color" />
            </div>
          </a-form-item>
        </a-form>
      </div>
    </a-modal>
  </div>
</template>

<script lang="ts" setup>
import { getTagList, createTag } from "@/api/article";
import { onMounted, ref, reactive } from "vue";
import { IconPlus } from "@arco-design/web-vue/es/icon";
import { tags } from "@/types/request";
const fromIn = ref();
const visible = ref<boolean>(false);
const tagFrom = reactive({
  tag: "",
  tag_color: "#37D4CF",
  tag_desc: "",
  ID: 0
});
const taglist = ref<tags[]>([]);
onMounted(async () => {
  try {
    const res = await getTagList(10, 0);
    taglist.value = res.tag_list;
  } catch (e) {
    console.error(e);
  }

});

const createTga = () => (visible.value = true);

const handleOk = async () => {
  const res = await createTag(tagFrom);
  console.log(res);
};
</script>

<style lang="scss" scoped>
.tags-container {
  padding: 15px;
  box-sizing: border-box;
}
.tag-header {
  margin-bottom: 15px;
}

.select_color {
  cursor: pointer;
}

.tag_color {
  width: 85px;
  height: 50px;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  border-radius: 5px;
}
</style>