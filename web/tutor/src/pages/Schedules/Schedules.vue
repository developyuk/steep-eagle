<template lang="pug">
  #schedules.mdc-typography
    header1
    spinner(v-if="!classes")
    .empty(v-if="!!classes && !classes.length") class not found
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
              span.mdc-list-item__secondary-text.tutor(v-if="!v._embedded.last_session || (!!v._embedded.last_session && !v._embedded.last_session.items.length)") Tutor : {{v._embedded.tutor.name}}
              span.mdc-list-item__secondary-text.tutor(v-if="!!v._embedded.last_session && !!v._embedded.last_session.items.length") Class started by {{parseLastSessionTutorName(v._embedded.last_session.items)}}
            button-status(:class_="v" :index="`${ii}.${i}`")

    aside#my-mdc-dialog.mdc-dialog(role='alertdialog' aria-labelledby='my-mdc-dialog-label' aria-describedby='my-mdc-dialog-description')
      form(@submit.prevent="checkPin($event)").mdc-dialog__surface
        header.mdc-dialog__header
          h2#my-mdc-dialog-label.mdc-dialog__header__title Start this class?
        section#my-mdc-dialog-description.mdc-dialog__body Insert 1234 to activate {{currentStartedClass._embedded.module.name.toUpperCase()}}.
          p
          input(type="text" name="username" v-model.trim="pin")
          .errMsg(v-if="errMsg") {{errMsg}}
        footer.mdc-dialog__footer
          button.mdc-button.mdc-dialog__footer__button.mdc-dialog__footer__button--cancel(type='button') No
          button.mdc-button.mdc-dialog__footer__button(type='submit') Yes
      .mdc-dialog__backdrop

    .mdc-snackbar(aria-live='assertive' aria-atomic='true' aria-hidden='true')
      .mdc-snackbar__text
      .mdc-snackbar__action-wrapper
        button.mdc-snackbar__action-button(type='button')
    tab-bottom

</template>

