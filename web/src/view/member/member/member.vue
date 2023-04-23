<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
      <!--代号搜索-->
      <el-form-item label="代号">
        <el-input v-model="searchInfo.nickname" placeholder="请输入代号"></el-input>
      </el-form-item>
      <!--姓名搜索-->
      <el-form-item label="姓名">
        <el-input v-model="searchInfo.name" placeholder="请输入姓名"></el-input>
      </el-form-item>

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
                <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="deleteVisible = true">删除</el-button>
            </template>
            </el-popover>
        </div> -->
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        >
        <!-- <el-table-column type="selection" width="55" /> -->
        <el-table-column align="left" label="日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="图片" prop="avatar" width="80">
            <template #default="scope">
            <el-image style="height: 50px;" :src="imagePath+scope.row.avatar" />
            </template>
        </el-table-column>
        
        <el-table-column align="left" label="代号" prop="nickname" width="120" />
        <el-table-column align="left" label="姓名" prop="name" width="120" />
        <el-table-column align="left" label="性别" prop="gender" width="60">
            <template #default="scope">
            {{ filterDict(scope.row.gender,genderOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="身高(厘米)" prop="height" width="120" />
        <el-table-column align="left" label="体重(千克)" prop="weight" width="120" />
        <el-table-column align="left" label="电话号" prop="phone" width="120" />
        <el-table-column align="left" label="按钮组">
            <template #default="scope">
            <el-button type="primary" link icon="edit" class="table-button" @click="updateMemberFunc(scope.row)">变更</el-button>

            <el-button type="primary" link icon="edit" class="table-button" @click="updateMemberMatchFunc(scope.row)">场次</el-button>

            <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
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
    <!--场次弹窗-->
    <el-dialog v-model="dialogMatchFormVisible" :before-close="closeMatchDialog" title="场次操作">
      <el-form :model="updateMatchData" label-position="right" ref="elFormRef" :rules="rule" label-width="80px">
        <el-form-item label="类型:"  prop="matchType" >
          <el-radio-group v-model="updateMatchData.matchType" class="ml-4">
            <el-radio :label=1 size="large">增加</el-radio>
            <el-radio :label=2 size="large">减少</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="场次:"  prop="match" >
          <el-input v-model.number="updateMatchData.match" type="number"  placeholder="请输入" />
        </el-form-item>

      </el-form>

      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeMatchDialog">取 消</el-button>
          <el-button type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>

    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="弹窗操作">
      <el-form :model="formData" label-position="right" ref="elFormRef" :rules="rule" label-width="80px">
        <el-form-item label="头像:"  prop="avatar" >
          <el-input v-model="formData.avatar" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="代号:"  prop="nickname" >
          <el-input v-model="formData.nickname" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="姓名:"  prop="name" >
          <el-input v-model="formData.name" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="性别:"  prop="gender" >
          <el-select v-model="formData.gender" placeholder="请选择" style="width:100%" :clearable="false" >
            <el-option v-for="(item,key) in genderOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="身高体重:"  prop="height" >
          <el-input v-model.number="formData.height" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="体重:"  prop="weight" >
          <el-input v-model.number="formData.weight" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="电话号:"  prop="phone" >
          <el-input v-model="formData.phone" :clearable="false"  placeholder="请输入" />
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
  name: 'Member'
}
</script>

<script setup>
import {
  createMember,
  deleteMember,
  deleteMemberByIds,
  updateMember,
  updateMemberMatch,
  findMember,
  getMemberList
} from '@/api/member'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

// 自动化生成的字典（可能为空）以及字段
const genderOptions = ref([])
const formData = ref({
        avatar: '',
        nickname: '',
        name: '',
        gender: undefined,
        height: 0,
        weight: 0,
        phone: '',
        match: 0,
        })
const updateMatchData = ref({
        match: 0,
        matchType: 1,
        userId:0,
        })

// 验证规则
const rule = reactive({
})

const elFormRef = ref()


// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
const imagePath = ref(import.meta.env.VITE_IMAGE_URL)

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
  const table = await getMemberList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
    genderOptions.value = await getDictFunc('gender')
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
            deleteMemberFunc(row)
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
      const res = await deleteMemberByIds({ ids })
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
const updateMemberFunc = async(row) => {
    const res = await findMember({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.remember
        dialogFormVisible.value = true
    }
}

//更新场次
const updateMemberMatchFunc = async(row) => {
  const res = await findMember({ ID: row.ID })
    type.value = 'updateMatch'
    if (res.code === 0) {
        updateMatchData.value.userId = res.data.remember.ID
        dialogMatchFormVisible.value = true
    }
}


// 删除行
const deleteMemberFunc = async (row) => {
    const res = await deleteMember({ ID: row.ID })
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

//场次弹窗标记
const dialogMatchFormVisible = ref(false)
// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        avatar: '',
        nickname: '',
        name: '',
        gender: undefined,
        height: 0,
        weight: 0,
        phone: '',
        }
}

//关闭场次弹窗
const closeMatchDialog = () => {
    dialogMatchFormVisible.value = false
}


// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createMember(formData.value)
                  break
                case 'update':
                  res = await updateMember(formData.value)
                  break
                case 'updateMatch':
                  res = await updateMemberMatch(updateMatchData.value)
                  break
                default:
                  res = await createMember(formData.value)
                  break
              }
              if (res.code === 0) {
                ElMessage({
                  type: 'success',
                  message: '创建/更改成功'
                })
                closeDialog()
                closeMatchDialog()
                getTableData()
              }
      })
}
</script>

<style>
</style>
