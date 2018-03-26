<template lang="pug">
  #buttonStatus
    button(v-if="status === 'start'" @click='start($event,v.id,ii,i)' data-mdc-auto-init="MDCRipple").mdc-button.mdc-button--raised.mdc-button--compact Start
    button(v-if="status === 'start-ongoing'" @click='start($event,v.id,ii,i)' data-mdc-auto-init="MDCRipple").mdc-button.mdc-button--raised.mdc-button--compact ongoing
    button(v-if="status === 'start-late-ongoing'" @click='start($event,v.id,ii,i)' data-mdc-auto-init="MDCRipple").mdc-button.mdc-button--raised.mdc-button--compact activated
    button(v-if="status === 'disabled'" disabled @click='start($event,v.id,ii,i)' data-mdc-auto-init="MDCRipple").mdc-button.mdc-button--raised.mdc-button--compact Start
    button(v-if="status === 'late'" @click='start($event,v.id,ii,i)' data-mdc-auto-init="MDCRipple").mdc-button.mdc-button--raised.mdc-button--compact Activate
    span.ongoing(v-if="status ==='ongoing'") ongoing
    span.late-ongoing(v-if="status ==='late-ongoing'") activated
</template>

<script>
  import moment from 'moment';

  export default {
    name: 'buttonStatus',
    props: ['class', 'index'],
    data() {
      return {
        msg: 'Welcome to Your Vue.js PWA',
        status: null
      }
    },
//    watch: {
//      index(val) {
//
//      }
//    },
    methods: {
      activate(cid) {
        const url = `${process.env.API}/classes/${cid}/sessions`;

        axios.post(url, {
          id: this.currentAuth.id
        })
          .then(response => {
            this.thisClassSession = response.data;
            this.getSchedules();
          })
          .catch(error => console.log(error))
      },
      start(e, id, ii, i) {
        this.thisClass = this.classes[ii].items[i];

        this.dialog.lastFocusedTarget = e.target;
        this.dialog.show();
      },

    },
    mounted() {
      const class_ = this.class;
      const msts = moment(class_.start_at_ts);
      const mfts = moment(class_.finish_at_ts);
      const mnow = moment();
      let status = 'disabled';
      let ls = class_._embedded.last_session;
//        console.log(class_._embedded.module.name, msts.diff(mnow, 'days') < 1);
      if (!!ls && !!ls.items.length && !!ls.items[0].created_at
        && msts.diff(mnow, 'days') < 1) {
        ls = ls.items;
        const mls = moment(ls[0].created_at);
//          console.log(mls.toISOString(), mfts.toISOString(), mnow.isAfter(mls), mls.isBefore(mfts), mls.isAfter(mfts));
//          console.log(this.currentAuth.id,ls[0]._embedded.tutor.id);
        if (!!this.currentAuth && !!ls[0]._embedded.tutor && !!ls[0]._embedded.tutor.id
          && !ls.filter(v => {
            console.log(v._embedded.tutor.id, this.currentAuth.id);
            return v._embedded.tutor.id === this.currentAuth.id
          }).length >= 1) {
          if (mls.isBefore(mfts)) {
            status = 'start-ongoing';
          }
          if (mls.isAfter(mfts)) {
            status = 'start-late-ongoing';
          }
          return status;
        } else {
          if (mls.isBefore(mfts)) {
            status = 'ongoing';
          }
          if (mls.isAfter(mfts)) {
            status = 'late-ongoing';
          }
        }
      } else {
//        console.log(msts.diff(mnow, 'minutes') < 5, mfts.diff(mnow, 'minutes') > 0);
        if (msts.diff(mnow, 'minutes') < 5 && mfts.diff(mnow, 'minutes') > 0) {
          status = 'start';
        }
        if (mnow.isAfter(mfts)) {
          status = 'late';
        }
      }

      this.status = status;
    }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style lang="scss" scoped>
  @import "../../assets/shared";

  .mdc-button, span.ongoing, span.late-ongoing {
    position: absolute;
    right: 1rem;
    font-size: .675rem;
    background-color: map-get($palettes, green);
    width: 5rem;
    top: 1.5rem;
  }

  .mdc-button {
    &:disabled {
      background-color: #999999;
      color: #fff;
    }
  }

  span.ongoing, span.late-ongoing {
    text-transform: uppercase;
    background-color: map-get($palettes, blue);
    color: #fff;
    line-height: 2.25rem;
    text-align: center;
  }
</style>
