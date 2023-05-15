<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="房间ID:" prop="roomId">
          <el-input v-model.number="formData.roomId" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="用户ID:" prop="userId">
          <el-input v-model.number="formData.userId" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="所属阵营:" prop="faction">
          <el-select v-model="formData.faction" placeholder="请选择" :clearable="false">
            <el-option v-for="(item,key) in redAndBlueOptions" :key="key" :label="item.label" :value="item.value" />
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
  name: 'MemberRoom'
}
</script>

<script setup>
import {
  createMemberRoom,
  updateMemberRoom,
  findMemberRoom
} from '@/api/memberRoom'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const redAndBlueOptions = ref([])
const formData = ref({
            roomId: 0,
            userId: 0,
            faction: undefined,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findMemberRoom({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rememberRoom
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    redAndBlueOptions.value = await getDictFunc('redAndBlue')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createMemberRoom(formData.value)
               break
             case 'update':
               res = await updateMemberRoom(formData.value)
               break
             default:
               res = await createMemberRoom(formData.value)
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
