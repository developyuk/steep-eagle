<template lang="pug">
  #schedules.mdc-typography
    header1
    .loading(v-if="Object.keys(classes).length === 0") loading...
    .mdc-list-group
      template(v-for="(vv,ii) in classes")
        h3.mdc-list-group__subheader
          span.date-day
            .date {{vv.date}}
            .day {{vv.day}}
          span.text {{vv.text}}
        ul.mdc-list
          li.mdc-list-item(v-for="(v,i) in vv.items")
            .mdc-list-item__graphic
              img(':src'="v._embedded.module.image")
            span.mdc-list-item__text {{v._embedded.module.name}}
              span.mdc-list-item__secondary-text {{v._embedded.branch.name}}
              span.mdc-list-item__secondary-text {{v.start_at}} - {{v.finish_at}}
            button(v-if="buttonStatus(v) === 'start'" @click='start($event,v.id,ii,i)').mdc-button.mdc-button--raised.mdc-button--compact Start
            button(v-if="buttonStatus(v) === 'disabled'" disabled @click='start($event,v.id,ii,i)').mdc-button.mdc-button--raised.mdc-button--compact Start
            button(v-if="buttonStatus(v) === 'late'" @click='start($event,v.id,ii,i)').mdc-button.mdc-button--raised.mdc-button--compact Activate
            span.ongoing(v-if="buttonStatus(v)==='ongoing'") ongoing
            span.late-ongoing(v-if="buttonStatus(v)==='late-ongoing'") activated

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
//      'tab-bottom': () => import('@/components/TabBottom'),
//      'header1': () => import('@/components/Header'),
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
        const msts = moment(class_.start_at_ts);
        const mfts = moment(class_.finish_at_ts);
        let ls = class_._embedded;
        ls = ls.last_session ? ls.last_session.created_at : ls.last_session;

        const mnow = moment();
//        console.log('now', mnow.toString());
//        console.log(mts.isAfter(mnow));

        if (ls && mnow.diff(mfts, 'days') === 0) {
          const mls = moment(ls);
//          console.log('ls',mls.toString(), '-', ls);
//          console.log('mlsdiff',mls.diff(mnow,'minutes'));
          if (mnow.isAfter(mls) && mls.isBefore(mfts)) {
            return 'ongoing';
          }
          if (mnow.isAfter(mls) && mls.isAfter(mfts)) {
            return 'late-ongoing';
          }
        }
        if (mnow.isAfter(mfts)) {
          return 'late';
        }
//        console.log('diff day', mts.diff(mnow, 'days'));
        if (msts.diff(mnow, 'minutes') < 5 && mfts.diff(mnow, 'minutes') > 0) {
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
      start(e, id, ii, i) {
//        this.classes.forEach(v => {
//          console.log(v.items.filter(v2 => v2.id = id));
//        });
//        console.log(this.classes[ii],this.classes[ii].items[i]);
        this.thisClass = this.classes[ii].items[i];

        this.dialog.lastFocusedTarget = e.target;
        this.dialog.show();
      },
      getSchedules(page = 1) {
        const url = `${process.env.API}/classes/group/date?sort=start_at_ts.asc`;

        axios.get(url)
          .then(response => {
            const data = response.data._embedded;

            data.forEach((v, i, a) => {
              v.items.forEach((v2, i2, a2) => {
                let image = v2._embedded.module.image.replace('https://', '').replace('http://', '');
                image = `//images.weserv.nl/?output=jpg&il&q=100&w=96&h=96&t=square&url=${image}`;
                this.$set(a[i]['items'][i2]['_embedded']['module'], 'image', image);
              })
            });
            this.classes = data;
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
  }

  .mdc-list-group__subheader {
    height: 1.5rem;
    .date-day {
      float: left;
      * {
        text-align: center;
        line-height: 1rem;
      }
      .date {
        font-size: 1.5rem;
      }
      .day {
        font-size: .5625rem;
      }
    }
    .text {
      float: right;
      font-size: .75rem;
      text-transform: capitalize;
    }
  }

  $size: 3rem;
  .mdc-list-item {
    height: $size+2rem;
    border-bottom: thin solid rgba(0, 0, 0, .12);
    .mdc-button, span.ongoing, span.late-ongoing {
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
    span.ongoing, span.late-ongoing {
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
