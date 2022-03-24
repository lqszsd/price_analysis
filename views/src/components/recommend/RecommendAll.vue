<template>
  <div>
    <el-table :data="tableData" height="800" border style="width: 100%">
      <el-table-column prop="股票名称" label="名称" width="180"> </el-table-column>
      <el-table-column prop="最新评级" label="最新评级"> </el-table-column>
      <el-table-column prop="目标价" label="目标价">
      </el-table-column>
      <el-table-column prop="评级日期" label="评级日期"> </el-table-column>
      <el-table-column prop="综合评级" label="综合评级"> </el-table-column>
      <el-table-column prop="平均涨幅" label="平均涨幅"> </el-table-column>
      <el-table-column prop="行业" label="行业"> </el-table-column>
      <el-table-column prop="股票代码" label="股票代码"> </el-table-column>
      <el-table-column fixed="right" label="操作" width="100">
        <template slot-scope="scope">
          <el-button @click="handleClick(scope.row)" type="text" size="small"
            >查看</el-button
          >
          <el-button @click="add_select(scope.row)" type="text" size="small"
            >加入自选</el-button
          >
          <el-button @click="add_strategy(scope.row)" type="text" size="small"
            >加入抄底</el-button
          >
          <el-button @click="view_notice(scope.row)" type="text" size="small"
            >查看公告</el-button
          >
        </template>
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
    <el-dialog
      title="提示"
      :visible.sync="dialogVisible"
      width="80%"
      :before-close="handleClose"
    >
      <HelloWorld :symbol="symbol"></HelloWorld>
    </el-dialog>
    <el-dialog
      title="提示"
      :visible.sync="add_select_dialog"
      width="80%"
      :before-close="handleClose"
    >
      <AddSelect :symbol="symbol" :price="price" :name="name"></AddSelect>
    </el-dialog>
    <el-dialog
      title="提示"
      :visible.sync="add_strategy_dialog"
      width="80%"
      :before-close="handleClose"
    >
      <AddStrategy :symbol="symbol" :price="price" :name="name"></AddStrategy>
    </el-dialog>
    <el-dialog
      title="新闻"
      :visible.sync="news_dialog"
      width="80%"
      :before-close="handleClose"
    >
      <NewsList :symbol="symbol"></NewsList>
    </el-dialog>
  </div>
</template>


<script>
import HelloWorld from "../HelloWorld.vue";
import AddSelect from "../userselect/AddSelect.vue";
import AddStrategy from "../strategy/AddStrategy.vue";
import NewsList from "../news/NewsList.vue";
var ws = new WebSocket("ws://127.0.0.1:9090/api/all_recommend");
export default {
  name: "RecommendAll",
  components: {
    HelloWorld,
    AddSelect,
    AddStrategy,
    NewsList,
  },
  data: function () {
    return {
      tableData: null,
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
        api: "recommend",
      };
      var that = this;
      if (ws.readyState == 1) {
        console.log("发送消息2222");
        ws.send(JSON.stringify(data));
        //接收到消息时触发
        ws.onmessage = function (evt) {
          that.tableData = JSON.parse(evt.data);
          console.log("Received Message: " + evt.data);
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
      api: "recommend",
    };
    var that = this;
    setTimeout(function () {
      if (ws.readyState == 1) {
        console.log("发送消息");
        ws.send(JSON.stringify(data));
        //接收到消息时触发
        ws.onmessage = function (evt) {
          that.tableData = JSON.parse(evt.data);
          console.log("Received Message: " + evt.data);
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