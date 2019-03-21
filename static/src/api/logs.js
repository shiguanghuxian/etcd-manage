import axios from './api'

const Logs = {
    /**
     * 日志列表
     */
    GetLogsList(dateStr, page, pageSize, user, logType) {
        return axios.get(`/v1/logs?date=${dateStr}&page=${page}&page_size=${pageSize}&user=${user}&log_type=${logType}`);
    },

    /**
     * 获取所有用户列表
     */
    GetAllUser(){
        return axios.get(`/v1/users`);
    },

    /**
     * 获取所有日志类型
     */
    GetLogTypes(){
        return axios.get(`/v1/logtypes`);
    }

}

export {
    Logs
}
