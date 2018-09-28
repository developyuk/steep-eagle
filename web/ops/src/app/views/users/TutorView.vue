<template lang="pug">
  .view
    .card
      .image
        img(:src="this.model.photo")
      .card-header
        h4 {{this.model.name}}
      .card-content
        .row
          .col-sm-3: strong Name
          .col-sm-9 {{this.model.name}}
          .col-sm-3: strong Email
          .col-sm-9 {{this.model.email}}
          .col-sm-3: strong Contact
          .col-sm-9 {{this.model.contact_no}}
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
      let config = {
        params: { where: { role: "tutor" } }
      };
      if (this.model._etag) {
        config.headers = { "If-None-Match": this.model._etag };
      }
      axios
        .get(`${process.env.API}/tutors/${id}`, config)
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