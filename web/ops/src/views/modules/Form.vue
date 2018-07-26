<template lang="pug">
  .card
    form.form-horizontal
      .card-header
        h4.card-title
          | Modules Form
      .card-content
        fieldset
          .form-group
            label.col-sm-2.control-label Name
            .col-sm-9
              input.form-control(type="text" name="name" v-validate="modelValidations.name" v-model="model.name")
              small.text-danger(v-show="name.invalid")
                | {{ getError('name') }}
        fieldset
          .form-group
            label.col-sm-2.control-label Image
            .col-sm-9
              input.form-control(type="file" name="image" v-validate="modelValidations.image")
              small.text-danger(v-show="image.invalid")
                | {{ getError('image') }}
        fieldset
          .form-group
            label.col-sm-2.control-label Program
            .col-sm-9
              el-select.select-danger(size="large" placeholder="Single Select" v-model="model.program")
                el-option.select-danger(v-for="option in selects.countries" :value="option.value" :label="option.label" :key="option.label")

              small.text-danger(v-show="program.invalid")
                | {{ getError('program') }}
      .card-footer.text-center
        .row
          .col-sm-4.col-sm-offset-2
            router-link.btn.btn-outline.btn-wd(to="/admin/modules") Cancel
          .col-sm-4
            button.btn.btn-fill.btn-wd(type="submit" @click.prevent="validate") Validate inputs

</template>
<script>
  import {mapFields} from 'vee-validate'
  import axios from 'axios'
  import mixinNotify from '../mixins/notify'
  import {Select, Option} from 'element-ui'

  export default {
    components: {
      [Option.name]: Option,
      [Select.name]: Select,
    },
    computed: {
      ...mapFields(['name', 'image','program'])
    },
    mixins: [mixinNotify],
    data() {
      return {
        selects: {
          simple: '',
          countries: [{value: 'Bahasa Indonesia', label: 'Bahasa Indonesia'},
            {value: 'Bahasa Melayu', label: 'Bahasa Melayu'},
            {value: 'Català', label: 'Català'},
            {value: 'Dansk', label: 'Dansk'},
            {value: 'Deutsch', label: 'Deutsch'},
            {value: 'English', label: 'English'},
            {value: 'Español', label: 'Español'},
            {value: 'Eλληνικά', label: 'Eλληνικά'},
            {value: 'Français', label: 'Français'},
            {value: 'Italiano', label: 'Italiano'},
            {value: 'Magyar', label: 'Magyar'},
            {value: 'Nederlands', label: 'Nederlands'},
            {value: 'Norsk', label: 'Norsk'},
            {value: 'Polski', label: 'Polski'},
            {value: 'Português', label: 'Português'},
            {value: 'Suomi', label: 'Suomi'},
            {value: 'Svenska', label: 'Svenska'},
            {value: 'Türkçe', label: 'Türkçe'},
            {value: 'Íslenska', label: 'Íslenska'},
            {value: 'Čeština', label: 'Čeština'},
            {value: 'Русский', label: 'Русский'},
            {value: 'ภาษาไทย', label: 'ภาษาไทย'},
            {value: '中文 (简体)', label: '中文 (简体)'},
            {value: 'W">中文 (繁體)', label: 'W">中文 (繁體)'},
            {value: '日本語', label: '日本語'},
            {value: '한국어', label: '한국어'}],
          multiple: 'ARS'
        },
        isCreate: true,
        model: {
          name: '',
          image: '',
          program: '',
        },
        modelValidations: {
          name: {
            required: true
          },
          image: {
            required: true
          },
          program: {
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
            name: this.model.name,
            image: this.model.image,
            program: this.model.program,
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
        axios.get(`${process.env.DBAPI}/modules/${id}`, {headers: {'If-None-Match': this.model._etag}})
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