<script>
  import axios from 'axios';
  import sharedHal from '../../mixins/hal';
  import sharedImage from '../../mixins/image';
  import _findIndex from 'lodash/findIndex';
  import {mapState, mapMutations} from 'vuex';
  import mqtt from "mqtt";

  export default {
    name: 'schedules',
    mixins: [sharedHal, sharedImage],
    components: {
      'spinner': () => import('@/components/Spinner'),
      'tab-bottom': () => import('@/components/TabBottom'),
      'header1': () => import('@/components/Header'),
      'button-status': () => import('./ButtonStatus'),
    },
    computed: {
      ...mapState(['currentAuth']),
    },
    data() {
      return {
        pin: null,
        classes: null,
        currentStartedClass: {
          id: 0,
          _embedded: {
            module: {name: "..."}
          }
        },
        dialog: null,
        snackbar: null,
        errMsg: null,
        ws: null
      }
    },
    methods: {
      ...mapMutations(['nextMqtt', 'nextDialog']),
      parseClass(isToday, _class) {
        const timeout = 1;
        setTimeout(() => this.parseEmbedded('module', _class._links.module.href, _class['_embedded'], item => {
          item.image = this.normalizeImage(item.image);
          return item;
        }), timeout);
        setTimeout(() => this.parseEmbedded('branch', _class._links.branch.href, _class['_embedded']), timeout);
        setTimeout(() => this.parseEmbedded('tutor', _class._links.tutor.href, _class['_embedded']), timeout);
        if (isToday) {
          setTimeout(() => this.parseEmbedded('last_session', _class._links.last_session.href, _class['_embedded'], item => {
            item.items.forEach((v, i, arr) => {
              if (!arr[i]['_embedded']) {
                this.$set(arr[i], '_embedded', {});
              }
              this.parseEmbedded('tutor', v._links.tutor.href, arr[i]['_embedded']);
            });
            return item
          }), timeout);
        }
      },
      checkPin(e) {
        if (this.pin === "1234") {
          const url = `${process.env.API}/classes/${this.currentStartedClass.id}/sessions`;

          axios.post(url, {
            id: this.currentAuth.id
          })
            .then(response => {
              this.pin = null;

              this.ws
                .publish('schedules', JSON.stringify({
                  action: "start-yes",
                  by: this.currentAuth,
                  class_: this.currentStartedClass,
                  session: response.data
                }));
            })
            .catch(error => console.log(error));

          this.dialog.close();
        } else {
          this.errMsg = "invalid. Check pin again!";
        }

      },
      parseLastSessionTutorName(array) {
//        console.log(array);
        return array.map(v => v['_embedded']['tutor'] ? v['_embedded']['tutor']['name'] : "").join(", ");
      },
      getSchedules(page = 1) {
        const url = `${process.env.API}/classes/group/date`;
        const params = {
          'sort': 'start_at_ts.asc',
        };
        if (!!this.q) {
          params['q'] = this.q;
        }

        axios.get(url, {params})
          .then(response => {
            let classes = response.data._embedded.items;
            if (!!classes) {
              classes = classes.map(v => {
                v.items = v.items.map(v2 => {
                  v2._embedded = {
                    module: {
                      name: "...",
                      image: "data:image/gif;base64,R0lGODdhAQABAPAAAMPDwwAAACwAAAAAAQABAAACAkQBADs="
                    },
                    branch: {name: "..."},
                    tutor: {name: "..."},
                  };
                  return v2
                });
                return v
              });
            } else {
              classes = []
            }
            this.classes = classes;

            this.classes.forEach((v, i, a) => {
              v.items.forEach((v2, i2, a2) => this.parseClass(v.text === 'today', a2[i2]))
            });
          })
          .catch(error => console.log(error))
      },
    },
    mounted() {
      this.dialog = mdc.dialog.MDCDialog.attachTo(document.querySelector('#my-mdc-dialog'));
      this.nextDialog(this.dialog);
      this.snackbar = mdc.snackbar.MDCSnackbar.attachTo(document.querySelector('.mdc-snackbar'));
      this.getSchedules();

      this.$bus.$on('onKeyupSearch', q => {
        this.q = `ilike.*${q}*`;
        this.getSchedules();
      });

      const _this = this;
      this.ws = mqtt.connect(process.env.WS);

      this.ws
        .on('connect', () => {
          const topic = 'schedules';
          this.nextMqtt({topic, mqtt: this.ws});
          this.ws.subscribe(topic);
        })
        .on('message', (topic, message) => {
//        console.log(topic, message.toString());
          const parsedMessage = JSON.parse(message.toString());

          console.log(parsedMessage);

          if (parsedMessage.action === 'start') {
            this.currentStartedClass = parsedMessage.class_;
          }
          if (parsedMessage.action === 'start-yes') {
            this.currentStartedClass = parsedMessage.class_;

            let i, ii;
            i = _findIndex(this.classes, v => {
              ii = _findIndex(v.items, {id: this.currentStartedClass.id,});
              return ii > -1;
            });
            console.log(i, ii, this.classes);

            setTimeout(() => this.parseClass(true, this.classes[i].items[ii]), 200);

            let snackbarOpts = {
              message: `Started class ${this.classes[i].items[ii]._embedded.module.name.toUpperCase()}`
            };
            if (parsedMessage.by.id === this.currentAuth.id) {
              console.log(parsedMessage.session);
              snackbarOpts = Object.assign(snackbarOpts, {
                actionText: 'Undo',
                actionHandler: () => {
                  const url = `${process.env.API}/sessions/${parsedMessage.session.id}`;

                  axios.delete(url)
                    .then(response => {
                      this.ws
                        .publish('schedules', JSON.stringify({
                          action: "undo",
                          class_: this.currentStartedClass,
                        }));
                    })
                    .catch(error => console.log(error));
                }
              })
            }
            this.snackbar.show(snackbarOpts);
          }
          if (parsedMessage.action === 'undo') {
            this.currentStartedClass = parsedMessage.class_;

            let i, ii;
            i = _findIndex(this.classes, v => {
              ii = _findIndex(v.items, {id: this.currentStartedClass.id,});
              return ii > -1;
            });
            console.log(i, ii, this.classes);

            setTimeout(() => this.parseClass(true, this.classes[i].items[ii]), 200);
            let snackbarOpts = {
              message: `Undo class ${this.classes[i].items[ii]._embedded.module.name.toUpperCase()}`
            };
            this.snackbar.show(snackbarOpts);
          }
        });

      window.mdc.autoInit();
    },
    beforeDestroy() {
      console.log('beforeDestroy');

      this.nextMqtt(null);
      this.ws.end()
    }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style lang="scss" scoped>
  @import "../../assets/shared";

  #schedules {
    position: relative;
    height: 100vh;
  }

  .mdc-list {
    padding: 0;
  }

  .empty {
    text-align: center;
    position: absolute;
    top: 50%;
    transform: translateY(-70%);
    width: 100%;
  }

  .errMsg {
    padding: .5rem 0;
    background-color: #fff;
    margin: .5rem 0;
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
      color: map-get($palettes, green-darken1);
    }
  }

  $size: 3rem;
  .mdc-list-item {
    height: $size+3rem;
    border-bottom: thin solid rgba(0, 0, 0, .12);
    /*box-sizing: content-box;*/
    padding: 0;
  }

  .mdc-list-item__graphic {
    margin-left: 1rem;
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
      color: map-get($palettes, purple);
      font-weight: bold;
    }
  }

  .mdc-dialog__header__title {
    color: #fff;
  }

  .mdc-dialog__header {
    background-color: $mdc-theme-primary;
  }

</style>
