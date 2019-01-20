import Vue from 'vue';
import VueI18n from 'vue-i18n';
import Vuex from 'vuex';
import App from './App';
import createRouter from './router/index';

import * as api from './api';

import enLocale from 'element-ui/lib/locale/lang/en';

import PaperDashboard from './plugins/paperDashboard';

import 'vue-notifyjs/themes/default.css';
import { stat } from 'fs';

Vue.use(Vuex);
Vue.use(VueI18n);
Vue.use(PaperDashboard);

Vue.config.lang = 'en';
Vue.locale('en', enLocale);

const store = new Vuex.Store({
  state: {
    isAuthenticated: !!localStorage.getItem('dashboard_auth'),
    dashboard: {
      statsCards: [],
      usersChartData: {
        labels: [
          '9:00AM',
          '12:00AM',
          '3:00PM',
          '6:00PM',
          '9:00PM',
          '12:00PM',
          '3:00AM',
          '6:00AM'
        ],
        series: []
      },
      activityChart: {
        data: {
          labels: [
            'Jan',
            'Feb',
            'Mar',
            'Apr',
            'Mai',
            'Jun',
            'Jul',
            'Aug',
            'Sep',
            'Oct',
            'Nov',
            'Dec'
          ],
          series: [
            [542, 543, 520, 680, 653, 753, 326, 434, 568, 610, 756, 895],
            [230, 293, 380, 480, 503, 553, 600, 664, 698, 710, 736, 795]
          ]
        },
        options: {
          seriesBarDistance: 10,
          axisX: {
            showGrid: false
          },
          height: '245px'
        }
      }
    }
  },

  mutations: {
    login(state) {
      state.isAuthenticated = true;
    },
    logout(state) {
      state.isAuthenticated = false;
    },
    setDashboardStats(state, statsCards) {
      state.dashboard.statsCards = statsCards;
    },

    setDashboardUsersChartData(state, series) {
      state.dashboard.usersChartData = Object.assign(
        state.dashboard.usersChartData,
        { series: series }
      );
    },

    setDashboardActivityChartData(state, series) {
      state.dashboard.activityChart.data.series = series;
    }
  },

  actions: {
    login(context, payload) {
      return new Promise((resolve, reject) => {
        api.login(payload).then(
          () => {
            context.commit('login');
            localStorage.setItem('dashboard_auth', true);
            resolve();
          },
          err => reject(err)
        );
      });
    },

    logout(context) {
      return new Promise((resolve, reject) => {
        context.commit('logout');
        localStorage.setItem('dashboard_auth', false);
        resolve();
      });
    },

    refreshDashboard(context, payload) {
      const [from, to] = payload;
      api.getDashboardStats(from, to).then(res => {
        context.commit('setDashboardStats', res.stats);
      });
      api.getDashboardUserChart(from, to).then(res => {
        context.commit('setDashboardUsersChartData', res.series);
      });
      api.getDashboardActivityChart(from, to).then(res => {
        context.commit('setDashboardActivityChartData', res.series);
      });
    }
  }
});

const router = createRouter(store);

/* eslint-disable no-new */
new Vue({
  router,
  store: store,
  render: h => h(App)
}).$mount('#app');
