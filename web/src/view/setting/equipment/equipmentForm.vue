<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="装备名称:" prop="name">
          <el-input v-model="formData.name" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="图标:" prop="icon">
          <el-input v-model="formData.icon" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="父级ID:" prop="parentId">
          <el-input v-model.number="formData.parentId" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="排序:" prop="sort">
          <el-input v-model.number="formData.sort" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="状态:" prop="status">
          <el-select v-model="formData.status" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in statusOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
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
  updateEquipment,
  findEquipment
} from '@/api/equipment'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const statusOptions = ref([])
const formData = ref({
            name: '',
            icon: '',
            parentId: 0,
            sort: 0,
            status: undefined,
        })
// 验证规则
const rule = reactive({
               name : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findEquipment({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reequipment
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    statusOptions.value = await getDictFunc('status')
}

init()
// 保存按钮
const save = async() => {
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
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
