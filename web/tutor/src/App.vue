<template lang="pug">
  #app.mdc-typography
    router-view
</template>
<script>
export default {
  mounted() {
    let deferredPrompt;
    window.addEventListener("beforeinstallprompt", e => {
      // Prevent Chrome 67 and earlier from automatically showing the prompt
      e.preventDefault();
      // Stash the event so it can be triggered later.
      deferredPrompt = e;
      // Update UI notify the user they can add to home screen
      btnAdd.style.display = "block";
    });
    // btnAdd.addEventListener("click", e => {
    //   // hide our user interface that shows our A2HS button
    //   btnAdd.style.display = "none";
    //   // Show the prompt
    //   deferredPrompt.prompt();
    //   // Wait for the user to respond to the prompt
    //   deferredPrompt.userChoice.then(choiceResult => {
    //     if (choiceResult.outcome === "accepted") {
    //       console.log("User accepted the A2HS prompt");
    //     } else {
    //       console.log("User dismissed the A2HS prompt");
    //     }
    //     deferredPrompt = null;
    //   });
    // });
  }
};
</script>
<style lang="scss">
@import "assets/shared";
@import "assets/material-components-web";
@import "~foundation-sites/scss/foundation";

@include foundation-xy-grid-classes;
@include foundation-float-classes;
@include foundation-flex-classes;
@include foundation-visibility-classes;

html,
body {
  font-family: "Roboto", sans-serif;
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
  padding: 0.25rem 1rem;
}

.mdc-list-group__subheader {
  background-color: #e0e0e0;
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
