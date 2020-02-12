<template>
    <div class="charts">
        <div>
            <ve-line :data="chartData" :grid="grid"></ve-line>
        </div>
        <div>
            <div>
                <label> 顯示方式 :</label>
                    <input type="radio" name="showType" value="room"> 房間
                    <input type="radio" name="showType" value="person"> 個人
                    <input type="radio" name="showType" value="all"> 整體
            </div>
            <div>
                <label> RTP設定值 :</label>
                    <input type="radio" name="rtp" value="60"> 60
                    <input type="radio" name="rtp" value="98"> 98
            </div>
        </div>
        <div>
            <button type="button" @click="getRTP('person')">每位玩家</button>
            <button type="button" @click="getRTP('room')">每個房間</button>
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
    getRTP : function (str) {

        let url = 'http://127.0.0.1:8800/api/rtp_person'
        if (str == 'room') {
            url = 'http://127.0.0.1:8800/api/rtp_room'
        }

        this.keys = []
        this.chartData.rows = []
        this.chartData.columns = []

        axios({
            methods:'get',
            url:url
        })
        .then((resp) => {
            this.info = resp.data.data
            this.chartData.columns.push("bullet_num")
            this.chartData.columns = this.chartData.columns.concat(resp.data.username)

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