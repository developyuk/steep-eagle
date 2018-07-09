.mqt
<template lang="pug">
  template-main#schedules
    .empty(v-if="!!classes && !classes.length") classes not found
    .mdc-list-group
      template(v-for="(v,i) in classes")
        h3.mdc-list-group__subheader
          span.date-day
            .date
              placeholder(:value="v.dateDay" val-empty="00")
            .day
              placeholder(:value="v.day" val-empty="lorem ipsum")
          span.text
            placeholder(:value="v.text")
        .mdc-list
          item(v-for="(vv,ii) in v._items" :item="vv" :key="`${v.date}-${vv.id}`" @click-start="onClickStart")

    form(@submit.prevent="checkPin($event)")
      my-dialog(@mounted="onMountedDialog")
        span Insert 1234 to activate&nbsp;
          placeholder(:value="currentClass.program_module.module.name.toUpperCase()" val-empty="lorem ipsum")
          | .
        p
        input(type="text" name="username" v-model.trim="pin")
        .errMsg(v-if="errMsg") {{errMsg}}
        template(slot="footer")
          button.mdc-button.mdc-dialog__footer__button.mdc-dialog__footer__button--cancel(type='button') No
          button.mdc-button.mdc-dialog__footer__button(type='submit') Yes
    snackbar(@mounted="onMountedSnackbar")
</template>

