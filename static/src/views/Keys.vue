<template>
    <div class="keys">
        <div class="path">
            <span @click="openKeyForPath({
                    key:'',
                    name:''
                })"><Icon type="ios-home-outline" /></span>
            <span v-show="showPathInput == false">
                <span v-for="(item,key) in paths" :key="key" @click="openKeyForPath(item)"><span v-if="key > 0 && item.name != '/' ">/</span> {{ item.name }} </span>
            </span>
            
            <!-- 路径文本框 -->
            <Input v-show="showPathInput == true" @on-enter="enterPath" v-model="currentPath" :placeholder="$t('key.currentPathTips')" style="width: 300px;margin-top:-9px;font-size:28px;" />
            <span @click="showPathInput = !showPathInput" style="font-size:15px;margin-left:5px;">
                <Button type="primary" v-show="showPathInput == false" size="small">{{$t('public.edit')}}</Button>
                <Button type="primary" v-show="showPathInput == true" size="small">{{$t('public.hide')}}</Button>
            </span>

            <Poptip style="float:right;margin-top:-3px;" placement="left-end"
                confirm
                :title="$t('key.delConfirm')"
                @on-ok="delKeys">
            <Button  type="error">{{$t('public.multiDelete')}}</Button>
            </Poptip>

            <Button type="success" @click="addKeyInfoModel = true;addKeyInfo = {kType:'KEY',value:''}" style="float:right;margin-right:6px;">{{$t('public.add')}}</Button>
            <!-- 语言 -->
            <RadioGroup v-model="lang" @on-change="changeLang" type="button" style="float:right;margin-right:6px;">
                <Radio label="en">EN</Radio>
                <Radio label="zh">ZH</Radio>
            </RadioGroup>
            <!-- 展示方式切换 -->
            <RadioGroup v-model="listType" @on-change="changeListType" type="button" style="float:right;margin-right:6px;">
                <Radio label="grid"><Icon type="md-grid" /></Radio>
                <Radio label="list"><Icon type="ios-list-box-outline" /></Radio>
                <Radio label="json"><img src="../assets/imgs/json.png" alt="json" style="width:16px;height:auto;margin-top:7px"></Radio>
            </RadioGroup>

        </div>

        <div class="search">
          <Form inline>
            <FormItem>
              <Select v-model="etcdName" style="text-align: left;width:300px" @on-change="changeEtcdName">
                <Option v-for="item in etcdServers" :label="item.Title" :value="item.Name" :key="item.Name">
                  <b>{{ item.Title }}</b> 
                  <span style="float:right;color:#ccc">{{ item.Name }}</span>
                </Option>
              </Select>
            </FormItem>
            <FormItem class="search-in">
              <Input v-model="searchVal" @on-keyup="onSearchLocal">
                <Button slot="prepend" type="primary">{{$t('public.screen')}}</Button>
                <Button slot="append" type="primary" icon="ios-search" @click="onSearchLocal"></Button>
              </Input>
            </FormItem>
          </Form>

        </div>

        <div v-if="keysList.length == 0 && listType == 'grid'">{{$t('key.notSubKey')}}</div>

        <div class="lists">
            <!-- grid方式展示 -->
            <div class="key-list-main" v-if="listType == 'grid'">
                <div v-for="(item,key) in keysList" :key="key" class="key-list" @click="checkKey(item)">
                    <Checkbox class="checkbox" v-model="item.check"></Checkbox>
                    <Tooltip :content="item.full_dir" placement="top-start">
                    <div @click.stop="openKey(item)">
                        <img v-if="item.is_dir==false" src="../assets/imgs/file.png" alt="file" class="key-icon">
                        <img v-if="item.is_dir==true" src="../assets/imgs/folder.png" alt="file" class="key-icon">
                    </div>
                    <div class="title">
                        {{item.value}}
                    </div>
                    </Tooltip>
                </div>
            </div>

            <!-- 表格形式展示 -->
            <div class="table-list-main" v-if="listType == 'list'">
                <Table border :columns="columnsKey" :data="keysList1" @on-selection-change="selectionChangeKeys" :loading="keyLoading"></Table>
                <div style="margin-top:10px; text-align: right;" v-if="pageShow == true">
                  <Page @on-change="changeListPage" @on-page-size-change="pageSizeChange" show-sizer :current="page" :page-size="pageSize" :total="listTotal" />
                </div>

            </div>

            <!-- json格式展示 -->
            <div class="key-json-main" v-show="listType == 'json' || listType == 'toml' || listType == 'yaml'">
              <codemirror v-model="mainConfig" :options="cmOptionShow" style="border: 1px solid #dcdee2;" ref="listEditor"></codemirror>
            </div>

            <div style="height:200px;" v-hide="listType == 'json' || listType == 'toml' || listType == 'yaml'"></div>

            <div style="clear:both"></div>
            
        </div>

        <!-- 查看弹框 -->
        <Drawer
            :width="60"
            v-model="showKeyInfoModel"
            :title="$t('key.editKey')">
            <Form :model="showKeyInfo" :label-width="80">
                <FormItem label="Key" prop="key">
                    <Input v-model="showKeyInfo.full_dir" disabled placeholder="key"></Input>
                </FormItem>
                <FormItem label="Version" prop="version">
                    <Input v-model="showKeyInfo.version" disabled placeholder="Version"></Input>
                </FormItem>
                <FormItem label="Value" prop="value" >
                  <codemirror v-model="showKeyInfo.value" :options="cmOption" style="line-height:20px;border: 1px solid #dcdee2;" ref="showEditor"></codemirror>
                </FormItem>
                <FormItem>
                    <Button @click="saveKey" type="primary">{{$t('public.save')}}</Button>
                    <Button @click="showKeyInfoModel = false" style="margin-left: 8px">{{$t('public.close')}}</Button>
                </FormItem>

            </Form>
        </Drawer>

        <!-- 添加弹框 -->
        <Drawer
            :width="60"
            v-model="addKeyInfoModel"
            :title="$t('key.addKey')">
            <Form :model="addKeyInfo" :label-width="80">
                <FormItem label="Key" prop="key">
                    <Input v-model="addKeyInfo.key" placeholder="key">
                        <span slot="prepend">{{currentPath}}</span>
                    </Input>
                </FormItem>
                <FormItem label="KeyType" prop="kType">
                    <RadioGroup v-model="addKeyInfo.kType">
                        <Radio label="KEY"></Radio>
                        <Radio label="DIR"></Radio>
                    </RadioGroup>
                </FormItem>
                
                <FormItem label="Value" prop="value" v-show="addKeyInfo.kType == 'KEY'">
                  <codemirror v-model="addKeyInfo.value" ref="addEditor" :options="cmOption" style="line-height:20px;border: 1px solid #dcdee2;"></codemirror>
                </FormItem>

                <FormItem>
                    <Button @click="addKey" type="primary">{{$t('public.save')}}</Button>
                    <Button @click="addKeyInfoModel = false" style="margin-left: 8px">{{$t('public.close')}}</Button>
                </FormItem>

            </Form>
        </Drawer>
    </div>
