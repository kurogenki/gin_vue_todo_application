<template>
  <div style="color: black; background-color: antiquewhite;" v-if="task">
    <h1>以下はタスク ID {{ task.ID }}のタイトルです</h1>
    <h2>タイトル：{{ task.Title }}</h2>
    <h2>詳細：{{ task.Description }}</h2>
  </div>
  <div @click="router.push({name: 'TaskEdit', params:{id: task.id}})">編集モードに変更する</div>
</template>
<script setup lang="ts">
import axios from "axios";
import { onMounted, ref } from "vue";
import { useRoute, useRouter } from "vue-router";
import {showTask} from "../../api/TaskApi/"

const route = useRoute()
const router = useRouter()
const task = ref<object>()

onMounted(async () => {
  task.value = await showTask(route.params.id)
})
</script>