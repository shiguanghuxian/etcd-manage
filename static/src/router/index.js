import Vue from 'vue'
import Router from 'vue-router'
import CloudHome from '@/page/CloudHome'

import KV from '@/page/KV'
import Members from '@/page/Members'


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
        component: KV,
      },
      {
        path: '/server/members',
        name: 'Members',
        component: Members,
      }
    ]
  }]
})
