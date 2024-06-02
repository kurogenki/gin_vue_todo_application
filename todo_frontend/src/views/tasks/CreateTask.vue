<template>
  <div style="color: black; background-color: antiquewhite;" >
    <h1 style="color: red;">タスクを新規作成する</h1>
    <form @submit.prevent="createTask">
      <div>
        <label>タスクのタイトル</label>
        <input type="text" v-model="task.title">
      </div>
      <div>
        <label>タスクの詳細</label>
        <input type="text" v-model="task.description">
      </div>
      <button type="submit">新規登録</button>
    </form>
  </div>
</template>
<script setup lang="ts">
import axios from "axios";
import { ref } from "vue";
import { useRouter } from "vue-router";

const router = useRouter()
const task = ref({
  title: "",
  description: ""
})

// TODO formData以外で送信する方法がないか検討
const formData = new FormData();
const createTask = () => {
  formData.append("title", task.value.title);
  formData.append("description", task.value.description);
  axios.post('http://localhost:8080/tasks', formData)
  .then(function (response) {
    router.push("/")
  })
  .catch(function (error) {
    console.log(error);
  })
  .finally(function () {
  });
}
</script>