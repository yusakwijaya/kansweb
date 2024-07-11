<template>
  <div class="ornament"></div>
  <div class="container">
    <div class="left">
      <h1 class="title">Kans</h1>
      <h2 class="subtitle">Simple web with register and login page</h2>
    </div>
    <div class="right">
      <h3 class="card-title">Register</h3>
      <form @submit.prevent="register">
        <div class="form-wrapper">
          <label for="name">Name</label>
          <input type="text" class="form-control" v-model="name" placeholder="Input your name here" required>
        </div>
        <div class="form-wrapper">
          <label for="email">Email</label>
          <input type="email" class="form-control" v-model="email" placeholder="Input your email here" required>
        </div>
        <div class="form-wrapper">
          <label for="password">Password</label>
          <input type="password" class="form-control" :class="{'input-error': passNotMatch}" v-model="password" placeholder="Input your password here" required>
          <p class="helper-text error" v-if="passNotMatch">Password doesn't match</p>
        </div>
        <div class="form-wrapper">
          <label for="confirmpassword">Confirm Password</label>
          <input type="password" class="form-control" :class="{'input-error': passNotMatch}" v-model="confPassword" placeholder="Input your confirm password here" required>
          <p class="helper-text error" v-if="passNotMatch">Password doesn't match</p>
        </div>
        <button type="submit" class="btn btn-primary">Register Now</button>
        <p class="other-way">Already have an account? <router-link to="/">Login Now</router-link></p>
      </form>
      <p v-if="message" class="notification" :class="{'success': isSuccess==true, 'error':isSuccess==false}">{{ message }}</p>
    </div>
  </div>
</template>

<script>
export default {
  name: 'RegisterPage',
  data() {
    return {
      name: '',
      email: '',
      password: '',
      confPassword: '',
      message: '',
      isSuccess: false
    }
  },
  computed: {
    passNotMatch() {
      return this.password !== this.confPassword && this.confPassword.length > 0;
    }
  },
  methods: {
    async register() {
      // set notif to empty
      this.message = '';

      // check kalau pass ngga sama sama confirm pass
      if (this.passNotMatch) {
        this.message = 'Password do not match';
        return
      }
      try {
        // bikin const fetch
        const response = await fetch('http://localhost:8000/api/register', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify({
            name: this.name,
            email: this.email,
            password: this.password,
          })
        });
        if (response.ok) {
          this.isSuccess = true;
          this.message = 'Registration Successful';
          
          // Kosongin value form input
          this.name = '';
          this.email = '';
          this.password = '';
          this.confPassword = '';
        } else {
          this.isSuccess = false;
          this.message = 'Registration failed'
        }
      } catch(error) {
        this.isSuccess = false
        this.message = 'An error occured, can not register'
      }
    }
  }
}
</script>
<style lang="scss" scoped>
.btn-primary {
  width: 100%;
}
</style>