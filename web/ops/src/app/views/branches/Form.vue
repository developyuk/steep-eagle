<template lang="pug">
  .card
    form.form-horizontal
      .card-header
        h4.card-title
          | Branch Form
      .card-content
        fieldset
          .form-group
            label.col-sm-2.control-label Name
            .col-sm-9
              input.form-control(type="text" name="name" v-validate="modelValidations.name" v-model="model.name")
              small.text-danger(v-show="name.invalid")
                | {{ getError('name') }}
      .card-footer.text-center
        .row
          .col-sm-4.col-sm-offset-2
            router-link.btn.btn-outline.btn-wd(to="/admin/branches") Back
          .col-sm-4
            button.btn.btn-fill.btn-wd(type="submit" @click.prevent="validate") Submit
</template>
<script>
  import {mapFields} from 'vee-validate'
  import axios from 'axios'
  import mixinNotify from 'src/app/mixins/notify'

  export default {
    computed: {
      ...mapFields(['name'])
    },
    mixins: [mixinNotify],
    data() {
      return {
        isCreate: true,
        model: {
          name: '',
        },
        modelValidations: {
          name: {
            required: true
          },
        }
      }
    },
    methods: {
      getError(fieldName) {
        return this.errors.first(fieldName)
      },
      validate() {
        this.$validator.validateAll().then(isValid => {
          const data = {
            name: this.model.name
          };
          if (this.isCreate) {
            axios.post(`${process.env.DBAPI}/branches`, data)
              .then(response => {
                this.model._etag = response.data._etag;
                this.notifyVue({
                  component: {
                    template: `<span>Success created</span>`
                  },
                  type: 'success'
                })
              })
              .catch(error => {
                console.log(error, error.response)
                this.notifyVue({
                  component: {
                    template: `<span>Fail created</span>`
                  },
                  type: 'danger'
                })
              })
          } else {
            const config = {
              headers: {'If-Match': this.model._etag}
            };
            axios.patch(`${process.env.DBAPI}/branches/${this.model.id}`, data, config)
              .then(response => {
                this.model._etag = response.data._etag;

                this.notifyVue({
                  component: {
                    template: `<span>Success updated</span>`
                  },
                  type: 'success'
                })
              })
              .catch(error => {
                console.log(error, error.response)
                this.notifyVue({
                  component: {
                    template: `<span>Fail updated</span>`
                  },
                  type: 'danger'
                })
              })
          }
        })
      }
    },
    mounted() {
      const id = this.$route.params.id;
      if (id) {
        this.isCreate = false;
        axios.get(`${process.env.DBAPI}/branches/${id}`, {headers: {'If-None-Match': this.model._etag}})
          .then(response => {
            this.model = response.data;
          })
          .catch(error => console.log(error, error.response))
      }
    }
  }
</script>
<style>
</style>
