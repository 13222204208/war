<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="房间名称:" prop="name">
          <el-input v-model="formData.name" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="最少人数:" prop="minPlayers">
          <el-input v-model.number="formData.minPlayers" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="最多人数:" prop="maxPlayers">
          <el-input v-model.number="formData.maxPlayers" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="倒计时分钟:" prop="countdown">
          <el-input v-model.number="formData.countdown" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="房间人数:" prop="numPlayers">
          <el-input v-model.number="formData.numPlayers" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="倒计时结束时间:" prop="endTime">
          <el-date-picker v-model="formData.endTime" type="date" placeholder="选择日期" :clearable="false"></el-date-picker>
        </el-form-item>
        <el-form-item label="游戏结束时间:" prop="gameOverTime">
          <el-date-picker v-model="formData.gameOverTime" type="date" placeholder="选择日期" :clearable="false"></el-date-picker>
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
  name: 'Room'
}
</script>

<script setup>
import {
  createRoom,
  updateRoom,
  findRoom
} from '@/api/room'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            name: '',
            minPlayers: 0,
            maxPlayers: 0,
            countdown: 0,
            numPlayers: 0,
            endTime: new Date(),
            gameOverTime: new Date(),
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findRoom({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reroom
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createRoom(formData.value)
               break
             case 'update':
               res = await updateRoom(formData.value)
               break
             default:
               res = await createRoom(formData.value)
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
