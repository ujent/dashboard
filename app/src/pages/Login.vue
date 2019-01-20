<template>
  <div class="login-page">
    <aside class="login-page-aside">
      <login-logo/>
    </aside>
    <main class="login-page-main">
      <div class="login-page-sign-up">
        Don't have an account?
        <a href="#">Sign up</a>
      </div>
      <div class="login-form">
        <header class="login-form-header">
          <img class="login-page-plant-1" src="@/assets/img/plants-left.svg" alt="plant">
          <h1>Sign in to DataCue</h1>
        </header>
        <div>
          <label for="emailInput">Email</label>
          <input id="emailInput" type="text">
        </div>
        <div>
          <label for="passwordInput">Password</label>
          <input id="passwordInput" type="text">
          
          <span>Show password</span>
        </div>

        <button class="login-button">Login</button>
        
        <img class="login-page-plant-2" src="@/assets/img/plants-right.svg" alt="plant">
      </div>
    </main>
  </div>
</template>

<script>
import LoginBackground from "./Login/LoginBackground.vue";
import LoginLogo from "./Login/LoginLogo.vue";

export default {
  components: {
    "login-logo": LoginLogo,
    "login-background": LoginBackground
  }
};
</script>

<style lang="scss">
.login-page {
  display: flex;
  height: 100vh;
}
.login-page-aside {
  flex: 1 0 auto;
  background-image: url("../assets/img/login-background.svg");
  background-repeat: no-repeat;
  background-size: cover;
}

.login-page-main {
  & {
    flex: 1 0 auto;
    padding: 1.2em;
  }

  .login-page-sign-up {
    text-align: right;
  }

  .login-form {
    & {
      max-width: 600px;
      min-width: 450px;
      margin: 180px auto 0 auto;
    }

    .login-form-header {
      & {
        position: relative;
        text-align: center;

        border-bottom: #e6e7e9 2px solid;
        margin-bottom: 32px;
      }

      h1 {
        font-size: 1.6rem;
      }
    }

    .login-page-plant-1 {
      position: absolute;
      left: 0;
      bottom: 4px;
    }

    .login-page-plant-2 {
      float: right;
    }

    input {
      display: block;
      width: 100%;
      height: 46px;
      line-height: 46px;
      border-radius: 6px;
      background-color: #f7f7f7;
      border: #e6e7e9 1px solid;
    }

    button {
      color: #ffffff;
      font-size: 1.2rem;
      background-color: #fdb80a;
      border-radius: 6px;
      border: #fdb80a 1px solid;
      min-width: 132px;
      height: 46px;
      margin-top: 70px;
    }
  }
}
</style>