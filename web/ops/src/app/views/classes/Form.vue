<template lang="pug">
  .card
    form.form-horizontal
      .card-header
        h4.card-title
          | Schedules Form
      .card-content
        fieldset
          .form-group
            label.col-sm-2.control-label Day
            .col-sm-9
              el-select.select-primary(size='large', placeholder='Select Day', v-model='model.day', v-validate="modelValidations.day")
                el-option.select-primary(v-for='option in selects.days', :value='option.value', :label='option.label', :key='option.label')
              small.text-danger(v-show="day.invalid")
                | {{ getError('day') }}
          .form-group
            label.col-sm-2.control-label Start at
            .col-sm-9
              el-time-select(v-model='model.start_at', :picker-options="{ start: '07:00', step: '00:15', end: '23:00' }", placeholder='Select time', v-validate="modelValidations.start_at")
              small.text-danger(v-show="start_at.invalid")
                | {{ getError('start_at') }}
          .form-group
            label.col-sm-2.control-label Finish at
            .col-sm-9
              el-time-select(v-model='model.finish_at', :picker-options="{ start: '07:00', step: '00:15', end: '23:00' }", placeholder='Select time', v-validate="modelValidations.finish_at")
              small.text-danger(v-show="finish_at.invalid")
                | {{ getError('finish_at') }}
          .form-group
            label.col-sm-2.control-label Module
            .col-sm-9
              el-select.select-primary(size='large', placeholder='Select Module', v-model='model.module', v-validate="modelValidations.module")
                el-option.select-primary(v-for='option in selects.modules', :value='option.value', :label='option.label', :key='option.label')
              small.text-danger(v-show="module.invalid")
                | {{ getError('module') }}
          .form-group
            label.col-sm-2.control-label Branch
            .col-sm-9
              el-select.select-primary(size='large', placeholder='Select Branch', v-model='model.branch', v-validate="modelValidations.branch")
                el-option.select-primary(v-for='option in selects.branches', :value='option.value', :label='option.label', :key='option.label')
              small.text-danger(v-show="branch.invalid")
                | {{ getError('branch') }}
          .form-group
            label.col-sm-2.control-label Tutor
            .col-sm-9
              el-select.select-primary(size='large', placeholder='Select Tutor', v-model='model.tutor', v-validate="modelValidations.tutor")
                el-option.select-primary(v-for='option in selects.tutors', :value='option.value', :label='option.label', :key='option.label')
              small.text-danger(v-show="tutor.invalid")
                | {{ getError('tutor') }}
          .form-group
            label.col-sm-2.control-label Students
            .col-sm-9
              el-select.select-primary.select-students(multiple, size='large', placeholder='Select Students', v-model='model.students', v-validate="modelValidations.students")
                el-option.select-primary(v-for='option in selects.students', :value='option.value', :label='option.label', :key='option.label')
              small.text-danger(v-show="students.invalid")
                | {{ getError('students') }}
      .card-footer.text-center
        .row
          .col-sm-4.col-sm-offset-2
            router-link.btn.btn-outline.btn-wd(to="/schedules") Back
          .col-sm-4
            button.btn.btn-fill.btn-wd(type="submit" @click.prevent="validate") Submit
</template>
<script>
import { mapFields } from "vee-validate";
import axios from "axios";
import mixinNotify from "src/app/mixins/notify";
import { TimeSelect, Select, Option } from "element-ui";

export default {
  components: {
    [TimeSelect.name]: TimeSelect,
    [Option.name]: Option,
    [Select.name]: Select
  },
  computed: {
    ...mapFields([
      "day",
      "start_at",
      "finish_at",
      "module",
      "branch",
      "tutor",
      "students"
    ])
  },
  mixins: [mixinNotify],
  data() {
    return {
      selects: {
        modules: [],
        branches: [],
        tutors: [],
        students: [],
        days: [
          { value: "monday", label: "Monday" },
          { value: "tuesday", label: "Tuesday" },
          { value: "wednesday", label: "Wednesday" },
          { value: "thursday", label: "Thursday" },
          { value: "friday", label: "Friday" },
          { value: "saturday", label: "Saturday" },
          { value: "sunday", label: "Sunday" }
        ]
      },
      isCreate: true,
      model: {
        day: "",
        start_at: "",
        finish_at: "",
        module: "",
        branch: "",
        tutor: "",
        students: []
      },
      modelValidations: {
        day: {
          required: true
        },
        start_at: {
          required: true
        },
        finish_at: {
          required: true
        },
        module: {
          required: true
        },
        branch: {
          required: true
        },
        tutor: {
          required: true
        },
        students: {}
      }
    };
  },
  methods: {
    getError(fieldName) {
      return this.errors.first(fieldName);
    },
    validate() {
      this.$validator.validateAll().then(isValid => {
        if (!isValid) {
          return false;
        }
        const data = {
          day: this.model.day,
          start_at: this.model.start_at,
          finish_at: this.model.finish_at,
          module: this.model.module,
          branch: this.model.branch,
          tutor: this.model.tutor,
          students_: this.model.students
        };
        if (this.isCreate) {
          axios
            .post(`${process.env.DBAPI}/classes`, data)
            .then(response => {
              this.model._etag = response.data._etag;

              this.$router.push("/schedules");
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
            .patch(
              `${process.env.DBAPI}/classes/${this.model.id}`,
              data,
              config
            )
            .then(response => {
              this.model._etag = response.data._etag;

              this.$router.push("/schedules");
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
    const params = { max_results: 999 };
    axios
      .get(`${process.env.DBAPI}/modules`, { params })
      .then(response => {
        this.selects.modules = response.data._items.map(v => {
          return { value: v.id, label: v.name.toUpperCase() };
        });
      })
      .catch(error => console.log(error, error.response));
    axios
      .get(`${process.env.DBAPI}/branches`, { params })
      .then(response => {
        this.selects.branches = response.data._items.map(v => {
          return { value: v.id, label: v.name.toUpperCase() };
        });
      })
      .catch(error => console.log(error, error.response));
    axios
      .get(`${process.env.DBAPI}/users`, {
        params: { where: { role: "tutor" }, sort: "name", max_results: 9999 }
      })
      .then(response => {
        this.selects.tutors = response.data._items.map(v => {
          return { value: v.id, label: v.name.toUpperCase() };
        });
      })
      .catch(error => console.log(error, error.response));
    axios
      .get(`${process.env.DBAPI}/users`, {
        params: { where: { role: "student" }, sort: "name", max_results: 9999 }
      })
      .then(response => {
        this.selects.students = response.data._items.map(v => {
          return { value: v.id, label: v.name.toUpperCase() };
        });
      })
      .catch(error => console.log(error, error.response));

    if (id) {
      this.isCreate = false;
      axios
        .get(`${process.env.DBAPI}/classes/${id}`, {
          params: { embedded: { students: 1, "students.student": 1 } },
          headers: { "If-None-Match": this.model._etag }
        })
        .then(response => {
          const data = response.data;
          this.model = data;
          this.model.students = data.students.map(v => v.student.id);
        })
        .catch(error => console.log(error, error.response));
    }
  }
};
</script>
<style scoped lang='scss'>
.select-students {
  width: 100%;
}
</style>
