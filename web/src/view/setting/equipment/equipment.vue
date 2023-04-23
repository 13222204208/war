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
        <div class="gva-btn-list">
            <el-button size="small" type="primary" icon="plus" @click="openDialog">新增分类</el-button>
     
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        >
        <el-table-column algin="left" label="ID" prop="ID" width="80"/>
        <el-table-column align="left" label="日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="装备名称" prop="name" width="120" />
        <el-table-column align="left" label="图标" prop="icon" width="120">
            <template #default="scope">
            <el-image style="height:50px" :src="imagePath+scope.row.icon" />
            </template>
        </el-table-column>
        <el-table-column align="left" label="父级ID" prop="parentId" width="120" />
        <el-table-column align="left" label="排序" prop="sort" width="120" />
        <el-table-column align="left" label="状态" prop="status" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.status,statusOptions) }}
            </template>
          </el-table-column>
          <el-table-column align="left" label="按钮组">
            <template #default="scope">
              <el-button
                size="small"
                type="primary" link
                icon="plus"
                @click="addMenu(scope.row.ID)"
              >添加子菜单</el-button>
              <el-button
                size="small"
                type="primary" link
                icon="edit"
                @click="updateEquipmentFunc(scope.row.ID)"
              >编辑</el-button>
              <el-button
                size="small"
                type="primary" link
                icon="delete"
                @click="deleteMenu(scope.row.ID)"
              >删除</el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="弹窗操作">
      <el-form :model="formData" label-position="right" ref="elFormRef" :rules="rule" label-width="120px">
        <el-form-item label="装备名称:"  prop="name" >
          <el-input v-model="formData.name" :clearable="false"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="图标:"  prop="icon" >
          <el-image style=" height: 100px" :src="imagePath+formData.icon" />
          <br />
          <el-button type="primary" @click="openChooseImg">选择图片</el-button>
          <ChooseImg ref="chooseImgRef" @enterImg="enterImg"></ChooseImg>
        </el-form-item>
        <el-form-item label="父级ID:"  prop="parentId" >
          <el-input v-model.number="formData.parentId" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="排序:"  prop="sort" >
          <el-input v-model.number="formData.sort" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="状态:"  prop="status" >
          <el-select v-model="formData.status" placeholder="请选择" style="width:100%" :clearable="true" >
            <el-option v-for="(item,key) in statusOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
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
  name: 'Equipment'
}
</script>

<script setup>
import {
  createEquipment,
  deleteEquipment,
  deleteEquipmentByIds,
  updateEquipment,
  findEquipment,
  getEquipmentList
} from '@/api/equipment'

import ChooseImg from '@/components/chooseImg/index.vue'
// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

// 自动化生成的字典（可能为空）以及字段
const statusOptions = ref([])
const formData = ref({
        name: '',
        icon: '',
        parentId: 0,
        sort: 0,
        status: 1,
        })

const chooseImgRef = ref()
const imagePath = ref(import.meta.env.VITE_IMAGE_URL)

const openChooseImg = () => {
  chooseImgRef.value.open()
}
const enterImg = (url) => {
  formData.value.icon = url
}

// 验证规则
const rule = reactive({
               name : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()


// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

const dialogTitle = ref('新增菜单')
const isEdit = ref(false)

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

// 查询
const getTableData = async() => {
  const table = await getEquipmentList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = arraytotree(table.data.list)
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
    statusOptions.value = await getDictFunc('status')
}

// 获取需要的字典 可能为空 按需保留
setOptions()


const addMenu = (id) => {
  dialogTitle.value = '新增菜单'
  formData.value.parentId = id
  console.log(formData.value)
  isEdit.value = false
  setOptions()
  dialogFormVisible.value = true
}
// 行为控制标记（弹窗内部需要增还是改）
// 修改菜单方法

const deleteMenu = (ID) => {
  ElMessageBox.confirm('此操作将删除菜单, 是否继续?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  })
    .then(async() => {
      const res = await deleteEquipment({ ID: ID })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功!'
        })
        if (tableData.value.length === 1 && page.value > 1) {
          page.value--
        }
        getTableData()
      }
    })
    .catch(() => {
      ElMessage({
        type: 'info',
        message: '已取消删除'
      })
    })
}


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
            deleteEquipmentFunc(row)
        })
    }


// 批量删除控制标记
const deleteVisible = ref(false)

// 多选删除
const onDelete = async() => {
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
      const res = await deleteEquipmentByIds({ ids })
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
const updateEquipmentFunc = async(id) => {
    const res = await findEquipment({ ID: id })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.reequipment
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteEquipmentFunc = async (row) => {
    const res = await deleteEquipment({ ID: row.ID })
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
        name: '',
        icon: '',
        parentId: 0,
        sort: 0,
        status: 1,
        }
}

const arraytotree = (arr)=> {
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
                    var parent = tempObj[item.parentId] || {'children': []}
                    // 把当前分类加入到父级的children中
                    parent.children.push(item)
                })
                return top
            }
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createEquipment(formData.value)
                  break
                case 'update':
                  res = await updateEquipment(formData.value)
                  break
                default:
                  res = await createEquipment(formData.value)
                  break
              }
              if (res.code === 0) {
                ElMessage({
                  type: 'success',
                  message: '创建/更改成功'
                })
                closeDialog()
                getTableData()
              }
      })
}
</script>

<style>
</style>
