<template>
  <div>
    <el-table :data="tableData" border style="width: 100%">
      <el-table-column prop="Symbol" label="代码" width="180">
      </el-table-column>
      <el-table-column prop="Cost" label="花费"> </el-table-column>
      <el-table-column prop="Price" label="价格"> </el-table-column>
      <el-table-column prop="Name" label="名称"> </el-table-column>
      <el-table-column prop="Estimate" label="当日金额"> </el-table-column>
      <el-table-column prop="Shareholding" label="持有"> </el-table-column>
      <el-table-column fixed="right" label="操作" width="100">
        <template slot-scope="scope">
          <el-button @click="handleClick(scope.row)" type="text" size="small"
            >查看</el-button
          >
          <el-button type="text" @click="change_shareholding(scope.row)"
            >修改持仓</el-button
          >
          <el-button type="text" @click="change_cost(scope.row)"
            >修改花费</el-button
          >
                  <el-button @click="view_notice(scope.row)" type="text" size="small"
            >查看公告</el-button>
                    <el-button @click="rm_select(scope.row)" type="text" size="small"
            >移除自选</el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-dialog
      title="提示"
      :visible.sync="dialogVisible"
      width="80%"
      :before-close="handleClose"
    >
      <HelloWorld :symbol="symbol"></HelloWorld>
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
import HelloWorld from "./HelloWorld.vue";
import NewsList from "./news/NewsList.vue";
export default {
  name: "Test",
  components: {
    HelloWorld,
    NewsList,
  },
  data: function () {
    return {
      tableData: null,
      dialogVisible: false,
      symbol: "",
      news_dialog:false,
    };
  },
  methods: {
    rm_select:function(row){
      this.$get("/api/rm_select", {
            symbol: row.Symbol,
          }).then((data) => {
            console.log(data);
          });
    },
     view_notice:function(row){
      this.symbol = row.Symbol;

      this.news_dialog=true;
    },
    handleClick: function (row) {
      this.symbol = row.Symbol;
      this.dialogVisible = true;
    },
     handleClose: function () {
       this.news_dialog=false;
      this.dialogVisible = false;
      this.add_select_dialog=false;
      this.add_strategy_dialog=false;
    },
    change_cost: function (row) {
      this.$prompt("请输入价格", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
      })
        .then(({ value }) => {
          this.$post("/api/add_select", {
            symbol: row.Symbol,
            price: row.Price,
            name: row.Name,
            cost: value,
            shareholding: row.Shareholding,
          }).then((data) => {
            console.log(data);
          });
        })
        .catch(() => {
          this.$message({
            type: "info",
            message: "取消输入",
          });
        });
    },
    change_shareholding: function (row) {
      this.$prompt("请输入价格", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
      })
        .then(({ value }) => {
          this.$post("/api/add_select", {
            symbol: row.Symbol,
            price: row.Price,
            name: row.Name,
            cost: row.Cost,
            shareholding: value,
          }).then((data) => {
            console.log(data);
          });
        })
        .catch(() => {
          this.$message({
            type: "info",
            message: "取消输入",
          });
        });
    },
  },
  mounted: function () {
    var that = this;
    that.$get("/api/user_select").then((data) => {
      console.log(data);
      that.tableData = data;
    });
    this.time_id = setInterval(function () {
      that.$get("/api/user_select").then((data) => {
        console.log(data);
        that.tableData = data;
      });
    }, 8000);
  },
};
</script>