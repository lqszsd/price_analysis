<template>
  <div>
    <el-form ref="form" :model="form" label-width="80px">
      <el-form-item label="抄底价格">
        <el-input type="number" v-model="form.plan_copy_the_bottom"></el-input>
      </el-form-item>
      <el-form-item label="卖出价格">
        <el-input type="number" v-model="form.plan_sale"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onSubmit">立即创建</el-button>
        <el-button>取消</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>
<script>
export default {
  name: "AddStrategy",
  components: {},
  data: function () {
    return {
      form: {
        cost: 0.00,
        sharehoshareholdinglding: 0.00,
      },
    };
  },
  methods: {
    onSubmit: function () {
        console.log(this.shareholding);
      this.$post("/api/add_strategy",{
          symbol:this.symbol,
          price:this.price,
          name:this.name,
          plan_copy_the_bottom:this.form.plan_copy_the_bottom,
          plan_sale:this.form.plan_sale
      }).then((data) => {
          console.log(data)
    });
    },
  },
  props: {
    symbol: String,
    price: Number,
    name: String,
  },
};
</script>
