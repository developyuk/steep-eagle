<template lang="pug">
  .card
    form.form-horizontal
      .card-header
        h4.card-title
          | Tutor Form
      .card-content
        fieldset
          .form-group
            label.col-sm-2.control-label Name
            .col-sm-9
              input.form-control(type="text" name="name" v-validate="modelValidations.name" v-model="model.name")
              small.text-danger(v-show="name.invalid")
                | {{ getError('name') }}
          .form-group
            label.col-sm-2.control-label Username
            .col-sm-9
              input.form-control(type="text" name="username" v-validate="modelValidations.username" v-model="model.username")
              small.text-danger(v-show="username.invalid")
                | {{ getError('username') }}
          .form-group
            label.col-sm-2.control-label Email
            .col-sm-9
              input.form-control(type="email" name="email" v-validate="modelValidations.email" v-model="model.email")
              small.text-danger(v-show="email.invalid")
                | {{ getError('email') }}
      .card-footer.text-center
        .row
          .col-sm-4.col-sm-offset-2
            router-link.btn.btn-outline.btn-wd(to="/admin/tutors") Back
          .col-sm-4
            button.btn.btn-fill.btn-wd(type="submit" @click.prevent="validate") Submit
</template>
<script>
import { mapFields } from "vee-validate";
import axios from "axios";
import mixinNotify from "src/app/mixins/notify";

export default {
  computed: {
    ...mapFields(["name", "username", "email"])
  },
  mixins: [mixinNotify],
  data() {
    return {
      isCreate: true,
      model: {
        name: "",
        username: "",
        email: ""
      },
      modelValidations: {
        name: {
          required: true
        },
        username: {
          required: true
        },
        email: {
          email: true
        }
      }
    };
  },
  methods: {
    getError(fieldName) {
      return this.errors.first(fieldName);
    },
    validate() {
      this.$validator.validateAll().then(isValid => {
        const data = {
          name: this.model.name,
          username: this.model.username,
          email: this.model.email,
          role: "tutor"
        };
        if (this.isCreate) {
          axios
            .post(`${process.env.DBAPI}/users`, data)
            .then(response => {
              this.model._etag = response.data._etag;

              this.$router.push("/admin/tutors");
              this.notifyVue({
                component: {
                  template: `<span>Success created</span>`
                },
                type: "success"
              });
            })
            .catch(error => {
              console.log(error, error.response);
              this.notifyVue({
                component: {
                  template: `<span>Fail created</span>`
                },
                type: "danger"
              });
            });
        } else {
          const config = {
            headers: { "If-Match": this.model._etag }
          };
          axios
            .patch(`${process.env.DBAPI}/users/${this.model.id}`, data, config)
            .then(response => {
              this.model._etag = response.data._etag;

              this.$router.push("/admin/tutors");
              this.notifyVue({
                component: {
                  template: `<span>Success updated</span>`
                },
                type: "success"
              });
            })
            .catch(error => {
              console.log(error, error.response);
              this.notifyVue({
                component: {
                  template: `<span>Fail updated</span>`
                },
                type: "danger"
              });
            });
        }
      });
    }
  },
  mounted() {
    const id = this.$route.params.id;
    if (id) {
      this.isCreate = false;
      axios
        .get(`${process.env.DBAPI}/users/${id}`, {
          headers: { "If-None-Match": this.model._etag }
        })
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
