<template>
  <div class="test">
   {{ msg }}
   <p>이게 뭐람</p>
  </div>
</template>

<script>
export default {
  name: 'Test',
  data () {
    return {
      msg: 'TEST'
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.test {
  font-size : 30x;
}
</style>
