<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
        <el-form-item label="创建时间">
          <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始时间"></el-date-picker>
          —
          <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束时间"></el-date-picker>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <!-- <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openDialog">新增</el-button>
        <el-popover v-model:visible="deleteVisible" placement="top" width="160">
          <p>确定要删除吗？</p>
          <div style="text-align: right; margin-top: 8px;">
            <el-button type="primary" link @click="deleteVisible = false">取消</el-button>
            <el-button type="primary" @click="onDelete">确定</el-button>
          </div>
          <template #reference>
            <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length"
              @click="deleteVisible = true">删除</el-button>
          </template>
        </el-popover>
      </div> -->
      <el-table ref="multipleTable" style="width: 100%" tooltip-effect="dark" :data="tableData" row-key="ID"
        @selection-change="handleSelectionChange">
        <!-- <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="日期" width="180">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column> -->
        <!-- <el-table-column align="left" label="用户ID" prop="userId" width="120" /> -->
        <!-- 用户名称 -->
        <!-- <el-table-column align="left" label="用户" prop="userId" width="120">
          <template #default="scope">
            {{ scope.row.user.nickname }}
          </template>
        </el-table-column> -->
        <!-- 装备分类名称 -->
        <el-table-column align="left" label="分类名称" prop="categoryId" width="120">
          <template #default="scope">
            {{ scope.row.name }}
          </template>
        </el-table-column>

        <!-- 装备名称 -->
        <el-table-column align="left" label="装备一" prop="" width="120">
          <template #default="scope">
            <div v-if="scope.row.children.length > 0">
              <div v-if="scope.row.children[0].status == 2">
                <el-select v-model="scope.row.children[0].ID" class="m-2" placeholder="无" size="large"
                  @change="selectEquipment">
                  <el-option v-for="item in scope.row.children" :key="item.ID" :label="item.name" :value="item.ID" />
                </el-select>
              </div>
              <div v-else>
                <el-select v-model="scope.row.children.ID" class="m-2" placeholder="无" size="large"
                  @change="selectEquipment">
                  <el-option v-for="item in scope.row.children" :key="item.ID" :label="item.name" :value="item.ID" />
                </el-select>
              </div>
            </div>
            <!-- <el-select v-model="scope.row.children[0].ID" class="m-2" placeholder="无" size="large">
              <el-option v-for="item in scope.row.children" :key="item.ID" :label="item.name" :value="item.ID" />
            </el-select> -->
          </template>
        </el-table-column>
        <!-- 装备图标 -->
        <el-table-column align="left" label="装备一图标" prop="equipmentId" width="120">
          <template #default="scope">
            <img v-if="scope.row.children.length > 0 && scope.row.children[0].status == 2"
              :src="imagePath + (scope.row.children[0].icon)" alt="" style="width: 50px; height: 50px;">
          </template>
        </el-table-column>

        <!-- 装备名称 -->
        <el-table-column align="left" label="装备二" prop="" width="120">
          <template #default="scope">
            <div v-if="scope.row.children.length > 0">
              <div v-if="scope.row.children[1].status == 2">
                <el-select v-model="scope.row.children[1].ID" class="m-2" placeholder="无" size="large"
                  @change="selectEquipment">
                  <el-option v-for="item in scope.row.children" :key="item.ID" :label="item.name" :value="item.ID" />
                </el-select>
              </div>
              <div v-else>
                <el-select v-model="scope.row.children.ID" class="m-2" placeholder="无" size="large"
                  @change="selectEquipment">
                  <el-option v-for="item in scope.row.children" :key="item.ID" :label="item.name" :value="item.ID" />
                </el-select>
              </div>
            </div>
          </template>
        </el-table-column>
        <!-- 装备图标 -->
        <el-table-column align="left" label="装备二图标" prop="equipmentId" width="120">
          <template #default="scope">
            <img v-if="scope.row.children.length > 1 && scope.row.children[1].status == 2"
              :src="imagePath + (scope.row.children[1].icon)" alt="" style="width: 50px; height: 50px;">
          </template>
        </el-table-column>
        <!-- 
        <el-table-column align="left" label="按钮组">
          <template #default="scope"> -->
        <!-- <el-button type="primary" link icon="edit" class="table-button"
              @click="updateUserEquipmentFunc(scope.row)">变更</el-button> -->
        <!-- <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
          </template> -->
        <!-- </el-table-column> -->
      </el-table>
      <div class="gva-pagination">
        <el-pagination layout="total, sizes, prev, pager, next, jumper" :current-page="page" :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]" :total="total" @current-change="handleCurrentChange"
          @size-change="handleSizeChange" />
      </div>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="弹窗操作">
      <el-form :model="formData" label-position="right" ref="elFormRef" :rules="rule" label-width="180px">
        <!-- <el-form-item label="用户ID:" prop="userId">
          <el-input v-model.number="formData.userId" :clearable="false" placeholder="请输入" />
        </el-form-item> -->
        <!-- <el-form-item label="装备分类:" prop="categoryId">
          <el-input v-model.number="formData.categoryId" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="装备:" prop="equipmentId">
          <el-input v-model.number="formData.equipmentId" :clearable="false" placeholder="请输入" />
        </el-form-item> -->
        <el-form-item label="选择装备:" prop="equipmentId">
          <el-tree :data="equipmentTableData" show-checkbox node-key="id" :default-expanded-keys="[2, 3]" ref="menuTree"
            @check="nodeChange" :default-checked-keys="[5]" :props="defaultProps" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">取 消</el-button>
          <el-button type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'UserEquipment'
}
</script>

