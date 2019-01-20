<template>
  <div>
    <div class="dashboard-datepicker-container clearfix">
      <el-date-picker
        class="dashboard-datepicker"
        v-model="dates"
        v-on:change="onDatesChange"
        type="daterange"
        align="right"
        start-placeholder="Start date"
        end-placeholder="End date"
      />
    </div>

    <!--Stats cards-->
    <div class="row">
      <div class="col-md-6 col-xl-3" v-for="stats in statsCards" :key="stats.title">
        <stats-card>
          <div class="icon-big text-center" :class="`icon-${stats.type}`" slot="header">
            <i :class="stats.icon"></i>
          </div>
          <div class="numbers" slot="content">
            <p>{{stats.title}}</p>
            {{stats.value}}
          </div>
          <div class="stats" slot="footer">
            <i :class="stats.footerIcon"></i>
            {{stats.footerText}}
          </div>
        </stats-card>
      </div>
    </div>

    <!--Charts-->
    <div class="row">
      <div class="col-12">
        <chart-card
          title="Users behavior"
          sub-title="24 Hours performance"
          :chart-data="usersChartData"
          :chart-options="usersChart.options"
        >
          <span slot="footer">
            <i class="ti-reload"></i> Updated 3 minutes ago
          </span>
          <div slot="legend">
            <i class="fa fa-circle text-info"></i> Open
            <i class="fa fa-circle text-danger"></i> Click
            <i class="fa fa-circle text-warning"></i> Click Second Time
          </div>
        </chart-card>
      </div>

      <div class="col-md-6 col-12">
        <chart-card
          title="Email Statistics"
          sub-title="Last campaign performance"
          :chart-data="preferencesChart.data"
          chart-type="Pie"
        >
          <span slot="footer">
            <i class="ti-timer"></i> Campaign set 2 days ago
          </span>
          <div slot="legend">
            <i class="fa fa-circle text-info"></i> Open
            <i class="fa fa-circle text-danger"></i> Bounce
            <i class="fa fa-circle text-warning"></i> Unsubscribe
          </div>
        </chart-card>
      </div>

      <div class="col-md-6 col-12">
        <chart-card
          title="2015 Sales"
          sub-title="All products including Taxes"
          :chart-data="activityChart.data"
          :chart-options="activityChart.options"
        >
          <span slot="footer">
            <i class="ti-check"></i> Data information certified
          </span>
          <div slot="legend">
            <i class="fa fa-circle text-info"></i> Tesla Model S
            <i class="fa fa-circle text-warning"></i> BMW 5 Series
          </div>
        </chart-card>
      </div>
    </div>
  </div>
</template>
<script>
import { StatsCard, ChartCard } from "@/components/index";
import Chartist from "chartist";
export default {
  components: {
    StatsCard,
    ChartCard
  },

  created() {
    this.$store.dispatch("refreshDashboard", this.dates);
  },

  /**
   * Chart data used to render stats, charts. Should be replaced with server data
   */
  methods: {
    onDatesChange: function(dates) {
      this.$store.dispatch("refreshDashboard", dates);
    }
  },

  computed: {
    statsCards() {
      return this.$store.state.dashboard.statsCards;
    },

    usersChartData() {
      return this.$store.state.dashboard.usersChartData;
    },

    activityChart() {
      return this.$store.state.dashboard.activityChart;
    }
  },

  data() {
    return {
      dates: [new Date(), new Date()],
      usersChart: {
        options: {
          low: 0,
          high: 1000,
          showArea: true,
          height: "245px",
          axisX: {
            showGrid: false
          },
          lineSmooth: Chartist.Interpolation.simple({
            divisor: 3
          }),
          showLine: true,
          showPoint: false
        }
      },
      preferencesChart: {
        data: {
          labels: ["62%", "32%", "6%"],
          series: [62, 32, 6]
        },
        options: {}
      }
    };
  }
};
</script>
<style>
.dashboard-datepicker-container {
  margin-bottom: 36px;
}
.dashboard-datepicker {
  float: right;
}
</style>
