import Vue from 'vue'
import Router from 'vue-router'
import Home from './views/Home.vue'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home,
      children:[{
        path: 'keys',
        name: 'Keys',
        component: () => import('./views/Keys.vue')
      },{
        path: 'members',
        name: 'Members',
        component: () => import('./views/Members.vue')
      },{
        path: 'servers',
        name: 'Servers',
        component: () => import('./views/EtcdServer.vue')
      },{
        path: 'Logs',
        name: 'Logs',
        component: () => import('./views/Logs.vue')
      }]
    },
    {
      path: '/about',
      name: 'about',
      // route level code-splitting
      // this generates a separate chunk (about.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import(/* webpackChunkName: "about" */ './views/About.vue')
    }
  ]
})
