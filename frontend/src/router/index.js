// import routernya dulu
import { createWebHistory, createRouter } from "vue-router"
// import view loginpage as a homepage
import LoginPage from "@/views/LoginPage.vue"
import RegisterPage from "@/views/RegisterPage.vue"
import DashboardPage from "@/views/Dashboard.vue"
import NotFoundPage from "@/views/404.vue"

// bikin routenya
const routes = [
  {
    path: "/",
    name: "Home",
    component: LoginPage
  },
  {
    path: "/register",
    name: "Register",
    component: RegisterPage
  },
  {
    path: "/dashboard",
    name: "Dashboard",
    component: DashboardPage,
    meta: {requiresAuth: true}
  },
  {
    path: "/:pathMatch(.*)*",
    name: "NotFound",
    component: NotFoundPage
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

// jagain urlnya
router.beforeEach((to, from, next) => {
  if (to.matched.some(record => record.meta.requiresAuth)) {
    const isAuthenticated = localStorage.getItem('token');
    if (!isAuthenticated) {
      next({name: 'Home'})
    } else {
      next()
    }
  } else {
    next()
  }
})

export default router