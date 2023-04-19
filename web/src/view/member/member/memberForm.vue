<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="头像:" prop="avatar">
          <el-input v-model="formData.avatar" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="代号:" prop="nickname">
          <el-input v-model="formData.nickname" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="姓名:" prop="name">
          <el-input v-model="formData.name" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="性别:" prop="gender">
          <el-select v-model="formData.gender" placeholder="请选择" :clearable="false">
            <el-option v-for="(item,key) in genderOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="身高体重:" prop="height">
          <el-input v-model.number="formData.height" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="体重:" prop="weight">
          <el-input v-model.number="formData.weight" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="电话号:" prop="phone">
          <el-input v-model="formData.phone" :clearable="false" placeholder="请输入" />
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
  name: 'Member'
}
</script>

<script setup>
import {
  createMember,
  updateMember,
  findMember
} from '@/api/member'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const genderOptions = ref([])
const formData = ref({
            avatar: '',
            nickname: '',
            name: '',
            gender: undefined,
            height: 0,
            weight: 0,
            phone: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findMember({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.remember
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    genderOptions.value = await getDictFunc('gender')
}

init()
// 保存按钮
const save = async() => {
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
             default:
               res = await createMember(formData.value)
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
