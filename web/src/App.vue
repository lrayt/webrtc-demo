<script setup lang="ts">
import {RouterView} from 'vue-router'

</script>

<template>
  <div class="app-header"></div>
  <div class="app-body">
    <div class="container">
      <RouterView/>
    </div>
  </div>
  <div class="app-footer"></div>
</template>

<style scoped lang="less">


.app-header {
  width: 100%;
  height: var(--header-height);
  background-color: var(--primary-color);
}

.app-body {
  width: 100%;
  height: calc(100vh - var(--header-height) - var(--footer-height));
  background-color: var(--bg-color);
  overflow-y: auto;

  .container {
    width: 1680px;
    margin: auto;
    padding: 32px 16px;
  }
}

.app-footer {
  width: 100%;
  height: var(--footer-height);
  background-color: var(--primary-color);
}
</style>
