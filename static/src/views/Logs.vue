<style scoped>
.logs{
  width: 100%;
  height: 100vh;
  overflow-y: scroll;
  overflow-x: hidden;
}
.search{
    width: 100%;
    height: 39px;
}
</style>

<template>
    <div class="logs">
        <div class="search">
            <Form :rules="ruleInline" inline>
                <FormItem prop="日期">
                    <DatePicker type="date" placeholder="Select date" style="width: 300px" @on-change="changeDate" format="yyyyMMdd"></DatePicker>
                </FormItem>
                <FormItem prop="用户">
                    <Select v-model="user" clearable style="width:200px">
                        <Option v-for="item in users" :value="item.name" :key="item.name">
                            {{ item.name }}
                            <!-- <span>{{ item.name }}</span>
                            <span style="float:right;color:#ccc">{{ item.role }}</span> -->
                        </Option>
                    </Select>
                </FormItem>
                <FormItem prop="类型">
                    <Select v-model="logType" clearable style="width:200px">
                        <Option v-for="item in logtypes" :value="item" :key="item">{{ item }}</Option>
                    </Select>
                </FormItem>
                <FormItem>
                    <Button type="primary" @click="getList">筛选</Button>
                </FormItem>
            </Form>
        </div>
        <Table border :columns="columns" :data="list" :loading="loading"></Table>
        <div style="margin-top:10px; text-align: right;">
            <Page @on-change="changeListPage" @on-page-size-change="pageSizeChange" show-sizer :current="page" :page-size="pageSize" :total="listTotal" />
        </div>
        <div style="height:200px;"></div>
    </div>
</template>

<script>
export default {
        data(){
        return {
            list:[],
            page:1,
            pageSize:10,
            listTotal:0,
            loading:false,
            date:'',
            logtypes:[], // 日志类型列表
            users:[], // 用户列表
            logType:'', // 日志类型
            user:'', // 用户名
            columns:[{
                    title: 'Date',
                    key: 'date'
                },
                {
                    title: 'User',
                    key: 'user'
                },
                {
                    title: 'Role',
                    key: 'role'
                },
                {
                    title: 'Action',
                    key: 'msg'
            }]
        }
    },
    mounted(){
        this.bindDate = new Date();
        // this.getList();
        this.$Message.info('选择日期查看日志');

        // 下拉框数据
        this.getUsers();
        this.getLogtypes();
    },
    methods:{
        // 切换每页数据大小
        pageSizeChange(pageSize){
            this.pageSize = pageSize;
            this.page = 1;
            this.getList();
        },

        // 点击页码
        changeListPage(page){
            this.page = page;
            this.getList();
        },

        // 切换日期
        changeDate(date){
            console.log(date);
            this.date = date;
            this.page = 1;
            this.getList();
        },

        // 获取数据
        getList(){
            if (this.date == '' || !this.date){
                this.$Message.info('请选择日期');
                return
            }
            this.loading = true;
            this.$Loading.start();
            this.list = [];
            this.listTotal = 0;

            this.user = this.user || '';
            this.logType = this.logType || '';

            this.$http.get(`/v1/logs?date=${this.date}&page=${this.page}&page_size=${this.pageSize}&user=${this.user}&log_type=${this.logType}`)
            .then(response => {
                console.log(response);
                if(response.status == 200){
                    this.list = response.data.list;
                    this.listTotal = response.data.total;
                }
                this.$Loading.finish();
                this.loading = false;
            }).catch(error=>{
                if (error.response){
                    this.$Message.error(error.response.data.msg);
                }
                this.loading = false;
                this.$Loading.error();
            });
        },

        // 获取用户列表
        getUsers(){
            this.$http.get(`/v1/users`)
            .then(response => {
                console.log(response);
                if(response.status == 200){
                    this.users = response.data || [];
                }
            }).catch(error=>{
                if (error.response){
                    this.$Message.error(error.response.data.msg);
                }
            });
        },

        // 获取日志类型列表
        getLogtypes(){
            this.$http.get(`/v1/logtypes`)
            .then(response => {
                console.log(response);
                if(response.status == 200){
                    this.logtypes = response.data || [];
                }
            }).catch(error=>{
                if (error.response){
                    this.$Message.error(error.response.data.msg);
                }
            });
        }
    }
}
</script>
