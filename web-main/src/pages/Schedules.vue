<template lang="pug">
  #schedules.mdc-typography
    header
      img.logo(src="static/img/logo.svg")
      span.search(@click="comingSoon($event)") search?
    ul.mdc-list
      li.mdc-list-item(v-for="(v,i) in classes")
        .mdc-list-item__graphic
          img(':src'="v.module.image")
        span.mdc-list-item__text {{v.module.name}}
          span.mdc-list-item__secondary-text {{v.branch.name}} {{v.day}} {{v.time}}
        button(v-if="buttonStatus(v)==='start'" data-mdc-auto-init="MDCRipple" @click='start($event,v.id,i)').mdc-button.mdc-button--raised.mdc-button--compact Start
        button(v-if="buttonStatus(v)==='disabled'" disabled data-mdc-auto-init="MDCRipple" @click='start($event,v.id,i)').mdc-button.mdc-button--raised.mdc-button--compact Start
        button(v-if="buttonStatus(v)==='late'" data-mdc-auto-init="MDCRipple" @click='start($event,v.id,i)').mdc-button.mdc-button--raised.mdc-button--compact Activate
        span.status(v-if="buttonStatus(v)==='ongoing'") ongoing
    aside#my-mdc-dialog.mdc-dialog(role='alertdialog' aria-labelledby='my-mdc-dialog-label' aria-describedby='my-mdc-dialog-description')
      .mdc-dialog__surface
        header.mdc-dialog__header
          h2#my-mdc-dialog-label.mdc-dialog__header__title
            | Start this class?
        section#my-mdc-dialog-description.mdc-dialog__body
          | Start this class
        footer.mdc-dialog__footer
          button.mdc-button.mdc-dialog__footer__button.mdc-dialog__footer__button--cancel(type='button') No
          button.mdc-button.mdc-dialog__footer__button.mdc-dialog__footer__button--accept(type='button') Yes
      .mdc-dialog__backdrop

    .mdc-snackbar(aria-live='assertive' aria-atomic='true' aria-hidden='true')
      .mdc-snackbar__text
      .mdc-snackbar__action-wrapper
        button.mdc-snackbar__action-button(type='button')
    tab-bottom

</template>

<script>
  import {MDCRipple} from '@material/ripple';
  import TabBottom from '@/components/TabBottom';
  import moment from 'moment';
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
        dialog: null,
        snackbar: null,
        thisClass: {id: 0},
      }
    },
    methods: {
      comingSoon(e){
        e.target.innerText = 'coming soon';
      },
      buttonStatus(thisClass) {
        console.log(moment(thisClass.ts).toString(),moment().toString(),moment(thisClass.ts).isAfter(moment()));
        if(moment(thisClass.ts).isAfter(moment())){
          return 'disabled';
        }
        return 'start';
      },
      activate(cid) {
        const url = `${process.env.API}/classes/${cid}/sessions`;
        axios.post(url, {
          id: 3
        })
          .then(response => {
            console.log(response);
          })
          .catch(error => console.log(error))
      },
      start(e, id, i) {
        this.thisClass = this.classes[i];

        this.dialog.lastFocusedTarget = e.target;
        this.dialog.show();
      },
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
      _mergeModule(classes) {
        classes.forEach((v, i, a) => {
          this.$set(a[i], 'module', {'name': ''});
          axios.get(`${process.env.API}${v._links.module.href}`)
            .then(response => {
              response.data.image = response.data.image.replace('https://', '').replace('http://', '');
              response.data.image = `//images.weserv.nl/?output=jpg&il&q=100&w=64&h=64&t=square&url=${response.data.image}`;
              this.$set(a[i], 'module', response.data);
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
            this._mergeModule(this.classes);
          })
          .catch(error => console.log(error))
      },
    },
    mounted() {
      const cls = this;
      Array.from(document.querySelectorAll('.mdc-tab, .mdc-button')).forEach(v => MDCRipple.attachTo(v));
      this.dialog = mdc.dialog.MDCDialog.attachTo(document.querySelector('#my-mdc-dialog'));
      this.snackbar = mdc.snackbar.MDCSnackbar.attachTo(document.querySelector('.mdc-snackbar'));
      this.dialog.listen('MDCDialog:accept', function () {
        cls.activate(cls.thisClass.id);
        const dataObj = {
          message: `Class ${cls.thisClass.id} has been started`,
          actionText: 'Undo',
          actionHandler: function () {
            console.log('my cool function');
          }
        };
        cls.snackbar.show(dataObj);
      });
      this.getSchedules();
    }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style lang="scss" scoped>
  #schedules {
    position: relative;
    height: 100vh;
  }

  header {
    position: relative;
  }

  span.search {
    color: #fff;
    text-shadow: 2px 2px 4px rgba(0, 0, 0, .5);
    position: absolute;
    right: 1rem;
    top: 50%;
    transform: translateY(-50%);
    font-weight: 500;
    text-transform: capitalize;
  }

  .mdc-list {
    padding: 0;
    height: calc(100vh - 8rem);
    overflow-y: auto;
  }

  $size: 3rem;
  .mdc-list-item {
    height: $size+2rem;
    border-bottom: thin solid rgba(0,0,0,.12);
    .mdc-button, span.status {
      position: absolute;
      right: 1rem;
      font-size: .675rem;
      background-color: #1FEEB2;
    }
    .mdc-button{
      &:disabled{
        background-color: #999999;
        color:#fff;
      }
    }
    span.status {
      padding: 1rem;
      text-transform: uppercase;
    }
  }

  .mdc-list-item__graphic {
    overflow: hidden;
    width: $size;
    height: $size;
    border-radius: .5rem;
    margin-right: 1rem;
  }

  .mdc-list-item__text, .mdc-list-item__secondary-text {
    text-transform: capitalize;
  }

  .mdc-list-item__secondary-text {
    font-size: .75rem;
  }

  .mdc-dialog__header__title {
    color: #fff;
  }

  .mdc-snackbar {
    z-index: 9;
  }
</style>
