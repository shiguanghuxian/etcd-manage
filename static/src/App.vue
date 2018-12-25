<template>
  <div id="app">
    <a href="https://github.com/shiguanghuxian/etcd-manage" target="_blank"><img style="position: absolute; bottom: 0px; right: 0; border: 0;z-index:900" src="https://raw.githubusercontent.com/shiguanghuxian/etcd-manage/master/images/forkme_green_007200.png" alt="Fork me on GitHub"></a>
    <router-view/>
  </div>
</template>

<script>
export default {
  name: 'app',
  components: {
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  /* text-align: center; */
  color: #2c3e50;
  height: 100vh;
  width: 100vw;
  overflow: hidden;
}
.ivu-poptip-body-message{
  display: inline-block !important;
}

</style>
