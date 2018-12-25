<template>
    <div class="members">
        <Table border :columns="columns" :data="list"></Table>
    </div>
</template>

<script>
export default {
    data(){
        return {
            etcdName: '', // etcd服务名
            list:[],
            columns:[{
                        title: 'ID',
                        key: 'ID'
                    },
                    {
                        title: 'Name',
                        key: 'name'
                    },
                    {
                        title: 'Role',
                        key: 'role'
                    },
                    {
                        title: 'Status',
                        key: 'status',
                        render: (h, params) => {
                            return h("Button", {
                                props:{
                                    type:params.row.status == 'healthy' ? 'success' :'error',
                                    size:'small'
                                }
                            }, params.row.status)
                        }
                    },
                    {
                        title: 'DB Size',
                        key: 'db_size'
                    },
                    {
                        title: 'PeerURLs',
                        key: 'peerURLs'
                    },
                    {
                        title: 'ClientURLs',
                        key: 'clientURLs'
                    }]
        }
    },
    mounted(){
        this.etcdName = localStorage.getItem("etcdName");
        this.getList();
    },
    methods:{
        getList(){
            this.$http.get(`/v1/members`,{
          headers:{
            "EtcdServerName":this.etcdName,
          }
        })
            .then(response => {
                console.log(response);
                if(response.status == 200){
                    this.list = response.data;
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

<style scoped>

</style>
