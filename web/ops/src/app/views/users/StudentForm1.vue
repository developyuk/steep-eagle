<template lang="pug">
  fieldset
    .form-group.row
      label.col-sm-2.control-label Name
      .col-sm-9
        input.form-control(type="text" name="name" v-validate="modelValidations.name" v-model="model.name")
        small.text-danger(v-show="name.invalid")
          | {{ getError('name') }}
    .form-group.row
      label.col-sm-2.control-label Email
      .col-sm-9
        input.form-control(type="email" name="email" v-validate="modelValidations.email" v-model="model.email")
        small.text-danger(v-show="email.invalid")
          | {{ getError('email') }}
    .form-group.row
      label.col-sm-2.control-label Contact no.
      .col-sm-9
        input.form-control(type="text" name="contact_no" v-validate="modelValidations.contact_no" v-model="model.contact_no")
        small.text-danger(v-show="contact_no.invalid")
          | {{ getError('contact_no') }}
</template>
<script>
import { mapFields } from "vee-validate";
import axios from "axios";

export default {
  computed: {
    ...mapFields(["name", "email", "contact_no"])
  },
  data() {
    return {
      isCreate: true,
      model: {
        name: "",
        email: "",
        contact_no: ""
      },
      modelValidations: {
        name: {
        },
        email: {
          email:true,
        },
        contact_no: {}
      }
    };
  },
  methods: {
    getError(fieldName) {
      return this.errors.first(fieldName);
    },
    validate() {
      return this.$validator.validateAll()
    }
  },
  mounted() {
    // const id = this.$route.params.id;
    // if (id) {
    //   this.isCreate = false;
    //   axios
    //     .get(`${process.env.API}/users/${id}`, {
    //       headers: { "If-None-Match": this.model._etag }
    //     })
    //     .then(response => {
    //       this.model = response.data;
    //       this.model.photo = null;
    //     })
    //     .catch(error => console.log(error, error.response));
    // }
  }
};
</script>
<style>
</style>
