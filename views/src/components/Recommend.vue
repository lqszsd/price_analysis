<template>
  <div>
    <el-table :data="tableData" height="800" border style="width: 100%">
      <el-table-column prop="Name" label="名称" width="180"> </el-table-column>
      <el-table-column prop="NowPrice" label="当前价格"> </el-table-column>
      <el-table-column prop="month_avg" label="最近两个月均价">
      </el-table-column>
      <el-table-column prop="month_max_money" label="月最高"> </el-table-column>
      <el-table-column prop="month_min_money" label="月最低"> </el-table-column>
      <el-table-column prop="month_msg" label="月建议"> </el-table-column>
      <el-table-column prop="seven_avg" label="周均价"> </el-table-column>
      <el-table-column prop="seven_max_money" label="周最高"> </el-table-column>
      <el-table-column prop="seven_min_money" label="周最低"> </el-table-column>
      <el-table-column prop="seven_msg" label="周建议"> </el-table-column>
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
import HelloWorld from "./HelloWorld.vue";
import AddSelect from "./userselect/AddSelect.vue";
import AddStrategy from "./strategy/AddStrategy.vue";
import NewsList from "./news/NewsList.vue";
export default {
  name: "Recommend",
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
      news_dialog:false,
      add_select_dialog: false,
      add_strategy_dialog:false,
      symbol: "",
      price: "",
      name: "",
    };
  },
  methods: {
    view_notice:function(row){
      this.symbol = row.symbol;

      this.news_dialog=true;
    },
    handleClick: function (row) {
      console.log(row);
      this.dialogVisible = true;
      this.symbol = row.symbol;
    },
    add_strategy:function(row){
       this.add_strategy_dialog = true;
      this.symbol = row.symbol;
      this.price = row.NowPrice;
      this.name = row.Name;
    },
    add_select: function (row) {
      this.add_select_dialog = true;
      this.symbol = row.symbol;
      this.price = row.NowPrice;
      this.name = row.Name;
    },
    handleClose: function () {
       this.news_dialog=false;
      this.dialogVisible = false;
      this.add_select_dialog=false;
      this.add_strategy_dialog=false;
    },
  },
  mounted: function () {
    var that = this;
    this.$get("/api/recommend").then((data) => {
      console.log(data);
      that.tableData = data;
    });
    console.log(this.$get);
  },
};
</script>
