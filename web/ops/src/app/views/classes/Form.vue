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
              el-select.select-primary(filterable, default-first-option, size='large', placeholder='Select Day', v-model='model._day', v-validate="modelValidations.schedule")
                el-option.select-primary(v-for='option in selects.days', :value='option.value', :label='option.label', :key='option.label')
              small.text-danger(v-show="schedule.invalid")
                | {{ getError('schedule') }}
          .form-group
            label.col-sm-2.control-label Start at
            .col-sm-9
              el-time-select(v-model='model._start', :picker-options="{ start: '07:00', step: '00:30', end: '23:00' }", placeholder='Select time', v-validate="modelValidations._start")
              small.text-danger(v-show="_start.invalid")
                | {{ getError('_start') }}
          .form-group
            label.col-sm-2.control-label Finish at
            .col-sm-9
              el-time-select(v-model='model._finish', :picker-options="{ start: '07:00', step: '00:30', end: '23:00' }", placeholder='Select time', v-validate="modelValidations._finish")
              small.text-danger(v-show="_finish.invalid")
                | {{ getError('_finish') }}
          .form-group
            label.col-sm-2.control-label Module
            .col-sm-9
              el-select.select-primary(filterable, default-first-option, size='large', placeholder='Select Module', v-model='model.module', v-validate="modelValidations.module")
                el-option.select-primary(v-for='option in selects.modules', :value='option.value', :label='option.label', :key='option.label')
                  .row.select-image
                    .col-sm-2
                      img(:src='option.image')
                    .col-sm-10 {{option.label}}
              small.text-danger(v-show="module.invalid")
                | {{ getError('module') }}
          .form-group
            label.col-sm-2.control-label Branch
            .col-sm-9
              el-select.select-primary(filterable, default-first-option, size='large', placeholder='Select Branch', v-model='model.branch', v-validate="modelValidations.branch")
                el-option.select-primary(v-for='option in selects.branches', :value='option.value', :label='option.label', :key='option.label')
              small.text-danger(v-show="branch.invalid")
                | {{ getError('branch') }}
          .form-group
            label.col-sm-2.control-label Tutor
            .col-sm-9
              el-select.select-primary(filterable, default-first-option, size='large', placeholder='Select Tutor', v-model='model.tutor', v-validate="modelValidations.tutor")
                el-option.select-primary(v-for='option in selects.tutors', :value='option.value', :label='option.label', :key='option.label')
                  .row.select-image
                    .col-sm-2
                      img(:src='option.image')
                    .col-sm-10 {{option.label}}
              small.text-danger(v-show="tutor.invalid")
                | {{ getError('tutor') }}
          .form-group
            label.col-sm-2.control-label Students
            .col-sm-9
              p-radio(label="false" v-model="is_all_students") available students
              p-radio(label="true" v-model="is_all_students") all students
              br
              el-select.select-primary.select-students(multiple, filterable, default-first-option, size='large', placeholder='Select Students', v-model='model.students', v-validate="modelValidations.students")
                el-option.select-primary(v-for='option in selects.students', :value='option.value', :label='option.label', :key='option.label')
                  .row.select-image
                    .col-sm-2
                      img(:src='option.image')
                    .col-sm-10 {{option.label}}
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
      "schedule",
      "_start",
      "_finish",
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
        schedule: "",
        _start: "",
        _finish: "",
        module: "",
        branch: "",
        tutor: "",
        students: []
      },
      is_all_students: false,
      modelValidations: {
        schedule: {
          required: true
        },
        _start: {
          required: true
        },
        _finish: {
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
  watch: {
    is_all_students(v, ov) {
      // console.log(v, eval(v));
      let url, params;
      if (eval(v)) {
        url = `${process.env.API}/users`;
        params = {
          sort: "_deleted,name",
          max_results: 9999,
          where: { role: "student" }
        };
      } else {
        url = `${process.env.API}/students/dormant`;
        params = {
          sort: "name",
          max_results: 9999
        };
      }
      axios
        .get(url, { params })
        .then(response => {
          this.selects.students = response.data._items.map(v => {
            name = v.name || v.username;
            return {
              value: v._id,
              label: name.toUpperCase(),
              image: v.photo
            };
          });
        })
        .catch(error => console.log(error, error.response));
    }
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
          day: this.model._day,
          startAt: this.model._start,
          finishAt: this.model._finish,
          module: this.model.module,
          branch: this.model.branch,
          tutor: this.model.tutor
        };
        const dataStudents = this.model.students;
        if (this.isCreate) {
          axios
            .post(`${process.env.API}/classes`, data)
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
            .patch(`${process.env.API}/classes/${this.model._id}`, data, config)
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
          const postStudents = () => {
            if (dataStudents.length) {
              const data = dataStudents.map(v => {
                return { class: this.model._id, student: v };
              });
              axios
                .post(
                  `${process.env.API}/classes/${this.model._id}/students`,
                  data,
                  config
                )
                .then(response => console.log(response))
                .catch(error => console.log(error, error.response));
            }
          };
          axios
            .delete(
              `${process.env.API}/classes/${this.model._id}/students`,
              config
            )
            .then(response => postStudents())
            .catch(error => postStudents());
        }
      });
    }
  },
  mounted() {
    const id = this.$route.params.id;
    const params = { max_results: 999 };
    axios
      .get(`${process.env.API}/modules`, { params })
      .then(response => {
        this.selects.modules = response.data._items.map(v => {
          return {
            value: v._id,
            label: v.name.toUpperCase(),
            image: v.image
          };
        });
      })
      .catch(error => console.log(error, error.response));
    axios
      .get(`${process.env.API}/branches`, { params })
      .then(response => {
        this.selects.branches = response.data._items.map(v => {
          return {
            value: v._id,
            label: v.name.toUpperCase()
          };
        });
      })
      .catch(error => console.log(error, error.response));
    axios
      .get(`${process.env.API}/users`, {
        params: {
          sort: "_deleted,name",
          max_results: 9999,
          where: { role: "tutor" }
        }
      })
      .then(response => {
        this.selects.tutors = response.data._items.map(v => {
          const name = v.name || v.username;
          return {
            value: v._id,
            label: name.toUpperCase(),
            image: v.photo
          };
        });
      })
      .catch(error => console.log(error, error.response));

    if (id) {
      this.is_all_students = true;
      this.isCreate = false;

      axios
        .get(`${process.env.API}/classes/${id}`, {
          params: {},
          headers: { "If-None-Match": this.model._etag }
        })
        .then(response => {
          this.model = Object.assign({}, this.model, response.data);

          axios
            .get(`${process.env.API}/classes/${id}/students`, {
              params: {},
              headers: { "If-None-Match": this.model._etag }
            })
            .then(response => {
              const items = response.data._items;
              if (items.length) {
                this.model.students = items.map(v => v.student);
                this.is_all_students = false;
              }
            })
            .catch(error => console.log(error, error.response));
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
