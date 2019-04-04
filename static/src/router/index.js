import Vue from 'vue'
import Router from 'vue-router'
import CloudHome from '@/page/CloudHome'

Vue.use(Router)

export default new Router({
  routes: [{
    path: '/',
    name: 'CloudHome',
    component: CloudHome,
    children: [
      {
        path: '/key/kv',
        name: 'KV',
        component: () => import('@/page/KV'),
      },
      {
        path: '/server/members',
        name: 'Members',
        component: () => import('@/page/Members'),
      },
      {
        path: '/other/EtcdServers',
        name: 'EtcdServers',
        component: () => import('@/page/EtcdServers'),
      },
      {
        path: '/other/Logs',
        name: 'Logs',
        component: () => import('@/page/Logs'),
      }
      
    ]
  }]
})
