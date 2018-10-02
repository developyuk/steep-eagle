<template lang="pug">
  #wizardCard.card.card-wizard
    form-wizard(shape="tab" @on-complete="wizardComplete" error-color="#EB5E28" color="#EE215B" title="Student Form" subtitle="")
      tab-content(title="Student details" :before-change="validateFirstStep" icon="ti-user")
        first-step(ref="firstStep")
      tab-content(title="Guardian details" :before-change="validateSecondStep" icon="ti-settings")
        second-step(ref="secondStep")
      button.btn.btn-default.btn-outline.btn-wd.btn-back(slot="prev") Back
      button.btn.btn-default.btn-outline.btn-info.btn-wd.btn-next(slot="next") Next
      button.btn.btn-default.btn-info.btn-fill.btn-wd(slot="finish") Finish
</template>
<script>
import axios from "axios";
import mixinNotify from "src/app/mixins/notify";
import { FormWizard, TabContent } from "vue-form-wizard";
import "vue-form-wizard/dist/vue-form-wizard.min.css";
import FirstStep from "./StudentForm0";
import SecondStep from "./StudentForm1";

export default {
  components: {
    FormWizard,
    TabContent,
    FirstStep,
    SecondStep
  },
  computed: {},
  mixins: [mixinNotify],
  data() {
    return {
      isCreate: true,
      model: {}
    };
  },
  methods: {
    validateFirstStep() {
      return this.$refs.firstStep.validate();
    },
    validateSecondStep() {
      return this.$refs.secondStep.validate();
    },
    wizardComplete() {
      this.model = {
        ...this.$refs.firstStep.model
      };
      this.model.guardians = [this.$refs.secondStep.model];
      this.model.guardians = this.model.guardians.map(v => {
        v.role = "guardian";
        return v;
      });

      const data = new FormData();
      data.append("name", this.model.name);
      data.append("role", "student");
      if (this.model.address) {
        data.append("address", this.model.address);
      }
      if (this.model.school) {
        data.append("school", this.model.school);
      }
      if (this.model.grade) {
        data.append("grade", this.model.grade);
      }
      if (this.model.photo) {
        data.append("photo", this.model.photo);
      }
      if (this.model.dob) {
        data.append("dob", this.model.dob);
      }
      if (this.isCreate) {
        axios
          .post(`${process.env.API}/users`, data)
          .then(response => {
            this.model._etag = response.data._etag;

            this.$router.push("/admin/students");
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
      }

      if (this.model.guardians.length) {
        const url = `${process.env.API}/users/${this.model._id}/guardians`;
        const postGuardians = () => {
          axios
            .post(url, this.model.guardians)
            .then(response => console.log(response));
        };
        axios
          .delete(url)
          .then(response => postGuardians())
          .catch(error => postGuardians());
      }
    }
  },
  mounted() {
    const id = this.$route.params.id;

    if (id) {
      this.isCreate = false;
      axios
        .get(`${process.env.API}/users/${id}`, {
          headers: { "If-None-Match": this.$refs.firstStep.model._etag }
        })
        .then(response => {
          this.$refs.firstStep.model = response.data;
          this.$refs.firstStep.model.is_active = !this.$refs.firstStep.model
            ._deleted;
          this.$refs.firstStep.model.photo = null;
        })
        .catch(error => console.log(error, error.response));

      axios
        .get(`${process.env.API}/users/${id}/guardians`, {
          headers: { "If-None-Match": this.$refs.firstStep.model._etag }
        })
        .then(response => {
          const items = response.data._items;
          if (items.length) {
            this.$refs.secondStep.model = items[0];
          }
        })
        .catch(error => console.log(error, error.response));
    }
  }
};
</script>
<style>
</style>
