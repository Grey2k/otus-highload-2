// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
// import 'bootstrap/dist/css/bootstrap.css';
import Vue from 'vue';
import VueMaterial from 'vue-material';
import Vuelidate from 'vuelidate';
import Paginate from 'vuejs-paginate';
import FlashMessage from '@smartweb/vue-flash-message';
import VueFilterDateFormat from '@vuejs-community/vue-filter-date-format';
import 'vue-material/dist/vue-material.css';
import 'vue-material/dist/theme/default.css';


import App from './App';
import router from './router';

Vue.config.productionTip = false;

Vue.use(VueMaterial);
Vue.use(Vuelidate);
Vue.use(FlashMessage);
Vue.use(VueFilterDateFormat);

Vue.component('paginate', Paginate);

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  components: { App },
  template: '<App/>',
});
