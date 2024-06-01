<template>
  <h1 style="color: red;">TaskIndex</h1>
  <RouterLink to="tasks/create">新規登録画面へ</RouterLink>
  <div style="background-color: aqua;">
    <h1 @click="getNumber" style="color: black;">ここを押すとバックエンドに通信！</h1>
    <h1 style="color: red;">{{ testString }}</h1>
  </div>
</template>
<script setup lang="ts">
import axios from "axios";
import { ref } from "vue";

const testString = ref<string>('');

const getNumber = () => {
  axios.get('http://localhost:8080/test')
  .then(function (response) {
 // handle success(axiosの処理が成功した場合に処理させたいことを記述)
    console.log(response.data);
    testString.value = response.data.message;
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