<template>
    <div class="menu">
        <div class="logo">
            Etcd-Manager
        </div>
        <Menu theme="light" active-name="1-1" class="menu-list">
            <MenuItem name="1-1" to="/keys">
                <Icon type="ios-home" style="font-size:26px;" />
                <span>KEYS</span>
            </MenuItem>
            <MenuItem name="1-2" to="/members">
                <Icon type="md-apps" style="font-size:26px;" />
                MEMBERS
            </MenuItem>
            <!-- <MenuItem name="1-3">
                <Icon type="ios-contact" style="font-size:26px;" />
                AUTH
            </MenuItem> -->
            <!-- <MenuItem name="1-4" to="/seting">
                <Icon type="ios-cog" style="font-size:26px;" />
                SETTING
            </MenuItem> -->
            <MenuItem name="1-5" to="/servers">
                <Icon type="ios-laptop" style="font-size:26px;" />
                SERVER
            </MenuItem>
            <MenuItem name="1-6" to="/Logs">
                <Icon type="logo-buffer" style="font-size:26px;" />
                Logs
            </MenuItem>
        </Menu>
    </div>
</template>

<script>
export default {
    methods: {
        openPage(name){
            console.log(name)
            this.$router.push(name);
        }
    }
};
</script>

<style scoped>
.menu .menu-list{
    font-size: 26px;
    line-height: 30px;
}
.menu .logo{
    width: 240px;
    height: 60px;
    line-height: 60px;
    font-size: 32px;
    text-align: center;
    /* background-color: #f00; */
}
</style>
