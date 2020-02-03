import Vue from 'vue'
import Router from 'vue-router'
import HelloWorld from '@/components/HelloWorld'

Vue.use(Router)

const NotFound = {template: '<div>Not Found</div>'}

export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/:keyword',
      name: 'Search',
      component: {
        HelloWorld: HelloWorld
      }
    },
    {
      path: '/vod',
      name: 'Test',
      component: NotFound
    }
  ]
})
