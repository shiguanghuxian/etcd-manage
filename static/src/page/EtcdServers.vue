
<template>
    <div class="etcd-servers">
        <Table border :columns="columns" :data="data1"></Table>
        <div class="page">
            <Page @on-change="changeListPage" @on-page-size-change="pageSizeChange" :total="total" :current="page" :page-size="pageSize" show-sizer />
        </div>
    </div>
</template>
<script>
import { SERVER } from "@/api/server.js";

export default {
    data() {
        return {
            page:1,
            total:0,
            pageSize:10,
            columns:[
                {
                    title: 'Title',
                    key: 'Title'
                },
                {
                    title: 'Name',
                    key: 'Name'
                },
                {
                    title: 'Address',
                    key: 'Address'
                },
                {
                    title: 'Roles',
                    key: 'Roles'
                },
                {
                    title: 'Desc',
                    key: 'Desc'
                },
                {
                    title: 'Action',
                    key: 'action',
                    width: 150,
                    align: 'center',
                    render: (h, params) => {
                        return h('div', [
                            h('Poptip', {
                                props: {
                                    confirm: true,
                                    title:this.$t('etcdServer.determineRepairDirectory')
                                },
                                on: {
                                    "on-ok": () => {
                                        this.restore(params.row);
                                    }
                                }
                            },[
                                h('Button', {
                                    props: {
                                        type: 'primary',
                                        size: 'small'
                                    },
                                    style: {
                                        marginRight: '5px'
                                    }
                                }, this.$t('etcdServer.repairDirectory')),
                            ])
                        ]);
                    }
                }
            ],
            data:[],
            data1:[] // 分页用
        }
    },
    methods:{
        getList(){
            SERVER.GetEtcdServerList().then(response => {
                console.log(response);
                this.data = response.data || [];
                this.changeListPage(this.page);
            });
        },
        // 切换页码
        changeListPage(page){
            // pageSize
            this.data1.splice(0, this.data1.length);
            this.data1 = this.data.slice((page - 1) * this.pageSize, page * this.pageSize);
            console.log(page);
        },
        // 页面打小
        pageSizeChange(pageSize){
            this.pageSize = pageSize;
            this.changeListPage(1);
            this.page = 1;
        },

        // 修复该服务目录问题
        restore(row){
            // console.log(row)
            SERVER.RestoreEtcdServer().then(response=>{
                if(response.status == 200){
                    this.$Message.info('OK');
                }
            })
        }
    },
    mounted(){
        this.getList();
    }
}
</script>

<style lang="scss" scoped>
    .etcd-servers{
        padding: 10px;
        .page{
            margin-top:10px; 
            text-align: right;
        }
    }
</style>