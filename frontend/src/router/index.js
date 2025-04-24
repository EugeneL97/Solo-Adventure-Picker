import { createRouter, createWebHistory } from "vue-router";
import HomePage from '../views/HomePage.vue'
import AdventurePage from "../views/AdventurePage.vue";

const routes = [
    { path: '/', component: HomePage },
    { path: '/adventure', component: AdventurePage },
]

const router = createRouter({
    history: createWebHistory(),
    routes
})

export default router