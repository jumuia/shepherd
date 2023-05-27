<script setup>
import { RouterView } from 'vue-router'
import Navbar from "./components/Navbar.vue";
</script>
<script>
export default {
  data() {
    return {
      isAuthenticated: true,
    }
  }
}
</script>

<template>
  <header>

    <div class="wrapper">

      <!-- <nav>
                          <RouterLink to="/">Home</RouterLink>
                          <RouterLink to="/about">About</RouterLink>
                        </nav> -->
      <Navbar v-if="isAuthenticated"></Navbar>
    </div>
  </header>

  <RouterView />
</template>

<style scoped>
</style>
