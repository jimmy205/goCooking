<template>
    <div class="charts">
        <div>
            <ve-line :data="chartData" :grid="grid"></ve-line>
        </div>
        <div>
            <button type="button" @click="getRTPInfo">Refresh!</button>
        </div>
    </div>
</template>

<script>
import VeLine from 'v-charts/lib/line.common'
import axios from 'axios'

export default {
  components: { VeLine },
  data () {
        this.grid = {
            top:110
        }
    return {
        keys:[],
        chartData: {
        columns: [],
        rows: []
        }
    }
  },
  methods:{
    getRTPInfo : function () {

        this.keys = []
        this.chartData.rows = []
        this.chartData.columns = []
        axios({
            methods:'get',
            url:'http://127.0.0.1:8800/api/rtp_chart'
        })
        .then((resp) => {
            this.info = resp.data.data
            this.chartData.columns.push("bullet_num")
            this.chartData.columns = this.chartData.columns.concat(resp.data.username.sort())

            this.keys = Object.keys(this.info)
            for (let i = 0; i < this.keys.length; i++) {
                const key = this.keys[i];
                var column = {"bullet_num":key}
                for (let i = 0; i < this.info[key].length; i++) {
                    const element = this.info[key][i];
                    Object.assign(column,element)
                }
                this.chartData.rows.push(column)
            }
        });
    }
  }
}
</script>