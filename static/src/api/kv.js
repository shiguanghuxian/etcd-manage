import axios from './api'

const KV = {
    /**
     * 查询指定目录的key列表
     * @param {string} dir 查询key的目录
     */
    GetKeyList(dir) {
        return axios.get(`/v1/list?key=${dir}`);
    },

    /**
     * 获取一个key的详细信息
     * @param {string} dir 
     */
    GetKeyInfo(dir) {
        return axios.get(`/v1/key?key=${dir}`)
    },

    /**
     * 添加key
     * @param {*} data 添加keybody信息
     */
    AddKey(data){
        return axios.post('/v1/key', data)
    },

    /**
     * 修改key
     * @param {} data 修改时的body
     */
    SaveKey(data){
        return axios.put(`/v1/key`, data)
    },

    /**
     * 删除key
     * @param {string} dir 要删除的key
     */
    DelKey(dir){
        return axios.delete(`/v1/key?key=${dir}`)
    }
}

// 获取指定目录的key列表
export {
    KV
}