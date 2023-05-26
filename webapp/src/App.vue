<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
import Button from "./components/widgets/Button.vue"
</script>

<template>
  <header>
    <div class="wrapper">
      <nav>
        <RouterLink to="/">Home</RouterLink>
        <RouterLink to="/about">About</RouterLink>
      </nav>
    </div>
  </header>

  <RouterView />
  <Button></Button>
</template>

<style scoped>
</style>
