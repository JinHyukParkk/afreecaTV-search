<template>
  <div id="search_content">
    <AfHeader></AfHeader>
    <AfList><AfList>
  </div>
</template>

<script>
import AfHeader from '@/components/AfHeader'
import AfList from '@/components/AfList'

export default {
  name: 'App',
  components: {
    'AfHeader': AfHeader,
    'AfList': AfList
  }
}
</script>

<style>
#app {

}
</style>
