<template lang="pug">
  #schedules.mdc-typography
    h1.title Home
    ul.mdc-list
      li.mdc-list-item(v-for="v in classes")
        .mdc-list-item__graphic
          img(':src'="v.image")
        span.mdc-list-item__text {{v.day}}
          span.mdc-list-item__secondary-text {{v.branch.name}} {{v.time}}
        button(data-mdc-auto-init="MDCRipple").mdc-button.mdc-button--raised Start
    tab-bottom

</template>

<script>
  import {MDCRipple} from '@material/ripple';
  import TabBottom from '@/components/TabBottom';
  import axios from 'axios';

  export default {
    name: 'schedules',
    components: {
      'tab-bottom': TabBottom
    },
    data() {
      return {
        msg: 'Welcome to Your Vue.js PWA',
        classes: [],
      }
    },
    methods: {
      _mergeBranch(classes) {
        classes.forEach((v, i, a) => {
          this.$set(a[i], 'branch', {'name': ''});
          axios.get(`${process.env.API}${v._links.branch.href}`)
            .then(response => {
              this.$set(a[i], 'branch', response.data);
            })
            .catch(error => console.log(error))
        });
      },
      getSchedules(page = 1) {
        const url = `${process.env.API}/classes?sort=datetime`;

        axios.get(url)
          .then(response => {
            this.classes = response.data._embedded;
            this._mergeBranch(this.classes);
          })
          .catch(error => console.log(error))
      },
    },
    mounted() {
      Array.from(document.querySelectorAll('.mdc-tab')).forEach(v => MDCRipple.attachTo(v));
      this.getSchedules()
    }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style lang="scss" scoped>
  #schedules {
    position: relative;
    height: calc(100vh - 2rem);
  }

  $size: 4rem;
  .mdc-list-item {
    height: $size+1rem;
  }

  .mdc-list-item__graphic {
    overflow: hidden;
    width: $size;
    height: $size;
    border-radius: .5rem;
  }

  .mdc-button {
    position: absolute;
    right: 1rem;
  }
</style>
