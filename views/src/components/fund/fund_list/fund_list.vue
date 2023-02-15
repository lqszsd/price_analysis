<template>
    <div>
        <el-form :inline="true"  ref="resetForm" class="demo-form-inline">
                                <el-row>
                                    <el-form-item label="基金代码:">
                                        <el-input size="mini" style="width: 280px;" v-model="fund_code"
                                                  placeholder="基金代码:"></el-input>
                                    </el-form-item>
                                    <el-form-item label="基金名称:">
                                        <el-input size="mini" style="width: 280px;" v-model="fund_name"
                                                  placeholder="基金名称:">
                                        </el-input>
                                    </el-form-item>
                                    <el-button size="small" type="primary" icon="el-icon-search" @click="onSubmit">查询
                                    </el-button>
                                    <el-button size="small" type="primary" icon="el-icon-delete"
                                               @click="resetForm('resetForm')">重置
                                    </el-button>
                                    <el-button size="small" type="primary" @click="flash()">刷新</el-button>
                                </el-row>
                            </el-form>

      <el-table :data="tableData" height="800" border style="width: 100%">
        <el-table-column prop="FundCode" label="代码" width="180"> </el-table-column>
        <el-table-column prop="FundName" label="名称"> </el-table-column>
        <el-table-column prop="AverageNav" label="价格">
        </el-table-column>
        <el-table-column prop="BeginDay" label="创建时间"> </el-table-column>
        <el-table-column prop="FundManager" label="管理人"> </el-table-column>
        <el-table-column prop="TotalPrice" label="总金额"> </el-table-column>
        <el-table-column prop="three_month" label="近3月"> </el-table-column>
        <el-table-column prop="six_month" label="近6月"> </el-table-column>
        <el-table-column prop="year" label="近1年"> </el-table-column>
        <el-table-column prop="two_year" label="近2年"> </el-table-column>
        <el-table-column fixed="right" label="操作" width="100">
          <!-- <template slot-scope="scope">
          </template> -->
        </el-table-column>
      </el-table>
      <el-row>
        <el-col :span="11" style="float: right !important">
          <el-pagination
            @size-change="handleSizeChange"
            @current-change="handleCurrentChange"
            :current-page="page"
            :page-sizes="[10, 100, 200, 500]"
            :page-size="10"
            layout="total, sizes, prev, pager, next, jumper"
            :total="count"
          >
          </el-pagination>
        </el-col>
      </el-row>
    </div>
  </template>
  
  
  <script>
  import ReconnectingWebSocket from 'reconnecting-websocket';

  var ws = new ReconnectingWebSocket("ws://127.0.0.1:9090/api/fund/fund_list");
  export default {
    name: "FundList",
    components: {
    },
    data: function () {
      return {
        tableData: null,
        fund_code:"",
        fund_name:"",
        dialogVisible: false,
        news_dialog: false,
        add_select_dialog: false,
        add_strategy_dialog: false,
        symbol: "",
        price: "",
        name: "",
        page: 1,
        limit: 10,
        count:4500,
      };
    },
    methods: {
        flash: function () {
                this.reload();
            },
            onSubmit: function () {
                this.reload()
            },
            resetForm() {
                this.fund_code="";
                this.fund_name = "";
                this.reload();
            },
      handleSizeChange: function (val) {
        this.limit = val;
        this.reload();
      },
      handleCurrentChange: function (val) {
        this.page = val;
        this.reload();
      },
      reload: function () {
        let data = {
          page: this.page,
          limit: this.limit,
          api: "fund_list",
          fund_code:this.fund_code,
          fund_name:this.fund_name
        };
        var that = this;
        if (ws.readyState == 1) {
          console.log("发送消息2222");
          ws.send(JSON.stringify(data));
          //接收到消息时触发
          ws.onmessage = function (evt) {
            let json_data = JSON.parse(evt.data);
            that.tableData=json_data.list;
            that.count=json_data.count;
            console.log("Received Message: " , json_data.list);
          };
        } else {
          console.log(ws);
          ws.send(JSON.stringify(data));
          //do something
        }
      },
      view_notice: function (row) {
        this.symbol = row.股票代码;
  
        this.news_dialog = true;
      },
      handleClick: function (row) {
        console.log(row);
        this.dialogVisible = true;
        this.symbol = row.股票代码;
      },
      add_strategy: function (row) {
        this.add_strategy_dialog = true;
        this.symbol = row.股票代码;
        this.price = 0.00;
        this.name = row.股票名称;
      },
      add_select: function (row) {
        this.add_select_dialog = true;
        this.symbol = row.股票代码;
        this.price = 0.00;
        this.name = row.股票名称;
      },
      handleClose: function () {
        this.news_dialog = false;
        this.dialogVisible = false;
        this.add_select_dialog = false;
        this.add_strategy_dialog = false;
      },
    },
    mounted: function () {
      console.log(3333333333);
      let data = {
        page: this.page,
        limit: this.limit,
        api: "fund_list",
        fund_code:this.fund_code,
        fund_name:this.fund_name
      };
      var that = this;
      setTimeout(function () {
        if (ws.readyState == 1) {
          console.log("发送消息");
          ws.send(JSON.stringify(data));
          //接收到消息时触发
          ws.onmessage = function (evt) {
            let json_data = JSON.parse(evt.data);
            that.tableData=json_data.list;
            that.count=json_data.count;
            console.log("Received Message: " ,json_data);
          };
        } else {
          console.log(ws);
          ws.send(JSON.stringify(data));
          //do something
        }
      }, 5000);
    },
  };
  </script>