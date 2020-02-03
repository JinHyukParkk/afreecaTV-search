// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import axios from 'axios'
import vueCookies from 'vue-cookies'

Vue.config.productionTip = false
Vue.prototype.$http = axios
Vue.prototype.$cookies = vueCookies

/* eslint-disable no-new */
new Vue({
  el: '#search_content',
  render: h => h(App),
  router
})