<script>
  import axios from 'axios';
  import moment from 'moment';
  import _findIndex from 'lodash/findIndex';
  import {mapState, mapMutations} from 'vuex';
  import mqtt from "mqtt";
  import {MDCDialog} from '@material/dialog';
  import {MDCRipple} from '@material/ripple';
  import TemplateMain from '@/components/views/Main';
  import MyDialog from '@/components/Dialog';

  const placeholderSchedules =
    [1, 2, 3].map(v => {
      const items = [1, 2].map(vv => {
        return {
          "id": vv,
          "start_at": "",
          "program_module": {
            "module": {
              "image": "",
              "name": "",
            },
          },
          "finish_at_ts": "",
          "start_at_ts": "",
          "finish_at": "",
          "tutor": {
            "profile": {
              "name": ""
            },
            "email": ""
          },
          "branch": {
            "name": "",
          },
        }
      });
      return {
        "date": v,
        "text": "",
        "dateDay": "",
        "day": "",
        "_items": items
      }
    });

  export default {
    components: {
      TemplateMain,
      MyDialog,
      'item': () => import('./Item'),
      'placeholder': () => import('@/components/Placeholder'),
      'snackbar': () => import('@/components/Snackbar'),
    },
    data() {
      return {
        pin: null,
        classes: placeholderSchedules,
        currentClass: {
          id: 0,
          program_module: {
            module: {name: ""}
          },
        },
        dialog: null,
        snackbar: null,
        errMsg: null,
        mqtt: null
      }
    },
    computed: {
      ...mapState(['currentAuth', 'currentSearch']),
    },
    watch: {
      currentSearch(v) {
        this.q = `ilike.*${v}*`;
        this.getSchedules();
      }
    },
    methods: {
      ...mapMutations(['nextSearch']),
      onMountedSnackbar(e) {
        this.snackbar = e
      },
      onMountedDialog(e) {
        this.dialog = e;
      },
      onClickStart(e) {
        this.dialog.show();
        const {i, ii} = this.findClassById(e.id);
        this.currentClass = this.classes[i]._items[ii];
      },
      checkPin(e) {
        if (this.pin === "1234") {
          let url = `${process.env.VUE_APP_DBAPI}/sessions`;
          let params = {
            where: {
              _created: `>="${moment.utc().format('YYYY-MM-DD')}"`,
              class_: this.currentClass.id
            }
          };

          const postSessionTutor = (session) => {
            const url = `${process.env.VUE_APP_DBAPI}/sessions_tutors`;
            const params = {session: session.id, tutor: this.currentAuth.id};

            axios.post(url, params)
              .then(response => {

                this.pin = null;

                this.mqtt.m
                  .publish(this.mqtt.topic, JSON.stringify({
                    on: "startYes",
                    by: this.currentAuth,
                    id: this.currentClass.id,
                    s: {
                      id: session.id,
                      et: session._etag,
                    },
                    st: {
                      id: response.data.id,
                      et: response.data._etag,
                    },
                  }));
              })
              .catch(error => console.log(error));
          };

          axios.get(url, {params})
            .then(response => {
              if (!!response.data._items.length) {
                const session = {
                  id: response.data._items[0].id,
                  _etag: response.data._items[0]._etag
                };
                postSessionTutor(session);
              } else {
                let params = {class_: this.currentClass.id};
                axios.post(url, params)
                  .then(response => {
                    const session = {
                      id: response.data.id,
                      _etag: response.data._etag
                    };
                    postSessionTutor(session);
                  })
                  .catch(error => console.log(error));
              }
            })
            .catch(error => console.log(error));
          this.dialog.close();
        } else {
          this.errMsg = "invalid. Check pin again!";
        }

      },
      getSchedules() {
        const url = `${process.env.VUE_APP_DBAPI}/schedules`;
        const params = {};
        if (!!this.q) {
          params['q'] = this.q;
        }

        axios.get(url, {params})
          .then(response => {
            this.classes = response.data._items
          })
          .catch(error => console.log(error))
      },
      findClassById(id) {
        let i, ii;
        i = _findIndex(this.classes, v => {
          ii = _findIndex(v._items, {id});
          return ii > -1;
        });
//        console.log(i, ii, this.classes);
        return {i, ii}
      }
    },
    mounted() {
//      new MDCRipple(this.$el.querySelector('.mdc-button'));
      this.dialog = new MDCDialog(this.$el.querySelector('#my-mdc-dialog'));

      this.getSchedules();

      this.mqtt = {
        topic: 'schedules',
        m: mqtt.connect(process.env.VUE_APP_WS),
      };

      this.mqtt.m
        .on('connect', () => {
          this.mqtt.m.subscribe(this.mqtt.topic);
        })
        .on('message', (topic, message) => {
          console.log(topic, message.toString());
          const parsedMessage = JSON.parse(message.toString());
          const {id: msgId, on: msgOn} = parsedMessage;

          if (msgOn === 'startYes') {
            const {i, ii} = this.findClassById(msgId);
            this.currentClass = this.classes[i]._items[ii];

            this.getSchedules();
            let snackbarOpts = {
              message: `Start ${this.currentClass.program_module.module.name.toUpperCase()}`
            };
            const {by: msgBy, s: MsgS, st: MsgSt} = parsedMessage;
            if (msgBy.id === this.currentAuth.id) {
//              console.log(MsgSid);
              snackbarOpts = Object.assign(snackbarOpts, {
                actionText: 'Undo',
                actionHandler: () => {
                  let url = `${process.env.VUE_APP_DBAPI}/sessions_tutors/${MsgSt.id}`;

                  axios.delete(url, {headers: {'If-Match': MsgSt.et}})
                    .then(response => {
                      url = `${process.env.VUE_APP_DBAPI}/sessions/${MsgS.id}`;
                      axios.delete(url, {headers: {'If-Match': MsgS.et}})
                        .then(response => {
                          this.mqtt.m
                            .publish(this.mqtt.topic, JSON.stringify({
                              on: "undo",
                              id: this.currentClass.id,
                            }));
                        })
                        .catch(error => console.log(error));
                    })
                    .catch(error => console.log(error));
                }
              })
            }
            this.snackbar.show(snackbarOpts);
          }
          if (msgOn === 'undo') {
            const {i, ii} = this.findClassById(msgId);
            this.currentClass = this.classes[i]._items[ii];

            this.getSchedules();
            let snackbarOpts = {
              message: `Undo ${this.classes[i]._items[ii].program_module.module.name.toUpperCase()}`
            };
            this.snackbar.show(snackbarOpts);
          }
        });

    },
    beforeDestroy() {
//      console.log('beforeDestroy');

      this.mqtt.m.end();
      this.nextSearch(null);
    }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style lang="scss" scoped>
  @import "../../assets/shared";
  @import "~sass-bem";

  #schedules {
    position: relative;
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

  .image {
    width: 4rem;
    height: 4rem;
    background-color: #C3C3C3;

    @include m(construct2 construct2-part-ii dummies-construct) {
      background-color: #4F4F4F;
    }
    @include m(unity-2d weblite project-based dummies-web) {
      background-color: #F2F2F2;
    }
  }


</style>
