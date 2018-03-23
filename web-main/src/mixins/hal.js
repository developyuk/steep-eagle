import axios from 'axios';

export default {
  methods: {
    parseEmbedded(name, url, array, onSuccesss) {
      axios.get(`${process.env.API}${url}`)
        .then(response => {
          let v = response.data._embedded ? response.data._embedded : response.data;
          if (!!onSuccesss) {
            v = onSuccesss(v, array);
          }
          this.$set(array, name, v);
        })
        .catch(error => console.log(error))
    }
  }
}


