<template lang="pug">
  ul.nav.nav-mobile-menu
    li
      a(@click.prevent="onClickSignOut")
        i.ti-face-sad
        span Sign Out
  //   li
  //     a.dropdown-toggle.btn-magnify(href='javascript:void(0)', data-toggle='dropdown')
  //       i.ti-panel
  //       p Stats
  //   drop-down(tag='li', title='5 Notificaitons', icon='ti-bell')
  //     a.dropdown-toggle.btn-rotate(slot='title', href='javascript:void(0)', data-toggle='dropdown')
  //       i.ti-bell
  //       span.notification 5
  //       p.hidden-md.hidden-lg
  //         | Notifications
  //         b.caret
  //     li
  //       a(href='javascript:void(0)') Notification 1
  //     li
  //       a(href='javascript:void(0)') Notification 2
  //     li
  //       a(href='javascript:void(0)') Notification 3
  //     li
  //       a(href='javascript:void(0)') Notification 4
  //     li
  //       a(href='javascript:void(0)') Another notification
  //   li
  //     a.btn-rotate(href='javascript:void(0)')
  //       i.ti-settings
  //       p.hidden-md.hidden-lg
  //         | Settings
</template>
<script>
export default {
  methods: {
    onClickSignOut(e) {
      localStorage.removeItem("token");
      window.location.reload();
    }
  }
};
</script>
<style>
</style>