</template>

<script>
require("codemirror/mode/javascript/javascript");
require("codemirror/mode/toml/toml");
require("codemirror/mode/yaml/yaml");
require("codemirror/mode/xml/xml");

// require('codemirror/addon/hint/show-hint.css');
// require('codemirror/addon/hint/show-hint');
// require('codemirror/addon/hint/javascript-hint');
// require('codemirror/addon/hint/anyword-hint');


export default {
  data() {
    return {
      // json 形式展示
      mainConfig:'',


      listTotal:0, // 总条数
      pageSize:10, // 默认10条
      page:1,
      pageShow:true, // 是否显示分页

      keyLoading: false, // 是否数据加载中

      etcdServers:[],

      lang:'en',
      etcdName: '', // etcd服务名
      keyPrefix: '', // key 前缀

      searchVal: '',   // 搜索内容

      showPathInput: false, // 是否显示路径文本框
      listType: "grid", // 显示方式
      paths: [
        {
          key: "/",
          name: ""
        }
      ], // 路径
      currentPath: "", // 当前key路径
      showKeyInfoModel: false, // 配置详情是否显示
      showKeyInfo: {}, // 要显示的配置值

      addKeyInfoModel: false, // 添加弹框
      addKeyInfo: {}, // 添加时对象

      selectionKeys: [], // 表格选中列表

      keysList: [],
      keysList1:[], // 分页使用

      columnsKey: [
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
                params.row.is_dir == true ? this.$t('key.open') : this.$t('key.show')
              ),
              h(
                "Poptip",
                {
                  props: {
                    confirm: true,
                    title: "确定删除"
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
                    this.$t('public.delete')
                  )
                ]
              )
            ]);
          }
        }
      ],
      // 代码编辑器配置
      cmOption: {
        tabSize: 4,
        smartIndent: true,
        styleActiveLine: true,
        lineNumbers: true,
        line: true,
        mode: 'text/javascript',
        lineWrapping: true,
        theme: 'default',
        // lint: true,
        // gutters: ['CodeMirror-lint-markers'],
      },
      // 显示指定格式
      cmOptionShow: {
        readOnly:'nocursor',
        tabSize: 4,
        smartIndent: true,
        styleActiveLine: true,
        lineNumbers: true,
        line: true,
        mode: 'text/javascript',
        lineWrapping: true,
        theme: 'default'
      }

    };
  },
  mounted() {
    this.getEtcdServers();

    this.lang = localStorage.getItem('lang') || 'en';
    this.listType = localStorage.getItem("list_type") || 'grid';
    this.currentPath = this.keyPrefix;
    this.etcdName = localStorage.getItem("etcdName") || '';

    // 编辑器高度
    this.$refs.addEditor.codemirror.setSize('auto','60vh');
    this.$refs.showEditor.codemirror.setSize('auto','60vh');
    this.$refs.listEditor.codemirror.setSize('auto','75vh');
    
  },
  methods: {
    // 路径文本框回车
    enterPath() {
      this.openKeyForPath({
        key: this.currentPath,
        name: ""
      });
    },
    // 选中key
    checkKey(item) {
      item.check = !item.check;
    },
    // 打开key
    openKey(item) {
      console.log(item);
      if (item.is_dir == false) {
        this.showKeyInfoModel = true;
        // 查询key的值
        this.$http.get(`/v1/key?key=${item.full_dir}`,{
          headers:{
            "EtcdServerName":this.etcdName,
          }
        }).then(response=>{
          if(response.status == 200){
            this.showKeyInfo = response.data;
            console.log(this.showKeyInfo)
          }
        }).catch(error=>{
          if (error.response){
            this.$Message.error(error.response.data.msg);
          }
        })
      } else {
        this.currentPath = item.full_dir;
        this.getKeyList();
      }
    },
    // 顶部路径打开目录
    openKeyForPath(item) {
      console.log(item)
      this.currentPath = item.key || this.keyPrefix;
      if (this.currentPath == "/" && this.currentPath != this.keyPrefix){
        this.currentPath = this.keyPrefix;
      }
      console.log(this.currentPath)

      this.getKeyList();
    },

    // 添加key
    addKey() {
      if (this.addKeyInfo.key == "" || typeof this.addKeyInfo.key == "undefined") {
        this.$Message.warning("key不能为空");
        return;
      }
      // console.log(this.currentPath.trim('/'))
      let fullDir = this.currentPath.trim('/');
      if (fullDir == '/'){
        fullDir = '';
      }
      // 请求参数
      let postData = {
        full_dir: fullDir + "/" + this.addKeyInfo.key,
        is_dir: this.addKeyInfo.kType == "DIR",
        value: this.addKeyInfo.value,
      };

      this.$http
        .post('/v1/key', postData,{
          headers:{
            "EtcdServerName":this.etcdName,
          }
        }).then(response => {
          console.log(response);
          if(response.status == 200){
            this.$Message.success("添加成功！");
            this.getKeyList();
            this.addKeyInfo = {};
            this.addKeyInfoModel = false;
          }
        }).catch(error=>{
          if (error.response){
            this.$Message.error(error.response.data.msg);
          }
        });
    },

    saveKey() {
      let putData = this.showKeyInfo;
      putData.is_dir = false;
      this.$http
        .put(`/v1/key`, putData,{
          headers:{
            "EtcdServerName":this.etcdName,
          }
        })
        .then(response => {
          console.log(response);
          if(response.status == 200){
            this.$Message.success("保存成功！");
            this.getKeyList();
            this.showKeyInfoModel = false;
          }
        }).catch(error=>{
          if (error.response){
            this.$Message.error(error.response.data.msg);
          }
        });
    },

    // 批量删除key
    delKeys() {
      this.keysList.forEach((val, index) => {
        // console.log(val._checked);
        if (val.check == true && this.listType == "grid") {
          this.delOneKey(val.full_dir);
        }
      });
      if (this.listType == "list") {
        this.selectionKeys.forEach(val => {
          this.delOneKey(val.full_dir);
        });
        this.selectionKeys = [];
      }
    },

    // 删除一个key
    delOneKey(key) {
      this.$http.delete(`/v1/key?key=${key}`,{
          headers:{
            "EtcdServerName":this.etcdName,
          }
        })
      .then(response => {
          console.log(response);
          if(response.status == 200){
            this.getKeyList();
          }
        }).catch(error=>{
          if (error.response){
            this.$Message.error(error.response.data.msg);
          }
        });
    },

    // 表格选中行
    selectionChangeKeys(selections) {
      this.selectionKeys = selections;
      console.log(selections);
    },

    // 获取key列表
    getKeyList() {
      if(this.listType == 'json' || this.listType == 'toml' || this.listType == 'yaml'){
        this.getKeyShowConfig();
      }

      this.keyLoading = true;
      this.$Loading.start();
      let k = this.currentPath;
      if (k == "" || k == "/"){
        k = this.keyPrefix;
      }
      console.log(this.keyPrefix);
      console.log(k);
      this.baseList = [];
      this.keysList = [];
      this.keysList1 = [];
      this.$http.get(`/v1/list?key=${k}`,{
          headers:{
            "EtcdServerName":this.etcdName,
          }
        }).then(response => {
        console.log(response);
        if(response.status == 200){
          this.baseList = response.data || [];
          this.keysList = response.data || [];
          this.listTotal = this.keysList.length;
          this.changeListPage(1);
          this.page = 1;
          this.$Loading.finish();
        }else{
          this.$Loading.error();
        }
        this.keyLoading = false;
      }).catch(error=>{
        if (error.response){
          this.$Message.error(error.response.data.msg);
        }
        this.keyLoading = false;
        this.$Loading.error();
      });
    },

    // 搜索本地
    onSearchLocal(){
      console.log(this.searchVal)
      if (this.searchVal == ''){
        this.pageShow = true;
      }else{
        this.pageShow = false;
      }
      let list = [];
      this.baseList.forEach(val => {
        let fullDir = val.full_dir.substring(this.currentPath.length+1);
        // console.log(fullDir)
        if(fullDir.indexOf(this.searchVal) == 0 || this.searchVal == ''){
          list.push(val);
        }
      });
      this.keysList1 = list;
      this.keysList = list;
      if (this.searchVal == '' && this.listType == 'list'){
        this.changeListPage(1);
        this.page = 1;
      }
    },

    // 展现方式
    changeListType(){
      localStorage.setItem("list_type", this.listType || 'grid');
      if(this.listType == 'json' || this.listType == 'toml' || this.listType == 'yaml'){        
        this.getKeyShowConfig();
        return
      }
      console.log(this.baseList)
      this.changeListPage(1);
      this.page = 1;

      this.onSearchLocal();
    },

    // 切换语言
    changeLang(){
      this.$i18n.locale = this.lang || 'en';
      localStorage.setItem("lang", this.lang || 'en');
    },

    // 切换页码
    changeListPage(page){
      // pageSize
      this.keysList1.splice(0, this.keysList1.length);
      this.keysList1=this.baseList.slice((page - 1) * this.pageSize,page * this.pageSize);
      console.log(page);
    },
    // 页面打小
    pageSizeChange(pageSize){
      this.pageSize = pageSize;
      this.changeListPage(1);
      this.page = 1;
    },

    // 获取etcd server列表
    getEtcdServers(){
      this.$http.get(`/v1/server`).then(response => {
          console.log(response);
          if(response.status == 200){
              this.etcdServers = response.data || [];
              if (this.etcdServers.length > 0){
                if (this.etcdName == ''){
                  this.etcdName = this.etcdServers[0].Name;
                  this.keyPrefix = this.etcdServers[0].KeyPrefix;
                  this.currentPath = this.keyPrefix;
                }else{
                  this.etcdServers.forEach(val => {
                    if(val.Name == this.etcdName){
                      this.keyPrefix = val.KeyPrefix;
                      this.currentPath = this.keyPrefix;
                    }
                  });
                }
                this.getKeyList();
              }
              localStorage.setItem("etcdName",this.etcdName)
              console.log(this.etcdServers)
          }
      }).catch(error=>{
          if (error.response){
              this.$Message.error(error.response.data.msg);
          }
      });
    },

    // 切换服务端
    changeEtcdName(val){
      console.log(val);
      this.etcdName = val;
      this.etcdServers.forEach(v => {
        if (v.Name == val){
          this.keyPrefix = v.KeyPrefix;
        }
      });
      this.currentPath = this.keyPrefix;
      this.getKeyList();
      localStorage.setItem("etcdName",this.etcdName)
    },

    // 获取当前key的json展现方式
    getKeyShowConfig(){
      this.$Loading.start();
      this.$http.get(`/v1/key/format?format=${this.listType}&key=${this.currentPath}`,{
          headers:{
            "EtcdServerName":this.etcdName,
          }
        }).then(response=>{
        if(response.status == 200){
          this.mainConfig = response.data;
          console.log(response)
          this.$Loading.finish();
        }else{
          this.$Loading.error();
        }
      }).catch(error=>{
          if (error.response){
              this.$Message.error(error.response.data.msg);
          }
          this.$Loading.error();
      });
    },


  },
  watch:{
    currentPath(newCurrentPath){
      let paths = newCurrentPath.split("/");
      if (paths.length == 0){
        paths = [];
        paths.push(this.keyPrefix);
      }
      console.log(paths)
      let fullDir = '';
      this.paths = [];
      paths.forEach(val => {
        console.log(this.keyPrefix.indexOf("/"))
        if (this.keyPrefix.indexOf("/") == 0){
          if (fullDir != '/'){
            fullDir += '/';
          }
        } else if (fullDir != ''){
          fullDir += '/';
        }
        fullDir += val;
        this.paths.push({
          key: fullDir, // .trim('/')
          name: val
        });
        console.log(this.paths)
      });

    }
  }
};
</script>

