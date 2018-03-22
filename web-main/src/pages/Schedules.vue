<template lang="pug">
  #schedules.mdc-typography
    header1
    spinner(v-if="!classes")
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
              span.mdc-list-item__secondary-text.tutor(v-if="!v._embedded.last_session.items.length") Tutor : {{v._embedded.tutor.name}}
              span.mdc-list-item__secondary-text.tutor(v-if="!!v._embedded.last_session.items.length") Class started by {{parseLastSessionTutorName(v._embedded.last_session.items)}}
            button(v-if="buttonStatus(v) === 'start'" @click='start($event,v.id,ii,i)'  data-mdc-auto-init="MDCRipple").mdc-button.mdc-button--raised.mdc-button--compact Start
            button(v-if="buttonStatus(v) === 'start-ongoing'" @click='start($event,v.id,ii,i)'  data-mdc-auto-init="MDCRipple").mdc-button.mdc-button--raised.mdc-button--compact ongoing
            button(v-if="buttonStatus(v) === 'start-late-ongoing'" @click='start($event,v.id,ii,i)'  data-mdc-auto-init="MDCRipple").mdc-button.mdc-button--raised.mdc-button--compact activated
            button(v-if="buttonStatus(v) === 'disabled'" disabled @click='start($event,v.id,ii,i)' data-mdc-auto-init="MDCRipple").mdc-button.mdc-button--raised.mdc-button--compact Start
            button(v-if="buttonStatus(v) === 'late'" @click='start($event,v.id,ii,i)' data-mdc-auto-init="MDCRipple").mdc-button.mdc-button--raised.mdc-button--compact Activate
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
  import moment from 'moment';
  import axios from 'axios';
  import sharedHal from '../mixins/hal';

  export default {
    name: 'schedules',
    mixins: [sharedHal],
    components: {
      'spinner': () => import('@/components/Spinner'),
      'tab-bottom': () => import('@/components/TabBottom'),
      'header1': () => import('@/components/Header'),
    },
    data() {
      return {
        msg: 'Welcome to Your Vue.js PWA',
        classes: null,
        dialog: null,
        snackbar: null,
        thisClass: {id: 0},
        currentAuth: null,
      }
    },
    methods: {
      buttonStatus(class_) {
        const msts = moment(class_.start_at_ts);
        const mfts = moment(class_.finish_at_ts);
        const mnow = moment();
        let status = 'disabled';
        let ls = class_._embedded.last_session;

        if (!!ls && !!ls.items.length && !!ls.items[0].created_at) {
          ls = ls.items;
          const mls = moment(ls[0].created_at);
//          console.log(mls.toISOString(), mfts.toISOString(), mnow.isAfter(mls), mls.isBefore(mfts), mls.isAfter(mfts));
//          console.log(this.currentAuth,ls[0]._embedded.tutor);
          if (!!this.currentAuth && !!ls[0]._embedded.tutor
            && !ls.filter(v => v._embedded.tutor.id === this.currentAuth.id).length >= 1) {
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

        return status;
      },
      activate(cid) {
        const url = `${process.env.API}/classes/${cid}/sessions`;

        axios.post(url, {
          id: this.currentAuth.id
        })
          .then(response => {
            this.getSchedules();
          })
          .catch(error => console.log(error))
      },
      start(e, id, ii, i) {
        this.thisClass = this.classes[ii].items[i];

        this.dialog.lastFocusedTarget = e.target;
        this.dialog.show();
      },
      parseLastSessionTutorName(array) {
        return array.map(v => v['_embedded']['tutor'] ? v['_embedded']['tutor']['name'] : "").join(", ");
      },
      getSchedules(page = 1) {
        const url = `${process.env.API}/classes/group/date`;

        axios.get(url, {
          params: {
            'sort': 'start_at_ts.asc',
          },
        })
          .then(response => {
            let classes = response.data._embedded.items;
            classes = classes.map(v => {
              v.items = v.items.map(v2 => {
                v2._embedded = {
                  module: {image: "data:image/gif;base64,R0lGODdhAQABAPAAAMPDwwAAACwAAAAAAQABAAACAkQBADs="},
                  branch: {name: "..."},
                  tutor: {name: "..."},
                  last_session: {
                    items: [
                      {
                        _embedded: {
                          tutor: {name: "..."}
                        }
                      }]
                  },
                };
                return v2
              });
              return v
            });
            this.classes = classes;
            let j = 0;
            const d = 200;

            this.classes.forEach((v, i, a) => {
              v.items.forEach((v2, i2, a2) => {
                j++;
                setTimeout(() => this.parseEmbedded('module', v2._links.module.href, a2[i2]['_embedded'], item => {
                  item.image = item.image ? item.image : "data:image/gif;base64,R0lGODdhAQABAPAAAMPDwwAAACwAAAAAAQABAAACAkQBADs=";
                  if (item.image.indexOf('data:image/gif') < 0) {
                    let image = item.image;
                    image = image.replace('https://', '').replace('http://', '');
                    image = `//images.weserv.nl/?output=jpg&il&q=100&w=96&h=96&t=square&url=${image}`;
                    item.image = image;
                  }
                  return item;
                }), j * d);
                setTimeout(() => this.parseEmbedded('branch', v2._links.branch.href, a2[i2]['_embedded']), j * d);
                setTimeout(() => this.parseEmbedded('tutor', v2._links.tutor.href, a2[i2]['_embedded']), j * d);

                j++;
                setTimeout(() => this.parseEmbedded('last_session', v2._links.last_session.href, a2[i2]['_embedded'], item => {
                  item.items.forEach((v3, i3, a3) => {
                    if (!a3[i3]['_embedded']) {
                      this.$set(a3[i3], '_embedded', {});
                    }
                    this.parseEmbedded('tutor', v3._links.tutor.href, a3[i3]['_embedded']);
                  });
                  return item
                }), j * d);

              })
            });
//            console.log(this.classes)
          })
          .catch(error => console.log(error))
      },
    },
    mounted() {
      const cls = this;
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

      window.mdc.autoInit();
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
        text-transform: capitalize;
      }
    }
    .text {
      float: right;
      font-size: .75rem;
      text-transform: capitalize;
      color: #6FCF97;
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

  .mdc-list-item__text {
    text-transform: uppercase;
  }

  .mdc-list-item__secondary-text {
    text-transform: capitalize;
    font-size: .75rem;
    &.tutor {
      color: #BB6BD9;
      font-weight: bold;
    }
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
