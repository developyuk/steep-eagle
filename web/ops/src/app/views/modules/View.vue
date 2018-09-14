<template lang="pug">
    .card
        .card-header
            h4 Modules
        .card-content
            .row
                .col-sm-3: strong Image
                .col-sm-9 
                  img(:src='this.model.image')
            .row
                .col-sm-3: strong Name
                .col-sm-9 {{this.model.name}}
</template>

<script>
import axios from "axios";
export default {
  data() {
    return {
      model: {}
    };
  },
  mounted() {
    const id = this.$route.params.id;
    if (id) {
      let config = {};
      if (this.model._etag) {
        config.headers = { "If-None-Match": this.model._etag };
      }
      axios
        .get(`${process.env.API}/modules/${id}`, config)
        .then(response => {
          this.model = response.data;
        })
        .catch(error => console.log(error, error.response));
    }
  }
};
</script>

<style>
</style>