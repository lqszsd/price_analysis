<template>

  <div id="app">
  <el-container style="height: 800px; border: 1px solid #eee">
  <el-aside width="200px" style="background-color: rgb(238, 241, 246)">
    <el-menu :default-openeds="['1', '3']">
       <el-menu-item index="0">
        <i class="el-icon-menu"></i>
        <span slot="title"><router-link to="/">热门股票</router-link></span>
      </el-menu-item>
       <el-menu-item index="1">
        <i class="el-icon-menu"></i>
        <span slot="title"> <router-link to="/recommend">推荐列表</router-link></span>
      </el-menu-item>
       <el-menu-item index="1">
        <i class="el-icon-menu"></i>
        <span slot="title"><router-link to="/my_choose">我的自选</router-link></span>
      </el-menu-item>
          <el-menu-item index="1">
        <i class="el-icon-menu"></i>
        <span slot="title"><router-link to="/my_strategy">我的策略</router-link></span>
      </el-menu-item>
       <el-menu-item index="1">
        <i class="el-icon-menu"></i>
        <span slot="title"><router-link to="/all_recommend">机构推荐</router-link></span>
      </el-menu-item>
    </el-menu>
  </el-aside>
  
  <el-container>
    <el-header style="text-align: right; font-size: 12px">
      <el-dropdown>
        <i class="el-icon-setting" style="margin-right: 15px"></i>
        <el-dropdown-menu slot="dropdown">
          <el-dropdown-item>查看</el-dropdown-item>
          <el-dropdown-item>新增</el-dropdown-item>
          <el-dropdown-item>删除</el-dropdown-item>
        </el-dropdown-menu>
      </el-dropdown>
      <span>王小</span>
    </el-header>
    
    <el-main>
      <router-view></router-view>
    </el-main>
  </el-container>
</el-container>
  </div>
</template>

<script>
export default {
  name: 'App',
  components: {
  
  },
  data() {
      const item = {
        date: '2016-05-02',
        name: '王小虎',
        address: '上海市普陀区金沙江路 1518 弄'
      };
      return {
        tableData: Array(20).fill(item)
      }
    }
}
</script>

<style>
body{
  height: 100%;
  width: 100%;
}
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
  height: 100%;
  width: 100%;
}
.el-table{

width:99.9%

}
.el-header {
    background-color: #B3C0D1;
    color: #333;
    line-height: 60px;
  }
  
  .el-aside {
    color: #333;
  }
</style>
