import {createRouter, createWebHistory} from 'vue-router'
import HomePage from '@/views/HomePage.vue'

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: '/',
            name: 'home',
            component: HomePage
        },
        {
            path: '/room/:id',
            name: 'room',
            component: () => import('@/views/RoomPage.vue')
        }
    ]
})

export default router
