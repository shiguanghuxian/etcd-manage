import axios from 'axios'
import iView from 'iview';
import { Message } from 'iview'

// 请求前缀
axios.defaults.baseURL = 'http://127.0.0.1:10280'

// 请求之前拦截
axios.interceptors.request.use(
    config => {
        // 判断本地token是否存在，
        // 如果不存在则跳转到登陆页，
        // 如果存在则加到头信息里
        // if (token) {
        //     config.headers.Authorization = token;
        // }
        // router.replace({
        //   path: '/login'
        // })

        let etcdSelectName = localStorage.getItem('etcd-name') || ''; // 读取当前选中的etcd server
        config.headers.Authorization = 'Basic YWRtaW46MTIzNDU2'
        config.headers.EtcdServerName = etcdSelectName;
        iView.LoadingBar.start();
        return config
    },
    err => {
        return Promise.reject(err)
    })


// 请求相应拦截器
axios.interceptors.response.use(
    response => {
        console.log(response)
        // let data = response.data;
        // // 未登录情况
        // if (response.status == 401) {
        //     // todo 删除本地用户缓存
        //     window.location.href = '/login'
        //     Message.error('未登录，或登录失效，请登录')
        // }
        if (response.status == 400) {
            Message.error('req err')
        }
        iView.LoadingBar.finish();
        return response
    },
    error => {
        iView.LoadingBar.error();
        console.log(error.response);
        if (error.response) {
            if (error.response.status == 400) {
                Message.error(error.response.data.msg);
            }
        } else {
            Message.error('请求错误');
        }

        // if (error.response) {
        //     switch (error.response.status) {
        //         case 401:
        //             // 清除本地token，并跳转到登陆页
        //             window.location.href = '/login';
        //     }
        // }


        return Promise.reject(error)
    });

export default axios