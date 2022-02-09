<script setup>
import { onMounted, ref } from "vue"
import { getArticles } from "@/api/article.js"
import { dateFormat } from "@/utils/index.js"

const page = ref({})
// CreatedAt: "2022-02-09T10:45:01+08:00"
// DeletedAt: null
// ID: 2
// UpdatedAt: "2022-02-09T10:45:01+08:00"
// article_author: "zzh"
// article_context: " "
// article_file_name: "blog"
// article_name: "上传服务器"
// article_storage_type: "services"
// article_tags: null
// article_title: "上传服务器"
// article_type: "Golang"
// article_url: "assets/upload/2022/February/"
// likes: 0
// number_views: 0
const data = ref([])
onMounted(async () => {
  getdata()
})

async function getdata () {
  const res = await getArticles(3,1)
  page.value = {
    limit_offset: res.limit_offset,
    count: res.count
  }
  data.value = res.article_list
  console.log(page.value, data.value)
}


</script>

<template>
  <div class="artic-container">
      <el-table
        :data="data"
        stripe
        style="width: 100%">
        <el-table-column
          prop="CreatedAt"
          label="创建时间"
          width="200"
        >
          <template #default="scope">
            <div class="table-icon">{{dateFormat(scope.row.CreatedAt)}}</div>
          </template>
        </el-table-column>
        <el-table-column
          prop="ID"
          label="ID"
          width="80">
        </el-table-column>
        <el-table-column
          prop="article_title"
          label="文章标题"
          width="200"
        >
        </el-table-column>
        <el-table-column
          prop="article_type"
          label="文章类型"
          width="100"
          >
        </el-table-column>
        <el-table-column
          prop="article_author"
          label="作者"
          width="100"
          >
        </el-table-column>
        <el-table-column
          prop="article_url"
          label="储存地址"
          >
        </el-table-column>
      </el-table>
      <el-pagination
        layout="prev, pager, next"
        :total="page.count">
      </el-pagination>
  </div>
</template>

<style lang="scss" scoped>
.artic-container {
  padding: 20px;
  box-sizing: border-box;
}
</style>