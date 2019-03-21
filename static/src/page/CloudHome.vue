<template>
  <div class="cloud-home">
    <div class="layout-header">
      <CloudHeader></CloudHeader>
    </div>
    <div class="layout-container">
      <div class="layout-left">
        <CloudSideBar></CloudSideBar>
      </div> 
      <div class="layout-right">
        <CloudContainer>
          <router-view></router-view>
        </CloudContainer>
      </div> 
    </div>
    <!-- 登录弹框 -->
    <Modal
        :title="$t('login.title')"
        v-model="showLogin"
        :mask-closable="false"
        :closable="false">
        <Form :model="loginForm" :label-width="80">
          
          <FormItem :label="$t('login.username')">
              <Input v-model="loginForm.username" :placeholder="$t('login.username')"></Input>
          </FormItem>
          <FormItem :label="$t('login.password')">
              <Input v-model="loginForm.password" type="password" :placeholder="$t('login.password')"></Input>
          </FormItem>
        </Form>
        <div slot="footer">
            <Button type="success" size="large" @click="onLogin">{{ $t('login.loginBtn') }}</Button>
          </div>
    </Modal>
    <!-- end 登录弹框 -->

  </div>
</template>

<script>
import CloudHeader from "../common/CloudHeader";
import CloudSideBar from "../common/CloudSideBar";
import CloudContainer from "../common/CloudContainer";
import { bus } from "@/page/bus.js";

export default {
  name: "CloudHome",
  data() {
    return {
      loginForm: {}, // 登录信息
      showLogin: false // 是否显示登录框
    };
  },
  components: {
    CloudHeader,
    CloudSideBar,
    CloudContainer
  },
  mounted(){
    // 选择etcd服务事件
    bus.$off("show-login");
    bus.$on("show-login", item => {
      // 显示登录框
      console.log('请登录', item);
      this.showLogin = item;
    });
  },
  methods:{
    onLogin(){
      console.log(this.loginForm);
      sessionStorage.setItem('login-info', JSON.stringify(this.loginForm));
      this.$router.go(0);
    }
  }
};
</script>

<style lang='scss'>
.cloud-home {
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;
  .layout-header {
    height: 50px;
  }
  .layout-container {
    flex: 1;
    display: flex;
    .layout-left {
    }
    .layout-right {
      flex: 1;
      overflow-y: auto;
    }
  }
}
</style>
