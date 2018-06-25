export default {
  methods: {
    insertAfter(el, referenceNode) {
      referenceNode.parentNode.insertBefore(el, referenceNode.nextSibling);
    },
    slug(text) {

      if (!!text) {
        return text
          .toLowerCase()
          .replace(/ /g, '-')
          .replace(/[^\w-]+/g, '')
      }
      return text;
    },
  }
}
