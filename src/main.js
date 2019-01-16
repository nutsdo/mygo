import '@babel/polyfill'
import Vue from 'vue'
import './plugins/bootstrap-vue'
import App from './App.vue'
import router from './router'
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'
import 'bootstrap/dist/js/bootstrap.js'
import axios from "axios"

Vue.config.productionTip = false
Vue.prototype.axios = axios
new Vue({
  router,
    methods:{
      axios:function () {
          return axios
      }
    },
  render: h => h(App)
}).$mount('#app')
