<template>
  <div id="container">
    <el-date-picker
      v-model="time_list"
      type="daterange"
       value-format="yyyyMMdd"
      range-separator="至"
      start-placeholder="开始日期"
      end-placeholder="结束日期">
    </el-date-picker>
  <div id="main"></div>
  </div>
</template>

<script>
import { Chart } from "@antv/g2";

function findMaxMin(data) {
  let maxValue = 0;
  let minValue = 50000;
  let maxObj = null;
  let minObj = null;
  for (const d of data) {
    if (d.收盘 > maxValue) {
      maxValue = d.收盘;
      maxObj = d;
    }
    if (d.收盘 < minValue) {
      minValue = d.收盘;
      minObj = d;
    }
  }
  return { max: maxObj, min: minObj };
}
export default {
  name: "HelloWorld",
  data:function(){
    return {
      time_list:"",
      start_date:"",
      end_date:"",
      time_id:null,
    }

  },
  mounted:function(){
    window.clearInterval()
var chart = new Chart({
  container: "main",
  autoFit: true,
  height: 500,
});
chart.scale({
  时间: {
    nice: true,
  },
  收盘: {
    nice: true,
  },
});
chart.axis("时间", {
  label: {
    formatter: (text) => {
      const dataStrings = text.split(".");
      return dataStrings[0];
    },
  },
});

chart.line().position("时间*收盘");
let that=this;
this.time_id=setInterval(function(){
  let thats=that;
  let start_date="";
  let end_date="";
  if(thats.time_list==""){
    let dd=new Date()
    dd.setDate(dd.getDate()-1);//获取AddDayCount天后的日期
    let y = dd.getFullYear();
    let m = dd.getMonth()+1;//获取当前月份的日期
    if(m<10){
      m="0"+m;
    }
    let d = dd.getDate();
    if(d<10){
      d="0"+d;
    }
    console.log(y,m,d)
    start_date=""+y+"-"+m+"-"+d;
    console.log(start_date)
    dd.setDate(dd.getDate()+2);//获取AddDayCount天后的日期
     y = dd.getFullYear();
     m = dd.getMonth()+1;//获取当前月份的日期
      if(m<10){
      m="0"+m;
    }
     d = dd.getDate();
      if(d<10){
      d="0"+d;
    }
    end_date=""+y+"-"+m+"-"+d;
  } else{
     start_date=thats.time_list[0];
     end_date=thats.time_list[1];
  }
fetch("http://127.0.0.1:9090/api/", {
  method: "POST",
  mode: "cors",
  body: JSON.stringify({ symbol: thats.symbol,start_date:start_date,end_date:end_date, }),
})
  .then((res) => res.json())
  .then((data) => {
    chart.changeData(data)
    const { min, max } = findMaxMin(data);
chart.annotation().dataMarker({
  top: true,
  position: [max.时间, max.收盘],
  text: {
    content: "全部峰值：" + max.收盘,
  },
  line: {
    length: 30,
  },
});
chart.annotation().dataMarker({
  top: true,
  position: [min.时间, min.收盘],
  text: {
    content: "全部谷值：" + min.收盘,
  },
  line: {
    length: 50,
  },
});
  });
},5000)
// console.log(data);
// annotation
  chart.render();
  },
  destroyed:function(){
    window.clearInterval(this.time_id)
  },
  props: {
    symbol: String,
  },
};

</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
