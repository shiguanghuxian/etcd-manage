<template>
  <div class="kv-grid">
    <div v-for="(item,key) in list" :key="key" class="key-list" @click="checkKey(item)">
      <Checkbox class="checkbox" v-model="item.check"></Checkbox>
      <Tooltip :content="item.full_dir" placement="top" transfer>
        <div @click.stop="openKey(item)">
          <img v-if="item.is_dir==false" src="../assets/imgs/file.png" alt="file" class="key-icon">
          <img v-if="item.is_dir==true" src="../assets/imgs/folder.png" alt="file" class="key-icon">
        </div>
        <div class="title">{{item.value}}</div>
      </Tooltip>
    </div>
  </div>
</template>
<script>
import { KV } from "@/api/kv.js";

// grid方式展示key
export default {
  data() {
    return {
      list: [],
      dir: "", // 要显示的目录
      selected: [], // 选中的key
    };
  },
  mounted() {},
  methods: {
    getList(dir) {
      this.list = [];
      this.dir = dir;
      KV.GetKeyList(dir).then(response => {
        this.list = response.data || [];
        // console.log(response);
      });
    },
    checkKey(item) {
      item.check = !item.check;
      
      let selected = [];
      this.list.forEach(val => {
        // console.log(val)
        if(val.check == true){
          selected.push(val);
        }
      });
      this.selected = selected;
      console.log(selected)
    },
    openKey(item) {
      this.$emit("openKey", item);
    },
    // 清除
    clearSelected(){
      this.selected = [];
    }
  }
};
</script>

<style lang="scss" scoped>
.key-list {
  position: relative;
  width: 110px;
  height: 110px;
  padding: 15px;
  margin: 10px;
  float: left;
  text-align: center;
  background-color: #eeeeee;
  .key-icon {
    width: 60px;
    height: 60px;
  }
  .title {
    font-size: 18px;
    width: 60px;
    overflow: hidden;
    white-space: nowrap;
  }
  .checkbox {
    position: absolute;
    left: 10px;
    top: 10px;
  }
  .title {
    position: relative;
    width: 80px;
  }
}
</style>
