<template>
  <div class="main_con">
    <p class="result"><em>"<span id="searchword">{{ keyword }}</span>"</em>에 대한 통합검색 결과 입니다. (<span id="total_cnt">{{ totalCnt }}</span>건)</p>
    <div id="broad_order_sec" class="tit_area" style="">
      <h4 id="broad_title_text"><a href="javascript:;">생방송</a><span id="broad_cnt">({{ liveCnt }})</span><img src="//res.afreecatv.com/images/aftv_search/ico_live.gif" alt="live"></h4>
      <div class="select_v">
          <a href="javascript:;" class="v" id="broad_order_text"><strong>정확도순</strong><em></em></a>
          <ul class="sub_list" style="display: none;">
          <li><a href="javascript:;">정확도순</a></li>
          <li><a href="javascript:;">시청인원순</a></li>
          <li><a href="javascript:;">최신순</a></li>
          <li><a href="javascript:;">고화질순</a></li>
          <li><a href="javascript:;" class="last">고해상도순</a></li>
          </ul>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'TotalSearch',
  data () {
    return {
      msg: 'Welcome to Your Vue.js Appqwe',
      totalCnt: 0,
      keyword: ''
    }
  },
  created: function () {
    this.setting()
  },
  methods: {
    setting () {
      this.keyword = this.$route.query.keyword
      this.$store.dispatch('callChangeTab', { tab: 'totalSearch' })
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1, h2 {
  font-weight: normal;
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
