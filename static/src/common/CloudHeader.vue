<template>
  <div class="cloud-header">
    <div class="logo" @mouseenter="hover(2, 1)" @mouseleave="hover(2, 2)">
      <!-- <img src="" alt="logo" /> -->
      Etcd-Manager
    </div>
    <div class="server" @mouseenter="hover(1, 1)" @mouseleave="hover(1, 2)">
      <Dropdown trigger="click" style="margin-left: 20px">
        <span>ETCD1
          <Icon type="ios-arrow-down"></Icon>
        </span>
        <DropdownMenu slot="list">
          <DropdownItem>ETCD1</DropdownItem>
          <DropdownItem divided>ETCD2</DropdownItem>
          <DropdownItem divided>ETCD3</DropdownItem>
        </DropdownMenu>
      </Dropdown>
    </div>
    <div class="user" @mouseenter="hover(3, 1)" @mouseleave="hover(3, 2)">
      <Avatar size="small" icon="ios-person"/>
      <Dropdown>
        <span>时光弧线
          <Icon type="ios-arrow-down"></Icon>
        </span>
        <DropdownMenu slot="list">
          <DropdownItem>{{$t('header.personalCenter')}}</DropdownItem>
          <DropdownItem style="color: #ed4014">{{$t('header.logout')}}</DropdownItem>
        </DropdownMenu>
      </Dropdown>
    </div>
    <div class="language" @mouseenter="hover(5, 1)" @mouseleave="hover(5, 2)">
      <Dropdown trigger="click" style="margin-left: 20px" @on-click="changeLanguage">
        <span>{{ showLang }}
          <Icon type="ios-arrow-down"></Icon>
        </span>
        <DropdownMenu slot="list">
          <DropdownItem name="zh" :selected="lang == 'zh'">简体中文</DropdownItem>
          <DropdownItem name="en" :selected="lang == 'en'">English</DropdownItem>
        </DropdownMenu>
      </Dropdown>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      showLang:'Lang',
      lang:'en',
    };
  },
  methods: {
    hover(typ, index) {
      if (typ == 1) {
        let obj = document.getElementsByClassName("server");
        if (index == 1) {
          obj[0].className = "server header-hover";
        } else {
          obj[0].className = "server";
        }
      } else if (typ == 2) {
        let obj = document.getElementsByClassName("logo");
        if (index == 1) {
          obj[0].className = "logo header-hover";
        } else {
          obj[0].className = "logo";
        }
      } else if (typ == 3) {
        let obj = document.getElementsByClassName("user");
        if (index == 1) {
          obj[0].className = "user header-hover";
        } else {
          obj[0].className = "user";
        }
      } else if (typ == 5) {
        let obj = document.getElementsByClassName("language");
        if (index == 1) {
          obj[0].className = "language header-hover";
        } else {
          obj[0].className = "language";
        }
      }
    },

    // 切换语言
    changeLanguage(name){
      this.lang = name;
      if (name == 'zh'){
        this.showLang = '语言';
      }else{
        this.showLang = 'Lang';
      }
      this.$i18n.locale = name;
      localStorage.setItem('etcd-language', name);
    }
  },
  created(){
    let lang = localStorage.getItem('etcd-language') || 'en';
    this.changeLanguage(lang);
  }
};
</script>

<style lang='scss'>
.cloud-header {
  width: 100%;
  height: 50px;
  line-height: 50px;
  background: #373d41;
  color: #fefefe;
  .header-hover {
    background: #2a2f32;
    cursor: pointer;
  }
  .logo {
    width: 180px;
    height: 100%;
    padding-right: 10px;
    float: left;
    line-height: 50px;
    text-align: center;
    font-size: 25px;
    font-weight: bold;
  }
  .server {
    width: auto;
    height: 100%;
    padding-right: 20px;
    float: left;
    line-height: 50px;
    text-align: left;
    font-size: 15px;
  }
  .user {
    width: auto;
    height: 100%;
    margin-right: 10px;
    float: right;
    padding: 0 20px 0 20px;
  }
  .language{
    width: auto;
    height: 100%;
    float: right;
    padding-right: 20px;
  }
}
</style>
