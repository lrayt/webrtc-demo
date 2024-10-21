<script setup lang="ts">
import {RouterView} from 'vue-router'
</script>

<template>
  <div class="container">
    <div class="header"></div>
    <div class="body">
      <router-view/>
    </div>
  </div>
</template>

<style scoped lang="less">
.container {
  height: 100%;
  width: 100%;
  background-color: #f8f8f8;

  .header {
    height: 54px;
    width: 100%;
    background-color: #181818;
  }

  .body {
    height: calc(100% - 54px);
    width: 100%;
    background-color: #2c3e50;
  }
}
</style>
