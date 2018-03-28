export default {
  /**
   * set image default
   *
   * @returns {{imageDefault: string}}
   */
  data(){
    return {
      imageDefault:"data:image/gif;base64,R0lGODdhAQABAPAAAMPDwwAAACwAAAAAAQABAAACAkQBADs="
    }
  },
  methods: {
    /**
     * if image is empty = set default image above
     * if image is set = modify image to use images.weserv.nl
     *
     * @param image image url
     * @param setting images.weserv.nl configuration
     */
    normalizeImage(image, settings="output=jpg&il&q=100&w=96&h=96&t=square") {
      image = !!image ? image : "data:image/gif;base64,R0lGODdhAQABAPAAAMPDwwAAACwAAAAAAQABAAACAkQBADs=";
      if (image.indexOf('data:image/gif') < 0) {
        image = image.replace('https://', '').replace('http://', '');
        image = `//images.weserv.nl/?${settings}&url=${image}`;
      }
      return image;
    },
  }
}
