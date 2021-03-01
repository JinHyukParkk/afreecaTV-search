<template>
  <div v-if="expose">
    <div class="search_no_txt" id="noSearchDiv2">
      <span>찾으시려는 검색어를 입력해주세요!</span>
    </div>
    <div class="tit_area bor" id="noSearchDiv4">
      <h4>추천 생방송<img src="//res.afreecatv.com/images/aftv_search/ico_live.gif" alt="live"></h4>
    </div>
    <div class="recommend" id="noSearchDiv5">
    </div>
    <div class="tit_area bor" id="noSearchDiv6">
      <h4>추천 VOD<span></span></h4>
    </div>
    <div class="recommend" id="noSearchDiv7">
    </div>
  </div>
</template>

<script>
export default {
  name: 'NoData',
  data () {
    return {
      msg: 'Vod 페이지',
      expose: false
    }
  },
  created: function () {
    this.setting()
  },
  methods: {
    setting () {
      if (!this.$route.query.keyword) {
        this.expose = true
      }
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
