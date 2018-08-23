<template lang="pug">
  li.mdc-list-item
    .mdc-list-item__graphic
      my-img(:src="item.module.image" myIs="module")
    span.mdc-list-item__text
      strong: placeholder(:value="item.module.name")
      placeholder(:value="item.branch.name").mdc-list-item__secondary-text
      span.mdc-list-item__secondary-text
        placeholder(:value="item.start_at")
        | &nbsp;-&nbsp;
        placeholder(:value="item.finish_at")
      span.mdc-list-item__secondary-text.tutor(v-if="!item.last_sessions || (!!item.last_sessions && !item.last_sessions.length)") Tutor:&nbsp;
        placeholder(:value="item.tutor.name")
      span.mdc-list-item__secondary-text.tutor(v-if="!!item.last_sessions && !!item.last_sessions.length") Class started by&nbsp;
        placeholder(:value="parseLastSessionTutorName(item.last_sessions)")
    button-status(:class_="item" @click-start="onClickStart")
</template>

<script>
export default {
  props: ["item"],
  components: {
    "button-status": () => import("./ButtonStatus"),
    "my-img": () => import("@/components/Img"),
    placeholder: () => import("@/components/Placeholder")
  },
  methods: {
    onClickStart(e) {
      this.$emit("click-start", e);
    },
    parseLastSessionTutorName(array) {
      return array
        .map(
          v =>
            !!v["session_tutors"].length
              ? v["session_tutors"].map(w => w["tutor"]["profile"]["name"])
              : ""
        )
        .join(", ");
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
@import "../../assets/shared";

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
    height: 4rem;
    border-radius: 0.5rem;
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
  &.tutor {
    color: map-get($palettes, purple);
    font-weight: bold;
  }
}
</style>
