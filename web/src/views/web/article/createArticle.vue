<script setup>
import { ref, reactive, watchEffect } from 'vue'
import Editor from '@/components/Editor/index.vue'
import { ElMessageBox, ElMessage } from 'element-plus'
import { uploadArticle } from '@/api/article.js'

const activeName = ref('Md')
const dialogVisible = ref(false)
const article = reactive({
  fileName: '',
  articleName: "",
  articleTitle: "",
  articleUrl: "",
  articleStorageType: "",
  articleAuthor: "",
  articleType: "",
  articleContext: "",
  articleTags: []
})
const osConfig = reactive({
  osPath: []
})
const osOptions = ref([
  {
    value: "bloghtml-1301735126",
    label: "blog",
    children: [
      {
        value: "shili",
        label: "shili"
      }
    ]
  },
])
const activeText = ref('')
const activeType = ref('')
const innerVisible = ref(false)

const handleClick = (tab, event) => {
  // console.log(tab, event)
}

const handleClose = (done) => {
  ElMessageBox.confirm('文章还未保存是否取消保存', {
      confirmButtonText: '确认',
      cancelButtonText: '取消',
    })
    .then(() => {
      done()
    })
    .catch(() => {
      // catch error
    })
}

const ooshandleClose = (done) => {
  ElMessageBox.confirm('对象存储还未选择配置是否取消选择？', {
      confirmButtonText: '确认',
      cancelButtonText: '取消',
    })
    .then(() => {
      done()
    })
    .catch(() => {
      // catch error
    })
}

const openSave = async (file, type) => {
  activeText.value = file
  activeType.value = type
  article.articleContext = file
  dialogVisible.value = true
}

const saveHandle = async () => {
  const res = await uploadArticle(article,{
    oos: osConfig.osPath
  })
  if (res) {
    dialogVisible.value = false
    ElMessage({
      message: '文章上传成功',
      type: 'success',
    })
  }
}

const selectChange = () => {
  if (article.articleStorageType === "oos") {
    innerVisible.value = true
  }
}

const handleChange = (value) => {
  osConfig.value = value.join("/")
}

const saveOsHandle = () => {
  innerVisible.value = false
}

</script>


<template>
  <div class="article-con">
      <el-tabs v-model="activeName" @tab-click="handleClick">
        <el-tab-pane label="Word" name="Word">
          <Editor EditorType="word"  />
        </el-tab-pane>
        <el-tab-pane label="Md" name="Md">
          <Editor EditorType="md"  @onSave="openSave" />
        </el-tab-pane>
      </el-tabs>
      <el-dialog
        v-model="dialogVisible"
        title="保存文章"
        width="40%"
        :before-close="handleClose"
      >
        <div class="content">
            <el-form
              label-position="right"
              label-width="100px"
              :model="article"
            >
              <el-form-item label="文件名字">
                <el-input v-model="article.fileName"></el-input>
              </el-form-item>
              <el-form-item label="文章标题">
                <el-input v-model="article.articleTitle"></el-input>
              </el-form-item>
              <el-form-item label="文章名字">
                <el-input v-model="article.articleName"></el-input>
              </el-form-item>
              <el-form-item label="文章作者">
                <el-input v-model="article.articleAuthor"></el-input>
              </el-form-item>
              <el-form-item label="保存类型">
                  <el-select v-model="article.articleStorageType" placeholder="选择保存类型" @change="selectChange">
                    <el-option label="服务器" value="services">服务器</el-option>
                    <el-option label="数据库" value="database">数据库</el-option>
                    <el-option label="对象存储" value="oos">对象存储</el-option>
                  </el-select>
              </el-form-item>
               <el-form-item label="标签">
                  <el-select v-model="article.articleTags" placeholder="选择保存类型" multiple>
                    <el-option label="Html" value="Html">Html</el-option>
                    <el-option label="Css" value="Css">Css</el-option>
                    <el-option label="Javascript" value="Javascript">Javascript</el-option>
                  </el-select>
              </el-form-item>
              <el-form-item label="文章类型">
                  <el-select v-model="article.articleType" placeholder="选择保存类型">
                    <el-option label="Html" value="Html">Html</el-option>
                    <el-option label="Css" value="Css">Css</el-option>
                    <el-option label="Javascript" value="Javascript">Javascript</el-option>
                    <el-option label="Golang" value="Golang">Golang</el-option>
                    <el-option label="Node.js" value="Node.js">Node.js</el-option>
                  </el-select>
              </el-form-item>
            </el-form>
            <el-dialog
              v-model="innerVisible"
              width="30%"
              title="cos选择"
              append-to-body
              :before-close="ooshandleClose"
            >
            <el-form
              label-position="right"
              label-width="200px"
              :model="osConfig"
            >
            <el-form-item label="对象存储文件路径">
                <el-cascader
                  v-model="osConfig.osPath"
                  :options="osOptions"
                  @change="handleChange"
                ></el-cascader>
              </el-form-item>
            </el-form>
            <template #footer>
              <span class="dialog-footer">
                <el-button @click="innerVisible = false">取消</el-button>
                <el-button type="primary" @click="saveOsHandle">确认</el-button>
              </span>
            </template>
            </el-dialog>
        </div>
        <template #footer>
          <span class="dialog-footer">
            <el-button @click="dialogVisible = false">取消</el-button>
            <el-button type="primary" @click="saveHandle">上传</el-button>
          </span>
        </template>
      </el-dialog>

  </div>
</template>

<style lang="scss" scoped>

.article-con {
  padding: 10px;
  box-sizing: border-box;
}

</style>