<template>
    <div class="cloud-container">
        <slot></slot>
    </div>
</template>

<script>
export default {

}
</script>

<style>
.cloud-container{
    width: 100%;
    height: 100%;
    overflow-y: scroll;
    /* padding: 10px; */
    /* background: #f5f7f9; */
}
</style>
