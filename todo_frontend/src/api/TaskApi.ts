import axios from "axios";
import { ref } from "vue";

const data = ref<object>({})

// タスクの一覧を取得
export const getTasks = async () => {
  await axios.get('http://localhost:8080/tasks')
  .then(function (response) {
    data.value = response.data.datas;
  })
  .catch(function (error) {
    console.log(error);
  })
  .finally(function () {
  });

  return data.value
}

// タスクの詳細を取得
export const showTask = async (id:number) => {
  await axios.get(`http://localhost:8080/tasks/${id}`)
  .then(function (response) {
    data.value = response.data.datas;
  })
  .catch(function (error) {
    console.log(error);
  })
  .finally(function () {
  });

  return data.value
}

// タスクを削除
export const deleteTask = async (id: number) => {
  await axios.delete(`http://localhost:8080/tasks/${id}`)
  .then(function (response) {
    data.value = response.data.datas;
  })
  .catch(function (error) {
    console.log(error);
  })
  .finally(function () {
  });

  return data.value
}