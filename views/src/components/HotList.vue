<template>
  <div>
    <el-table :data="tableData" border style="width: 100%">
      <el-table-column prop="代码" label="代码" width="180"> </el-table-column>
      <el-table-column prop="当前排名" label="当前排名" width="180">
      </el-table-column>
      <el-table-column prop="最新价" label="最新价"> </el-table-column>
      <el-table-column prop="涨跌幅" label="涨跌幅"> </el-table-column>
      <el-table-column prop="股票名称" label="股票名称"> </el-table-column>
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
      :visible.sync="add_strategy_dialog"
      width="80%"
      :before-close="handleClose"
    >
      <AddStrategy :symbol="symbol" :price="price" :name="name"></AddStrategy>
    
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
  name: "HotList",
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
      this.symbol = row.代码;

      this.news_dialog=true;
    },
    add_strategy:function(row){
       this.add_strategy_dialog = true;
      this.symbol = row.代码;
      this.price = row.最新价;
      this.name = row.股票名称;
    },
    add_select: function (row) {
      this.add_select_dialog = true;
      this.symbol = row.代码;
      this.price = row.最新价;
      this.name = row.股票名称;
    },
    handleClick: function (row) {
      console.log(2222222222, row);
      this.symbol = row.代码;
      this.dialogVisible = true;
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
    this.$get("/api/hot/list").then((data) => {
      console.log(data);
      that.tableData = data;
    });
    console.log(this.$get);
  },
};
</script>
