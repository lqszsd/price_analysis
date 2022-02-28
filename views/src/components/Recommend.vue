<template>
<div>
<el-table
    :data="tableData"
    border
    style="width: 100%">
    <el-table-column
      prop="Name"
      label="名称"
      width="180">
    </el-table-column>
    <el-table-column
      prop="NowPrice"
      label="当前价格">
    </el-table-column>
    <el-table-column
      prop="month_avg"
      label="最近两个月均价">
    </el-table-column>
     <el-table-column
      prop="month_max_money"
      label="月最高">
    </el-table-column>
     <el-table-column
      prop="month_min_money"
      label="月最低">
    </el-table-column>
         <el-table-column
      prop="month_msg"
      label="月建议">
    </el-table-column>
       <el-table-column
      prop="seven_avg"
      label="周均价">
    </el-table-column>
    <el-table-column
      prop="seven_max_money"
      label="周最高">
    </el-table-column>
     <el-table-column
      prop="seven_min_money"
      label="周最低">
    </el-table-column>
     <el-table-column
      prop="seven_msg"
      label="周建议">
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
import HelloWorld from './HelloWorld.vue'
export default {
  name: "Recommend",
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
          console.log(row)
          this.dialogVisible=true;
          this.symbol=row.symbol;
      },
  },
  mounted:function(){
      var that=this;
      this.$get("/api/recommend").then(data=>{
          console.log(data)
          that.tableData=data;
      })
    console.log(this.$get);
  }
}
</script>