<script setup>
import {
  createUserEquipment,
  deleteUserEquipment,
  deleteUserEquipmentByIds,
  updateUserEquipment,
  findUserEquipment,
  getUserEquipmentList,
  getEquipmentList
} from '@/api/userEquipment'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useUserStore } from '@/pinia/modules/user'
import { async } from 'q'

const imagePath = ref(import.meta.env.VITE_IMAGE_URL)
const userStore = useUserStore()
const path = ref("")
const route = useRoute()
const router = useRouter()
// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  userId: 0,
  categoryId: 0,
  equipmentId: 0,
})

router.isReady().then(() => {
  formData.value.userId = parseInt(route.query.id)
})

const defaultProps = {
  children: 'children',
  label: 'name',
}

// 验证规则
const rule = reactive({
})

const elFormRef = ref()


// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(50)
const tableData = ref([])
const searchInfo = ref({})

const needConfirm = ref(false)
// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  getTableData()
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

const nodeChange = () => {
  needConfirm.value = true
}
//查询装备
// const getEquipmentTableData = async () => {
//   const table = await getEquipmentList({ page: page.value, pageSize: 1000, ...searchInfo.value })
//   if (table.code === 0) {
//     equipmentTableData.value = arraytotree(table.data.list)
//     console.log(equipmentTableData.value)
//   }
// }

// getEquipmentTableData()

const selectEquipment = async (value) => {
  console.log('用户的Id', parseInt(route.query.id))
  console.log('选中的值', value)
  formData.value.equipment = value
  formData.value.userId = parseInt(route.query.id)
  const res = await createUserEquipment(formData.value)
  console.log("返回的数据", res)
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '更改成功'
    })
    getTableData()
  }
}

// 查询
const getTableData = async () => {
  const table = await getUserEquipmentList({ page: page.value, pageSize: pageSize.value, userId: route.query.id, ...searchInfo.value })
  console.log("查询装备", table.data.list)
  if (table.code === 0) {
    tableData.value = arraytotree(table.data.list)
    console.log("树头结构", tableData.value)
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

const arraytotree = (arr) => {
  var top = [], sub = [], tempObj = {};
  arr.forEach(function (item) {
    if (item.parentId == 0) { // 顶级分类
      top.push(item)
    } else {
      sub.push(item) // 其他分类
    }
    item.children = []; // 默然添加children属性
    tempObj[item.ID] = item // 用当前分类的id做key，存储在tempObj中
  })
  sub.forEach(function (item) {
    // 取父级
    var parent = tempObj[item.parentId] || { 'children': [] }
    // 把当前分类加入到父级的children中
    parent.children.push(item)
  })
  return top
}

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () => {
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
  multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    deleteUserEquipmentFunc(row)
  })
}


// 批量删除控制标记
const deleteVisible = ref(false)

// 多选删除
const onDelete = async () => {
  const ids = []
  if (multipleSelection.value.length === 0) {
    ElMessage({
      type: 'warning',
      message: '请选择要删除的数据'
    })
    return
  }
  multipleSelection.value &&
    multipleSelection.value.map(item => {
      ids.push(item.ID)
    })
  const res = await deleteUserEquipmentByIds({ ids })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    if (tableData.value.length === ids.length && page.value > 1) {
      page.value--
    }
    deleteVisible.value = false
    getTableData()
  }
}

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateUserEquipmentFunc = async (row) => {
  const res = await findUserEquipment({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.reuserEquipment
    dialogFormVisible.value = true
  }
}


// 删除行
const deleteUserEquipmentFunc = async (row) => {
  const res = await deleteUserEquipment({ ID: row.ID })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    if (tableData.value.length === 1 && page.value > 1) {
      page.value--
    }
    getTableData()
  }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
  type.value = 'create'
  dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    userId: 0,
    categoryId: 0,
    equipmentId: 0,
  }
}
const menuTree = ref(null)
// 弹窗确定
const enterDialog = async () => {
  const checkArr = menuTree.value.getCheckedNodes(false, false)
  if (checkArr.length === 0) {
    ElMessage({
      type: 'warning',
      message: '请选择装备'
    })
    return
  }
  //循环判断parentId不为0的数据，并将自己的id和parentId 放入一个数组中
  const arr = []
  checkArr.forEach(item => {
    if (item.parentId !== 0) {
      arr.push({ equipmentId: item.ID, categoryId: item.parentId, userId: formData.value.userId })
    }
  })
  console.log(arr)
  //数组转为json字符串
  formData.value.equipment = JSON.stringify(arr)

  elFormRef.value?.validate(async (valid) => {
    if (!valid) return
    let res
    switch (type.value) {
      case 'create':
        res = await createUserEquipment(formData.value)
        break
      case 'update':
        res = await updateUserEquipment(formData.value)
        break
      default:
        res = await createUserEquipment(formData.value)
        break
    }
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '更改成功'
      })
      closeDialog()
      getTableData()
    }
  })
}
</script>

<style></style>
