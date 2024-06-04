import { createRouter, createWebHistory } from 'vue-router'
import TaskIndex from '../views/tasks/TaskIndex.vue'
import TaskCreate from '../views/tasks/CreateTask.vue'
import TaskShow from '../views/tasks/ShowTask.vue'
import TaskEdit from '../views/tasks/EditTask.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {path: '/', name: 'home', component: TaskIndex},
    {path: '/tasks',
      children: [
      {path: '', name: 'TaskIndex', component: TaskIndex},
      {path: ':id', name: 'TaskShow', component: TaskShow},
      {path: ':id/edit', name: 'TaskEdit', component: TaskEdit},
      {path: 'create', name: 'TaskCreate', component: TaskCreate}
    ]},
  ]
})

export default router
