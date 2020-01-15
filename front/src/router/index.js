import Vue from 'vue'
import Router from 'vue-router'
import HelloWorld from '@/components/HelloWorld'
import Test from '@/components/Test'

Vue.use(Router)

export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      name: 'HelloWorld',
      component: {
        HelloWorld: HelloWorld
      }
    },
    {
      path: '/test',
      name: 'Test',
      component: {
        Test: Test
      }
    }
  ]
})
