<script setup>
import { ref, reactive } from 'vue'
import { createRole } from '@/api/role.js'
import { ElMessageBox, ElMessage } from 'element-plus'

const dialogTitle = ref('新增角色')
const dialogVisible = ref(false)
const role = reactive({
  parent_role: 0,
  role_id: 0,
  role_name: '',
  role_remarks: ''
})

const handleClose = (done) => {
  ElMessageBox.confirm(`是否取消${dialogTitle.value}`)
    .then(() => {
      done()
    })
    .catch(() => {
      // catch error
    })
}

const addRootRole = () => {
  dialogVisible.value = true;
}

const confirmHandler = () => {

    createRole(role)
    .then(data => {
        ElMessage({
            message: '角色创建成功',
            type: 'success',
        })
    })
    .catch(error => {
        console.log('error', error);
    })
}

</script>


<template>
  <div class="auto-container">
    <div class="user-header">
      <el-button @click="addRootRole" type="primary">新增角色</el-button>
    </div>

      <el-dialog
        v-model="dialogVisible"
        :title="dialogTitle"
        width="40%"
        :before-close="handleClose"
      >
        <div class="content">
            <el-form
              label-position="right"
              label-width="100px"
              :model="role"
            >
              <el-form-item label="父角色ID">
                <el-input v-model.number="role.parent_role"></el-input>
              </el-form-item>
              <el-form-item label="角色ID">
                <el-input v-model.number="role.role_id"></el-input>
              </el-form-item>
              <el-form-item label="角色名">
                <el-input v-model="role.role_name"></el-input>
              </el-form-item>
              <el-form-item label="备注">
                <el-input v-model="role.role_remarks"></el-input>
              </el-form-item>
            </el-form>
        </div>
        <template #footer>
          <span class="dialog-footer">
            <el-button @click="dialogVisible = false">取消</el-button>
            <el-button type="primary" @click="confirmHandler">添加</el-button>
          </span>
        </template>
      </el-dialog>
  </div>
</template>

<style lang="scss" scoped>

.auto-container {
  box-sizing: border-box;
  padding: 10px;
}
</style>