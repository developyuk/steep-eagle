<template lang="pug">
  template-main#progress
    tab-top
    .empty(v-if="!!progress && !progress.length")
      //- img(src="https://images.weserv.nl/?il&crop=400,100,0,80&url=dl.dropboxusercontent.com/s/n7l2woy0sqh09f8/no-progress.png")
      h1.mdc-typography--headline6 No record
      p.mdc-typography--body2 It seems you haven't rated your students from your class. Please make sure to do so
      router-link(to="/students").mdc-button.mdc-button--raised Rate Students

    .mdc-list-group
      template(v-for="(v,i) in progress")
        .mdc-list
          item(:item="v" :key="v.id")
</template>

<script>
import axios from "axios";
import TemplateMain from "@/components/views/Main";

export default {
  components: {
    TemplateMain,
    "tab-top": () => import("../TabTop"),
    item: () => import("./Item")
  },
  data() {
    return {
      progress: null
    };
  },
  methods: {
    getData() {
      const url = `${process.env.VUE_APP_API}/caches`;
      const config = {
        params: {
          where: { key: "progress_classes" }
        }
      };
      axios
        .get(url, config)
        .then(response => {
          this.progress = response.data._items[0]["value"];
          this.progress = JSON.parse(this.progress);
          this.progress = this.progress._items;
          console.log(this.progress);
        })
        .catch(error => console.log(error));
    }
  },
  mounted() {
    this.getData();
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style lang="scss" scoped>
@import "src/assets/shared";

.empty {
  position: absolute;
  top: 50%;
  /*left: 50%;
    transform: translate(-50%, -50%);*/
  transform: translateY(-60%);
  margin: 2rem;
  width: calc(100% - 4rem);
  text-align: center;

  h1 {
    color: #58595b;
  }
  p {
    color: #bdbdbd;
  }
  img {
    width: 100%;
  }
}
.mdc-button {
  background-color: map-get($palettes, green);
  width: 100%;
}
</style>
