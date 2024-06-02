<template>
  <h1 style="color: red;">TaskIndex</h1>
  <RouterLink to="tasks/create">新規登録画面へ</RouterLink>
  <div style="background-color: aqua;">
    <h1 @click="getTest" style="color: black;">ここを押すと/testのエンドポイント</h1>
    <h1 style="color: red;">{{ testString }}</h1>
    <h1 @click="getTasks" style="color: black;">ここを押すと/tasksのエンドポイント</h1>
    <div v-for="task in tasks" :key="task.id">
      <div style="color: black;">{{ task.Title }}</div>
    </div>
  </div>
</template>
<script setup lang="ts">
import axios from "axios";
import { ref } from "vue";

const testString = ref<string>('');

const getTest = () => {
  axios.get('http://localhost:8080/test')
  .then(function (response) {
    // handle success(axiosの処理が成功した場合に処理させたいことを記述)
    console.log("ok")
    console.log(response.data);
    testString.value = response.data.message;
  })
  .catch(function (error) {
  // handle error(axiosの処理にエラーが発生した場合に処理させたいことを記述)
    console.log("error")
    console.log(error);
  })
  .finally(function () {
 // always executed(axiosの処理結果によらずいつも実行させたい処理を記述)
  });
}

const tasks = ref<Array>([])
const getTasks = () => {
  axios.get('http://localhost:8080/tasks')
  .then(function (response) {
   // handle success(axiosの処理が成功した場合に処理させたいことを記述)
  tasks.value = response.data.datas;
  })
  .catch(function (error) {
 // handle error(axiosの処理にエラーが発生した場合に処理させたいことを記述)
    console.log(error);
  })
  .finally(function () {
 // always executed(axiosの処理結果によらずいつも実行させたい処理を記述)
  });
}
</script>