<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="房间ID:" prop="roomId">
          <el-input v-model.number="formData.roomId" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="投诉人:" prop="complainant">
          <el-input v-model.number="formData.complainant" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="被投诉人:" prop="complainee">
          <el-input v-model.number="formData.complainee" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="投诉原因:" prop="reason">
          <el-input v-model="formData.reason" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="投诉状态:" prop="status">
          <el-select v-model="formData.status" placeholder="请选择" :clearable="false">
            <el-option v-for="(item,key) in ComplaintOptions" :key="key" :label="item.label" :value="item.value" />
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
  name: 'Complaint'
}
</script>

<script setup>
import {
  createComplaint,
  updateComplaint,
  findComplaint
} from '@/api/complaint'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const ComplaintOptions = ref([])
const formData = ref({
            roomId: 0,
            complainant: 0,
            complainee: 0,
            reason: '',
            status: undefined,
        })
// 验证规则
const rule = reactive({
               roomId : [{
                   required: true,
                   message: '房间ID必须填写',
                   trigger: ['input','blur'],
               }],
               complainee : [{
                   required: true,
                   message: '被投诉人',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findComplaint({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.recomplaint
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    ComplaintOptions.value = await getDictFunc('Complaint')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createComplaint(formData.value)
               break
             case 'update':
               res = await updateComplaint(formData.value)
               break
             default:
               res = await createComplaint(formData.value)
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
