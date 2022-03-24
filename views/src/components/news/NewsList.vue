<template>
  <div>
    <el-table :data="tableData" border style="width: 100%">
      <el-table-column prop="code" label="代码" width="80"> </el-table-column>
            <el-table-column prop="title" label="标题" min-width="180" show-overflow-tooltip> </el-table-column>
      <el-table-column prop="content" label="内容" show-overflow-tooltip min-width="180">
      </el-table-column>
      <el-table-column prop="public_time" label="发布时间"> </el-table-column>
      <el-table-column prop="url" label="链接"> </el-table-column>
    </el-table>


  </div>
</template>
<script>
export default {
  name: "NewList",
  components: {
  },
  data: function () {
    return {
      tableData: null,
      dialogVisible: false,
      add_select_dialog: false,
      add_strategy_dialog:false,
      price: "",
      name: "",
    };
  },
  watch: {
    symbol (val) {
      var that = this;
      this.$get("/api/news_list",{symbol:val}).then((data) => {
      console.log(data);
      that.tableData = data;
    });
    }
  },
  methods: {
  },
    props: {
    symbol: String,
  },
  mounted: function () {
    var that = this;
    console.log(this.symbol);
    this.$get("/api/news_list",{symbol:that.symbol}).then((data) => {
      console.log(data);
      that.tableData = data;
    });
    console.log(this.$get);
  },
};
</script>
