<template>
  <div id="search_content">
    <router-view></router-view>
    <div>
      <AfHeader></AfHeader>
    </div>
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
  },
  data () {
    return {
      msg: 'Vue 태스트'
    }
  }
}
</script>

<style>
#app {

}
</style>
