<template lang="pug">
  li.mdc-list-item
    .mdc-list-item__graphic
      my-img(:src="item.class.module.image")
    span.mdc-list-item__text
      span.mdc-list-item__primary-text
        strong: placeholder(:value="item.class.module.name")
      placeholder(:value="item.class.branch.name").mdc-list-item__secondary-text
      span.mdc-list-item__secondary-text
        placeholder(:value="nextStartTime")
        | &nbsp;-&nbsp;
        placeholder(:value="nextFinishTime")
      span.mdc-list-item__secondary-text.tutor(v-if="!lastAttendances._items.length") Tutor:&nbsp;
        placeholder(:value="item.class.tutor.name || item.class.tutor.username")
      span.mdc-list-item__secondary-text.tutor(v-if="!!lastAttendances._items.length") Class started by&nbsp;
        placeholder(:value="parseTutorName(lastAttendances._items)")
    button-status(:item="item" v-model="currentItem" :lastAttendances="lastAttendances")
</template>

<script>
import ButtonStatus from "./ButtonStatus";
import MyImg from "@/components/Img";
import Placeholder from "@/components/Placeholder";
import axios from "axios";
import moment from "moment";

const defaultLastAttendance = { _items: [] };

export default {
  props: ["item"],
  components: {
    ButtonStatus,
    MyImg,
    Placeholder
  },
  data() {
    return {
      currentItem: null
    };
  },
  watch: {
    currentItem(v, ov) {
      this.$emit("input", v);
    }
  },
  methods: {
    parseTutorName(list) {
      return list.length
        ? list.map(v => v.tutor.name || v.tutor.username).join(", ")
        : "";
    }
  },
  computed: {
    nextStartTime() {
      if (!this.item.nextStart) {
        return "00:00";
      }
      return moment(this.item.nextStart).format("HH:mm");
    },
    nextFinishTime() {
      if (!this.item.nextFinish) {
        return "00:00";
      }
      return moment(this.item.nextFinish).format("HH:mm");
    }
  },
  asyncComputed: {
    lastAttendances: {
      get() {
        if (!this.item.class._id) {
          return defaultLastAttendance;
        }
        const config = {
          params: {
            embedded: JSON.stringify({
              tutor: 1
            })
          }
        };
        return axios.get(
          `${process.env.VUE_APP_API}/schedules/${
            this.item.class._id
          }/last_attendances`,
          config
        );
      },

      watch: "item",
      default: defaultLastAttendance
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
@import "src/assets/shared";

$size: 3rem;
.mdc-list-item {
  height: $size + 3rem;
  border-bottom: thin solid rgba(0, 0, 0, 0.12);
  /*box-sizing: content-box;*/
  padding: 0;
  font-size: 0.75rem;
}

.mdc-list-item__graphic {
  margin-left: 1rem;
  width: 64px;
  height: 64px;

  margin-right: 1rem;
  img {
    width: 4rem;
    // height: 4rem;
    border-radius: 0.5rem;
    background-color: map-get($palettes, gainsboro);
  }
}

.mdc-list-item__text {
  text-transform: uppercase;
  strong {
    color: #333333;
    font-weight: 500;
  }
}

.mdc-list-item__secondary-text {
  text-transform: capitalize;
  font-size: 0.65rem;
  &:before {
    height: 0.875rem;
  }
  &.tutor {
    color: map-get($palettes, purple);
    font-weight: bold;
  }
}
.mdc-list-item__primary-text::before {
  height: 0;
}
</style>
