import Vue from "vue"

Vue.mixin({
  methods: {
    jsonldBreadcrumbs(list) {
      let innerList = `
      {
        "@type": "ListItem",
        "position": 1,
        "item": {
          "@id": "${this.base_url()}",
          "name": "kik4.work"
        }
      }
      `
      list.forEach((item, index) => {
        innerList += `
        ,{
          "@type": "ListItem",
          "position": ${index + 2},
          "item": {
            "@id": "${this.base_url() + item.path}",
            "name": "${item.text}"
          }
        }
        `
      })
      return {
        hid: "jsonld",
        innerHTML: `{
            "@context": "http://schema.org",
            "@type": "BreadcrumbList",
            "itemListElement": [ ${innerList} ]
          }`,
        type: "application/ld+json",
      }
    },
  },
})
