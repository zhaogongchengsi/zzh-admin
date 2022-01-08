<script setup>
import { ref, reactive } from 'vue'
// import { ElMessageBox } from 'element-plus'

const dialogVisible = ref(false)
const formLabelAlign = reactive({
  name: '',
  region: '',
  type: '',
})

const handleClose = (done) => {
  ElMessageBox.confirm('Are you sure to close this dialog?')
    .then(() => {
      done()
    })
    .catch(() => {
      // catch error
    })
}

const addRootMenu = () => {
  dialogVisible.value = true;
}

</script>


<template>
  <div class="user-container">
    <div class="user-header">
      <el-button @click="addRootMenu" type="primary">新增用户</el-button>
    </div>

      <el-dialog
        v-model="dialogVisible"
        title="添加用户"
        width="40%"
        :before-close="handleClose"
      >
        <div class="content">
            <el-form
              label-position="right"
              label-width="100px"
              :model="formLabelAlign"
            >
              <el-form-item label="用户名">
                <el-input v-model="formLabelAlign.name"></el-input>
              </el-form-item>
              <el-form-item label="密码">
                <el-input v-model="formLabelAlign.region"></el-input>
              </el-form-item>
              <el-form-item label="别名">
                <el-input v-model="formLabelAlign.type"></el-input>
              </el-form-item>
              <el-form-item label="用户角色">
                <el-input v-model="formLabelAlign.type"></el-input>
              </el-form-item>
              <el-form-item label="头像">
                <el-input v-model="formLabelAlign.type"></el-input>
              </el-form-item>
            </el-form>
        </div>
        <template #footer>
          <span class="dialog-footer">
            <el-button @click="dialogVisible = false">取消</el-button>
            <el-button type="primary" @click="dialogVisible = false">添加</el-button>
          </span>
        </template>
      </el-dialog>
  </div>
</template>

<style lang="scss" scoped>

.user-container {
  box-sizing: border-box;
  padding: 10px;
}
</style>