<template>
  <div id="app">
    <p>Vue 테스트 서버</p>
    <AfHeader></AfHeader>
    <AfList><AfList>
    <router-view></router-view>
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
