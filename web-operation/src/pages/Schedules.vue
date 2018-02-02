<template lang="pug">
  section.row
    section.col.s2.card-panel
      nav-side
    section.col.s10
      header-menu
      section
        h1 Schedules
        #calendar
</template>

<script>
  import NavSide from '@/components/NavSide'
  import Header from '@/components/Header'
  require("fullcalendar/dist/fullcalendar.min.js");
  require("fullcalendar/dist/fullcalendar.min.css");

  export default {
    name: 'schedules',
    components: {
      'nav-side': NavSide,
      'header-menu': Header
    },
    data() {
      return {
        msg: 'Welcome to Your Vue.js PWA'
      }
    },
    mounted(){
      const $calendar = $('#calendar').fullCalendar({
        header: {
          left: 'title',
          right: 'today,month,agendaDay,agendaWeek prev,next'
        },
        // Make possible to respond to clicks and selections
        selectable: true,

        // This is the callback that will be triggered when a selection is made.
        // It gets start and end date/time as part of its arguments
        select: function(start, end, jsEvent, view) {

          // Ask for a title. If empty it will default to "New event"
          var title = prompt("Enter a title for this event", "New event");

          // If did not pressed Cancel button
          if (title != null) {
            // Create event
            var event = {
              title: title.trim() != "" ? title : "New event",
              start: start,
              end: end
            };

            // Push event into fullCalendar's array of events
            // and displays it. The last argument is the
            // "stick" value. If set to true the event
            // will "stick" even after you move to other
            // year, month, day or week.

            $calendar.fullCalendar("renderEvent", event, true);
          };
          // Whatever happens, unselect selection
          $calendar.fullCalendar("unselect");

        }, // End select callback

        // Make events editable, globally
        editable : true,

        // Callback triggered when we click on an event

        eventClick: function(event, jsEvent, view){
          // Ask for a title. If empty it will default to "New event"
          var newTitle = prompt("Enter a new title for this event", event.title);

          // If did not pressed Cancel button
          if (newTitle != null) {
            // Update event
            event.title = newTitle.trim() != "" ? newTitle : event.title;

            // Call the "updateEvent" method
            $calendar.fullCalendar("updateEvent", event);

          }
        } // End callback eventClick
      })
    }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
  section.col {
    padding: 0 !important;
  }

  section.row {
    margin: 0;
  }

  .card-panel.col {
    margin: 0;
    height: 100vh;
  }
</style>
