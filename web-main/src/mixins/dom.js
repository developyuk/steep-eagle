export default {
  methods: {
    insertAfter(el, referenceNode) {
      referenceNode.parentNode.insertBefore(el, referenceNode.nextSibling);
    }
  }
}
