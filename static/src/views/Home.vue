<template>
  <div class="home">
    <div class="left">
      <LeftMenu></LeftMenu>
    </div>
    <div class="right">
      <router-view/>
    </div>
  </div>
</template>

<script>
// @ is an alias to /src
import LeftMenu from './Menu.vue'

export default {
  name: 'home',
  components: {
    LeftMenu
  },
  mounted(){
    this.$router.push("keys");
  }
}
</script>

<style lang="css" scoped>
.home{
  widows: 100vw;
  height: 100vh;
  display: flex;
}
.home .left{
  flex: 1;
  width: 240px;
  height: 100vh;
  border-right:1px solid #cecece;
  /* background-color: #f00; */
}
.home .right{
  flex: 12;
  height: 100vh;
  padding: 10px;
  /* background-color: aqua; */
}
</style>
