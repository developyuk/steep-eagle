import axios from 'axios';

export default {
  methods: {
    /**
     * embed neccesary link
     *
     * @param name key name
     * @param url url to get data
     * @param array will be modified array
     * @param onSuccesss callback before modify array
     */
    parseEmbedded(name, url, array, onSuccesss) {
      axios.get(`${process.env.VUE_APP_DBAPI}${url}`)
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


