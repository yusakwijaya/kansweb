<template>
  <div class="container">
    <div class="wrapper">
      <h1 class="title">Hellow {{name}},</h1>
      <p class="subtitle">welcome to your dashboard</p>
      <button @click="logout" class="btn btn-primary">Logout</button>
    </div>
  </div>
</template>

<script>
export default {
  name: 'DashboardPage',
  data() {
    return {
      name: ''
    }
  },
  created() {
    //get username dari local storage
    this.name = localStorage.getItem('name')
  },
  methods: {
    logout() {
      localStorage.removeItem('token');
      localStorage.removeItem('name');
      this.$router.push({name: 'Home'});
    }
  }
}
</script>

<style lang="scss" scoped>
.btn-primary {
  max-width: 120px;
}

.subtitle {
  font-size: 1.25rem;
  margin-bottom: 1rem;
}
.container {
  justify-content: center;
  text-align: center;
}
</style>