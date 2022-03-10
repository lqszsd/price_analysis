<template>
<div>
<el-table
    :data="tableData"
    border
    style="width: 100%">
    <el-table-column
      prop="symbol"
      label="代码"
      width="180">
    </el-table-column>
    <el-table-column
      prop="plan_copy_the_bottom"
      label="抄底价"
    >
    </el-table-column>
    <el-table-column
      prop="plan_sale"
      label="卖出价">
    </el-table-column>
     <el-table-column
      prop="name"
      label="名称">
    </el-table-column>
       <el-table-column
      prop="status"
      label="状态">
      <template slot-scope="scope">
       <el-switch
        @change="change_strategy_status(scope.row)"
        v-model="scope.row.status"
        active-text="已买入"
         inactive-text="待买入">
        </el-switch>
      </template>
    </el-table-column>
      <el-table-column
      prop="email_count"
      label="发送邮件次数">
    </el-table-column>
    <el-table-column
      fixed="right"
      label="操作"
      width="100">
      <template slot-scope="scope">
        <el-button @click="handleClick(scope.row)" type="text" size="small">查看</el-button>
      </template>
    </el-table-column>
  </el-table>
      <el-dialog
  title="提示"
  :visible.sync="dialogVisible"
  width="80%"
  :before-close="handleClose">
  <HelloWorld :symbol="symbol"></HelloWorld>
</el-dialog>
</div>
</template>


<script>
import HelloWorld from "../HelloWorld.vue";
export default {
  name: "StrategyList",
  components: {
    HelloWorld,
  },
  data: function () {
    return {
      tableData: null,
      dialogVisible: false,
      symbol: "",
    };
  },
  methods: {
    handleClick: function (row) {
      this.symbol = row.symbol;
      this.dialogVisible = true;
    },
    handleClose: function () {
      this.dialogVisible = false;
    },
    change_strategy_status: function (row) {
       this.$post("/api/change_strategy_status",row).then((data) => {
        console.log(data);
      });
    },
  },
  mounted: function () {
    var that = this;
    that.$get("/api/user_strategy").then((data) => {
      console.log(data);
      that.tableData = data;
    });
    this.time_id = setInterval(function () {
      that.$get("/api/user_strategy").then((data) => {
        console.log(data);
        that.tableData = data;
      });
    }, 5000);
  },
};
</script>