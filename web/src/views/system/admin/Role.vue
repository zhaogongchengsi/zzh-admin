<script setup>
import { ref, reactive } from 'vue'
import { createRole, getRole } from '@/api/role.js'
import { ElMessageBox, ElMessage } from 'element-plus'
import { onMounted } from 'vue'
import { DadLookSon } from '@/utils/index'

const dialogTitle = ref('新增角色')
const dialogVisible = ref(false)
const role = reactive({
  parent_role: 0,
  role_id: 0,
  role_name: '',
  role_remarks: ''
})
const roleList = ref([])

onMounted(async () => {
    try {
        let roles  = await getRole()
        let newdata = datahandler(roles)
        console.table(newdata)
        roleList.value = roles
    } catch {
    
    }
})


function datahandler(data) {
  let _data = JSON.parse(JSON.stringify(data));

   return _data.map(function (item) {
        return getSon(item, _data)
    })


  function getSon (pra, sonArr) {
      sonArr.forEach((sItem, si, arr) => {
          if (pra.AuthRoleID === sItem.ParentID) {
            let sonChild =  getSon(sItem, sonArr)
            if (sonChild.length > 0) {
                sItem.children = sonChild
              
            }
            if (pra.children) {
                pra.children.push(sItem)
            } else {
                pra.children = [sItem];
            }
              sonArr.splice(si, arr)
          }
      });
    return pra
  }


}

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

const handleEdit = () => {}
const handleAddMenu = () => {}
const handleDelete = () => {}

</script>


<template>
  <div class="auto-container">
    <div class="user-header">
      <el-button @click="addRootRole" type="primary">新增角色</el-button>
    </div>

    <div class="role-container">
        <el-table :data="roleList" stripe style="width: 100%" border>
            <el-table-column prop="AuthRoleID" label="角色ID" width="300" />
            <el-table-column prop="AuthRoleName" label="角色名称" width="300" />
            <el-table-column prop="AuthRoleRemarks" label="角色备注" width="300" />
            <el-table-column prop="address" label="操作" >
                <template #default="scope">
                    <el-button size="mini" type="text" @click="handleEdit(scope.$index, scope.row.ID)">分配权限</el-button>
                    <el-button size="mini" type="text" @click="handleAddMenu(scope.$index, scope.row.ID)">新增子角色</el-button>
                    <el-button
                        size="mini"
                        type="text"
                        @click="handleDelete(scope.$index, scope.row.ID)"
                    >删除角色</el-button>
                </template>
            </el-table-column>
        </el-table>
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

.role-container {
    margin: 10px 0;
}
</style>