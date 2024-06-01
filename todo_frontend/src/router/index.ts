import { createRouter, createWebHistory } from 'vue-router'
// import HomeView from '../views/HomeView.vue'
import TaskIndex from '../views/tasks/TaskIndex.vue'
import TaskCreate from '../views/tasks/CreateTask.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {path: '/', name: 'home', component: TaskIndex},
    {path: '/tasks',
      children: [
      {path: '',  component: TaskIndex},
      {path: 'create',  component: TaskCreate}
    ]},
  ]
})

export default router
