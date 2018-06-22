<template lang="pug">
  img(:src="newSrc")
</template>

<script>
  const imgDefault = 'data:image/gif;base64,R0lGODdhAQABAPAAAMPDwwAAACwAAAAAAQABAAACAkQBADs=';
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
      newSrc() {
        let newSrc = !!this.src ? this.src : imgDefault;
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
