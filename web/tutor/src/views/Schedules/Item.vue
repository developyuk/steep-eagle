<template lang="pug">
  li.mdc-list-item
    .mdc-list-item__graphic
      my-img(:src="item.class.module.image")
    span.mdc-list-item__text
      span.mdc-list-item__primary-text
        strong: placeholder(:value="item.class.module.name")
      placeholder(:value="item.class.branch.name").mdc-list-item__secondary-text
      span.mdc-list-item__secondary-text
        placeholder(:value="item.startTime")
        | &nbsp;-&nbsp;
        placeholder(:value="item.finishTime")
      span.mdc-list-item__secondary-text.tutor(v-if="!item.last_attendances._items || (!!item.last_attendances._items && !item.last_attendances._items.length)") Tutor:&nbsp;
        placeholder(:value="item.class.tutor.name || item.class.tutor.username")
      span.mdc-list-item__secondary-text.tutor(v-if="!!item.last_attendances._items && !!item.last_attendances._items.length") Class started by&nbsp;
        placeholder(:value="parseLastAttendanceTutorName(item.last_attendances._items)")
    button-status(:item="item" v-model="currentClass")
</template>

<script>
import ButtonStatus from "./ButtonStatus";
import MyImg from "@/components/Img";
import Placeholder from "@/components/Placeholder";

export default {
  props: ["item"],
  components: {
    ButtonStatus,
    MyImg,
    Placeholder
  },
  data() {
    return {
      currentClass: null
    };
  },
  watch: {
    currentClass(v, ov) {
      this.$emit("input", v);
    }
  },
  methods: {
    parseLastAttendanceTutorName(list) {
      return list.length ? list.map(v => v["tutor"]["name"]).join(", ") : "";
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
