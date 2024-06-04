<template>
  <div style="color: black;">
    <h1 style="color: red;">TaskIndex</h1>
    <div style="background-color: antiquewhite;" v-if="tasks">
      <h1>下記はタスクの一覧です。</h1>
      <div v-for="task in tasks" :key="task.id">
        <h2 @click=" router.push({name: 'TaskShow', params:{id: task.ID}})">
          <span>{{task.ID}}, </span>
          <span>タイトル：{{ task.Title }}</span>
          <button type="button" @click.stop="drop(task.ID)" style="color: red; margin-left: 30px;">削除</button>
        </h2>
      </div>
    </div>
  </div>
  <RouterLink to="tasks/create">新規登録画面へ</RouterLink>
</template>
<script setup lang="ts">
import { onMounted, ref } from "vue";
import { useRouter } from "vue-router";
import {getTasks, deleteTask} from "../../api/TaskApi/";

const router = useRouter()

const tasks = ref<object>({})

onMounted(async () => {
  tasks.value = await getTasks();
})

const drop = async (id:number) => {
  await deleteTask(id)
  tasks.value = await getTasks()
}
</script>