<template>
  <div class="ornament"></div>
  <div class="container">
    <div class="left">
      <h1 class="title">Kans</h1>
      <h2 class="subtitle">Simple web with register and login page</h2>
    </div>
    <div class="right">
      <h3 class="card-title">Login</h3>
      <form @submit.prevent="login">
        <div class="form-wrapper">
          <label for="email">Email</label>
          <input type="email" v-model="email" class="form-control" placeholder="Input your email here" required>
        </div>
        <div class="form-wrapper">
          <label for="password">Password</label>
          <input type="password" v-model="password" class="form-control" placeholder="Input your password here" required>
        </div>
        <button type="submit" class="btn btn-primary">Login Now</button>
        <p class="other-way">Don't have an account? <router-link to="/register">Register Now</router-link></p>
      </form>
      <p v-if="message" class="notification" :class="{'success': isSuccess==true, 'error':isSuccess==false}"><span class="icon icon-info"></span>{{ message }}</p>
    </div>
  </div>
</template>

<script>
  export default {
    name: "LoginPage",
    data() {
      return {
        email: '',
        password: '',
        message: ''
      }
    },
    methods: {
      async login() {
        try {
          const response = await fetch('http://localhost:8000/api/login', {
            method: 'POST',
            headers: {
              'Content-Type': 'application/json'
            },
            body: JSON.stringify({
              email: this.email,
              password: this.password
            })
          });
          if (response.ok) {
            this.isSuccess = true;
            this.message = 'Login Success';
            
            // get token n simpen di localstorage
            const data = await response.json();
            console.log('isi token', data.token, data.name);
            localStorage.setItem('token', data.token);
            localStorage.setItem('name', data.name)
            
            // Redirect ke dashboard
            this.$router.push({name: 'Dashboard'})
          } else {
            this.isSuccess = false;
            this.message = 'Login Failed';
          }
        } catch (error) {
          console.log('error', error);
          this.isSuccess = false
          this.message = "An error occured, can\n't login"
        }
      }
    }
  }
</script>

<style scoped>
.btn-primary {
  width: 100%;
}
</style>