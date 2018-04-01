<template lang="pug">
  #buttonStatus
    button(v-if="status === 'start'" @click='start($event)' data-mdc-auto-init="MDCRipple").mdc-button.mdc-button--raised.mdc-button--compact Start
    button(v-if="status === 'start-ongoing'" @click='start($event)' data-mdc-auto-init="MDCRipple").mdc-button.mdc-button--raised.mdc-button--compact Ongoing
    button(v-if="status === 'start-late-ongoing'" @click='start($event)' data-mdc-auto-init="MDCRipple").mdc-button.mdc-button--raised.mdc-button--compact Activated
    button(v-if="status === 'disabled'" disabled @click='start($event)' data-mdc-auto-init="MDCRipple").mdc-button.mdc-button--raised.mdc-button--compact Start
    button(v-if="status === 'late'" @click='start($event)' data-mdc-auto-init="MDCRipple").mdc-button.mdc-button--raised.mdc-button--compact Activate
    span.ongoing(v-if="status ==='ongoing'") Ongoing
    span.late-ongoing(v-if="status ==='late-ongoing'") Activated
</template>

<script>
  import moment from 'moment';

  export default {
    name: 'buttonStatus',
    props: ['class_', 'index'],
    data() {
      return {
//        status: null,
        currentAuth: null,
      }
    },
//    watch: {
//      class_(val) {
//        this.calculateStatus(val);
//      }
//    },
    computed: {
      status() {
        const class_ = this.class_;
        const msts = moment(class_.start_at_ts);
        const mfts = moment(class_.finish_at_ts);
        const mnow = moment();
        let status = 'disabled';
        const ls = class_._embedded.last_session;

        if (!!ls && !!ls.items.length && !!ls.items[0].created_at
          && msts.diff(mnow, 'days') < 1) {
          const lsItems = ls.items;
          const mls = moment(lsItems[0].created_at);

          if (!!this.currentAuth && !!lsItems[0]._embedded.tutor && !!lsItems[0]._embedded.tutor.id
            && !lsItems.filter(v => {

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

          if (msts.diff(mnow, 'minutes') < 5 && mfts.diff(mnow, 'minutes') > 0) {
            status = 'start';
          }
          if (mnow.isAfter(mfts)) {
            status = 'late';
          }
        }

        return status;
      }
    },
    methods: {
      activate(cid) {
      },
      start(e) {
        this.$bus.$emit('onStartClass', {e, index: this.index, class_: this.class_,});
      }
    },
    mounted() {
      this.$bus.$on('currentAuth', auth => {
        if (!!this.currentAuth) {
          return;
        }
        this.currentAuth = auth;
      });

      window.mdc.autoInit();
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
