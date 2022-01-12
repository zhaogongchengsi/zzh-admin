<script setup>
import { ref, reactive } from 'vue'
import { ElMessageBox, ElMessage } from 'element-plus'
import { createUser } from '@/api/user.js'

const dialogVisible = ref(false)
const adminUser = reactive({
  useradmin: '',
  username: '',
  password: '',
  nickname: '',
  authorityId: '',
  headerImg:''
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

const addRootUser = () => {
  dialogVisible.value = true;
}

const Register = () => {
  createUser(adminUser)
  .then((result) => {
    ElMessage({
      message: `添加成功`,
      type: 'success',
    })
    dialogVisible.value = false
    console.log(result)
  })
  .catch((err) =>{
    ElMessage({
      message: `添加失败`,
      type: 'error',
    })
    dialogVisible.value = false
  })
}

</script>


<template>
  <div class="user-container">
    <div class="user-header">
      <el-button @click="addRootUser" type="primary">新增用户</el-button>
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
              :model="adminUser"
            >
              <el-form-item label="账号">
                <el-input v-model="adminUser.useradmin"></el-input>
              </el-form-item>
              <el-form-item label="用户名">
                <el-input v-model="adminUser.username"></el-input>
              </el-form-item>
              <el-form-item label="昵称">
                <el-input v-model="adminUser.nickname"></el-input>
              </el-form-item>
              <el-form-item label="密码">
                <el-input v-model="adminUser.password"></el-input>
              </el-form-item>
              <el-form-item label="用户角色">
                <el-input v-model="adminUser.authorityId"></el-input>
              </el-form-item>
              <el-form-item label="头像">
                <el-input v-model="adminUser.headerImg"></el-input>
              </el-form-item>
            </el-form>
        </div>
        <template #footer>
          <span class="dialog-footer">
            <el-button @click="dialogVisible = false">取消</el-button>
            <el-button type="primary" @click="Register">添加</el-button>
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