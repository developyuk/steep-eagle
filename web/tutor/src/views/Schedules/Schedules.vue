<template lang="pug">
  template-main#schedules
    .mdc-list-group
      template(v-for="(v,i) in classes")
        h3.mdc-list-group__subheader
          span.date-day
            .date
              placeholder(:value="v.dateDay")
            .day
              placeholder(:value="v.day")
          span.text
            placeholder(:value="v.text")

        transition-group(tag="ul" leave-active-class="animated slideOutLeft").mdc-list
          item(v-for="(vv,ii) in v._items" :item="vv" :key="`${i}-${ii}`")

    form(@submit.prevent="checkPin($event)")
      my-dialog(@mounted="onDialogMounted")
        span Insert 1234 to activate {{currentStartedClass._embedded.module.name.toUpperCase()}}.
        p
        input(type="text" name="username" v-model.trim="pin")
        .errMsg(v-if="errMsg") {{errMsg}}
        template(slot="footer")
          button.mdc-button.mdc-dialog__footer__button.mdc-dialog__footer__button--cancel(type='button') No
          button.mdc-button.mdc-dialog__footer__button(type='submit') Yes

    .mdc-snackbar(aria-live='assertive' aria-atomic='true' aria-hidden='true')
      .mdc-snackbar__text
      .mdc-snackbar__action-wrapper
        button.mdc-snackbar__action-button(type='button')
</template>

<script>
  import axios from 'axios';
  import mixinHal from '../../mixins/hal';
  import mixinDom from '../../mixins/dom';
  import _findIndex from 'lodash/findIndex';
  import {mapState, mapMutations} from 'vuex';
  import mqtt from "mqtt";
  import {MDCDialog} from '@material/dialog';
  import {MDCSnackbar} from '@material/snackbar';
  import {MDCRipple} from '@material/ripple';
  import TemplateMain from '@/components/views/Main';
  import MyDialog from '@/components/Dialog';

  const defaultClass =
    [1, 2].map(v => {
      const items = [1, 2, 3].map(vv => {
        return {
          "start_at": "",
          "program_module": {
            "module": {
              "image": "",
              "name": "",
            },
          },
          "finish_at_ts": "Sun, 08 Jul 2018 03:30:00 GMT",
          "start_at_ts": "Sun, 08 Jul 2018 02:00:00 GMT",
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
          "last_session": [],
        }
      });
      return {
        "date": "2018-07-07",
        "text": "",
        "dateDay": null,
        "day": "",
        "_items": items
      }
    });
  export default {
    mixins: [mixinHal, mixinDom],
    components: {
      TemplateMain,
      MyDialog,
      'item': () => import('./Item'),
      'placeholder': () => import('@/components/Placeholder'),
    },
    data() {
      return {
        pin: null,
        classes: defaultClass,
        currentStartedClass: {
          id: 0,
          _embedded: {
            module: {name: "..."}
          }
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
      ...mapMutations(['nextMqtt', 'nextDialog', 'nextSearch']),
      onDialogMounted(e) {
        this.dialog = e;
      },
      _parseClass(isToday, _class) {
        const timeout = 1;

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
          const url = `${process.env.VUE_APP_DBAPI}/classes/${this.currentStartedClass.id}/sessions`;

          axios.post(url, {
            id: this.currentAuth.id
          })
            .then(response => {
              this.pin = null;

              this.mqtt
                .publish('schedules', JSON.stringify({
                  on: "startYes",
                  by: this.currentAuth,
                  id: this.currentStartedClass.id,
                  sid: response.data.id
                }));
            })
            .catch(error => console.log(error));

          this.dialog.close();
        } else {
          this.errMsg = "invalid. Check pin again!";
        }

      },
      getSchedules(page = 1) {
        const url = `${process.env.VUE_APP_DBAPI}/schedules`;
        const params = {};
        if (!!this.q) {
          params['q'] = this.q;
        }

        axios.get(url, {params})
          .then(response => {
            this.classes = Object.assign({}, this.classes, response.data._items);
            console.log(this.classes)

//            this.classes.forEach((v, i, a) => {
//              v.items.forEach((v2, i2, a2) => this._parseClass(v.text === 'today', a2[i2]))
//            });
          })
          .catch(error => console.log(error))
      },
      findClassById(id) {
        let i, ii;
        i = _findIndex(this.classes, v => {
          ii = _findIndex(v.items, {id});
          return ii > -1;
        });
//        console.log(i, ii, this.classes);
        return {i, ii}
      }
    },
    mounted() {
//      new MDCRipple(this.$el.querySelector('.mdc-button'));
      this.dialog = new MDCDialog(this.$el.querySelector('#my-mdc-dialog'));
      this.nextDialog(this.dialog);
      this.snackbar = new MDCSnackbar(this.$el.querySelector('.mdc-snackbar'));
      this.getSchedules();

      this.mqtt = mqtt.connect(process.env.VUE_APP_WS);

      this.mqtt
        .on('connect', () => {
          const topic = 'schedules';
          this.nextMqtt({topic, mqtt: this.mqtt});
          this.mqtt.subscribe(topic);
        })
        .on('message', (topic, message) => {
          console.log(topic, message.toString());
          const parsedMessage = JSON.parse(message.toString());
          const {id: msgId, on: msgOn} = parsedMessage;

          if (msgOn === 'start') {
            const {i, ii} = this.findClassById(msgId);
            this.currentStartedClass = this.classes[i].items[ii];
          }
          if (msgOn === 'startYes') {
            const {i, ii} = this.findClassById(msgId);
            this.currentStartedClass = this.classes[i].items[ii];

            setTimeout(() => this._parseClass(true, this.classes[i].items[ii]), 200);

            let snackbarOpts = {
              message: `Start ${this.currentStartedClass._embedded.module.name.toUpperCase()}`
            };
            const {by: msgBy, sid: MsgSid} = parsedMessage;
            if (msgBy.id === this.currentAuth.id) {
//              console.log(MsgSid);
              snackbarOpts = Object.assign(snackbarOpts, {
                actionText: 'Undo',
                actionHandler: () => {
                  const url = `${process.env.VUE_APP_DBAPI}/sessions/${MsgSid}`;

                  axios.delete(url)
                    .then(response => {
                      this.mqtt
                        .publish('schedules', JSON.stringify({
                          on: "undo",
                          id: this.currentStartedClass.id,
                        }));
                    })
                    .catch(error => console.log(error));
                }
              })
            }
            this.snackbar.show(snackbarOpts);
          }
          if (msgOn === 'undo') {
            const {i, ii} = this.findClassById(msgId);
            this.currentStartedClass = this.classes[i].items[ii];

            setTimeout(() => this._parseClass(true, this.classes[i].items[ii]), 200);
            let snackbarOpts = {
              message: `Undo ${this.classes[i].items[ii]._embedded.module.name.toUpperCase()}`
            };
            this.snackbar.show(snackbarOpts);
          }
        });

    },
    beforeDestroy() {
//      console.log('beforeDestroy');

      this.mqtt.end();
      this.nextMqtt(null);
      this.nextSearch(null);
      this.nextDialog(null);
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
