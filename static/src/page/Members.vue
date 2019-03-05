<template>
  <div class="members">
    <Table border :columns="columns" :data="list"></Table>
  </div>
</template>

<script>
import { SERVER } from "@/api/server.js";
import { bus } from "@/page/bus.js";

export default {
  data() {
    return {
      list: [],
      columns: [
        {
          title: "ID",
          key: "ID"
        },
        {
          title: "Name",
          key: "name"
        },
        {
          title: "Role",
          key: "role"
        },
        {
          title: "Status",
          key: "status",
          render: (h, params) => {
            return h(
              "Button",
              {
                props: {
                  type: params.row.status == "healthy" ? "success" : "error",
                  size: "small"
                }
              },
              params.row.status
            );
          }
        },
        {
          title: "DB Size",
          key: "db_size"
        },
        {
          title: "PeerURLs",
          key: "peerURLs"
        },
        {
          title: "ClientURLs",
          key: "clientURLs"
        }
      ]
    };
  },
  mounted(){
    this.getList();

    bus.$off('etcd-server-selected');
    bus.$on('etcd-server-selected', (item) => {
        console.log(item);
        this.getList();
    })
  },
  methods:{
      getList(){
          SERVER.GetMemberList().then(response => {
              this.list = response.data || [];
          });
      }
  }
};
</script>
<style lang='scss' scoped>
.members{
    padding: 10px;
}
</style>
