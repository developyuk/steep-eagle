<template lang="pug">
  .card
    form.form-horizontal
      .card-header
        h4.card-title
          | Admin Form
      .card-content
        fieldset
          .form-group
            label.col-sm-2.control-label Username
            .col-sm-9
              input.form-control(type="text" name="username" v-validate="modelValidations.username" v-model="model.username")
              small.text-danger(v-show="username.invalid")
                | {{ getError('username') }}
          .form-group
            label.col-sm-2.control-label Password
            .col-sm-9
              button(@click.prevent="onClickResetPassword") Reset Password
          .form-group
            label.col-sm-2.control-label Name
            .col-sm-9
              input.form-control(type="text" name="name" v-validate="modelValidations.name" v-model="model.name")
              small.text-danger(v-show="name.invalid")
                | {{ getError('name') }}
          .form-group
            label.col-sm-2.control-label Email
            .col-sm-9
              input.form-control(type="email" name="email" v-validate="modelValidations.email" v-model="model.email")
              small.text-danger(v-show="email.invalid")
                | {{ getError('email') }}
          .form-group
            label.col-sm-2.control-label Photo
            .col-sm-9
              input.form-control(type="file" name="photo" v-validate="modelValidations.photo" @change="onChangePhoto")
              small.text-danger(v-show="photo.invalid")
                | {{ getError('photo') }}
          .form-group
            label.col-sm-2.control-label Is leaving ?
            .col-sm-9
              p-switch(v-model="model.is_deleted" @input="onChangeLeaving")
                i.fa.fa-check(slot="on")
                i.fa.fa-times(slot="off")
      .card-footer.text-center
        .row
          .col-sm-4.col-sm-offset-2
            router-link.btn.btn-outline.btn-wd(to="/admin/operations") Back
          .col-sm-4
            button.btn.btn-fill.btn-wd(type="submit" @click.prevent="validate") Submit
</template>
<script>
import { mapFields } from "vee-validate";
import axios from "axios";
import mixinNotify from "src/app/mixins/notify";
import PSwitch from "src/components/UIComponents/Switch.vue";

export default {
  components: {
    PSwitch
  },
  computed: {
    ...mapFields(["name", "username", "email", "photo"])
  },
  mixins: [mixinNotify],
  data() {
    return {
      isCreate: true,
      model: {
        name: "",
        username: "",
        email: "",
        photo: ""
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
        },
        photo: {
          image: true
        }
      }
    };
  },
  methods: {
    onChangeLeaving(e) {
      const config = {
        headers: { "If-Match": this.model._etag }
      };
      const data = {
        is_deleted: e
      };
      axios
        .patch(`${process.env.API}/users/${this.model.id}`, data, config)
        .then(response => {
          this.model._etag = response.data._etag;

          this.$router.push("/admin/operations");
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
    },
    onClickResetPassword(e) {},
    onChangePhoto(e) {
      this.model.photo = e.target.files[0];
    },
    getError(fieldName) {
      return this.errors.first(fieldName);
    },
    validate() {
      this.$validator.validateAll().then(isValid => {
        if (!isValid) {
          return false;
        }
        const data = new FormData();
        data.append("name", this.model.name);
        data.append("username", this.model.username);
        data.append("email", this.model.email);
        data.append("role", "operation");
        if (this.model.photo) {
          data.append("photo", this.model.photo);
        }
        if (this.isCreate) {
          axios
            .post(`${process.env.API}/users`, data)
            .then(response => {
              this.model._etag = response.data._etag;

              this.$router.push("/admin/operations");
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
            .patch(`${process.env.API}/users/${this.model.id}`, data, config)
            .then(response => {
              this.model._etag = response.data._etag;

              this.$router.push("/admin/operations");
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
        .get(`${process.env.API}/users/${id}`, {
          headers: { "If-None-Match": this.model._etag }
        })
        .then(response => {
          this.model = response.data;
          this.model.photo = null;
        })
        .catch(error => console.log(error, error.response));
    }
  }
};
</script>
<style>
</style>
