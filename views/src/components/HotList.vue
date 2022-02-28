<template>
<div>
<el-table
    :data="tableData"
    border
    style="width: 100%">
    <el-table-column
      prop="代码"
      label="代码"
      width="180">
    </el-table-column>
    <el-table-column
      prop="当前排名"
      label="当前排名"
      width="180">
    </el-table-column>
    <el-table-column
      prop="最新价"
      label="最新价">
    </el-table-column>
     <el-table-column
      prop="涨跌幅"
      label="涨跌幅">
    </el-table-column>
     <el-table-column
      prop="股票名称"
      label="股票名称">
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
  {{symbol}}
  <HelloWorld :symbol="symbol"></HelloWorld>
</el-dialog>
</div>
</template>
<script>
import HelloWorld from './HelloWorld.vue'
export default {
  name: "HotList",
  components: {
    HelloWorld
  },
   data:function(){
    return {
        tableData:null,
        dialogVisible:false,
        symbol:"",
    }
  },
    methods:{
      handleClick:function(row){
          console.log(2222222222,row)
          this.symbol=row.代码
          this.dialogVisible=true;
      },
      handleClose:function(){
        this.dialogVisible=false;
      },
  },
  mounted:function(){
      var that=this;
      this.$get("/api/hot/list").then(data=>{
          console.log(data)
          that.tableData=data;
      })
    console.log(this.$get);
  }
}
</script>
