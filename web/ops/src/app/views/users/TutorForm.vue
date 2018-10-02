<template lang="pug">
  .card
    form.form-horizontal
      .card-header
        h4.card-title
          | Tutor Form
      .card-content
        fieldset
          .form-group
            label.col-sm-2.control-label Username
            .col-sm-9
              input.form-control(type="text" name="username" v-validate="modelValidations.username" v-model="model.username")
              small.text-danger(v-show="username.invalid")
                | {{ getError('username') }}
        fieldset
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
          .form-group.row
            label.col-sm-2.control-label Contact no.
            .col-sm-9
              input.form-control(type="text" name="contact" v-validate="modelValidations.contact" v-model="model.contact")
              small.text-danger(v-show="contact.invalid")
                | {{ getError('contact') }}
          .form-group
            label.col-sm-2.control-label Photo
            .col-sm-9
              input.form-control(type="file" name="photo" v-validate="modelValidations.photo" @change="onChangePhoto")
              small.text-danger(v-show="photo.invalid")
                | {{ getError('photo') }}
          .form-group(v-if="!isCreate")
            label.col-sm-2.control-label Active
            .col-sm-9
              p-switch(v-model="model.is_active" @input="onChangeLeaving")
                i.fa.fa-check(slot="on")
                i.fa.fa-times(slot="off")

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
import PSwitch from "src/components/UIComponents/Switch.vue";

export default {
  components: {
    PSwitch
  },
  computed: {
    ...mapFields(["name", "username", "email", "photo", "contact"])
  },
  mixins: [mixinNotify],
  data() {
    return {
      isCreate: true,
      model: {
        name: "",
        username: "",
        email: "",
        photo: "",
        contact: ""
      },
      modelValidations: {
        name: {
          required: true
        },
        username: {
          required: true
        },
        contact: {},
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
      const is_deleted = !e;
      let _axios = axios;
      const url = `${process.env.API}/users/${this.model._id}`;
      if (is_deleted) {
        _axios = _axios.delete(url, config);
      } else {
        _axios = _axios.patch(url, {}, config);
      }
      _axios
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
        data.append("contact", this.model.contact);
        data.append("role", "tutor");
        if (this.model.photo) {
          data.append("photo", this.model.photo);
        }
        if (this.isCreate) {
          axios
            .post(`${process.env.API}/users`, data)
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
            .patch(`${process.env.API}/users/${this.model._id}`, data, config)
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
        .get(`${process.env.API}/users/${id}`, {
          headers: { "If-None-Match": this.model._etag },
          params: {
            // show_deleted: true
          }
        })
        .then(response => {
          this.model = response.data;
          this.model.is_active = !this.model._deleted;
          this.model.photo = null;
        })
        .catch(error => {
          this.model = error.response.data;
          this.model.is_active = !this.model._deleted;
          this.model.photo = null;
          console.log(error, error.response);
        });
    }
  }
};
</script>
<style>
</style>
