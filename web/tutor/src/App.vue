<template lang="pug">
  #app.mdc-typography
    router-view
</template>

<style lang="scss">
  @import "~normalize.css";
  @import "assets/shared";
  @import "assets/material-components-web";

  html, body {
    font-family: 'Roboto', sans-serif;
  }

  #app {
    background-color: #f9f9f9;
    /*background-color: #fff;*/
    max-width: 30rem;
    padding: 0;
    height: 100vh;
    margin: 0 auto;
  }

  .errMsg {
    background-color: pink;
    color: red;
    padding: 1rem;
    margin: 1rem 0;
  }

  .mdc-list-item {
    padding: .25rem 1rem;
  }

  .mdc-list-group__subheader {
    background-color: #E0E0E0;
    margin: 0;
    padding: 0.75rem 16px;
  }

  /*h1.title {*/
  /*background-color: var(--mdc-theme-primary, #6200ee);*/
  /*color: #fff;*/
  /*text-align: center;*/
  /*padding: 1rem 0;*/
  /*margin: 0;*/
  /*}*/
</style>
