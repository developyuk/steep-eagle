<template lang="pug">
  #schedules.mdc-typography
    header1
    ul.mdc-list
      li.mdc-list-item(v-for="(v,i) in classes")
        .mdc-list-item__graphic
          img(':src'="v._embedded.module.image")
        span.mdc-list-item__text {{v._embedded.module.name}}
          span.mdc-list-item__secondary-text {{v._embedded.branch.name}} {{v.day}} {{v.time}}
        button(v-if="buttonStatus(v) === 'start'" data-mdc-auto-init="MDCRipple" @click='start($event,v.id,i)').mdc-button.mdc-button--raised.mdc-button--compact Start
        button(v-if="buttonStatus(v) === 'disabled'" disabled data-mdc-auto-init="MDCRipple" @click='start($event,v.id,i)').mdc-button.mdc-button--raised.mdc-button--compact Start
        button(v-if="buttonStatus(v) === 'late'" data-mdc-auto-init="MDCRipple" @click='start($event,v.id,i)').mdc-button.mdc-button--raised.mdc-button--compact Activate
        span.ongoing(v-if="buttonStatus(v)==='ongoing'") ongoing

    aside#my-mdc-dialog.mdc-dialog(role='alertdialog' aria-labelledby='my-mdc-dialog-label' aria-describedby='my-mdc-dialog-description')
      .mdc-dialog__surface
        header.mdc-dialog__header
          h2#my-mdc-dialog-label.mdc-dialog__header__title
            | Start this class?
        section#my-mdc-dialog-description.mdc-dialog__body
          | Class will be start when you click yes.
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
  //  import {MDCRipple} from '@material/ripple';
  import TabBottom from '@/components/TabBottom';
  import Header from '@/components/Header';
  import moment from 'moment';
  import axios from 'axios';

  export default {
    name: 'schedules',
    components: {
      TabBottom,
      'header1': Header
    },
    data() {
      return {
        msg: 'Welcome to Your Vue.js PWA',
        classes: [],
        dialog: null,
        snackbar: null,
        thisClass: {id: 0},
        currentAuth: null,
      }
    },
    methods: {
      comingSoon(e) {
        e.target.innerText = 'coming soon';
      },
      buttonStatus(class_) {
        const ts = class_.ts;
        let ls = class_._embedded;
        ls = ls.last_session ? ls.last_session.created_at : ls.last_session;

        const mts = moment(ts);
//        console.log('ts', mts.toString(), '---', ts);
        const mnow = moment();
        const mdiff = mts.diff(mnow, 'minutes');
//        console.log('now', mnow.toString());
//        console.log(mts.isAfter(mnow));
//        console.log('diff', mdiff);
        if (ls) {
          const mls = moment(ls);
//          console.log('ls',mls.toString(), '-', ls);
//          console.log('mlsdiff',mls.diff(mnow,'minutes'));
          if (mnow.isAfter(mls) && mls.diff(mnow,'minutes') > -1 * (2 * 60)) {
            return 'ongoing';
          }
        }
        if (mdiff < -5) {
          return 'late';
        }
//        console.log('diff day', mts.diff(mnow, 'days'));
        if (mts.diff(mnow, 'days') === 0) {
          return 'start';
        }
        return 'disabled';
      },
      activate(cid) {
        const url = `${process.env.API}/classes/${cid}/sessions`;
//        console.log(cid,this.classes[this.classes.findIndex(el => el.id === cid)]);

        axios.post(url, {
          id: this.currentAuth.id
        })
          .then(response => {
//            console.log(response);
//            this.$delete(this.classes, this.classes.findIndex(el => el.id === cid));
            this.getSchedules();
          })
          .catch(error => console.log(error))
      },
      start(e, id, i) {
        this.thisClass = this.classes[i];

        this.dialog.lastFocusedTarget = e.target;
        this.dialog.show();
      },
      getSchedules(page = 1) {
        const url = `${process.env.API}/classes?sort=ts.asc`;

        axios.get(url)
          .then(response => {
            this.classes = response.data._embedded;

            this.classes.forEach((v, i, a) => {
              let image = v._embedded.module.image.replace('https://', '').replace('http://', '');
              image = `//images.weserv.nl/?output=jpg&il&q=100&w=96&h=96&t=square&url=${image}`;
              this.$set(a[i]['_embedded']['module'], 'image', image);
            });
          })
          .catch(error => console.log(error))
      },
    },
    mounted() {
      const cls = this;
      Array.from(document.querySelectorAll('.mdc-tab, .mdc-button')).forEach(v => mdc.ripple.MDCRipple.attachTo(v));
      this.dialog = mdc.dialog.MDCDialog.attachTo(document.querySelector('#my-mdc-dialog'));
      this.snackbar = mdc.snackbar.MDCSnackbar.attachTo(document.querySelector('.mdc-snackbar'));
      this.dialog.listen('MDCDialog:accept', () => {
        cls.activate(cls.thisClass.id);
        const dataObj = {
          message: `Class ${cls.thisClass.id} has been started`,
//          actionText: 'Undo',
//          actionHandler: function () {
//            console.log('my cool function');
//          }
        };
        cls.snackbar.show(dataObj);
      });
      this.getSchedules();
      this.$bus.$on('currentAuth', (auth) => {
        this.currentAuth = auth;
      });
    }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style lang="scss" scoped>
  #schedules {
    position: relative;
    height: 100vh;
  }

  .mdc-list {
    padding: 0;
    height: calc(100vh - 8rem);
    overflow-y: auto;
  }

  $size: 3rem;
  .mdc-list-item {
    height: $size+2rem;
    border-bottom: thin solid rgba(0, 0, 0, .12);
    .mdc-button, span.ongoing {
      position: absolute;
      right: 1rem;
      font-size: .675rem;
      background-color: #1FEEB2;
      width: 5rem;
    }
    .mdc-button {
      &:disabled {
        background-color: #999999;
        color: #fff;
      }
    }
    span.ongoing {
      text-transform: uppercase;
      background-color: #56CCF2;
      color: #fff;
      line-height: 2.25rem;
      text-align: center;
    }
  }

  .mdc-list-item__graphic {
    /*overflow: hidden;*/
    width: 64px;
    height: 64px;
    img {
      width: 64px;
      height: 64px;
      border-radius: .5rem;
    }

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

  .mdc-dialog__header {
    background-color: var(--mdc-theme-primary);
  }

  .mdc-snackbar {
    z-index: 9;
  }
</style>
