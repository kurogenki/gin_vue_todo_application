<template>
  <div style="color: black;">
    <h1 style="color: red;">TaskIndex</h1>
    <div style="background-color: aqua;">
      <h1 @click="getTest">ここを押すと/testのエンドポイント</h1>
      <h1 style="color: red;">{{ testString }}</h1>
      <h1 @click="getTasks">ここを押すと/tasks(get)エンドポイント</h1>
      <div v-for="task in tasks" :key="task.id">
        <div>{{ task.Title }}</div>
      </div>
      <h1 @click="createTask">ここを押すと/tasks(post)のエンドポイント</h1>
    </div>
    <div style="background-color: antiquewhite;">
      <h1>下記はタスクの一覧です。</h1>
    </div>
  </div>
  <RouterLink to="tasks/create">新規登録画面へ</RouterLink>
</template>
<script setup lang="ts">
import axios from "axios";
import { ref } from "vue";

const testString = ref<string>('');

const getTest = () => {
  axios.get('http://localhost:8080/test')
  .then(function (response) {
    testString.value = response.data.message;
  })
  .catch(function (error) {
    console.log(error);
  })
  .finally(function () {
  });
}

const tasks = ref<Array>([])
const getTasks = () => {
  axios.get('http://localhost:8080/tasks')
  .then(function (response) {
  console.log(response)
  tasks.value = response.data.datas;
  })
  .catch(function (error) {
    console.log(error);
  })
  .finally(function () {
  });
}

const createTask = () => {
  axios.post('http://localhost:8080/tasks')
  .then(function (response) {
    console.log(response.data)
  // tasks.value = response.data.datas;
  })
  .catch(function (error) {
    console.log(error);
  })
  .finally(function () {
  });
}
</script>