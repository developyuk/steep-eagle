<template lang="pug">
  fieldset
    .form-group.row
      label.col-sm-2.control-label Name
      .col-sm-9
        input.form-control(type="text" name="name" v-validate="modelValidations.name" v-model="model.name")
        small.text-danger(v-show="name.invalid")
          | {{ getError('name') }}
    .form-group.row
      label.col-sm-2.control-label Date of Birth
      .col-sm-9
        el-date-picker(v-model="model.dob" type="date" placeholder="Pick date of birt" :picker-options="pickerOptions1")

    .form-group.row
      label.col-sm-2.control-label Address
      .col-sm-9
        textarea.form-control(placeholder="" rows="3" name="address" v-validate="modelValidations.address" v-model="model.address")
        small.text-danger(v-show="address.invalid")
          | {{ getError('address') }}

    .form-group.row
      label.col-sm-2.control-label School
      .col-sm-9
        input.form-control(type="text" name="school" v-validate="modelValidations.school" v-model="model.school")
        small.text-danger(v-show="school.invalid")
          | {{ getError('school') }}
    .form-group.row
      label.col-sm-2.control-label Grade
      .col-sm-9
        input.form-control(type="text" name="grade" v-validate="modelValidations.grade" v-model="model.grade")
        small.text-danger(v-show="grade.invalid")
          | {{ getError('grade') }}
    .form-group.row
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
</template>
<script>
import { mapFields } from "vee-validate";
import axios from "axios";
import { DatePicker } from "element-ui";
import PSwitch from "src/components/UIComponents/Switch.vue";
import mixinNotify from "src/app/mixins/notify";

export default {
  components: {
    [DatePicker.name]: DatePicker,
    PSwitch
  },
  computed: {
    ...mapFields(["name", "photo", "school", "grade", "address"])
  },
  mixins: [mixinNotify],
  data() {
    return {
      pickerOptions1: {
        // shortcuts: [
        //   {
        //     text: "Today",
        //     onClick(picker) {
        //       picker.$emit("pick", new Date());
        //     }
        //   },
        //   {
        //     text: "Yesterday",
        //     onClick(picker) {
        //       const date = new Date();
        //       date.setTime(date.getTime() - 3600 * 1000 * 24);
        //       picker.$emit("pick", date);
        //     }
        //   },
        //   {
        //     text: "A week ago",
        //     onClick(picker) {
        //       const date = new Date();
        //       date.setTime(date.getTime() - 3600 * 1000 * 24 * 7);
        //       picker.$emit("pick", date);
        //     }
        //   }
        // ]
      },
      isCreate: true,
      model: {
        name: "",
        photo: "",
        school: "",
        grade: "",
        address: ""
      },
      modelValidations: {
        name: {
          required: true
        },
        school: {},
        grade: {},
        address: {},
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

          this.$router.push("/admin/students");
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
    onChangePhoto(e) {
      this.model.photo = e.target.files[0];
    },
    getError(fieldName) {
      return this.errors.first(fieldName);
    },
    validate() {
      return this.$validator.validateAll();
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
