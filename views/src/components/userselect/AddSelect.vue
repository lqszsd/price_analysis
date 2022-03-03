<template>
  <div>
    <el-form ref="form" :model="form" label-width="80px">
      <el-form-item label="成本价">
        <el-input type="number" v-model="form.cost"></el-input>
      </el-form-item>
      <el-form-item label="持有数量">
        <el-input type="number" v-model="form.shareholding"></el-input>
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
  name: "AddSelect",
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
      this.$post("/api/add_select",{
          symbol:this.symbol,
          price:this.price,
          name:this.name,
          cost:this.form.cost,
          shareholding:this.form.shareholding
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
