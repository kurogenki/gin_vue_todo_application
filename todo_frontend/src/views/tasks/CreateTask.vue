<template>
  <div style="color: black; background-color: antiquewhite;" >
    <h1 style="color: red;">タスクを新規作成する</h1>
    <form @submit.prevent="create">
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
import { ref } from "vue";
import { useRouter } from "vue-router";
import {createTask} from "../../api/TaskApi";

const router = useRouter()
const task = ref({
  title: "",
  description: ""
})

// TODO formData以外で送信する方法がないか検討
const formData = new FormData();
const create = async () => {
  formData.append("title", task.value.title);
  formData.append("description", task.value.description);
  await createTask(formData)
  router.push("/")
}

</script>