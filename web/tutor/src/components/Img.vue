<template lang="pug">
  transition(enter-active-class="animated fadeIn")
    img(:src="newSrc" :key="newSrc")
</template>

<script>
  const imgDefault = 'data:image/gif;base64,R0lGODdhAQABAPAAAMPDwwAAACwAAAAAAQABAAACAkQBADs=';
  const tutorDefault = 'https://dl.dropboxusercontent.com/s/xdy592qcgjvbbnp/Teacher.png';
  const moduleOpts = 'il&q=100&w=96&h=96&t=square';
  const profileOpts = 'il&q=100&w=64&h=64&t=square&shape=circle';
  export default {
    props: ['src', 'myIs'],
    computed: {
      opts() {
        let opts;
        switch (this.myIs) {
          case 'profile':
            opts = profileOpts;
            break;
          case 'module':
            opts = moduleOpts;
            break;
        }
        return opts;
      },
      newSrcType() {
        let src;
        switch (this.myIs) {
          case 'tutor':
            src = tutorDefault;
            break;
          default:
            src = this.src
        }
        return src;
      },
      newSrc() {
        let newSrc = !!this.newSrcType ? this.newSrcType : imgDefault;
        if (newSrc.indexOf('data:image/') < 0) {
          newSrc = newSrc.replace('https://', '').replace('http://', '');
          newSrc = `//images.weserv.nl/?${this.opts}&url=${newSrc}`;
        }
        return newSrc;
      }
    },
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
</style>