<style scoped>
.keys {
  width: 100%;
  height: 100vh;
  overflow-y: scroll;
  overflow-x: hidden;
}
.keys .path {
  width: 100%;
  font-size: 24px;
  border-bottom: 1px solid #cecece;
}
.keys .path span {
  margin: 5px 0px 13px 0px;
  color: #333333;
  cursor: pointer;
}

.keys .lists {
  position: relative;
  width: 100%;
  height: 100vh;
  /* background-color: aqua; */
}

.keys .lists .key-list-main {
  position: absolute;
  left: 0;
  top: 0;
}

.keys .lists .key-list {
  position: relative;
  width: 110px;
  height: 110px;
  padding: 15px;
  margin: 10px;
  float: left;
  text-align: center;
  background-color: #eeeeee;
}

.keys .lists .key-list .title {
  font-size: 18px;
  width: 60px;
  overflow: hidden;
  /* text-overflow:ellipsis; */
  white-space: nowrap;
}

.keys .lists .key-list .checkbox {
  position: absolute;
  left: 10px;
  top: 10px;
}

.keys .lists .key-icon {
  width: 60px;
  height: 60px;
}

.keys .lists .table-list-main {
  margin-top: 10px;
}

.ivu-poptip-body-message{
  display: inline-block !important;
}

.search{
  margin-top: 10px;
  width: 100%;
  height: 36px;
  text-align: center;
}

.search .search-in{
  width: 50%;
  min-width: 300px;
  margin: 0 auto;
}

</style>
