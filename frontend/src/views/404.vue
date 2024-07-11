<template>
  <div class="container">
    <div class="wrapper">
      <h1>Page Not Found</h1>
      <p class="subtitle">The page you are looking for does not exist.</p>
      <router-link to="/">Back to Home</router-link>
    </div>
  </div>
</template>

<script>
export default {
  name: "NotFoundPage"
}
</script>

<style scoped>
.subtitle {
  font-size: 1.25rem;
  margin-bottom: 1rem;
}
.container {
  justify-content: center;
  text-align: center;
}
</style>