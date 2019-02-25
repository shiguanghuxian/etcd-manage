import Vue from 'vue'
import Router from 'vue-router'
import CloudHome from '@/page/CloudHome'

import Page1 from '@/components/Page1'
import Page2 from '@/components/Page2'
import Page3 from '@/components/Page3'
import Page4 from '@/components/Page4'
import Page5 from '@/components/Page5'
import Page6 from '@/components/Page6'
import Page7 from '@/components/Page7'
import Page8 from '@/components/Page8'

import KV from '@/page/KV'


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
        path: '/cloud/page1',
        name: 'Page1',
        component: Page1,
      }, 
      {
        path: '/cloud/page2',
        name: 'Page2',
        component: Page2,
      },
      {
        path: '/cloud/page3',
        name: 'Page3',
        component: Page3,
      }, 
      {
        path: '/cloud/page4',
        name: 'Page4',
        component: Page4,
      },
      {
        path: '/cloud/page5',
        name: 'Page5',
        component: Page5,
      },
      {
        path: '/cloud/page6',
        name: 'Page6',
        component: Page6,
      },
      {
        path: '/cloud/page7',
        name: 'Page7',
        component: Page7,
      },
      {
        path: '/cloud/page8',
        name: 'Page8',
        component: Page8,
      },
    ]
  }]
})
