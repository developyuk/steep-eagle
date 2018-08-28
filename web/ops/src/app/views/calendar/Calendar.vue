<template lang="pug">
  .row
    .col-md-12
      h4.title Calendar
    .col-md-12.card
      .card-header
        .buttons
          .row
            .col-sm-6
              .btn-group
                router-link(to="/schedules").btn.btn-primary.btn-icon
                  i.fa.fa-list-ul
                router-link(to="/schedules/calendar").btn.btn-primary.btn-icon.btn-fill
                  i.fa.fa-th-large
            .col-sm-6.text-right
              router-link(to="/schedules/create").btn.btn-primary.btn-fill
                span.btn-label: i.fa.fa-plus
                span Create
      .card-content.row
        .container-fluid
          .card.card-calendar
            .card-content
              #fullCalendar
</template>
<script>
import swal from "sweetalert2";
import $ from "jquery";
import "fullcalendar";
import axios from "axios";

var today = new Date();
var y = today.getFullYear();
var m = today.getMonth();
var d = today.getDate();
export default {
  data() {
    return {
      events: []
      // events: [
      //   {
      //     title: "All Day Event",
      //     start: new Date(y, m, 1),
      //     className: "event-default"
      //   },
      //   {
      //     id: 999,
      //     title: "Repeating Event",
      //     start: new Date(y, m, d - 4, 6, 0),
      //     allDay: false,
      //     className: "event-rose"
      //   },
      //   {
      //     id: 999,
      //     title: "Repeating Event",
      //     start: new Date(y, m, d + 3, 6, 0),
      //     allDay: false,
      //     className: "event-rose"
      //   },
      //   {
      //     title: "Meeting",
      //     start: new Date(y, m, d - 1, 10, 30),
      //     allDay: false,
      //     className: "event-green"
      //   },
      //   {
      //     title: "Lunch",
      //     start: new Date(y, m, d + 7, 12, 0),
      //     end: new Date(y, m, d + 7, 14, 0),
      //     allDay: false,
      //     className: "event-red"
      //   },
      //   {
      //     title: "Md-pro Launch",
      //     start: new Date(y, m, d - 2, 12, 0),
      //     allDay: true,
      //     className: "event-azure"
      //   },
      //   {
      //     title: "Birthday Party",
      //     start: new Date(y, m, d + 1, 19, 0),
      //     end: new Date(y, m, d + 1, 22, 30),
      //     allDay: false,
      //     className: "event-azure"
      //   },
      //   {
      //     title: "Click for Creative Tim",
      //     start: new Date(y, m, 21),
      //     end: new Date(y, m, 22),
      //     url: "http://www.creative-tim.com/",
      //     className: "event-orange"
      //   },
      //   {
      //     title: "Click for Google",
      //     start: new Date(y, m, 21),
      //     end: new Date(y, m, 22),
      //     url: "http://www.creative-tim.com/",
      //     className: "event-orange"
      //   }
      // ]
    };
  },
  methods: {
    initCalendar($) {
      var self = this;
      var $calendar = $("#fullCalendar");
      $calendar.fullCalendar({
        header: {
          left: "title",
          center: "month,agendaWeek,agendaDay",
          right: "prev,next,today"
        },
        defaultDate: today,
        selectable: true,
        selectHelper: true,
        views: {
          month: {
            // name of view
            titleFormat: "MMMM YYYY"
            // other view-specific options here
          },
          week: {
            titleFormat: " MMMM D YYYY"
          },
          day: {
            titleFormat: "D MMM, YYYY"
          }
        },
        select: function(start, end) {
          // on select we show the Sweet Alert modal with an input
          swal({
            title: "Create an Event",
            html:
              '<div class="form-group">' +
              '<input class="form-control" placeholder="Event Title" id="input-field">' +
              "</div>",
            showCancelButton: true,
            confirmButtonClass: "btn btn-success",
            cancelButtonClass: "btn btn-danger",
            buttonsStyling: false
          }).then(function(result) {
            var eventData;
            var eventTitle = $("#input-field").val();
            if (eventTitle) {
              eventData = {
                title: eventTitle,
                start: start,
                end: end
              };
              $calendar.fullCalendar("renderEvent", eventData, true); // stick? = true
            }
            $calendar.fullCalendar("unselect");
          });
        },
        editable: true,
        eventLimit: true, // allow "more" link when too many events

        // color classes: [ event-blue | event-azure | event-green | event-orange | event-red ]
        events: self.events
      });
    }
  },
  mounted() {
    window.$ = window.jQuery = $;
    axios
      .get(`${process.env.DBAPI}/calendar`)
      .then(response => {
        this.events = response.data._items;
        this.initCalendar($);
      })
      .catch(error => console.log(error, error.response));
  }
};
</script>
<style>
#fullCalendar {
  min-height: 300px;
}

.el-loading-spinner .path {
  stroke: #66615b !important;
}
</style>
