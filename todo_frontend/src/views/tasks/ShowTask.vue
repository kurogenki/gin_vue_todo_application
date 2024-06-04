<template>
  <div style="color: black; background-color: antiquewhite;" v-if="task">
    <h1>以下はタスク ID {{ task.ID }}のタイトルです</h1>
    <h2>タイトル：{{ task.Title }}</h2>
    <h2>詳細：{{ task.Description }}</h2>
  </div>
</template>
<script setup lang="ts">
import axios from "axios";
import { onMounted, ref } from "vue";
import { useRoute } from "vue-router";

const route = useRoute()
const task = ref<object>()

onMounted(() => {
  showTask()
})

const showTask = () => {
  axios.get(`http://localhost:8080/tasks/${route.params.id}`)
  .then(function (response) {
    task.value = response.data.datas;
  })
  .catch(function (error) {
    console.log(error);
  })
  .finally(function () {
  });
}
</script>