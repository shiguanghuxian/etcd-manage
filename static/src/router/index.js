import Vue from 'vue'
import Router from 'vue-router'
import CloudHome from '@/page/CloudHome'

import KV from '@/page/KV'
import Members from '@/page/Members'
import EtcdServers from '@/page/EtcdServers'
import Logs from '@/page/Logs'


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
      },
      {
        path: '/other/EtcdServers',
        name: 'EtcdServers',
        component: EtcdServers,
      },
      {
        path: '/other/Logs',
        name: 'Logs',
        component: Logs,
      }
      
    ]
  }]
})
