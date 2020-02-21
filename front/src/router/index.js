import Vue from 'vue'
import Router from 'vue-router'
import Vuex from 'vuex'
import VueCookie from 'vue-cookie'
import 'es6-promise/auto'
import TotalSearch from '@/components/tab/TotalSearch'
import Live from '@/components/tab/Live'
import Vod from '@/components/tab/Vod'
import Post from '@/components/tab/Post'
import Bj from '@/components/tab/Bj'

Vue.use(Router)
Vue.use(Vuex)
Vue.use(VueCookie)

export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      component: TotalSearch
    },
    {
      path: '/live',
      component: Live
    },
    {
      path: '/vod',
      component: Vod
    },
    {
      path: '/post',
      component: Post
    },
    {
      path: '/bj',
      component: Bj
    }
  ]
})
