<template>
  <div class="logo">
    <img class="logo-img" src="@/assets/img/logo.svg" alt="datacue logo">
    <p class="logo-text">Automatically personalize your website &amp; email to each visitor</p>
  </div>
</template>

<script>
export default {};
</script>

<style lang="scss">
.logo {
  & {
    width: 240px;
    margin: 0 auto;
    margin-top: 50%;
  }

  .logo-img {
    display: block;
    margin: 0 auto;
  }

  .logo-text {
    margin-top: 1.5em;
    text-align: justify;

    font-size: 1.1rem;
  }
}

@media all and (max-width: 768px) {
  .logo {
    & {
      margin-top: 36px;
    }
  }
}
</style>

