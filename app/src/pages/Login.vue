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
      <div class="login-form" v-on:keypress="handleKeyPress">
        <header class="login-form-header">
          <img class="login-page-plant-1" src="@/assets/img/plants-left.svg" alt="plant">
          <h1>Sign in to DataCue</h1>
        </header>
        <div class="login-error" v-text="error"></div>
        <div>
          <label for="emailInput">Email</label>
          <input id="emailInput" v-model="email" type="email">
        </div>
        <div class="password-section">
          <label for="passwordInput">Password</label>
          
          <input id="passwordInput" v-model="password" :type="showPassword ? 'text' : 'password'">
          
          <span
            class="show-password-button"
            v-on:click="toggleShowPassword"
          >👁&nbsp;&nbsp;Show password</span>
        </div>
        <button class="login-button" v-on:click="login" :disabled="!email || !password">Login</button>
        
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
  },

  methods: {
    toggleShowPassword: function(e) {
      this.showPassword = !this.showPassword;
    },

    handleKeyPress: function(e) {
      this.error = "";
      if (e.key === "Enter" && this.email && this.password) {
        this.login();
      }
    },

    login: function() {
      this.error = "";

      this.$store
        .dispatch("login", {
          email: this.email,
          password: this.password
        })
        .then(
          () => {
            this.$router.push("/");
          },
          err => {
            if (err.status == 500) {
              this.error = "Internal Server Error";
            } else {
              this.error = "Invalid credentials";
            }
          }
        );
    }
  },

  data() {
    return {
      email: "",
      password: "",
      date: "",
      error: "",
      showPassword: ""
    };
  }
};
</script>

<style lang="scss">
.login-page {
  display: flex;
  height: 100vh;
}
.login-page-aside {
  min-width: 320px;
  flex: 0 1 640px;
  background-image: url("../assets/img/login-background.svg");
  background-repeat: no-repeat;
  background-size: contain;
  background-color: #fcf3e4;
}

.login-page-main {
  & {
    flex: 1 0 auto;
    padding: 1.2em;
  }

  a {
    color: #fdb80a;
  }

  a:hover {
    color: darkgray;
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

    .login-error {
      color: red;
      text-align: right;
      height: 21px;
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
      margin-top: 16px;
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
      padding: 0 12px;
    }

    input:focus {
      background-color: #ffffff;
      border-color: #fdb80a;
    }

    .password-section {
      margin-top: 32px;
      position: relative;
    }

    .show-password-icon {
      margin-top: 2px;
      margin-right: 8px;
    }

    .show-password-button {
      position: absolute;
      right: 0;
      top: 2px;
      cursor: pointer;
    }

    .show-password-button:hover {
      color: #fdb80a;
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

    button:hover {
      background-color: darkgray;
      border-color: darkgray;
      cursor: pointer;
    }

    button:disabled {
      background-color: #e6e7e9;
      border-color: #e6e7e9;
    }
  }
}

@media all and (max-width: 480px) {
  .login-page {
    display: block;
  }

  .login-page-aside {
    height: 320px;
    padding-top: 120px;
    background-image: url("../assets/img/login-background-mobile.svg");
    background-repeat: no-repeat;
    background-color: #fcf3e4;
  }

  .login-page-plant-1,
  .login-page-plant-2 {
    display: none;
  }

  .login-page-main {
    .login-page-sign-up {
      text-align: center;
      font-size: 1rem;
    }

    .login-form {
      & {
        margin: 60px 0 0 0;
        max-width: unset;
        min-width: unset;
      }

      .login-form-header h1 {
        font-weight: bold;
      }

      .login-button {
        width: 100%;
        height: 52px;
        font-size: 1.6rem;
      }

      .show-password-button {
        font-size: 1rem;
        top: 90px;
      }

      label {
        font-size: 1.2rem;
      }
    }
  }
}
</style>