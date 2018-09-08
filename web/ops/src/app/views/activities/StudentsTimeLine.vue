<template lang='pug'>
  .row
    .col-md-12
      h4.title Tutors Timeline
    .col-md-12.card
      .card-header
        .buttons
          .row
            .col-sm-6
              .btn-group
                router-link(to="/activities/tutors").btn.btn-primary.btn-icon
                  i.fa.fa-list-ul
                router-link(to="/activities/tutors/timeline").btn.btn-primary.btn-icon.btn-fill
                  i.fa.fa-th-large
      .card-content.row
        .col-md-12
          time-line
            time-line-item.timeline-inverted(badgetype='danger' badgeicon='ti-gallery')
              span.label.label-danger(slot='header') Some title
              p(slot='body')
                | Wifey made the best Father&apos;s Day meal ever. So thankful so happy so blessed. Thank you for making my family We just had fun with the &ldquo;future&rdquo; theme !!!   It was a fun night all together ... The always rude Kanye Show at 2am Sold Out Famous viewing @ Figueroa and 12th in downtown.
              h6(slot='footer')
                i.ti-time
                |           11 hours ago via Twitter
            time-line-item(badgetype='success' badgeicon='ti-check-box')
              span.label.label-success(slot='header') Another Title
              p(slot='body')
                | Thank God for the support of my wife and real friends. I also wanted to point out that it&rsquo;s the first album to go number 1 off of streaming!!! I love you Ellen and also my number one design rule of anything I do from shoes to music to homes is that Kim has to like it....
            time-line-item.timeline-inverted(badgetype='info' badgeicon='ti-credit-card')
              span.label.label-success(slot='header') Another Title
              div(slot='body')
                p
                  | Called I Miss the Old Kanye That&rsquo;s all it was Kanye And I love you like Kanye loves Kanye Famous viewing @ Figueroa and 12th in downtown LA 11:10PM
                p
                  | What if Kanye made a song about Kanye Roy&egrave;re doesn&apos;t make a Polar bear bed but the Polar bear couch is my favorite piece of furniture we own It wasn&rsquo;t any Kanyes Set on his goals Kanye
                hr
                drop-down.btn-group
                  button.btn.btn-info.dropdown-toggle(slot='title' href='#' type='button' data-toggle='dropdown')
                    i.ti-settings
                    span.caret
                  li
                    a(href='#action') Action
                  li
                    a(href='#action') Another action
                  li
                    a(href='#here') Something else here
                  li.divider
                  li
                    a(href='#link') Separated link
            time-line-item(badgetype='warning' badgeicon='ti-gallery')
              span.label.label-warning(slot='header') Another Title
              p(slot='body')
                | Tune into Big Boy&apos;s 92.3 I&apos;m about to play the first single from Cruel Winter Tune into Big Boy&apos;s 92.3 I&apos;m about to play the first single from Cruel Winter also to Kim&rsquo;s hair and makeup Lorraine jewelry and the whole style squad at Balmain and the Yeezy team. Thank you Anna for the invite thank you to the whole Vogue team
</template>
<script>
import TimeLine from "src/components/Dashboard/Views/Dashboard/Stats/TimeLine.vue";
import TimeLineItem from "src/components/Dashboard/Views/Dashboard/Stats/TimeLineItem.vue";

export default {
  components: {
    TimeLine,
    TimeLineItem
  }
};
</script>
<style scoped lang='scss'>
.col-md-12.card {
  background-color: transparent;
}
</style>
