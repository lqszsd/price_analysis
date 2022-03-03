import Vue from 'vue'
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
import App from './App.vue'
import routes from './router'
import VueRouter from 'vue-router'
import axios from 'axios';
import {get} from './utils/http.js'
import {post} from './utils/http.js'
Vue.use(VueRouter)
axios.defaults.headers.common['X-Requested-With'] = 'XMLHttpRequest';
axios.defaults.headers.post['Content-Type'] = 'application/x-www-form-urlencoded';
axios.defaults.headers.get['Content-Type'] = 'application/x-www-form-urlencoded';
Vue.prototype.$get = get;
Vue.prototype.$post = post;
const router = new VueRouter({
  mode: 'history',
  routes
})
// 3. 创建路由实例并传递 `routes` 配置
// 你可以在这里输入更多的配置，但我们在这里
// 暂时保持简单
Vue.use(ElementUI);

Vue.config.productionTip = false;
new Vue({
  router,
  render: h => h(App),
}).$mount('#app');

