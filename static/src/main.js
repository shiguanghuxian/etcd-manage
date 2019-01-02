import Vue from 'vue'
import App from './App.vue'
import router from './router'
import axios from 'axios'
import './plugins/iview.js'

import VueI18n from 'vue-i18n';
import en from 'iview/dist/locale/en-US';
import zh from 'iview/dist/locale/zh-CN';

import myen from './i18n/en-US';
import myzh from './i18n/zh-CN';

Vue.use(VueI18n);

// 编辑器
import VueCodemirror from 'vue-codemirror'

// require styles
import 'codemirror/lib/codemirror.css'

Vue.use(VueCodemirror)


Vue.config.productionTip = false;

/* axios */
// axios.defaults.baseURL = 'http://127.0.0.1:10280';

// 请求拦截器
axios.interceptors.request.use(function (config) {
  // 读取本地是否存在用户信息
  let etcdUserStr = localStorage.getItem('etcdUser');
  // console.log(etcdUserStr)
  if(etcdUserStr != null && typeof etcdUserStr != 'undefined'){
    let etcdUser = JSON.parse(etcdUserStr);
    if (typeof etcdUser.username != 'undefined' && etcdUser.username != ''){
      config.headers['X-Etcd-Username'] = etcdUser.username;
    }
    if (typeof etcdUser.password != 'undefined' && etcdUser.password != ''){
      config.headers['X-Etcd-Password'] = etcdUser.password;
    }
  }
  return config;
}, function (error) {
  return Promise.reject(error);
})


Vue.prototype.$http = axios;

// 语言
Vue.locale = () => {};

const messages = {
    en: Object.assign(myen, en),
    zh: Object.assign(myzh, zh)
};

let mylang = localStorage.getItem("lang") || 'en';

// Create VueI18n instance with options
const i18n = new VueI18n({
    locale: mylang,  // set locale
    messages  // set locale messages
});

new Vue({
  router,
  i18n,
  render: h => h(App)
}).$mount('#app')
