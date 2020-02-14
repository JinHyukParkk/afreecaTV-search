<template>
  <p>{{ msg }}</p>
</template>

<script>
export default {
  name: 'Bj',
  data () {
    return {
      msg: 'Bj 페이지'
    }
  },
  created: function () {
    this.setting()
  },
  methods: {
    setting () {
      this.$store.dispatch('callChangeTab', { tab: 'bj' })
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1, h2 {
  font-weight: normal;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
