
//这里只写了index，没有写扩展名.vue， 因为webpack会自动去找index.vue，这样我们的代码就显得比较简洁
//import HelloWorld from './components/HelloWorld.vue'
import HotList from './components/HotList.vue'
import Recommend from './components/Recommend.vue'
import Test from './components/Test.vue'
import StrategyList from './components/strategy/StrategyList.vue'
import RecommendAll from './components/recommend/RecommendAll.vue'
const routes = [
  // 定义一个数组，里面存放我们路由和组件的映射关系
  {
    path: '/', // 比如这里 / 表示根路由， 对应index这个组件
    name: 'Home',
    component: HotList,
  },
  {
    path: '/recommend',
    name: 'test',
    component: Recommend,
  },
  {
    path: '/my_choose',
    name: 'test',
    component: Test,
  },
  {
    path: '/my_strategy',
    name: 'test',
    component: StrategyList,
  },
  {
    path:"/all_recommend",
    name: 'test',
    component: RecommendAll,
  }
]

export default routes // 最后使用原生的js导出我们创建的router对象，方便在main.js中使用