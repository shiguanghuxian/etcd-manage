<template>
  <div class="kv-list">
    <Table
      border
      :columns="columns"
      :data="list"
      @on-selection-change="selectionChangeKeys"
    ></Table>
    <div style="margin-top:10px; text-align: right;" v-if="pageShow == true">
      <Page
        @on-change="changeListPage"
        @on-page-size-change="pageSizeChange"
        show-sizer
        :current="page"
        :page-size="pageSize"
        :total="listTotal"
      />
    </div>
  </div>
</template>
<script>
import { KV } from "@/api/kv.js";

// grid方式展示key
export default {
  data() {
    return {
      columns: [
        // 表格形式展示表头
        {
          type: "selection",
          width: 60,
          align: "center"
        },
        {
          title: "TYPE",
          width: 70,
          render: (h, params) => {
            return h("img", {
              attrs: {
                src:
                  params.row.is_dir == true
                    ? "data:image/png;base64,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"
                    : "data:image/png;base64,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",
                style: "width:35px;height:35px"
              },
              on: {
                click: () => {
                  this.openKey(params.row);
                }
              }
            });
          }
        },
        {
          title: "NAME",
          key: "value"
        },
        {
          title: "KEY",
          key: "full_dir"
        },
        {
          title: "Version",
          key: "version"
        },
        {
          title: " ",
          align: "center",
          render: (h, params) => {
            return h("div", [
              h(
                "Button",
                {
                  props: {
                    type: "primary",
                    size: "small"
                  },
                  style: {
                    marginRight: "5px"
                  },
                  on: {
                    click: () => {
                      this.openKey(params.row);
                    }
                  }
                },
                params.row.is_dir == true
                  ? this.$t("key.open")
                  : this.$t("key.show")
              ),
              h(
                "Poptip",
                {
                  props: {
                    confirm: true,
                    title: this.$t("public.confirmDelete")
                  },
                  on: {
                    "on-ok": () => {
                      this.delOneKey(params.row.full_dir);
                    }
                  }
                },
                [
                  h(
                    "Button",
                    {
                      props: {
                        type: "error",
                        size: "small"
                      }
                    },
                    this.$t("public.delete")
                  )
                ]
              )
            ]);
          }
        }
      ],
      baseList: [], // 永远存储者key列表
      list: [], // 当前显示列表
      page: 1,
      pageSize: 10, // 每页大小
      listTotal: 0, // 总行数
      pageShow: true, // 是否显示分页

      dir: "", // 要显示的目录
      selected: [] // 选中的key
    };
  },
  mounted() {},
  methods: {
    getList(dir) {
      this.list = [];
      this.dir = dir;
      KV.GetKeyList(dir).then(response => {
        this.baseList = JSON.parse(JSON.stringify(response.data || []));
        this.listTotal = this.baseList.length;
        this.changeListPage(1);
        // console.log(response);
      });
    },
    checkKey(item) {
      item.check = !item.check;

      let selected = [];
      this.list.forEach(val => {
        // console.log(val)
        if (val.check == true) {
          selected.push(val);
        }
      });
      this.selected = selected;
      console.log(selected);
    },
    openKey(item) {
      this.$emit("openKey", item);
    },
    // 选择key
    selectionChangeKeys(selections){
      this.selected = selections;
      console.log(selected);
    },
    // 清除
    clearSelected() {
      this.selected = [];
    },

    changeListPage(page){
      this.page = page || 0;
      // pageSize
      this.list=this.baseList.slice((page - 1) * this.pageSize,page * this.pageSize);
      console.log(page);
    },
    pageSizeChange(pageSize){
      this.pageSize = pageSize;
      this.changeListPage(1);
      this.page = 1;
    },



  }
};
</script>

<style lang="scss" scoped>
.kv-list{
  padding: 10px;
}
</style>
