<template>
  <div id="app">

    <CreateResume></CreateResume>
  </div>
</template>

<script>
import CreateResume from "@/components/CreateResume";

export default {
  name: 'App',
  components: {
    CreateResume,
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
