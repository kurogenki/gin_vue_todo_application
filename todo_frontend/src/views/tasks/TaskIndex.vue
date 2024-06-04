<template>
  <div style="color: black;">
    <h1 style="color: red;">TaskIndex</h1>
    <div style="background-color: antiquewhite;" v-if="tasks">
      <h1>下記はタスクの一覧です。</h1>
      <div v-for="task in tasks" :key="task.id">
        <h2 @click=" router.push({name: 'TaskShow', params:{id: task.ID}})">
          <span>{{task.ID}}, </span>
          <span>タイトル：{{ task.Title }}</span>
          <button type="button" @click.stop="deleteTask(task.ID)" style="color: red; margin-left: 30px;">削除</button>
        </h2>
      </div>
    </div>
  </div>
  <RouterLink to="tasks/create">新規登録画面へ</RouterLink>
</template>
<script setup lang="ts">
import axios from "axios";
import { onMounted, ref } from "vue";
import { useRouter } from "vue-router";

const router = useRouter()
const testString = ref<string>('');

const tasks = ref<object>({})
onMounted(() => {
  getTasks()
})

const getTasks = () => {
  axios.get('http://localhost:8080/tasks')
  .then(function (response) {
    tasks.value = response.data.datas;
  })
  .catch(function (error) {
    console.log(error);
  })
  .finally(function () {
  });
}

const deleteTask = (id) => {
  axios.delete(`http://localhost:8080/tasks/${id}`)
  .then(function (response) {
    tasks.value = response.data.datas;
    getTasks()
  })
  .catch(function (error) {
    console.log(error);
  })
  .finally(function () {
  });
}
</script>