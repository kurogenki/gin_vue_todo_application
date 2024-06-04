<template>
  <h1>編集画面</h1>
  <form @submit.prevent="update(task)">
    <label>タイトル</label>
    <input type="text" v-model="task.Title">
    <label>詳細</label>
    <input type="text" v-model="task.Description">
    <button type="submit">更新</button>
  </form>
</template>
<script setup lang="ts">
import { onMounted, ref } from "vue";
import { useRoute, useRouter } from "vue-router";
import {showTask, updateTask} from "../../api/TaskApi";

const route = useRoute()
const router = useRouter()

const task = ref<object>({})

onMounted(async () => {
  task.value = await showTask(route.params.id)
})

const formData = new FormData()
const update = async (task: object) => {
  formData.append("title", task.Title);
  formData.append("description", task.Description);
  await updateTask(route.params.id, formData)
  router.push({name: "TaskShow", params:{id: route.params.id}})
}
</script>