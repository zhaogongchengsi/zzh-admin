<script setup>
import { ref, reactive, onMounted } from 'vue'
import { ElMessageBox, ElMessage } from 'element-plus'
import icons from "@/utils/Icons"
import { menuRules } from '@/utils/validate'
import { createMenu } from '@/api/menus'
import { useStore } from 'vuex'
const store = useStore()
const { solid, outline } = icons
    const menuFrom = ref(null)
    const dialogVisible = ref(false)
    const menuData = reactive({
        parent_id: 0,
        component: "",
        menu_path: "",
        menu_name: "",
        menu_icon: "",
        menu_disable: true,
        Label: "",
        sort: 0,
        remarks: ""
    })
    
    const iconArray = ref([])
    const MenuDataTree = ref([])
    const drawer = ref(false)
    onMounted(() => {
       const originlRoutData = store.state.router.OriginlRoutData
       MenuDataTree.value = originlRoutData
       console.log(originlRoutData)
    })

    const handleClose = (done) => {
      ElMessageBox.confirm('是否取消添加菜单',{
            cancelButtonText: "继续添加",
            confirmButtonText: "取消"
        })
        .then(() => {
          done()
        })
        .catch(() => {
          // catch error
        })
    }

    const handleCloseIcon = (done) => {
        ElMessageBox.confirm( "是否放弃添加菜单图标？", {
            cancelButtonText: "添加",
            confirmButtonText: "取消"
        })
        .then(() => {
          done()
        })
        .catch(() => {
          // catch error
        })
    }

    const openDialog = (done) => {
        dialogVisible.value = true
    }

    const addMenu = () => {
        menuFrom.value.validate((fromState) => {
            if (fromState) {
                createMenu(menuData)
                .then(res => {


                    ElMessage.success('菜单创建成功')
                    dialogVisible.value = false
                })
                .catch(err => {
                    ElMessage.error('菜单创建失败')
                })
            } else {
                ElMessage.error('用户信息不完整')
            }
        })
    }

    const openDrawer = (type) => {
        iconArray.value = icons[type]
        drawer.value = true;
    }

    const seleckIcon = (name) => {
        menuData.menu_icon = name;
        drawer.value = false;
    }

    const handleEdit = (index, Id) => {
        console.log("编辑的数据",index, Id)
    }

    const handleDelete = (index, Id) => {
        

        ElMessageBox.confirm(
            '删除该节点后不可恢复且连所属子节点一起删除，是否继续删除',
            'Warning',
            {
                confirmButtonText: '继续删除',
                cancelButtonText: '取消删除',
                type: 'warning',
            }
        )
        .then(() => {
          ElMessage({
            type: 'success',
            message: '删除成功',
          })
        })
        .catch(() => {
          // 取消删除
        })

    }

    const handleAddMenu = (index, Id) => {
        console.log("增加子节点")
    }
</script>


