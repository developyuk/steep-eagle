<template lang="pug">
  .rating
    label(v-for="i in [1,2,3,4,5]")
      input(type="radio" :value="i" :name="newName" @change="onChange").hide
      i.material-icons(:data-value="i")
</template>

<script>
export default {
  props: ["name"],
  computed: {
    newName() {
      return !!this.name ? this.name : "default";
    }
  },
  methods: {
    onChange(e) {
      //        console.log(this.name, e.target.value);
      this.$emit("input", e.target.value);

      const icons = e.target
        .closest(".rating")
        .querySelectorAll("i.material-icons");

      [...icons].forEach((v, i) => {
        v.classList.remove("is-active");
        if (i < e.target.value) {
          v.classList.add("is-active");
        }
      });
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
@import "src/assets/shared";

.rating {
  /*unicode-bidi: bidi-override;*/
  /*direction: rtl;*/
}

i.material-icons {
  display: inline-block;
  position: relative;
  width: 1.1em;
  color: $mdc-theme-primary;
  &.is-active {
    &:before {
      content: "star";
    }
  }
  &:before {
    content: "star_outline";
  }
  /*&:hover:before,*/
  /*&:hover ~ i:before {*/
  /*content: 'star'*/
  /*}*/
}
</style>
