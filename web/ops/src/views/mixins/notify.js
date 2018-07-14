export default {

  data() {
    return {
      notifyOpts: {
        component: {
          template: `<span>Success create</span>`
        },
        icon: 'ti-gift',
        horizontalAlign: 'right',
        verticalAlign: 'top',
        type: 'success'
      },
    }
  },
  methods: {
    notifyVue(opts) {
      this.$notify(Object.assign(this.notifyOpts, opts))
    }
  }
}