<template>
    <div class="menus-container">
          <div class="menus-header">
              <el-button @click="openDialog" type="primary">新增根菜单</el-button>
          </div>

        <el-table
            :data="MenuDataTree"
            style="width: 100%; margin-bottom: 20px"
            row-key="ID"
            border
            default-expand-all
            :stripe="true"
        >   
            <el-table-column prop="Sort" label="排序" sortable width="80" align="center" />
            <el-table-column prop="ParentId" label="父节点" width="80" align="center" />
            <el-table-column prop="ID" label="ID" width="80" align="center" />
            <el-table-column prop="Path" label="路由路径"  width="180" />
            <el-table-column prop="Name" label="路由名称"  width="180" />
            <el-table-column prop="Component" label="文件路径"  width="250" />
            <el-table-column prop="Icon" label="菜单图标"  width="80" >
                <template #default="scope">
                   <div class="table-icon-cont">
                        <div class="table-icon"><el-icon> <component :is="solid[scope.row.Icon]"></component> </el-icon></div>
                        <div class="table-icon"><el-icon> <component :is="outline[scope.row.Icon]"></component> </el-icon></div>
                   </div>
                </template>
            </el-table-column>
            <el-table-column prop="Disabled" label="是否开启"  width="80" >
                <template #default="scope">
                     <el-switch v-model="scope.row.Disabled" />
                </template>
            </el-table-column>
            <el-table-column prop="Label" label="展示的名称"  width="100" />
            <el-table-column prop="Remarks" label="备注"  width="auto" />
                <el-table-column align="center"  label="操作"  width="300" >
                <template #default="scope">
                    <el-button size="mini" @click="handleEdit(scope.$index, scope.row.ID)">编辑</el-button>
                    <el-button size="mini" @click="handleAddMenu(scope.$index, scope.row.ID)">添加子节点</el-button>
                    <el-button
                        size="mini"
                        type="danger"
                        @click="handleDelete(scope.$index, scope.row.ID)"
                    >删除节点</el-button>
                </template>
            </el-table-column>
        </el-table>

        <el-dialog
            v-model="dialogVisible"
            title="编辑菜单"
            width="40%"
            :before-close="handleClose"
        >
              <el-form
                label-position="right"
                label-width="100px"
                :model="menuData"
                :rules="menuRules"
                ref="menuFrom"
            >
                <el-form-item label="父节点">
                    <el-input v-model.number="menuData.parent_id" placeholder="请输入父节点ID" ></el-input>
                </el-form-item>
                <el-form-item label="文件路径" prop="component">
                    <el-input v-model="menuData.component" placeholder="请输入组件的文件路径(输入英文字母)"></el-input>
                </el-form-item>
                <el-form-item label="路由路径" prop="menu_path">
                    <el-input v-model="menuData.menu_path" placeholder="请输入组件的路由路径(输入英文字母)"></el-input>
                </el-form-item>
                <el-form-item label="路由名称" prop="menu_path">
                    <el-input v-model="menuData.menu_name" placeholder="请输入组件的路由名字"></el-input>
                </el-form-item>
                <el-form-item label="图标">
                    <el-row :gutter="24">
                            <el-col :span="14"><el-input v-model="menuData.menu_icon" disabled  placeholder="请输入组件展示的名称"></el-input></el-col>
                            <el-col :span="5"><el-button type="primary" @click="openDrawer('solids')" >实体图标</el-button></el-col>
                            <el-col :span="5"><el-button type="primary" @click="openDrawer('outlines')" >线体图标</el-button></el-col>
                      </el-row>
                </el-form-item>
                <el-form-item label="展示名" prop="Label">
                      <el-input v-model="menuData.Label" placeholder="请输入组件的展示名称"></el-input>
                </el-form-item>
                <el-form-item label="是否禁用" prop="menu_disable">
                     <el-switch v-model="menuData.menu_disable" />
                </el-form-item>
                <el-form-item label="排序标记" prop="sort">
                     <el-input v-model.number="menuData.sort" placeholder="请输入排序标记"></el-input>
                </el-form-item>
                <el-form-item label="备注">
                    <el-input v-model="menuData.remarks" placeholder="有什么需要备注的吗？"></el-input>
                </el-form-item>
            </el-form>
            <template #footer>
            <span class="dialog-footer">
                <el-button @click="dialogVisible = false">取消</el-button>
                <el-button type="primary" @click="addMenu">添加</el-button>
            </span>
            </template>
        </el-dialog>

          <el-drawer
            v-model="drawer"
            title="请选择图标"
            direction="rtl"
            :before-close="handleCloseIcon"
        >
          <ul class="icon-drawer">
              <li v-for="(icon, index) of iconArray" :key="index" class="icon-li" @click="seleckIcon(icon.iconName)">
                    <el-popover
                        placement="top-start"
                        trigger="hover"
                        :content="icon.iconName"
                    >
                        <template #reference>
                            <div class="icon-component"><el-icon> <component :is="icon.component"></component> </el-icon></div>
                        </template>
                    </el-popover>
              </li>
          </ul>
        </el-drawer>
    </div>
</template>

<style scoped lang="scss">


.menus-container:deep(.el-drawer) {
    overflow:auto !important;
}
.icon-drawer {
    display: flex;
    flex-wrap: wrap;
    justify-content: space-around;
    overflow: auto;
    .icon-li {
        box-sizing: border-box;
        padding: 10px;
        .icon-component {
            width: 50px;
            height: 50px;
            display: flex;
            align-items: center;
            justify-content: center;
            border: 1px solid #ccc;
            border-radius:10px;
            cursor: pointer;
        }
        .icon-name {
            text-align: center;
        }
    }
}
.menus-container {
    padding: 10px;
    box-sizing: border-box;
    .menus-header {
        margin-bottom: 10px;
    }
}

.table-icon-cont {
    display: flex;
    justify-content:space-around;
}
</style>